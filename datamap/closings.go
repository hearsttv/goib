package datamap

type ClosingsFilter string

const (
	ClosingsAll    ClosingsFilter = "all"
	ClosingsClosed ClosingsFilter = "closed"
	ClosingsCount  ClosingsFilter = "count"
	ClosingsInst   ClosingsFilter = "institution"
)

type ClosingsResponse struct {
	Count              ClsCount                    `json:"count"`
	Institutions       map[string][]ClsInstitution `json:"institutions,omitempty"`
	ClosedInstitutions map[string][]ClsInstitution `json:"closed_institutions,omitempty"`
	Institution        ClsInstitution              `json:"institution"`
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
	Status          string `json:"status"`
}
