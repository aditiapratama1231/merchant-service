package request

//GetUsersRequest contain request
type GetUsersRequest struct{}

//GetUsersResponse contain response
type GetUsersResponse struct {
	Users string `json:"users"`
	Err   string `json:"err,omitempty"`
}
