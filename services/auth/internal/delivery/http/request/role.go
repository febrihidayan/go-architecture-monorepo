package request

type RoleCreateRequest struct {
	Name        string `json:"name" validate:"required|min:3"`
	DisplayName string `json:"display_name" validate:"required|min:3"`
	Description string `json:"description"`
}
