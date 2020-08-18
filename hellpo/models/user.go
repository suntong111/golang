package models

import (
	_ "github.com/go-sql-driver/mysql"
)
type User struct {
	Id int `json:"id"orm:"pk;auto"`
	Username string `json:"username"orm:"NOT NULL;unique"`
	Password string `json:"password"orm:"NULL:size(255)"`
	Token string `json:"token"orm:"NULL"`

}
