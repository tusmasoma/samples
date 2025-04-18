// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameUser = "Users"

// User mapped from table <Users>
type User struct {
	ID       string `gorm:"column:id;type:char(36);primaryKey" json:"id"`
	Name     string `gorm:"column:name;type:varchar(255);not null" json:"name"`
	Email    string `gorm:"column:email;type:varchar(255);not null;uniqueIndex:email,priority:1" json:"email"`
	Password string `gorm:"column:password;type:varchar(255);not null" json:"password"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
