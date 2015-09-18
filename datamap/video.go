package datamap

// Video represents an IB video
type Video struct {
	fullBase

	ExternalID     string        `json:"external_id"`
	Caption        string        `json:"caption"`
	Subheadline    string        `json:"subheadline"`
	Flavors        []VideoFlavor `json:"flavors"`
	StreamURL      string        `json:"m3u8"`
	ShowAds        bool          `json:"show_ads"`
	SharingEnabled bool          `json:"sharing_enabled"`
}

func (v *Video) GetType() ItemType {
	return VideoType
}

func (v *Video) GetContentID() int {
	return v.ContentID
}

func (v *Video) GetTeaserTitle() string {
	return v.TeaserTitle
}

func (v *Video) GetTeaserText() string {
	return v.TeaserText
}

func (v *Video) GetPublicationDate() string {
	return v.PublicationDate
}

func (v *Video) Unmarshal(r *Receiver) error {
	if err := v.unmarshalFullBase(r); err != nil {
		return err
	}

	v.ExternalID = r.ExternalID
	v.Caption = r.Caption
	v.Subheadline = r.Subheadline
	v.Flavors = r.Flavors
	v.StreamURL = r.StreamURL
	v.ShowAds = r.ShowAds
	v.SharingEnabled = r.SharingEnabled

	return nil
}

// VideoFlavor represents a flavor (i.e. resolution) of an IB Video
type VideoFlavor struct {
	Format   string `json:"video_type"`
	URL      string `json:"url"`
	Bitrate  int    `json:"bitrate"`
	Duration int    `json:"duration"`
	FileSize int    `json:"file_size"`
	Codec    string `json:"codec"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	Tags     string `json:"tags"`
}
