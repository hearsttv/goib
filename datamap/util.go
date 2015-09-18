package datamap

import l5g "github.com/neocortical/log5go"

var log = l5g.Logger(l5g.LogAll)

func getTeaserTitle(r *Receiver) (title string) {
	title = r.TeaserTitle
	if title == "" {
		title = r.Title
	}
	return
}

func unmarshalItems(r []Receiver) (items []Item, err error) {
	items = []Item{}

	for _, rInner := range r {
		if i, err := rInner.Unmarshal(); err != nil {
			log.Warn("error unmarshalling sub-object: %v", err)
		} else {
			items = append(items, i)
		}
	}

	return
}

func unmarshalCollections(r []Receiver) ([]Collection, error) {
	collections := []Collection{}

	for _, rInner := range r {
		c := Collection{}
		if err := c.Unmarshal(&rInner); err != nil {
			log.Warn("error unmarshalling collection: %v", err)
			return nil, err
		}
		collections = append(collections, c)
	}

	return collections, nil
}
