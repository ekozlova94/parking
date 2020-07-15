package forms

type Response struct {
	Success bool `json:"success,omitempty"`
	Pass    bool `json:"pass,omitempty"`
}

func NewResponse(s, p bool) *Response {
	return &Response{
		Success: s,
		Pass:    p,
	}
}
