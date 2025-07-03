package utils

type PageInfo struct {
	Total      int `json:"total"`
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	TotalPages int `json:"total_pages"`
}

type Response struct {
	Success  bool      `json:"success"`
	Message  string    `json:"message"`
	Errors   any       `json:"errors,omitempty"`
	Results  any       `json:"results,omitempty"`
	PageInfo *PageInfo `json:"page_info,omitempty"`
}
