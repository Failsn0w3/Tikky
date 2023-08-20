package controller

import "github.com/jinzhu/gorm"

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}
type User struct {
	Name     string `json:"name,omitempty"`
	Password string `json:"password,omitempty"`
	gorm.Model
}
