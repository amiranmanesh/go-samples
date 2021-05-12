package models

import (
	"awesome_webkits/database"
	"awesome_webkits/http/response"
	"awesome_webkits/utils/env"
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
	"os"
)

// for generating token
type ProjectClaims struct {
	//Username string
	ProjectId uint `json:"project_id"`
	jwt.StandardClaims
}

type OauthProjectToken struct {
	ModelID
	Project      Project `gorm:"foreignkey:project_id;association_foreignkey:id"` // use UserRefer as foreign key
	ProjectID    uint
	ProjectToken string `gorm:"type:varchar(255);unique_index;not null" json:"project_token"`
	ModelTimeStamps
}

func (o OauthProjectToken) GenerateToken(projectID uint) (string, *response.Data) {
	claims := &ProjectClaims{
		ProjectId:      projectID,
		StandardClaims: jwt.StandardClaims{},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(env.GetEnvItem("SECRET_KEY")))

	if err != nil {
		return "", response.NewInternalServerError("Project make jwt error", err)
	}

	o.ProjectID = projectID
	o.ProjectToken = tokenString

	//UpdateOrCreate
	err = database.DB.GetDB().Scopes(o.scopeProject(o.ProjectID)).Assign(OauthAccessToken{AccessToken: o.ProjectToken}).FirstOrCreate(&o).Error
	if err != nil {
		return "", response.NewInternalServerError("Project_token update user's token error", err)
	}

	return tokenString, nil

}

func (o OauthProjectToken) VerifyToken(token string) (uint, *response.Data) {
	claims := &ProjectClaims{}

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
		return 0, response.NewUnauthorizedError("Token doesn't match with Project", nil)
	}

	return claims.ProjectId, nil
}

func (o OauthProjectToken) GetAuthProjectId(token string) (uint, *response.Data) {
	//token := c.GetHeader("Project-Token")
	projectId, err2 := o.VerifyToken(token)
	if err2 != nil {
		return 0, response.NewUnauthorizedError("Token doesn't match with Project", nil)
	}
	return projectId, nil

}

func (o OauthProjectToken) GetAuthProject(token string) (*Project, *response.Data) {
	projectId, err := o.GetAuthProjectId(token)
	if err != nil {
		return nil, err
	}

	var project Project
	project.ID = projectId
	err = project.FindWithId()
	if err != nil {
		return nil, err
	}

	return &project, nil

}

//scopes
func (o *OauthProjectToken) scopeProject(projectID uint) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("project_id = ?", projectID)
	}
}

func (o *OauthProjectToken) scopeToken(token string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("project_token = ?", token)
	}
}
