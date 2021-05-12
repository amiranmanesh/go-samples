package graph_controller

import (
	"awesome_webkits/database/models"
	"awesome_webkits/graph/middleware"
	"awesome_webkits/graph/model"
	"context"
	"fmt"
)

func GetProjects(ctx context.Context) ([]*model.Project, error) {
	user := middleware.GraphQLForContext(ctx)
	if user == nil {
		return []*model.Project{}, fmt.Errorf("access denied")
	}
	var projectModelsTemp []*model.Project

	projectModels, err := models.Project{}.GetAllProjects(user)
	if err != nil {
		return []*model.Project{}, fmt.Errorf("access denied")
	}

	for _, projectItem := range projectModels {
		projectModelsTemp = append(projectModelsTemp, &model.Project{
			ID:   fmt.Sprint(projectItem.ID),
			Name: projectItem.Name,
		})
	}

	return projectModelsTemp, nil
}
