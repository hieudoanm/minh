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

type Status struct {
	Status string `json:"status"`
}

func GetHealth() Status {
	return Status{Status: "healthy"}
}
