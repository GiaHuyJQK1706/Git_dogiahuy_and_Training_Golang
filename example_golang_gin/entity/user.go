package entity

import "time"

// Co the dung ORM- GORM
type User struct {
	ID        int64
	FirstName string
	LastName  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
