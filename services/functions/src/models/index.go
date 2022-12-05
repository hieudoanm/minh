package models

// swagger:response errorResponse
type ErrorResponse struct {
	// Error Response
	// in: body
	Body struct {
		Message string
	}
}
