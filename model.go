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
}

// Item is the base type of all items. It is not used outside the IB package, as
// we return full objects, partially populated
type Item struct {
	TeaserTitle     string `json:"teaser_title"`
	TeaserText      string `json:"teaser_text"`
	TeaserImage     string `json:"teaser_image"`
	ContentID       int    `json:"content_id"`
	PublicationDate int64  `json:"publication_date"`
}

// Collection represents a collection of IB Items and metadata about those items
type Collection struct {
	ContentID      int           `json:"content_id"`
	TeaserTitle    string        `json:"teaser_title"`
	CollectionName string        `json:"collection_name"`
	TotalCount     int           `json:"total_count"`
	StartIndex     int           `json:"start_index"`
	Items          []interface{} `json:"items"`
}

// Article represents an IB article
type Article struct {
	Item
	Title  string `json:"title"`
	Text   string `json:"article_text"`
	Author string `json:"author"`
}

// Video represents an IB video
type Video struct {
	Item
	Title   string        `json:"title"`
	Flavors []VideoFlavor `json:"flavors"`
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

// SearchResult represents the results of an API search
type SearchResult struct {
	Type       ItemType      `json:"type"`
	StartIndex int           `json:"start_index"`
	TotalCount int           `json:"total_count"`
	Keywords   string        `json:"keywords"`
	Items      []interface{} `json:"items"`
}
