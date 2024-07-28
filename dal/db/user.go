package db 

import (
	"context"
)


func GetUserByName(name string) (*User, error) {
	var user User
	err := DB.Select("id, user_name, password").Where("user_name = ?", name).First(&user).Error
	return &user, err
}

func CreateUser(ctx context.Context, usr *User) error {
	return DB.WithContext(ctx).Create(usr).Error
}