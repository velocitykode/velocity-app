package models

import (
	"github.com/velocitykode/velocity/pkg/orm"
)

// User model represents a user in the application
type User struct {
	orm.Model[User]
	Name     string `orm:"column:name;type:varchar(255);not_null" json:"name"`
	Email    string `orm:"column:email;type:varchar(255);unique;not_null" json:"email"`
	Password string `orm:"column:password;type:varchar(255);not_null" json:"-"`
}

// TableName specifies the table name for the User model
func (User) TableName() string {
	return "users"
}
