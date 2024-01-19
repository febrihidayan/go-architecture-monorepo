package entities

type AclQueryParams struct {
}

type AclMeta struct {
	Roles       []*Role
	Permissions []*Permission
}

type AclUserDto struct {
	UserId      string
	Roles       []string
	Permissions []string
}

type AclPermissionDto struct {
	RoleId      string
	Permissions []string
}
