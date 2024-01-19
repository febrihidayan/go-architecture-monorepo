package response

import (
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"
)

type AclAccessListResponse struct {
	Permissions []string `json:"permissions"`
	Roles       []string `json:"roles"`
}

type AclListResponse struct {
	Permissions []PermissionListResponse `json:"permissions"`
	Roles       []RoleListResponse       `json:"roles"`
}

func MapAclAccessListResponse(x *entities.AclMeta) AclAccessListResponse {
	var result AclAccessListResponse

	for _, item := range x.Permissions {
		result.Permissions = append(result.Permissions, item.Name)
	}

	for _, item := range x.Roles {
		result.Roles = append(result.Roles, item.Name)
	}

	return result
}

func MapAclListResponse(x *entities.AclMeta) AclListResponse {
	return AclListResponse{
		Permissions: MapPermissionListResponses(x.Permissions),
		Roles:       MapRoleListResponses(x.Roles),
	}
}
