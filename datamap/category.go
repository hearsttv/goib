package datamap

type Category struct {
	bareBase

	Identifier string              `json:"identifier"`
	Children   []Category          `json:"children"`
	Settings   []map[string]string `json:"settings"`
}

func (c *Category) GetType() ItemType {
	return CategoryType
}

func (c *Category) GetContentID() int {
	return c.ContentID
}

func (c *Category) GetTeaserTitle() string {
	return ""
}

func (c *Category) GetTeaserText() string {
	return ""
}

func (c *Category) GetPublicationDate() string {
	return c.PublicationDate
}

func (c *Category) Unmarshal(r *Receiver) error {
	c.unmarshalBareBase(r)
	c.Identifier = r.Identifier
	c.Children = r.Children
	c.Settings = r.Settings

	return nil
}
