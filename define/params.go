package define

type Params struct {
	Url string `json:"url"`
	UA string `json:"ua"`
	TimeOut int `json:"time_out"`
	Cookie bool `json:"cookie"`
}

type Response struct {
	Header string `json:"header"`
	Body string `json:"body"`
	StatusCode int `json:"status_code"`
}