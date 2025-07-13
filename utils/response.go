package utils

type PageInfo struct {
	Total      int `json:"total_data"`
	Page       int `json:"current_page"`
	Limit      int `json:"limit"`
	TotalPages int `json:"total_pages"`
}

type Response struct {
	Success  bool      `json:"success"`
	Message  string    `json:"message"`
	Errors   any       `json:"errors,omitempty"`
	PageInfo *PageInfo `json:"page_info,omitempty"`
	Results  any       `json:"results,omitempty"`
}
