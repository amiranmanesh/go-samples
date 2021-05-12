package requests

type NewProjectRequest struct {
	Name string `json:"name" binding:"required,min=3,max=20"`
}

type EditProjectRequest struct {
	Name string `json:"name" binding:"required,min=3,max=20"`
}

type ProjectApiSaveRequest struct {
	Name   string `json:"name" binding:"required,min=1,db_exists='project_apis:name'"`
	Path   string `json:"path" binding:"required,min=1"`
	Type   string `json:"type" binding:"required,oneof='static' 'dynamic'"`
	Result string `json:"result" binding:"required"`
}

type ProjectApiEditRequest struct {
	ProjectApiID uint   `json:"project_api_id" binding:"required"`
	Name         string `json:"name" `
	Path         string `json:"path" `
	Type         string `json:"type" `
	Result       string `json:"result" `
}
