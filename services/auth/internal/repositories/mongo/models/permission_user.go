package models

type PermissionUser struct {
	PermissionId string `bson:"permission_id"`
	UserId       string `bson:"user_id"`
}

func (PermissionUser) TableName() string {
	return "permission_user"
}
