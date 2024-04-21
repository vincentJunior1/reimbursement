package models

// MetaData ..
type MetaData struct {
	Code    int    `json:"code"`
	Title   string `json:"title"`
	Message string `json:"message"`
}

// MappingErrorCodes models
type MappingErrorCodes struct {
	Key     string           `json:"key"`
	Content ContentErrorCode `json:"content"`
}

// ContentErrorCode models
type ContentErrorCode struct {
	Code    string `json:"code"`
	Title   string `json:"title"`
	Message string `json:"message"`
}

// Response ..
type Response struct {
	Data  interface{} `json:"data,omitempty"`
	Meta  MetaData    `json:"meta"`
	Count int64       `json:"count,omitempty"`
	Page  *Page       `json:"page,omitempty"`
}

// Page ..
type Page struct {
	CurrentPage  int `json:"curPage,omitempty"`
	PreviousPage int `json:"prevPage,omitempty"`
	NextPage     int `json:"nextPage,omitempty"`
}
