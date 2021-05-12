package models

import (
	"awesome_webkits/database"
	"awesome_webkits/http/response"
	"awesome_webkits/otp"
	"awesome_webkits/utils/random"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
)

const (
	EmailTypeRegister  string = "register"
	EmailTypeReset     string = "reset"
	EmailStatusSent    string = "sent"
	EmailStatusNotSent string = "not sent"
	EmailStatusExpired string = "expired"
	EmailStatusUsed    string = "used"
)

type Email struct {
	gorm.Model
	Email    string `gorm:"type:varchar(100); unique_index; not null" json:"email"`
	Type     string `gorm:"type:varchar(20); not null" json:"type"`
	Token    string `gorm:"type:varchar(100); not null" json:"token"`
	Status   string `gorm:"type:varchar(20)" json:"status"`
	ExpireAt time.Time
}

func (e *Email) Register() *response.Data {
	e.Type = EmailTypeRegister
	e.Token = random.RandString36Bytes()
	e.Status = EmailStatusSent
	e.ExpireAt = time.Now().Add(48 * time.Hour)

	if err := database.DB.GetDB().Create(&e).Error; err != nil {
		return response.NewInternalServerError(fmt.Sprintf("Error in create Item:"), err.Error())
	}

	tokenModel := otp.EmailTokenModel{
		Email: e.Email,
		Token: e.Token,
	}

	//send activation token
	go func() {
		if err := tokenModel.SendActivationToken(); err != nil {
			if err := UpdateEmailStatus(tokenModel.Email, EmailTypeRegister, EmailStatusNotSent); err != nil {
				//todo handle error
				logrus.WithFields(logrus.Fields{
					"email":  tokenModel.Email,
					"status": EmailStatusNotSent,
				}).Error("Error in update status in email", err)
			}
		}
	}()

	return nil
}

func UpdateEmailStatus(email, emailType, statusNew string) error {
	model := &Email{
		Email: email,
		Type:  emailType,
	}
	if err := model.FindWithEmail(); err != nil {
		return err
	}

	if model.ExpireAt.Unix() > time.Now().Unix() {
		model.Status = statusNew
	} else {
		model.Status = EmailStatusExpired
	}

	if result := database.DB.GetDB().Save(&model); result.Error != nil {
		return result.Error
	}
	return nil
}

func (e *Email) FindWithEmail() error {
	result := database.DB.GetDB().Where("type = ? AND email = ?", e.Type, e.Email).First(&e)
	//result := database.GetDB().First(&e, "type = ?", e.Type, "email = ?", e.Email)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}
	return nil
}

func (e *Email) FindWithToken() *response.Data {
	result := database.DB.GetDB().Where("type = ? AND token = ? and status = ?", e.Type, e.Token, e.Status).First(&e)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return response.NewUnauthorizedError(fmt.Sprintf("Email not found : %s", e.Token), result.Error)
	}
	return nil
}

func GetAllUnsentRegisteredUsers() ([]Email, error) {
	var emails []Email
	if result := database.DB.GetDB().Where("type = ? AND status = ?", EmailTypeRegister, EmailStatusNotSent).Find(&emails); result.Error != nil {
		return nil, result.Error
	}
	return emails, nil
}
