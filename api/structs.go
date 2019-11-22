package api

const (
	MethodNotAllow       = "Method not allow"
	BadRequest           = "Bad request"
	InternalServiceError = "Internal Service Error"
)

// ListRequest - входящий запроса для метода request
type ListRequest struct {
	Limit  int64 `json:"limit"`
	Offset int64 `json:"offset"`
}

// IDRequest - входящий запрос для методов Get и Delete
type IDRequest struct {
	ID int64 `json:"id"`
}

// InsertRequest - вхоядщий запрос для метода insert
type InsertRequest struct {
	Company  string `json:"company,omitempty"`
	Model    string `json:"model,omitempty"`
	Capacity int    `json:"capacity,omitempty"`
}

// UpdateRequest - входящий запрос для метода update
type UpdateRequest struct {
	ID       int64  `json:"id,omitempty"`
	Company  string `json:"company,omitempty"`
	Model    string `json:"model,omitempty"`
	Capacity int    `json:"capacity,omitempty"`
}
