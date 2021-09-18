package requests

type ChangeGroupDetailRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateGroupRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
