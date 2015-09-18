package datamap

import (
	"fmt"
)

// Receiver captures a type-agnostic representation of an API response as a
// step in processing a response. Its fields are a superset of all content fields,
// so any type can be captured and derived from this struct.
// This object should not be used outside of the IB API classes. It is exposed
// only to facilitate JSON unmarshalling.
//
// - Text was article_text but now there are 2 fields (article_text and only text)
//   so they got separated inside Receiver
// - maybe the outer publication date is now a string
//   and only the inner kept as integers (timestamp) what now?
type Receiver struct {
	Type                ItemType      `json:"type"`
	Hash                string        `json:"hash"`
	ContentID           int           `json:"content_id"`
	ContentName         string        `json:"content_name"`
	Title               string        `json:"title"`
	PublicationDate     string        `json:"publication_date"`
	CreationDate        string        `json:"creation_date"`
	NavContext          []string      `json:"navigation_context"`
	ValidFrom           string        `json:"valid_from"`
	ValidTo             string        `json:"valid_to"`
	AnalyticsCategory   string        `json:"analytics_category"`
	AdvertisingCategory string        `json:"advertising_category"`
	TeaserTitle         string        `json:"teaser_title"`
	TeaserText          string        `json:"teaser_text"`
	TeaserImage         string        `json:"teaser_image"`
	Author              string        `json:"author"`
	Authors             []Person      `json:"author_objects"`
	EditorialComment    string        `json:"editorial_comment"`
	Copyright           string        `json:"copyright"`
	Copyrights          []Copyright   `json:"copyright_objects"`
	Media               []Receiver    `json:"media"`
	CanonicalURL        string        `json:"canonical_url"`
	URL                 string        `json:"url"`
	Categories          []interface{} `json:"categories"`
	Struct              []interface{} `json:"struct"`
	CommentingEnabled   bool          `json:"commenting_enabled"`
	NoFollow            bool          `json:"no_follow"`
	NotSearchable       bool          `json:"not_searchable"`
	Period              string        `json:"period"`
	Keywords            string        `json:"keywords"`

	//shared at least by 2 structs
	ExternalID      string              `json:"external_id"`
	ExternalContent string              `json:"external_content"`
	Subheadline     string              `json:"subheadline"`
	StreamURL       string              `json:"m3u8"`
	ShowAds         bool                `json:"show_ads"`
	Settings        []map[string]string `json:"settings"`
	StartIndex      int                 `json:"start_index"`
	LinkText        string              `json:"link_text"`
	Items           []Receiver          `json:"items"`
	Caption         string              `json:"caption"`

	//Article only
	Dateline     string     `json:"author_location"`
	ArticleText  string     `json:"article_text"`
	RelatedMedia []Receiver `json:"related_media"`

	//Category only
	Identifier string     `json:"identifier"`
	Children   []Category `json:"children"`

	//Collection only
	CollectionName string `json:"collection_name"`
	Dynamic        bool   `json:"dynamic"`
	ViewType       string `json:"view_type"`
	TotalCount     int    `json:"total_count"`

	//Copyright only
	Name string `json:"name"`
	Text string `json:"text"`

	//ExternalContent only
	TransformedExternalContent string `json:"transformed_external_content"`
	TransformedType            string `json:"transformed_type"`

	//Image only
	URLs    []ImageURL `json:"urls"`
	AltText string     `json:"alt_text"`

	//LiveStream only
	PrimaryURL       string `json:"primaryUrl"`
	BackupURL        string `json:"backupUrl"`
	StreamName       string `json:"streamName"`
	PrimaryEncoder   string `json:"primaryEncoder"`
	SecondaryEncoder string `json:"secondaryEncoder"`
	UserName         string `json:"userName"`
	Password         string `json:"password"`
	StreamTranscoded bool   `json:"streamTranscoded"`

	//LiveVideo only
	LiveStream LiveStream `json:"livestream"`
	Stream     string     `json:"stream"`

	//Person only
	FullName         string `json:"full_name"`
	Email            string `json:"email"`
	Bio              string `json:"biography"`
	Photo            Image  `json:"photo"`
	FacebookURL      string `json:"facebook_url"`
	GooglePlusURL    string `json:"google_plus_url"`
	TwitterURL       string `json:"twitter_url"`
	FacebookUsername string `json:"facebook_username"`
	FacebookUID      string `json:"facebook_uid"`
	TwitterUsername  string `json:"twitter_username"`
	GPlusUID         string `json:"gplus_uid"`
	StoriesCOID      int    `json:"recent_stories"`

	//Video only
	Flavors        []VideoFlavor `json:"flavors"`
	SharingEnabled bool          `json:"sharing_enabled"`

	//MODELS WITH NO SPECIFIC PROPERTIES
	//Audio
	//DownloadFile
	//ExternalLink

	//receiver internal submodules
	Code           string            `json:"code"`
	StaticMap      string            `json:"static_map"`
	InteractiveMap string            `json:"interactive_map"`
	Target         *Receiver         `json:"target"`
	Captions       map[string]string `json:"captions"` // not from IB, but needed for UnmarshalReceiver()
}

//Maps the data of Receiver to the specific struct
func (r *Receiver) Unmarshal() (Item, error) {
	var item Item

	//maps string type to the correct struct
	switch r.Type {
	// migratin types
	case ArticleType:
		item = &Article{}
	case AudioType:
		item = &Audio{}
	case CategoryType:
		item = &Category{}
	case CollectionType:
		item = &Collection{}
	case CopyrightType:
		item = &Copyright{}
	case DownloadFileType:
		item = &DownloadFile{}
	case ExternalContentType:
		item = &ExternalContent{}
	case ExternalLinkType:
		item = &ExternalLink{}
	case GalleryType:
		item = &Gallery{}
	case ImageType:
		item = &Image{}
	case LiveStreamType:
		item = &LiveStream{}
	case LiveVideoType:
		item = &LiveVideo{}
	case PersonType:
		item = &Person{}
	case VideoType:
		item = &Video{}
	// api types
	case SearchType:
		item = &Collection{}
	case MapType:
		item = &Map{}
	case HTMLType:
		item = &HTMLContent{}
	case SettingsType:
		item = &Settings{}
	case TeaserType:
		item = &Teaser{}
	default:
		return nil, fmt.Errorf("unknonwn response type for obj %d: %s", r.ContentID, r.Type)
	}

	err := item.Unmarshal(r)

	return item, err
}

// this struct here is a list of the common properties shared
// among all content types
type bareBase struct {
	Type                ItemType `json:"type"`
	Hash                string   `json:"hash"`
	ContentID           int      `json:"content_id"`
	ContentName         string   `json:"content_name"`
	Title               string   `json:"title"`
	CreationDate        string   `json:"creation_date"`
	PublicationDate     string   `json:"publication_date"`
	NavContext          []string `json:"navigation_context"`
	ValidFrom           string   `json:"valid_from"`
	ValidTo             string   `json:"valid_to"`
	AnalyticsCategory   string   `json:"analytics_category"`
	AdvertisingCategory string   `json:"advertising_category"`
	StationName         string   `json:"station_name"` //new api adition that seems to be set/get locally
}

func (b *bareBase) unmarshalBareBase(r *Receiver) {
	b.Hash = r.Hash
	b.ContentID = r.ContentID
	b.ContentName = r.ContentName
	b.Title = r.Title
	b.PublicationDate = r.PublicationDate
	b.CreationDate = r.CreationDate
	b.NavContext = r.NavContext
	b.ValidFrom = r.ValidFrom
	b.ValidTo = r.ValidTo
	b.AnalyticsCategory = r.AnalyticsCategory
	b.AdvertisingCategory = r.AdvertisingCategory
}

// this struct here is a list of the most common properties shared
// among some content types (the barebone properties are in bareBase)
type fullBase struct {
	bareBase

	TeaserTitle       string        `json:"teaser_title"`
	TeaserText        string        `json:"teaser_text"`
	TeaserImage       string        `json:"teaser_image"`
	Author            string        `json:"author"`
	Authors           []Person      `json:"author_objects"`
	EditorialComment  string        `json:"editorial_comment"`
	Copyright         string        `json:"copyright"`
	Copyrights        []Copyright   `json:"copyright_objects"`
	Media             []Item        `json:"media"`
	CanonicalURL      string        `json:"canonical_url"`
	URL               string        `json:"url"`
	Categories        []interface{} `json:"categories"`
	Struct            []interface{} `json:"struct"`
	CommentingEnabled bool          `json:"commenting_enabled"`
	NoFollow          bool          `json:"no_follow"`
	NotSearchable     bool          `json:"not_searchable"`
	Period            string        `json:"period"`
	Keywords          string        `json:"keywords"`
}

func (f *fullBase) unmarshalFullBase(r *Receiver) error {
	f.unmarshalBareBase(r)

	f.TeaserTitle = getTeaserTitle(r)
	f.TeaserText = r.TeaserText
	f.TeaserImage = r.TeaserImage
	f.Author = r.Author
	f.Authors = r.Authors
	f.EditorialComment = r.EditorialComment
	f.Copyright = r.Copyright
	f.CanonicalURL = r.CanonicalURL
	f.URL = r.URL
	f.Categories = r.Categories
	f.Struct = r.Struct
	f.CommentingEnabled = r.CommentingEnabled
	f.NoFollow = r.NoFollow
	f.NotSearchable = r.NotSearchable
	f.Period = r.Period
	f.Keywords = r.Keywords
	medias, err := unmarshalItems(r.Media)
	f.Media = medias

	return err
}

// Map represents a map
type Map struct {
	Type                ItemType `json:"type"`
	ContentID           int      `json:"content_id"`
	PublicationDate     string   `json:"publication_date"`
	TeaserTitle         string   `json:"teaser_title"`
	TeaserText          string   `json:"teaser_text"`
	Title               string   `json:"title"`
	Subheadline         string   `json:"subheadline"`
	StaticMap           string   `json:"static_map"`
	InteractiveMap      string   `json:"interactive_map"`
	CanonicalURL        string   `json:"canonical_url"`
	URL                 string   `json:"url"`
	NavContext          []string `json:"navigation_context"`
	AnalyticsCategory   string   `json:"analytics_category"`
	AdvertisingCategory string   `json:"advertising_category"`
	StationName         string   `json:"station_name"`
}

func (m *Map) GetType() ItemType {
	return MapType
}

func (m *Map) GetContentID() int {
	return m.ContentID
}

func (m *Map) GetTeaserTitle() string {
	return m.TeaserTitle
}

func (m *Map) GetTeaserText() string {
	return ""
}

func (m *Map) GetPublicationDate() string {
	return m.PublicationDate
}

func (m *Map) Unmarshal(r *Receiver) error {
	m.Type = r.Type
	m.ContentID = r.ContentID
	m.PublicationDate = r.PublicationDate
	m.Title = r.Title
	m.Subheadline = r.Subheadline
	m.TeaserTitle = getTeaserTitle(r)
	m.TeaserText = r.TeaserText
	m.StaticMap = r.StaticMap
	m.InteractiveMap = r.InteractiveMap
	m.CanonicalURL = r.CanonicalURL
	m.URL = r.URL
	m.NavContext = r.NavContext
	m.AnalyticsCategory = r.AnalyticsCategory
	m.AdvertisingCategory = r.AdvertisingCategory

	return nil
}

// HTMLContent represents a content object that contains a raw HTML payload
type HTMLContent struct {
	Type                ItemType `json:"type"`
	ContentID           int      `json:"content_id"`
	PublicationDate     string   `json:"publication_date"`
	Code                string   `json:"code"`
	TeaserTitle         string   `json:"teaser_title"`
	NavContext          []string `json:"navigation_context"`
	AnalyticsCategory   string   `json:"analytics_category"`
	AdvertisingCategory string   `json:"advertising_category"`
	StationName         string   `json:"station_name"`
}

func (h *HTMLContent) GetType() ItemType {
	return HTMLType
}

func (h *HTMLContent) GetContentID() int {
	return h.ContentID
}

func (h *HTMLContent) GetTeaserTitle() string {
	return h.TeaserTitle
}

func (h *HTMLContent) GetTeaserText() string {
	return ""
}

func (h *HTMLContent) GetPublicationDate() string {
	return h.PublicationDate
}

func (h *HTMLContent) Unmarshal(r *Receiver) error {
	h.Type = r.Type
	h.ContentID = r.ContentID
	h.PublicationDate = r.PublicationDate
	h.Code = r.Code
	h.TeaserTitle = getTeaserTitle(r)
	h.NavContext = r.NavContext
	h.AnalyticsCategory = r.AnalyticsCategory
	h.AdvertisingCategory = r.AdvertisingCategory

	return nil
}

// Settings represents a collection of settings
type Settings struct {
	ContentID   int               `json:"content_id"`
	Settings    map[string]string `json:"settings"`
	StationName string            `json:"station_name"`
}

func (s *Settings) GetType() ItemType {
	return SettingsType
}

func (s *Settings) GetContentID() int {
	return s.ContentID
}

func (s *Settings) GetTeaserTitle() string {
	return ""
}

func (s *Settings) GetTeaserText() string {
	return ""
}

func (s *Settings) GetPublicationDate() string {
	return ""
}

func (s *Settings) Unmarshal(r *Receiver) error {
	s.ContentID = r.ContentID
	if len(r.Settings) > 0 {
		s.Settings = r.Settings[0]
	}

	return nil
}

// Teaser represents ... something
type Teaser struct {
	Type                ItemType `json:"type"`
	ContentID           int      `json:"content_id"`
	Title               string   `json:"title"`
	TeaserTitle         string   `json:"teaser_title"`
	TeaserText          string   `json:"teaser_text"`
	PublicationDate     string   `json:"publication_date"`
	Authors             []Person `json:"author_objects"`
	Media               []Item   `json:"media"`
	NavContext          []string `json:"navigation_context"`
	AnalyticsCategory   string   `json:"analytics_category"`
	AdvertisingCategory string   `json:"advertising_category"`
	Target              Item     `json:"target"`
	StationName         string   `json:"station_name"`
}

func (t *Teaser) GetType() ItemType {
	return TeaserType
}

func (t *Teaser) GetContentID() int {
	return t.ContentID
}

func (t *Teaser) GetTeaserTitle() string {
	return t.TeaserTitle
}

func (t *Teaser) GetTeaserText() string {
	return t.TeaserText
}

func (t *Teaser) GetPublicationDate() string {
	return t.PublicationDate
}

func (t *Teaser) Unmarshal(r *Receiver) error {
	t.Type = r.Type
	t.ContentID = r.ContentID
	t.Title = r.Title
	t.TeaserTitle = getTeaserTitle(r)
	t.TeaserText = r.TeaserText
	t.PublicationDate = r.PublicationDate
	t.Authors = r.Authors
	t.NavContext = r.NavContext
	t.AnalyticsCategory = r.AnalyticsCategory
	t.AdvertisingCategory = r.AdvertisingCategory

	if r.Target == nil {
		return fmt.Errorf("teaser object missing target %d", t.ContentID)
	}

	if target, err := r.Target.Unmarshal(); err != nil {
		return err
	} else {
		t.Target = target
	}

	media, err := unmarshalItems(r.Media)
	t.Media = media

	return err
}
