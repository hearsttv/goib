package goib

import "strconv"

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
		Log.Warn("unexpected type: %s", input.GetType())
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

	switch node.GetType() {
	case ArticleType:
		ch <- &MediaNode{node.(*Article), parent}
	case VideoType:
		ch <- &MediaNode{node.(*Video), parent}
	case ImageType:
		ch <- &MediaNode{node.(*Image), parent}
	case GalleryType:
		ch <- &MediaNode{node.(*Gallery), parent}
	case ExternalType:
		ch <- &MediaNode{node.(*ExternalContent), parent}
	case CollectionType:
		c := node.(*Collection)
		if c.Items == nil || len(c.Items) == 0 {
			return
		}
		for _, item := range c.Items {
			iterateMediaRecursive(item, c, ch)
		}

	default:
		Log.Warn("unexpected type found during iteration: %s", node.GetType())
		break
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
func GetSettings(i Item) map[string]string {
	if i.GetType() == CollectionType {
		return i.(*Collection).Settings
	}
	return make(map[string]string)
}

// SettingIsTrue returns true/false for the specified setting
func SettingIsTrue(settings map[string]string, val string) bool {
	result, err := strconv.ParseBool(settings[val])
	if err != nil {
		return false
	}
	return result
}
