package models

import (
	"awesome_webkits/database"
	"awesome_webkits/http/response"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type ProjectApi struct {
	gorm.Model
	Name      string `gorm:"type:varchar(100); not null" json:"name"`
	ProjectID uint   `gorm:"index:idx_show"`
	Project   Project
	Path      string `gorm:"type:varchar(100); not null; index:idx_show" json:"path"`
	Type      string `gorm:"type:varchar(100); not null" json:"type"`
	Result    string `gorm:"type:text" json:"result"`
	Number    uint   `gorm:"type:int" json:"type"`
	Json      string `gorm:"type:text" json:"result_json"`
}

func (p *ProjectApi) Save() *response.Data {

	// Check that user with this email exists
	var projectApi ProjectApi
	result := database.DB.GetDB().Scopes(projectApi.ScopePathAndProjectID(p.Path, p.ProjectID)).First(&projectApi)
	if err := result.Error; !errors.Is(err, gorm.ErrRecordNotFound) {
		return response.NewBadRequestErr("This path is already used for your project", err)
	}

	if err := database.DB.GetDB().Create(&p).Error; err != nil {
		return response.NewInternalServerError(fmt.Sprintf("Error in create project api- %s", err.Error()), err)
	}

	return nil

}

func (p *ProjectApi) Update(name, path, types, result string) *response.Data {
	if err := p.FindWithId(); err != nil {
		return err
	}
	if name != "" {
		p.Name = name
	}
	if path != "" {
		p.Path = path
	}
	if types != "" {
		p.Type = types
	}
	if result != "" {
		p.Result = result
	}

	if result := database.DB.GetDB().Save(&p); result.Error != nil {
		return response.NewUnauthorizedError(fmt.Sprintf("Error in update project_api info: %d", p.ID), result.Error)
	}
	return nil

}

func (p *ProjectApi) FindWithId() *response.Data {

	if result := database.DB.GetDB().First(&p, p.ID); result.Error != nil {
		return response.NewUnauthorizedError(fmt.Sprintf("ProjectApi not found with id: %d", p.ID), result.Error)
	}
	return nil

}
func (p *ProjectApi) FindWithPath() *response.Data {
	if result := database.DB.GetDB().Where("path = ?", p.Path).First(&p); result.Error != nil {
		return response.NewUnauthorizedError(fmt.Sprintf("ProjectApi not found with path: %s", p.Path), result.Error)
	}
	return nil
}

//scopes
func (p *ProjectApi) ScopePathAndProjectID(path string, projectID uint) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("path = ? and Project_id = ?", path, projectID)
	}
}
