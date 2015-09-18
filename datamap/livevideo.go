package datamap

// LiveVideo represents a stream
type LiveVideo struct {
	fullBase

	ExternalID  string     `json:"external_id"`
	Subheadline string     `json:"subheadline"`
	LiveStream  LiveStream `json:"livestream"`
	StreamURL   string     `json:"m3u8"`
	ShowAds     bool       `json:"show_ads"`
	// backward compatibility
	Stream string `json:"stream"`
}

func (l *LiveVideo) GetType() ItemType {
	return LiveVideoType
}

func (l *LiveVideo) GetContentID() int {
	return l.ContentID
}

func (l *LiveVideo) GetTeaserTitle() string {
	return l.TeaserTitle
}

func (l *LiveVideo) GetTeaserText() string {
	return l.TeaserText
}

func (l *LiveVideo) GetPublicationDate() string {
	return l.PublicationDate
}

func (l *LiveVideo) Unmarshal(r *Receiver) error {
	if err := l.unmarshalFullBase(r); err != nil {
		return err
	}

	l.ExternalID = r.ExternalID
	l.Subheadline = r.Subheadline
	l.LiveStream = r.LiveStream
	l.StreamURL = r.StreamURL
	l.ShowAds = r.ShowAds

	l.Stream = r.Stream

	return nil
}
