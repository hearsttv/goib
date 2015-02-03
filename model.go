package goib

// ItemType is the type of content encapsulated by the object
type ItemType string

const (
	// ArticleType item type
	ArticleType ItemType = "ARTICLE"
	// VideoType item type
	VideoType ItemType = "VIDEO"
	// CollectionType item type
	CollectionType ItemType = "COLLECTION"
	// SearchType item type
	SearchType ItemType = "SEARCH"
	// ImageType item type
	ImageType ItemType = "IMAGE"
	// GalleryType item type
	GalleryType ItemType = "GALLERY"
	// MapType item type
	MapType ItemType = "MAP"
	// LivevideoType item type
	LivevideoType = "LIVEVIDEO"
	// AudioType item type
	AudioType = "AUDIO"
	// ExternalType external content item type
	ExternalType = "EXTERNAL_CONTENT"
	// ClosingsType item type
	ClosingsType = "CLOSINGS"
	// HTMLType HTML item type
	HTMLType = "HTML"
	// PersonType Perosn (AKA Author) item type
	PersonType = "PERSON"
	// SettingsType is someone's idiot idea of a joke
	SettingsType = ""
)

type ClosingsFilter string

const (
	ClosingsAll    ClosingsFilter = "all"
	ClosingsClosed ClosingsFilter = "closed"
	ClosingsCount  ClosingsFilter = "count"
	ClosingsInst   ClosingsFilter = "institution"
)

// Receiver captures a type-agnostic representation of an API response as a
// step in processing a response. Its fields are a superset of all content fields,
// so any type can be captured and derived from this struct.
// This object should not be used outside of the IB API classes. It is exposed
// only to facilitate JSON unmarshalling.
type Receiver struct {
	Type                ItemType            `json:"type"`
	ContentID           int                 `json:"content_id"`
	ContentName         string              `json:"content_name"`
	CollectionName      string              `json:"collection_name"`
	Items               []Receiver          `json:"items"`
	TeaserTitle         string              `json:"teaser_title"`
	TeaserText          string              `json:"teaser_text"`
	TeaserImage         string              `json:"teaser_image"`
	PublicationDate     int64               `json:"publication_date"`
	Title               string              `json:"title"`
	Subheadline         string              `json:"subheadline"`
	Text                string              `json:"article_text"`
	Author              string              `json:"author"`
	Flavors             []VideoFlavor       `json:"flavors"`
	StartIndex          int                 `json:"start_index"`
	TotalCount          int                 `json:"total_count"`
	Keywords            string              `json:"keywords"`
	AltText             string              `json:"alt_text"`
	Caption             string              `json:"caption"`
	URLs                []ImageURL          `json:"urls"`
	Media               []Receiver          `json:"media"`
	AuthorObjects       []Person            `json:"author_objects"`
	Settings            []map[string]string `json:"settings"`
	Copyright           string              `json:"copyright"`
	CopyrightObjects    []CopyrightObject   `json:"copyright_objects"`
	ExternalContent     string              `json:"external_content"`
	Code                string              `json:"code"`
	CanonicalURL        string              `json:"canonical_url"`
	URL                 string              `json:"url"`
	StaticMap           string              `json:"static_map"`
	InteractiveMap      string              `json:"interactive_map"`
	Email               string              `json:"email"`
	Bio                 string              `json:"biography"`
	FullName            string              `json:"full_name"`
	Struct              []interface{}       `json:"struct"`
	Photo               []Image             `json:"photo"`
	LivevideoStream     string              `json:"m3u8"`
	NavContext          []string            `json:"navigation_context"`
	AnalyticsCategory   string              `json:"analytics_category"`
	AdvertisingCategory string              `json:"advertising_category"`
}

// Item is the base type of all items. It is not used outside the IB package, as
// we return full objects, partially populated
type Item interface {
	GetType() ItemType
	GetContentID() int
	GetTeaserTitle() string
	GetTeaserText() string
	GetPublicationDate() int64
}

// Collection represents a collection of IB Items and metadata about those items
type Collection struct {
	ContentID           int               `json:"content_id"`
	TeaserTitle         string            `json:"teaser_title"`
	TeaserText          string            `json:"teaser_text"`
	TeaserImage         string            `json:"teaser_image"`
	CollectionName      string            `json:"collection_name"`
	ContentName         string            `json:"content_name"`
	TotalCount          int               `json:"total_count"`
	StartIndex          int               `json:"start_index"`
	Keywords            string            `json:"keywords"` // populated only in search results
	Items               []Item            `json:"items"`
	Settings            map[string]string `json:"settings"`
	NavContext          []string          `json:"navigation_context"`
	AnalyticsCategory   string            `json:"analytics_category"`
	AdvertisingCategory string            `json:"advertising_category"`
}

func (c *Collection) GetType() ItemType {
	return CollectionType
}

func (c *Collection) GetContentID() int {
	return c.ContentID
}

func (c *Collection) GetTeaserTitle() string {
	return c.TeaserTitle
}

func (c *Collection) GetTeaserText() string {
	return c.TeaserText
}

func (c *Collection) GetPublicationDate() int64 {
	return 0 // collections do not have pub dates
}

// Article represents an IB article
type Article struct {
	ContentID           int      `json:"content_id"`
	TeaserTitle         string   `json:"teaser_title"`
	TeaserText          string   `json:"teaser_text"`
	TeaserImage         string   `json:"teaser_image"`
	PublicationDate     int64    `json:"publication_date"`
	Title               string   `json:"title"`
	Subheadline         string   `json:"subheadline"`
	Text                string   `json:"article_text"`
	Authors             []Person `json:"author_objects"`
	Media               []Item   `json:"media"`
	CanonicalURL        string   `json:"canonical_url"`
	URL                 string   `json:"url"`
	NavContext          []string `json:"navigation_context"`
	AnalyticsCategory   string   `json:"analytics_category"`
	AdvertisingCategory string   `json:"advertising_category"`
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

func (a *Article) GetPublicationDate() int64 {
	return a.PublicationDate
}

// Video represents an IB video
type Video struct {
	ContentID           int           `json:"content_id"`
	TeaserTitle         string        `json:"teaser_title"`
	TeaserText          string        `json:"teaser_text"`
	TeaserImage         string        `json:"teaser_image"`
	PublicationDate     int64         `json:"publication_date"`
	Authors             []Person      `json:"author_objects"`
	Title               string        `json:"title"`
	Subheadline         string        `json:"subheadline"`
	Flavors             []VideoFlavor `json:"flavors"`
	CanonicalURL        string        `json:"canonical_url"`
	URL                 string        `json:"url"`
	NavContext          []string      `json:"navigation_context"`
	AnalyticsCategory   string        `json:"analytics_category"`
	AdvertisingCategory string        `json:"advertising_category"`
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

func (v *Video) GetPublicationDate() int64 {
	return v.PublicationDate
}

// VideoFlavor represents a flavor (i.e. resolution) of an IB Video
type VideoFlavor struct {
	Type     string `json:"video_type"`
	URL      string `json:"url"`
	Bitrate  int    `json:"bitrate"`
	Duration int    `json:"duration"`
	FileSize int    `json:"file_size"`
	Codec    string `json:"codec"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
}

// Image represents an IB image content piece
type Image struct {
	ContentID           int               `json:"content_id"`
	TeaserTitle         string            `json:"teaser_title"`
	TeaserText          string            `json:"teaser_text"`
	TeaserImage         string            `json:"teaser_image"`
	PublicationDate     int64             `json:"publication_date"`
	AltText             string            `json:"alt_text"`
	Caption             string            `json:"caption"`
	Author              string            `json:"author"`
	Authors             []Person          `json:"author_objects"`
	Title               string            `json:"title"`
	Keywords            string            `json:"keywords"`
	URLs                []ImageURL        `json:"urls"`
	Copyright           string            `json:"copyright"`
	CopyrightObjects    []CopyrightObject `json:"copyright_objects"`
	CanonicalURL        string            `json:"canonical_url"`
	NavContext          []string          `json:"navigation_context"`
	AnalyticsCategory   string            `json:"analytics_category"`
	AdvertisingCategory string            `json:"advertising_category"`
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

func (i *Image) GetPublicationDate() int64 {
	return i.PublicationDate
}

// ImageURL is a URL flavor for an image
type ImageURL struct {
	Version string `json:"version"`
	Height  int    `json:"height"`
	Width   int    `json:"width"`
	URL     string `json:"url"`
	Mime    string `json:"mime"`
}

// Gallery represents an image gallery
type Gallery struct {
	ContentID           int      `json:"content_id"`
	TeaserTitle         string   `json:"teaser_title"`
	TeaserText          string   `json:"teaser_text"`
	TeaserImage         string   `json:"teaser_image"`
	PublicationDate     int64    `json:"publication_date"`
	Authors             []Person `json:"author_objects"`
	Keywords            string   `json:"keywords"`
	Title               string   `json:"title"`
	Subheadline         string   `json:"subheadline"`
	Media               []Item   `json:"media"`
	Items               []Item   `json:"items"`
	CanonicalURL        string   `json:"canonical_url"`
	URL                 string   `json:"url"`
	NavContext          []string `json:"navigation_context"`
	AnalyticsCategory   string   `json:"analytics_category"`
	AdvertisingCategory string   `json:"advertising_category"`
}

func (g *Gallery) GetType() ItemType {
	return GalleryType
}

func (g *Gallery) GetContentID() int {
	return g.ContentID
}

func (g *Gallery) GetTeaserTitle() string {
	return g.TeaserTitle
}

func (g *Gallery) GetTeaserText() string {
	return g.TeaserText
}

func (g *Gallery) GetPublicationDate() int64 {
	return g.PublicationDate
}

// Audio represents an audio clip
// TODO: no idea what this looks like ATM
type Audio struct {
	ContentID           int      `json:"content_id"`
	TeaserTitle         string   `json:"teaser_title"`
	Authors             []Person `json:"author_objects"`
	CanonicalURL        string   `json:"canonical_url"`
	URL                 string   `json:"url"`
	NavContext          []string `json:"navigation_context"`
	AnalyticsCategory   string   `json:"analytics_category"`
	AdvertisingCategory string   `json:"advertising_category"`
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
	return ""
}

func (a *Audio) GetPublicationDate() int64 {
	return 0
}

// Livevideo represents a live stream
type Livevideo struct {
	ContentID           int      `json:"content_id"`
	Title               string   `json:"title"`
	Subheadline         string   `json:"subheadline"`
	TeaserTitle         string   `json:"teaser_title"`
	TeaserText          string   `json:"teaser_text"`
	PublicationDate     int64    `json:"publication_date"`
	Authors             []Person `json:"author_objects"`
	CanonicalURL        string   `json:"canonical_url"`
	URL                 string   `json:"url"`
	Media               []Item   `json:"media"`
	Stream              string   `json:"stream"`
	NavContext          []string `json:"navigation_context"`
	AnalyticsCategory   string   `json:"analytics_category"`
	AdvertisingCategory string   `json:"advertising_category"`
}

func (l *Livevideo) GetType() ItemType {
	return LivevideoType
}

func (l *Livevideo) GetContentID() int {
	return l.ContentID
}

func (l *Livevideo) GetTeaserTitle() string {
	return l.TeaserTitle
}

func (l *Livevideo) GetTeaserText() string {
	return l.TeaserText
}

func (l *Livevideo) GetPublicationDate() int64 {
	return l.PublicationDate
}

// Map represents a map
type Map struct {
	ContentID           int      `json:"content_id"`
	TeaserTitle         string   `json:"teaser_title"`
	StaticMap           string   `json:"static_map"`
	InteractiveMap      string   `json:"interactive_map"`
	NavContext          []string `json:"navigation_context"`
	AnalyticsCategory   string   `json:"analytics_category"`
	AdvertisingCategory string   `json:"advertising_category"`
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

func (m *Map) GetPublicationDate() int64 {
	return 0
}

// ExternalContent represents an external content object
type ExternalContent struct {
	ContentID       int           `json:"content_id"`
	TeaserTitle     string        `json:"teaser_title"`
	ExternalContent string        `json:"external_content"`
	Struct          []interface{} `json:"struct"`
}

func (e *ExternalContent) GetType() ItemType {
	return ExternalType
}

func (e *ExternalContent) GetContentID() int {
	return e.ContentID
}

func (e *ExternalContent) GetTeaserTitle() string {
	return e.TeaserTitle
}

func (e *ExternalContent) GetTeaserText() string {
	return ""
}

func (e *ExternalContent) GetPublicationDate() int64 {
	return 0
}

// HTMLContent represents a content object that contains a raw HTML payload
type HTMLContent struct {
	ContentID int    `json:"content_id"`
	Code      string `json:"code"`
}

func (h *HTMLContent) GetType() ItemType {
	return HTMLType
}

func (h *HTMLContent) GetContentID() int {
	return h.ContentID
}

func (h *HTMLContent) GetTeaserTitle() string {
	return ""
}

func (h *HTMLContent) GetTeaserText() string {
	return ""
}

func (h *HTMLContent) GetPublicationDate() int64 {
	return 0
}

// Person represents an IB person
type Person struct {
	ContentID           int      `json:"content_id"`
	Blurb               string   `json:"teaser_text"`
	FullName            string   `json:"full_name"`
	Title               string   `json:"title"`
	TeaserImage         string   `json:"teaser_image"`
	PublicationDate     int64    `json:"publication_date"`
	Bio                 string   `json:"biography"`
	Photo               *Image   `json:"photo,omitempty"`
	Email               string   `json:"email"`
	FacebookUsername    string   `json:"facebook_username"`
	TwitterUsername     string   `json:"twitter_username"`
	GPlusUID            string   `json:"gplus_uid"`
	StoriesCOID         int      `json:"recent_stories"`
	CanonicalURL        string   `json:"canonical_url"`
	URL                 string   `json:"url"`
	NavContext          []string `json:"navigation_context"`
	AnalyticsCategory   string   `json:"analytics_category"`
	AdvertisingCategory string   `json:"advertising_category"`
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

func (p *Person) GetPublicationDate() int64 {
	return p.PublicationDate
}

type CopyrightObject struct {
	Name string `json:"name"`
	Text string `json:"text"`
}

type ClosingsResponse struct {
	Count              ClsCount         `json:"count"`
	Institutions       []ClsInstitution `json:"institutions"`
	ClosedInstitutions []ClsInstitution `json:"closed_institutions"`
}

type ClsCount struct {
	Total           int   `json:"total"`
	PublicationDate int64 `json:"publication_date"`
}

type ClsInstitution struct {
	Name            string `json:"name"`
	PublicationDate int64  `json:"publication_date"`
	City            string `json:"city"`
	County          string `json:"county"`
	State           string `json:"state"`
	ProviderID      string `json:"provider_id"`
}

// Settings represents a collection of settings
type Settings struct {
	ContentID int               `json:"content_id"`
	Settings  map[string]string `json:"settings"`
}

func (c *Settings) GetType() ItemType {
	return SettingsType
}

func (c *Settings) GetContentID() int {
	return c.ContentID
}

func (c *Settings) GetTeaserTitle() string {
	return ""
}

func (c *Settings) GetTeaserText() string {
	return ""
}

func (c *Settings) GetPublicationDate() int64 {
	return 0 // collections do not have pub dates
}
