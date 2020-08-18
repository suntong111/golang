package models

type Upload struct {
	Id int `json:"id"orm:"pk;auto"`
	File string `json:"file"orm:"NULL"`

}