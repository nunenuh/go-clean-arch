package model

// GORM model for user
type User struct {
	ID       uint64 `gorm:"column:id;primary_key;auto_increment"`
	Name     string `gorm:"column:name;not null"`
	Username string `gorm:"column:username;unique_index;not null"`
	Password string `gorm:"column:password;not null"`
	Email    string `gorm:"column:email;unique_index;not null"`
	Phone    string `gorm:"column:phone;unique_index;not null"`
	Created  int64  `gorm:"column:created;not null"`
	Modified int64  `gorm:"column:modified;not null"`
	Deleted  int64  `gorm:"column:deleted;not null"`
}
