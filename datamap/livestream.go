package datamap

//Live Stream is a stream that can be used by the Live Video
type LiveStream struct {
	fullBase

	ExternalID       string `json:"external_id"`
	PrimaryURL       string `json:"primaryUrl"`
	BackupURL        string `json:"backupUrl"`
	Name             string `json:"streamName"`
	PrimaryEncoder   string `json:"primaryEncoder"`
	SecondaryEncoder string `json:"secondaryEncoder"`
	UserName         string `json:"userName"`
	Password         string `json:"password"`
	Transcoded       bool   `json:"streamTranscoded"`
}

func (l *LiveStream) GetType() ItemType {
	return LiveStreamType
}

func (l *LiveStream) GetContentID() int {
	return l.ContentID
}

func (l *LiveStream) GetTeaserTitle() string {
	return l.TeaserTitle
}

func (l *LiveStream) GetTeaserText() string {
	return l.TeaserText
}

func (l *LiveStream) GetPublicationDate() string {
	return l.PublicationDate
}

func (l *LiveStream) Unmarshal(r *Receiver) error {
	if err := l.unmarshalFullBase(r); err != nil {
		return err
	}

	l.ExternalID = r.ExternalID
	l.PrimaryURL = r.PrimaryURL
	l.BackupURL = r.BackupURL
	l.Name = r.StreamName
	l.PrimaryEncoder = r.PrimaryEncoder
	l.SecondaryEncoder = r.SecondaryEncoder
	l.UserName = r.UserName
	l.Password = r.Password
	l.Transcoded = r.StreamTranscoded

	return nil
}
