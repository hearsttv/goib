package datamap

//Copyright type
type Copyright struct {
	fullBase

	Name string `json:"name"`
	Text string `json:"text"`
}

func (c *Copyright) GetType() ItemType {
	return CollectionType
}

func (c *Copyright) GetContentID() int {
	return c.ContentID
}

func (c *Copyright) GetTeaserTitle() string {
	return c.TeaserTitle
}

func (c *Copyright) GetTeaserText() string {
	return c.TeaserText
}

func (c *Copyright) GetPublicationDate() string {
	return c.PublicationDate
}

func (c *Copyright) Unmarshal(r *Receiver) error {
	if err := c.unmarshalFullBase(r); err != nil {
		return err
	}

	c.Name = r.Name
	c.Text = r.Text

	return nil
}
