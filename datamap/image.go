package datamap

// Image represents an IB image content piece
type Image struct {
	fullBase

	Caption string     `json:"caption"`
	URLs    []ImageURL `json:"urls"`
	//backward compatibility
	AltText string `json:"alt_text"`
}

func (i *Image) GetType() ItemType {
	return ImageType
}

func (i *Image) GetContentID() int {
	return i.ContentID
}

func (i *Image) GetTeaserTitle() string {
	return i.TeaserTitle
}

func (i *Image) GetTeaserText() string {
	return i.TeaserText
}

func (i *Image) GetPublicationDate() string {
	return i.PublicationDate
}

func (i *Image) Unmarshal(r *Receiver) error {
	if err := i.unmarshalFullBase(r); err != nil {
		return err
	}

	i.Caption = r.Caption
	i.URLs = r.URLs

	i.AltText = r.AltText

	return nil
}

// ImageURL is a URL flavor for an image
type ImageURL struct {
	Version string `json:"version"`
	Height  int    `json:"height"`
	Width   int    `json:"width"`
	URL     string `json:"url"`
	Mime    string `json:"mime"`
}
