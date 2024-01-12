package request

type DeviceTokenCreateRequest struct {
	Token  string `json:"token" validate:"required"`
	OsName string `json:"os_name" validate:"required"`
}
