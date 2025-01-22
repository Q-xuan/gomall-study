package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email          string `gorm:"uniqueIndex:idx_user_email;type:varchar(255) not null"`
	PasswordHashed string `gorm:"type:varchar(255) not null"`
}

func (User) TableName() string {
	return "user"
}

func Create(db *gorm.DB, user *User) error {
	return db.Create(user).Error
}

func GetUserByEmail(db *gorm.DB, email string) (*User, error) {
	user := &User{}
	err := db.Where("email =?", email).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
