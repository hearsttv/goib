package goib

// ExtractMedia recursively scans the supplied object for simple (displayable) media types
// A simple media item will produce a list containing that item. Collections and SearchResults
// are recursively scanned for inner media
func ExtractMedia(input interface{}) (media []interface{}) {
	media = []interface{}{}

	extractMediaRecursive(input, &media)

	return media
}

func extractMediaRecursive(input interface{}, media *[]interface{}) {
	if input == nil {
		return
	}

	switch t := input.(type) {
	case Article:
		*media = append(*media, input)
	case Video:
		*media = append(*media, input)
	case Image:
		*media = append(*media, input)
	case Gallery:
		*media = append(*media, input)
	case Collection:
		if t.Items == nil || len(t.Items) == 0 {
			return
		}
		for _, item := range t.Items {
			extractMediaRecursive(item, media)
		}
	case SearchResult:
		if t.Items == nil || len(t.Items) == 0 {
			return
		}
		for _, item := range t.Items {
			extractMediaRecursive(item, media)
		}

	default:
		log.Warn("unexpected type: %t", t)
		break
	}

	return
}

type MediaNode struct {
	Media            interface{}
	ParentCollection interface{}
}

func Iterator(root interface{}) chan *MediaNode {
	ch := make(chan *MediaNode)
	go iterateMedia(root, ch)
	return ch
}

func iterateMedia(root interface{}, ch chan *MediaNode) {
	iterateMediaRecursive(root, nil, ch)
	close(ch)
}

func iterateMediaRecursive(node interface{}, parent interface{}, ch chan *MediaNode) {
	if node == nil {
		return
	}

	switch t := node.(type) {
	case Article:
		ch <- &MediaNode{t, parent}
	case Video:
		ch <- &MediaNode{t, parent}
	case Image:
		ch <- &MediaNode{t, parent}
	case Gallery:
		ch <- &MediaNode{t, parent}
	case Collection:
		if t.Items == nil || len(t.Items) == 0 {
			return
		}
		for _, item := range t.Items {
			iterateMediaRecursive(item, t, ch)
		}

	default:
		log.Warn("unexpected type found during iteration: %t", t)
		break
	}

	return
}
