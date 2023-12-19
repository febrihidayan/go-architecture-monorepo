package request

import "github.com/febrihidayan/go-architecture-monorepo/pkg/common"

type PermissionCreateRequest struct {
	Name        string `json:"name" validate:"required|min:3"`
	DisplayName string `json:"display_name" validate:"required|min:3"`
	Description string `json:"description"`
}

type PermissionUpdateRequest struct {
	ID          *common.ID `param:"id" validate:"required"`
	Name        string     `json:"name" validate:"required|min:3"`
	DisplayName string     `json:"display_name" validate:"required|min:3"`
	Description string     `json:"description"`
}

type PermissionQueryParams struct {
	Search  string `query:"search"`
	Page    int    `query:"page" default:"1"`
	PerPage int    `query:"per_page" default:"10"`
}
