package datamap

// ItemType is the type of content encapsulated by the object
type ItemType string

const (
	//Migration item types
	// ArticleType item type
	ArticleType ItemType = "ARTICLE"
	// AudioType item type
	AudioType ItemType = "AUDIO"
	// CategoryType item type
	CategoryType ItemType = "CATEGORY"
	// CollectionType item type
	CollectionType ItemType = "COLLECTION"
	// CopyrightType item type
	CopyrightType ItemType = "COPYRIGHT"
	// DownloadFileType item type
	DownloadFileType ItemType = "DOWNLOAD_FILE"
	// ExternalContentType external content item type
	ExternalContentType ItemType = "EXTERNAL_CONTENT"
	// ExternalLinkType external content item type
	ExternalLinkType ItemType = "EXTERNAL_LINK"
	// GalleryType item type
	GalleryType ItemType = "GALLERY"
	// ImageType item type
	ImageType ItemType = "IMAGE"
	// LiveStreamType item type
	LiveStreamType ItemType = "LIVESTREAM"
	// LiveVideoType item type
	LiveVideoType ItemType = "LIVEVIDEO"
	// PersonType Person (AKA Author) item type
	PersonType ItemType = "PERSON"
	// VideoType item type
	VideoType ItemType = "VIDEO"

	//API item types
	// SearchType item type
	SearchType ItemType = "SEARCH"
	// MapType item type
	MapType ItemType = "MAP"
	// ClosingsType item type
	ClosingsType ItemType = "CLOSINGS"
	// HTMLType HTML item type
	HTMLType ItemType = "HTML"
	// TeaserType teasy teasy tease.
	TeaserType ItemType = "TEASER"
	// SettingsType is someone's idiot idea of a joke
	SettingsType ItemType = ""
)

// Item is the base type of all items. It is not used outside the IB package, as
// we return full objects, partially populated
type Item interface {
	GetType() ItemType
	GetContentID() int
	GetTeaserTitle() string
	GetTeaserText() string
	GetPublicationDate() string
	Unmarshal(r *Receiver) error
	//NOTE: Despite StationName added in the types
	//      I don't know if this model will be goib compatible
	//		because of the GetPublicationDate type.
	//		so, I'll have to check with an API dev if
	//		all that makes sense
	//GetStationName() string
	//SetStationName(string)
}
