package models

type RoleUser struct {
	UserId string `bson:"user_id"`
	RoleId string `bson:"role_id"`
}

func (RoleUser) TableName() string {
	return "role_user"
}
