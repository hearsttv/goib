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
	// LivestreamType item type
	LivestreamType = "LIVESTREAM"
	// AudioType item type
	AudioType = "AUDIO"
)

// Receiver captures a type-agnostic representation of an API response as a
// step in processing a response. Its fields are a superset of all content fields,
// so any type can be captured and derived from this struct.
// This object should not be used outside of the IB API classes. It is exposed
// only to facilitate JSON unmarshalling.
type Receiver struct {
	Type            ItemType      `json:"type"`
	ContentID       int           `json:"content_id"`
	ContentName     string        `json:"content_name"`
	CollectionName  string        `json:"collection_name"`
	Items           []Receiver    `json:"items"`
	TeaserTitle     string        `json:"teaser_title"`
	TeaserText      string        `json:"teaser_text"`
	TeaserImage     string        `json:"teaser_image"`
	PublicationDate int64         `json:"publication_date"`
	Title           string        `json:"title"`
	Text            string        `json:"article_text"`
	Author          string        `json:"author"`
	Flavors         []VideoFlavor `json:"flavors"`
	StartIndex      int           `json:"start_index"`
	TotalCount      int           `json:"total_count"`
	Keywords        string        `json:"keywords"`
	AltText         string        `json:"alt_text"`
	Caption         string        `json:"caption"`
	URLs            []ImageURL    `json:"urls"`
	Media           []Receiver    `json:"media"`
	AuthorObjects   []Author      `json:"author_objects"`
}

// Item is the base type of all items. It is not used outside the IB package, as
// we return full objects, partially populated
type Item interface {
	GetType() ItemType
	GetContentID() int
	GetTeaserTitle() string
}

// Collection represents a collection of IB Items and metadata about those items
type Collection struct {
	ContentID      int    `json:"content_id"`
	TeaserTitle    string `json:"teaser_title"`
	TeaserText     string `json:"teaser_text"`
	TeaserImage    string `json:"teaser_image"`
	CollectionName string `json:"collection_name"`
	ContentName    string `json:"content_name"`
	TotalCount     int    `json:"total_count"`
	StartIndex     int    `json:"start_index"`
	Keywords       string `json:"keywords"` // populated only in search results
	Items          []Item `json:"items"`
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

// Article represents an IB article
type Article struct {
	ContentID       int      `json:"content_id"`
	TeaserTitle     string   `json:"teaser_title"`
	TeaserText      string   `json:"teaser_text"`
	TeaserImage     string   `json:"teaser_image"`
	PublicationDate int64    `json:"publication_date"`
	Title           string   `json:"title"`
	Text            string   `json:"article_text"`
	Authors         []Author `json:"author_objects"`
	Media           []Item   `json:"media"`
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

// Video represents an IB video
type Video struct {
	ContentID       int           `json:"content_id"`
	TeaserTitle     string        `json:"teaser_title"`
	TeaserText      string        `json:"teaser_text"`
	TeaserImage     string        `json:"teaser_image"`
	PublicationDate int64         `json:"publication_date"`
	Title           string        `json:"title"`
	Flavors         []VideoFlavor `json:"flavors"`
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
	ContentID       int        `json:"content_id"`
	TeaserTitle     string     `json:"teaser_title"`
	TeaserText      string     `json:"teaser_text"`
	TeaserImage     string     `json:"teaser_image"`
	PublicationDate int64      `json:"publication_date"`
	AltText         string     `json:"alt_text"`
	Caption         string     `json:"caption"`
	Author          string     `json:"author"`
	Title           string     `json:"title"`
	Keywords        string     `json:"keywords"`
	URLs            []ImageURL `json:"urls"`
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
	ContentID       int           `json:"content_id"`
	TeaserTitle     string        `json:"teaser_title"`
	TeaserText      string        `json:"teaser_text"`
	TeaserImage     string        `json:"teaser_image"`
	PublicationDate int64         `json:"publication_date"`
	Keywords        string        `json:"keywords"`
	Title           string        `json:"title"`
	Media           []interface{} `json:"media"`
	Items           []interface{} `json:"items"`
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

type Author struct {
	FullName string `json:"full_name"`
	Title    string `json:"title"`
	Email    string `json:"email"`
}

// Audio represents an audio clip
// TODO: no idea what this looks like ATM
type Audio struct {
	ContentID   int    `json:"content_id"`
	TeaserTitle string `json:"teaser_title"`
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

// Livestream represents a live stream
// TODO: no idea what this looks like ATM
type Livestream struct {
	ContentID   int    `json:"content_id"`
	TeaserTitle string `json:"teaser_title"`
}

func (l *Livestream) GetType() ItemType {
	return AudioType
}

func (l *Livestream) GetContentID() int {
	return l.ContentID
}

func (l *Livestream) GetTeaserTitle() string {
	return l.TeaserTitle
}

// Map represents a map
type Map struct {
	ContentID   int    `json:"content_id"`
	TeaserTitle string `json:"teaser_title"`
}

func (m *Map) GetType() ItemType {
	return AudioType
}

func (m *Map) GetContentID() int {
	return m.ContentID
}

func (m *Map) GetTeaserTitle() string {
	return m.TeaserTitle
}
