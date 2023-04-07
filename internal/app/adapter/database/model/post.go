package model

// GORM model for post with relation one to many to User
type Post struct {
	ID       uint64 `gorm:"column:id;primary_key;auto_increment"`
	Title    string `gorm:"column:title;not null"`
	Content  string `gorm:"column:content;not null"`
	Created  int64  `gorm:"column:created;not null"`
	Modified int64  `gorm:"column:modified;not null"`
	Deleted  int64  `gorm:"column:deleted;not null"`
	User     uint64 `gorm:"column:user;not null"`
}
