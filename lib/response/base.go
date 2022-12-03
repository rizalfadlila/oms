package response

type BaseResponse struct {
	Message    string                 `json:"message"`
	Code       string                 `json:"code"`
	StatusCode int                    `json:"status_code"`
	Error      string                 `json:"error,omitempty"`
	Log        map[string]interface{} `json:"-"`
}
