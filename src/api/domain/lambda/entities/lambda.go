package entities

import "fmt"

// Entities por lambda function Request and Response.
type (
	Request struct {
		ID    float64 `json:"id"`
		Value string  `json:"value"`
	}

	Response struct {
		Message string `json:"message"`
		Ok      bool   `json:"ok"`
	}
)

// Handler for lambda function, we use this name to show the difference between handlers.
func HandlerLambda(request Request) (Response, error) {
	return Response{
		Message: fmt.Sprintf("Process request ID %f", request.ID),
		Ok:      true,
	}, nil
}
