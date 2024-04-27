package models

import "time"

type User struct{
	Id int
	Username string 
	Role  string 
	CreatedAt time.Time
	Password string
}