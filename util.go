package goib

import (
	"strconv"

	l5g "github.com/neocortical/log5go"
)

// ExtractMedia recursively scans the supplied object for simple (displayable) media types
// A simple media item will produce a list containing that item. Collections and SearchResults
// are recursively scanned for inner media
func ExtractMedia(input Item) (media []Item) {
	media = []Item{}

	extractMediaRecursive(input, &media)

	return media
}

func extractMediaRecursive(input Item, media *[]Item) {
	if input == nil {
		return
	}

	switch input.GetType() {
	case ArticleType:
		*media = append(*media, input)
	case VideoType:
		*media = append(*media, input)
	case ImageType:
		*media = append(*media, input)
	case GalleryType:
		*media = append(*media, input)
	case CollectionType:
		c := input.(*Collection)
		if c.Items == nil || len(c.Items) == 0 {
			return
		}
		for _, item := range c.Items {
			extractMediaRecursive(item, media)
		}
	default:
		log.Warn("unexpected type: %s", input.GetType())
		break
	}

	return
}

// MediaNode holds a media item and a reference to the collection that item appears in
type MediaNode struct {
	Media            Item
	ParentCollection *Collection
}

// MediaIterator iterates over an IB response, recursively returning all media referenced
// therein in depth-first order
func MediaIterator(root Item) chan *MediaNode {
	ch := make(chan *MediaNode)
	go iterateMedia(root, ch)
	return ch
}

func iterateMedia(root Item, ch chan *MediaNode) {
	iterateMediaRecursive(root, nil, ch)
	close(ch)
}

func iterateMediaRecursive(node Item, parent *Collection, ch chan *MediaNode) {
	if node == nil {
		return
	}

	switch t := node.(type) {
	case *Article, *Video, *Livevideo, *Image, *Gallery, *Map, *Audio, *ExternalContent, *HTMLContent, *Person, *Teaser:
		ch <- &MediaNode{node, parent}
	case *Collection:
		if t.Items == nil || len(t.Items) == 0 {
			return
		}
		for _, item := range t.Items {
			iterateMediaRecursive(item, t, ch)
		}
	default:
		log.Warn("unexpected type found during iteration: %s", node.GetType())

	}

	return
}

// CollectionIterator returns a range of first-level subcollections of the supplied collection
func CollectionIterator(root Item) chan *Collection {
	ch := make(chan *Collection)
	go iterateCollections(root, ch)
	return ch
}

func iterateCollections(root Item, ch chan *Collection) {
	defer close(ch)

	if root == nil {
		return
	}

	switch root.GetType() {
	case CollectionType:
		for _, item := range root.(*Collection).Items {
			if item.GetType() == CollectionType {
				ch <- item.(*Collection)
			}
		}
	default:
		return
	}

	return
}

// GetSubcollections returns a slice of all first-level subcollections for the given IB API result
func GetSubcollections(root Item) (result []*Collection) {
	result = make([]*Collection, 0)

	switch root.GetType() {
	case CollectionType:
		for _, item := range root.(*Collection).Items {
			switch item.GetType() {
			case CollectionType:
				result = append(result, item.(*Collection))
			default:
				continue
			}
		}
	default:
		return result
	}

	return result
}

// GetSettings returns a map of settings if the the passed item is a collection
func GetSettings(i Item) (result map[string]string) {
	if i.GetType() == CollectionType {
		result = i.(*Collection).Settings
	} else if i.GetType() == SettingsType {
		result = i.(*Settings).Settings
	}
	if result == nil {
		result = make(map[string]string)
	}
	return result
}

// GetBooleanSetting returns the explicit value of a setting by its key, or a default if not set
func GetBooleanSetting(settings map[string]string, key string, dflt bool) bool {
	if settings == nil {
		return dflt
	}
	result, err := strconv.ParseBool(settings[key])
	if err != nil {
		return dflt
	}
	return result
}

// GetIntSetting returns the explicit value of a setting by its key, or a default if not set
func GetIntSetting(settings map[string]string, key string, dflt int) int {
	if settings == nil {
		return dflt
	}
	result, err := strconv.Atoi(settings[key])
	if err != nil {
		return dflt
	}
	return result
}

// SetLog sets the package-level logger
func SetLog(newLog l5g.Log5Go) {
	log = newLog
}
