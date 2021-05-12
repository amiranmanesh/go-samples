package cronjob

import (
	"awesome_webkits/database/models"
	"awesome_webkits/otp"
	"github.com/jasonlvhit/gocron"
	"github.com/sirupsen/logrus"
	"time"
)

func initJobs() {
	if err := gocron.Every(15).Minute().Do(sendRegisterEmails); err != nil {
		logrus.Error("Error in running Cronjob", err)
		return
	}
}

func sendRegisterEmails() {
	logrus.WithFields(logrus.Fields{
		"time": time.Now,
	}).Debug("Run Sending Register Emails")

	emails, err := models.GetAllUnsentRegisteredUsers()
	if err != nil {
		logrus.Error("Error in getting emails in Cronjob", err)
		return
	}
	for _, email := range emails {
		tokenModel := otp.EmailTokenModel{
			Email: email.Email,
			Token: email.Token,
		}
		if err := tokenModel.SendActivationToken(); err != nil {
			if err := models.UpdateEmailStatus(tokenModel.Email, models.EmailTypeRegister, models.EmailStatusNotSent); err != nil {
				//todo handle error
				logrus.WithFields(logrus.Fields{
					"email":  tokenModel.Email,
					"status": models.EmailStatusNotSent,
				}).Error("Error in update email in Cronjob", err)
			}
		}
	}
}
