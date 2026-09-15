package http

type greetingResponse struct {
	Message string `json:"message"`
}

func newGreetingResponse(message string) greetingResponse {
	return greetingResponse{Message: message}
}
