package models

import (
	DB "gin/database"
)

type User struct {
	ID    uint   `gorm:"primaryKey" json:"id,string,omitempty"`
	Name  string `gorm:"type:string" json:"name,string,omitempty"`
	Phone string `gorm:"type:string" json:"phone,string"`
}

func init() {

	err := DB.Con().AutoMigrate(&User{})
	if err != nil {
		panic(err.Error())
	}

}
