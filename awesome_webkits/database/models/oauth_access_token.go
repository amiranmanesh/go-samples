package models

import (
	"awesome_webkits/database"
	"awesome_webkits/http/response"
	"awesome_webkits/utils/env"
	"awesome_webkits/utils/parser"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

const (
	ExpirationTimeDays = 90
)

// for generating token
type Claims struct {
	//Username string
	UserId uint `json:"user_id"`
	jwt.StandardClaims
}

type OauthAccessToken struct {
	ModelID

	User        User `gorm:"foreignkey:user_id;association_foreignkey:id"` // use UserRefer as foreign key
	UserID      uint
	AccessToken string `gorm:"type:varchar(255);unique_index;not null" json:"access_token"`
	ModelTimeStamps
}

func (o OauthAccessToken) GenerateToken(userID uint) (string, *response.Data) {
	claims := &Claims{
		UserId: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour * ExpirationTimeDays).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(env.GetEnvItem("SECRET_KEY")))

	if err != nil {
		return "", response.NewInternalServerError("Project make jwt error", err)
	}

	o.UserID = userID
	o.AccessToken = tokenString

	//UpdateOrCreate
	err = database.DB.GetDB().Scopes(o.scopeUser(o.UserID)).Assign(OauthAccessToken{AccessToken: o.AccessToken}).FirstOrCreate(&o).Error
	if err != nil {
		logrus.Error("access_token update user's token error", err)
		return "", response.NewInternalServerError("Project_token update user's token error", err)
	}

	return tokenString, nil

}

func (o OauthAccessToken) VerifyToken(token string) (uint, *response.Data) {
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil {

		return 0, response.NewUnauthorizedError("Error in parse token", err)
	}

	if !tkn.Valid {

		return 0, response.NewUnauthorizedError("Token is not valid", nil)
	}

	//check user token in db
	if result := database.DB.GetDB().Scopes(o.scopeToken(token)).Find(&o); result.Error != nil {

		return 0, response.NewUnauthorizedError("Token doesn't match with User", nil)
	}

	return claims.UserId, nil
}

func (o OauthAccessToken) GetAuthUserId(tokenHeader string) (uint, *response.Data) {
	//tokenHeader := c.GetHeader("Authorization")

	token, err := parser.ReadBearerToken(tokenHeader)
	if err != nil {

		return 0, response.NewUnauthorizedError("Token doesn't match with User", nil)
	}

	return o.VerifyToken(token)

}

func (o OauthAccessToken) GetAuthUser(token string) (*User, *response.Data) {
	userId, err := o.GetAuthUserId(token)
	if err != nil {
		return nil, err
	}

	var user User
	user.ID = userId
	err = user.FindWithId()
	if err != nil {
		return nil, err
	}

	return &user, nil

}

//scopes
func (o *OauthAccessToken) scopeUser(userId uint) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("user_id = ?", userId)
	}
}

func (o *OauthAccessToken) scopeToken(token string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("access_token = ?", token)
	}
}
