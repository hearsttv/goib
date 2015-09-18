package datamap

// ExternalContent represents an external content object
type ExternalContent struct {
	fullBase

	Content         string `json:"external_content"`
	Transformed     string `json:"transformed_external_content"`
	TransformerType string `json:"transformer_type"`
	LinkText        string `string:"link_text"`
}

func (e *ExternalContent) GetType() ItemType {
	return ExternalContentType
}

func (e *ExternalContent) GetContentID() int {
	return e.ContentID
}

func (e *ExternalContent) GetTeaserTitle() string {
	return e.TeaserTitle
}

func (e *ExternalContent) GetTeaserText() string {
	return e.TeaserText
}

func (e *ExternalContent) GetPublicationDate() string {
	return e.PublicationDate
}

func (e *ExternalContent) Unmarshal(r *Receiver) error {
	if err := e.unmarshalFullBase(r); err != nil {
		return err
	}

	e.Content = r.ExternalContent
	e.Transformed = r.TransformedExternalContent
	e.TransformerType = r.TransformedType
	e.LinkText = e.LinkText

	return nil
}
