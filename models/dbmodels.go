package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DbUser struct {
	Id       uuid.UUID `gorm:"primaryKey,column:id"`
	Username string    `gorm:"column:username;varchar(255);not null"`
	Email    string    `gorm:"column:email;type:varchar(100);not null"`
	Fname    string    `gorm:"column:first_name;type:varchar(40);not null"`
	Lname    string    `gorm:"column:last_name;type:varchar(40);not null"`
	VoterID  uuid.UUID `gorm:"column:voter_id"`
}

func (DbUser) TableName() string {
	return "user_tbl"
}

func (*DbUser) BeforeCreate(tx *gorm.DB) error {
	uuid := uuid.New().String()
	tx.Statement.SetColumn("Id", uuid)
	return nil
}
