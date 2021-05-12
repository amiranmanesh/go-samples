package models

import (
	"awesome_webkits/database"
	"awesome_webkits/http/response"
	"awesome_webkits/utils/encrypting"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name            string    `gorm:"type:varchar(100); not null" json:"name"`
	Email           string    `gorm:"type:varchar(100);unique_index; not null" json:"email"`
	Password        string    `json:"-"`
	EmailVerifiedAt time.Time `gorm:"type:varchar(100)" json:"-"`
}

func (u *User) Save() (string /* token */, *response.Data) {

	//email validates in request

	// Check that user with this email exists
	var user User
	result := database.DB.GetDB().Scopes(user.ScopeEmail(u.Email)).First(&user)
	if err := result.Error; !errors.Is(err, gorm.ErrRecordNotFound) {
		return "", response.NewBadRequestErr("email in use", err)
	}

	u.Password = encrypting.GetHashedPassword(u.Password)

	if err := database.DB.GetDB().Create(&u).Error; err != nil {
		return "", response.NewInternalServerError(fmt.Sprintf("Error in create Item:"), err.Error())
	}

	token, err := OauthAccessToken{}.GenerateToken(u.ID)
	if err != nil {
		return "", err
	}

	return token, nil

}

func (u *User) Login(email string, password string) (string /* token */, *response.Data) {

	isFounded := database.DB.GetDB().Scopes(u.ScopeEmail(email)).Find(&u)
	if isFounded.Error != nil {
		return "", response.NewNotFoundRequestErr(fmt.Sprintf("User not found with email: %s", email), isFounded.Error)
	}

	if !encrypting.CheckPassword(u.Password, password) {
		return "", response.NewBadRequestErr("wrong email or password", nil)
	}

	token, err := OauthAccessToken{}.GenerateToken(u.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

//scopes
func (u *User) ScopeEmail(email string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("Email = ?", email)
	}
}

func (u *User) FindWithId() *response.Data {

	result := database.DB.GetDB().First(&u, u.ID)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return response.NewUnauthorizedError(fmt.Sprintf("User not found with id: %d", u.ID), result.Error)
	}
	return nil

}

func (u *User) FindWithEmail() *response.Data {
	result := database.DB.GetDB().Where("email = ?", u.Email).First(&u)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return response.NewUnauthorizedError(fmt.Sprintf("User not found with email: %d", u.ID), result.Error)
	}
	return nil

}

func (u *User) Update(name, email string) *response.Data {
	if err := u.FindWithId(); err != nil {
		return err
	}
	if name != "" {
		u.Name = name
	}
	if email != "" {
		u.Email = email
	}
	if result := database.DB.GetDB().Save(&u); result.Error != nil {
		return response.NewUnauthorizedError(fmt.Sprintf("Error in update user info: %d", u.ID), result.Error)
	}
	return nil
}

func (u *User) Verified() *response.Data {
	//user has already founded
	u.EmailVerifiedAt = time.Now()
	if result := database.DB.GetDB().Save(&u); result.Error != nil {
		return response.NewUnauthorizedError(fmt.Sprintf("Error in update user verify at: %d", u.ID), result.Error)
	}
	return nil
}

func (u *User) GetAuthUserId(token string) (uint, *response.Data) {
	return OauthAccessToken{}.GetAuthUserId(token)
}

func (u *User) GetAuthUser(token string) (*User, *response.Data) {
	return OauthAccessToken{}.GetAuthUser(token)
}
