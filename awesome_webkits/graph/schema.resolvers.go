package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	graph_controller "awesome_webkits/graph/controller"
	"awesome_webkits/graph/generated"
	"awesome_webkits/graph/model"
	"context"
	"fmt"
)

func (r *queryResolver) Projects(ctx context.Context) ([]*model.Project, error) {
	return graph_controller.GetProjects(ctx)
}

func (r *queryResolver) ProjectApis(ctx context.Context) ([]*model.ProjectAPI, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
