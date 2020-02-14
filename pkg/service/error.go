package service

type ConnectionError struct {
	Message string `json:"message"`
}

func (c ConnectionError) Error() string {
	return c.Message
}

type errorWrapper struct {
	Error string `json:"error"`
}
