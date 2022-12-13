package model

// User users information
type User struct {
	ID       int64  //key
	Username string `gorm:"username"`
	Password string `gorm:"password"`
}
