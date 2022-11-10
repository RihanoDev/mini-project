package users

import (
	"api-mini-project/businesses/users"
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) users.Repository {
	return &userRepository{
		conn: conn,
	}
}

func (ur *userRepository) Register(userDomain *users.Domain) users.Domain {
	password, _ := bcrypt.GenerateFromPassword([]byte(userDomain.Password), bcrypt.DefaultCost)

	rec := FromDomain(userDomain)

	rec.Password = string(password)

	result := ur.conn.Create(&rec)

	result.Last(&rec)

	return rec.ToDomain()
}

func (ur *userRepository) GetByEmail(userDomain *users.Domain) users.Domain {
	var user User

	ur.conn.First(&user, "email = ?", userDomain.Email)

	if user.ID == 0 {
		fmt.Println("User Not Found")
		return users.Domain{}
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userDomain.Password))

	if err != nil {
		fmt.Println("Passowrd Failed!")
		return users.Domain{}
	}

	return user.ToDomain()
}

func (ur *userRepository) CheckData(userDomain *users.Domain) users.Domain {
	var data User

	ur.conn.Find(&data, "email = ?", userDomain.Email)

	return data.ToDomain()
}
