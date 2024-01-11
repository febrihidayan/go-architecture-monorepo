package request

type TemplateCreateRequest struct {
	Name string      `json:"name" validate:"required|min:3"`
	Data interface{} `json:"data" validate:"required"`
}
