package health_service

// swagger:response statusResponse
type StatusResponse struct {
	// Status Response
	// in: body
	Body struct {
		// Required: true
		// Status of the Service
		Status string `json:"status"`
	}
}

type HealthResponse struct {
	Status string `json:"status"`
}

func GetHealth() HealthResponse {
	return HealthResponse{Status: "healthy"}
}
