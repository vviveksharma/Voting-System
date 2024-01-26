package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DbUser struct {
	Id         uuid.UUID `gorm:"primaryKey,column:id"`
	Username   string    `gorm:"column:username;varchar(255);not null"`
	SecretKey  string    `gorm:"column:secretkey;varchar(500);not null"`
	Email      string    `gorm:"column:email;type:varchar(100);not null"`
	Fname      string    `gorm:"column:first_name;type:varchar(40);not null"`
	Lname      string    `gorm:"column:last_name;type:varchar(40);not null"`
	IsVoted    bool      `gorm:"column:is_voted;type:bool"`
	IsLoggedIn bool      `gorm:"column:is_logged_in;type:bool"`
	IsValidate bool      `gorm:"column:is_validated;type:bool"`
	VoterID    uuid.UUID `gorm:"column:voter_id"`
}

func (DbUser) TableName() string {
	return "user_tbl"
}

func (*DbUser) BeforeCreate(tx *gorm.DB) error {
	uuid := uuid.New().String()
	tx.Statement.SetColumn("Id", uuid)
	return nil
}

type DbAdmin struct {
	Id   uuid.UUID `gorm:"primaryKey,column:id"`
	Role string    `gorm:"column:role;varchar(255);not null"`
}

func (DbAdmin) TableName() string {
	return "admin_tbl"
}

func (*DbAdmin) BeforeCreate(tx *gorm.DB) error {
	uuid := uuid.New().String()
	tx.Statement.SetColumn("Id", uuid)
	return nil
}

type DbCandidate struct {
	Id    uuid.UUID `gorm:"primaryKey,column:id"`
	Name  string    `gorm:"column:name;varchar(255);not null"`
	Count int       `gorm:"column:role;index"`
	Image string    `gorm:"column:Image;varchar(500)"`
}

func (DbCandidate) TableName() string {
	return "candidate_tbl"
}

func (*DbCandidate) BeforeCreate(tx *gorm.DB) error {
	uuid := uuid.New().String()
	tx.Statement.SetColumn("Id", uuid)
	return nil
}
