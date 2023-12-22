package request

import "github.com/febrihidayan/go-architecture-monorepo/pkg/common"

type UserCreateRequest struct {
	Name     string `json:"name" validate:"required|min:3"`
	Email    string `json:"email" validate:"required|email|min:3"`
	Password string `json:"password" validate:"required|min:6"`
}

type UserUpdateRequest struct {
	ID       *common.ID `param:"id" validate:"required"`
	Name     string     `json:"name" validate:"required|min:3"`
	Email    string     `json:"email" validate:"required|email|min:3"`
	Password string     `json:"password" validate:"nullable|min:6"`
}

type UserQueryParams struct {
	Search  string `query:"search"`
	Page    int    `query:"page" default:"1"`
	PerPage int    `query:"per_page" default:"10"`
}
