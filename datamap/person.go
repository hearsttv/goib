package datamap

// Person represents an IB person
type Person struct {
	fullBase

	Blurb         string `json:"teaser_text"`
	FullName      string `json:"full_name"`
	Email         string `json:"email"`
	Bio           string `json:"biography"`
	Photo         Image  `json:"photo"`
	FacebookURL   string `json:"facebook_url"`
	GooglePlusURL string `json:"google_plus_url"`
	TwitterURL    string `json:"twitter_url"`
	// backward compatibility
	FacebookUsername string `json:"facebook_username"`
	FacebookUID      string `json:"facebook_uid"`
	TwitterUsername  string `json:"twitter_username"`
	GPlusUID         string `json:"gplus_uid"`
	StoriesCOID      int    `json:"recent_stories"`
}

func (p *Person) GetType() ItemType {
	return PersonType
}

func (p *Person) GetContentID() int {
	return p.ContentID
}

func (p *Person) GetTeaserTitle() string {
	return p.FullName
}

func (p *Person) GetTeaserText() string {
	return p.Blurb
}

func (p *Person) GetPublicationDate() string {
	return p.PublicationDate
}

func (p *Person) Unmarshal(r *Receiver) error {
	if err := p.unmarshalFullBase(r); err != nil {
		return err
	}

	p.Blurb = r.TeaserText
	p.FullName = r.FullName
	p.Email = r.Email
	p.Bio = r.Bio
	p.Photo = r.Photo
	p.FacebookURL = r.FacebookURL
	p.GooglePlusURL = r.GooglePlusURL
	p.TwitterURL = r.TwitterURL

	p.FacebookUsername = r.FacebookUsername
	p.FacebookUID = r.FacebookUID
	p.TwitterUsername = r.TwitterUsername
	p.GPlusUID = r.GPlusUID
	p.StoriesCOID = r.StoriesCOID

	return nil
}
