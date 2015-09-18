package datamap

type Gallery struct {
	fullBase

	Subheadline string `json:"subheadline"`
	Items       []Item `json:"items"` //IMAGES
}

func (g *Gallery) GetType() ItemType {
	return GalleryType
}

func (g *Gallery) GetContentID() int {
	return g.ContentID
}

func (g *Gallery) GetTeaserTitle() string {
	return g.TeaserTitle
}

func (g *Gallery) GetTeaserText() string {
	return g.TeaserText
}

func (g *Gallery) GetPublicationDate() string {
	return g.PublicationDate
}

func (g *Gallery) Unmarshal(r *Receiver) error {
	if err := g.unmarshalFullBase(r); err != nil {
		return err
	}

	g.Subheadline = r.Subheadline
	items, err := unmarshalItems(r.Items)
	g.Items = items

	return err
}
