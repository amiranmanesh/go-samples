package models

import (
	"awesome_webkits/database"
	"awesome_webkits/http/response"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	Name   string `gorm:"type:varchar(100); not null" json:"name"`
	User   User   `gorm:"foreignkey:user_id;association_foreignkey:id"` // use UserRefer as foreign key
	UserID uint
}

func (p *Project) Save() (string /* token */, *response.Data) {

	if err := database.DB.GetDB().Create(&p).Error; err != nil {
		return "", response.NewInternalServerError("Error in save project in db", err.Error())
	}

	//@todo change project token
	token, err := OauthProjectToken{}.GenerateToken(p.ID)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (p *Project) FindWithId() *response.Data {

	result := database.DB.GetDB().First(&p, p.ID)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return response.NewUnauthorizedError(fmt.Sprintf("Project not found with id: %d", p.ID), result.Error)
	}
	return nil

}

func (p *Project) UpdateName(name string) *response.Data {

	if result := database.DB.GetDB().Model(&p).Update("name", name); result.Error != nil {
		return response.NewUnauthorizedError(fmt.Sprintf("Error in update name project_id: %d", p.ID), result.Error)
	}
	return nil

}

func (p *Project) GetAuthProjectId(token string) (uint, *response.Data) {
	return OauthProjectToken{}.GetAuthProjectId(token)
}

func (p *Project) GetAuthProject(token string) (*Project, *response.Data) {
	return OauthProjectToken{}.GetAuthProject(token)
}

func (p Project) GetAllProjects(user *User) ([]*Project, error) {
	var projects []*Project
	if result := database.DB.GetDB().Where("user_id = ?", user.ID).Find(&projects); result.Error != nil {
		return nil, result.Error
	}
	return projects, nil
}
