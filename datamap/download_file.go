package datamap

//DownloadFile represents a link to download a file
type DownloadFile struct {
	fullBase

	LinkText string `json:"link_text"`
}

func (d *DownloadFile) GetType() ItemType {
	return DownloadFileType
}

func (d *DownloadFile) GetContentID() int {
	return d.ContentID
}

func (d *DownloadFile) GetTeaserTitle() string {
	return d.TeaserTitle
}

func (d *DownloadFile) GetTeaserText() string {
	return d.TeaserText
}

func (d *DownloadFile) GetPublicationDate() string {
	return d.PublicationDate
}

func (d *DownloadFile) Unmarshal(r *Receiver) error {
	if err := d.unmarshalFullBase(r); err != nil {
		return err
	}

	d.LinkText = r.LinkText

	return nil
}
