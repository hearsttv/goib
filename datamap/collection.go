package datamap

type Collection struct {
	fullBase

	Name            string              `json:"collection_name"`
	ExternalContent string              `json:"external_content"`
	Dynamic         bool                `json:"dynamic"`
	ViewType        string              `json:"view_type"`
	Items           []Collection        `json:"items"`
	StartIndex      int                 `json:"start_index"`
	TotalCount      int                 `json:"total_count"`
	Settings        []map[string]string `json:"settings"`
}

func (c *Collection) GetType() ItemType {
	return CollectionType
}

func (c *Collection) GetContentID() int {
	return c.ContentID
}

func (c *Collection) GetTeaserTitle() string {
	return c.TeaserTitle
}

func (c *Collection) GetTeaserText() string {
	return c.TeaserText
}

func (c *Collection) GetPublicationDate() string {
	return c.PublicationDate
}

func (c *Collection) Unmarshal(r *Receiver) error {
	if err := c.unmarshalFullBase(r); err != nil {
		return err
	}

	c.Name = r.CollectionName
	c.ExternalContent = r.ExternalContent
	c.Dynamic = r.Dynamic
	c.ViewType = r.ViewType
	c.StartIndex = r.StartIndex
	c.TotalCount = r.TotalCount
	c.Settings = r.Settings

	if collections, err := unmarshalCollections(r.Items); err != nil {
		return err
	} else {
		c.Items = collections
	}

	return nil
}
