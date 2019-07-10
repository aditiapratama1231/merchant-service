package request

// PaginationRequest struct contain variable to save query
type PaginationRequest struct {
	Include []string `json:"include,omitempty"`
	Page    uint32   `json:"offset,omitempty"`
	Limit   uint32   `json:"limit,omitempty"`
}

// PaginationResponse struct contain variable to response query
type PaginationResponse struct {
	CurrentPage uint32 `json:"current_page,omitempty"`
	From        uint32 `json:"from,omitempty"`
	PerPage     uint32 `json:"per_page,omitempty"`
	Total       uint32 `json:"total,omitempty"`
}