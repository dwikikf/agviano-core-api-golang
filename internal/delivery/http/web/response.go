package web

type WebResponse struct {
	Code    int            `json:"code"`
	Message string         `json:"message"`
	Error   any            `json:"error,omitempty"`
	Data    any            `json:"data,omitempty"`
	Meta    PaginationMeta `json:"meta"`
}
