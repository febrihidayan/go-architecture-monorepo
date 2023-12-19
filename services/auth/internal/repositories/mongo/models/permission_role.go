package models

type PermissionRole struct {
	PermissionId string `bson:"permission_id"`
	RoleId       string `bson:"role_id"`
}

func (PermissionRole) TableName() string {
	return "permission_role"
}
