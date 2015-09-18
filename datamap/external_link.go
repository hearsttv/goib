package datamap

type ExternalLink struct {
	fullBase

	LinkText string `json:"link_text"`
}

func (e *ExternalLink) GetType() ItemType {
	return ExternalLinkType
}

func (e *ExternalLink) GetContentID() int {
	return e.ContentID
}

func (e *ExternalLink) GetTeaserTitle() string {
	return e.TeaserTitle
}

func (e *ExternalLink) GetTeaserText() string {
	return e.TeaserText
}

func (e *ExternalLink) GetPublicationDate() string {
	return e.PublicationDate
}

func (e *ExternalLink) Unmarshal(r *Receiver) error {
	if err := e.unmarshalFullBase(r); err != nil {
		return err
	}

	e.LinkText = e.LinkText

	return nil
}
