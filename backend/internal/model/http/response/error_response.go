package response

type ErrorResponse struct {
	Errors interface{} `json:"errors"`
}
