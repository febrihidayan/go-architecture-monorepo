package request

type AclUserUpdateRequest struct {
	Permissions []string `json:"permissions"`
	Roles       []string `json:"roles"`
}

type AclPermissionUpdateRequest struct {
	Permissions []string `json:"permissions"`
}
