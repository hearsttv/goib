package datamap

// Audio represents an audio clip
type Audio struct {
	fullBase

	ExternalID  string `json:"external_id"`
	Subheadline string `json:"subheadline"`
	StreamURL   string `json:"m3u8"`
	ShowAds     bool   `json:"show_ads"`
}

func (a *Audio) GetType() ItemType {
	return AudioType
}

func (a *Audio) GetContentID() int {
	return a.ContentID
}

func (a *Audio) GetTeaserTitle() string {
	return a.TeaserTitle
}

func (a *Audio) GetTeaserText() string {
	return a.TeaserText
}

func (a *Audio) GetPublicationDate() string {
	return a.PublicationDate
}

func (a *Audio) Unmarshal(r *Receiver) error {
	if err := a.unmarshalFullBase(r); err != nil {
		return err
	}

	a.ExternalID = r.ExternalID
	a.Subheadline = r.Subheadline
	a.StreamURL = r.StreamURL

	return nil
}
