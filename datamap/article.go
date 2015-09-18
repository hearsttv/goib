package datamap

// Article represents an IB article
type Article struct {
	fullBase

	Dateline     string `json:"author_location"`
	Text         string `json:"article_text"`
	Subheadline  string `json:"subheadline"`
	RelatedMedia []Item `json:"related_media"`
}

func (a *Article) GetType() ItemType {
	return ArticleType
}

func (a *Article) GetContentID() int {
	return a.ContentID
}

func (a *Article) GetTeaserTitle() string {
	return a.TeaserTitle
}

func (a *Article) GetTeaserText() string {
	return a.TeaserText
}

func (a *Article) GetPublicationDate() string {
	return a.PublicationDate
}

func (a *Article) Unmarshal(r *Receiver) error {
	if err := a.unmarshalFullBase(r); err != nil {
		return err
	}

	a.Dateline = r.Dateline
	a.Text = r.ArticleText
	a.Subheadline = r.Subheadline
	relatedMedia, err := unmarshalItems(r.RelatedMedia)
	a.RelatedMedia = relatedMedia

	return err
}
