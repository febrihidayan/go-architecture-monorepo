package request

type AclUserUpdateRequest struct {
	Permissions []string `json:"permissions"`
	Roles       []string `json:"roles"`
}
