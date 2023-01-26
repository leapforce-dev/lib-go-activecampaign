package activecampaign

// ErrorResponse stores general ActiveCampaign error response
//
type ErrorResponse struct {
	Message string `json:"message"`
	Errors  []struct {
		Title  string `json:"title"`
		Detail string `json:"detail"`
		Code   string `json:"code"`
		Error  string `json:"error"`
		Source struct {
			Pointer string `json:"pointer"`
		} `json:"source"`
	} `json:"errors"`
	StatusCode int
}
