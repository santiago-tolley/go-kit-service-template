package service

type DoStuffRequest struct {
	Value int `json:"value"`
}

type DoStuffResponse struct {
	Result int   `json:"result"`
	Err    error `json:"err"`
}

func (r *DoStuffResponse) Failed() error {
	return r.Err
}
