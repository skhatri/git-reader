package controller

import (
	"github.com/skhatri/api-router-go/router"
	"github.com/skhatri/api-router-go/router/model"
	"github.com/skhatri/git-reader/gitapi/bitbucket"
)

//GetProjects - list git projects
func GetProjects(_ *router.WebRequest) *model.Container {
	bit := bitbucket.NewBitBucketClient()
	list, err := bit.ListProjects()
	if err != nil {
		return model.ErrorResponse(model.MessageItem{
			Code:    "list-error",
			Message: err.Error(),
		}, 500)
	}
	data := make([]interface{}, 0)
	for _, item := range list {
		data = append(data, item)
	}
	return model.ListResponse(data)
}

//GetRepositories - list git repositories inside a project
func GetRepositories(web *router.WebRequest) *model.Container {
	projectName := web.GetQueryParam("project")
	if projectName == "" {
		return model.ErrorResponse(model.MessageItem{
			Code:    "invalid-request",
			Message: "project name is required",
		}, 500)
	}
	bit := bitbucket.NewBitBucketClient()
	list, err := bit.ListRepositories(projectName)
	if err != nil {
		return model.ErrorResponse(model.MessageItem{
			Code:    "list-repo-error",
			Message: err.Error(),
		}, 500)
	}
	data := make([]interface{}, 0)
	for _, item := range list {
		data = append(data, item)
	}
	return model.ListResponse(data)
}

//GetRepository - get repository detail by slug name inside a project
func GetRepository(web *router.WebRequest) *model.Container {
	projectName := web.GetQueryParam("project")
	name := web.GetQueryParam("name")
	if projectName == "" || name == "" {
		return model.ErrorResponse(model.MessageItem{
			Code:    "invalid-request",
			Message: "project name and repository name are required",
		}, 500)
	}
	bit := bitbucket.NewBitBucketClient()
	item, err := bit.GetRepositoryDetail(projectName, name)
	if err != nil {
		return model.ErrorResponse(model.MessageItem{
			Code:    "list-repo-error",
			Message: err.Error(),
		}, 500)
	}
	return model.Response(item)
}


//GetTags - list git tags inside a repository
func GetTags(web *router.WebRequest) *model.Container {
	projectName := web.GetQueryParam("project")
	name := web.GetQueryParam("name")
	if projectName == "" || name == "" {
		return model.ErrorResponse(model.MessageItem{
			Code:    "invalid-request",
			Message: "project name is required",
		}, 500)
	}
	bit := bitbucket.NewBitBucketClient()
	list, err := bit.ListTags(projectName, name)
	if err != nil {
		return model.ErrorResponse(model.MessageItem{
			Code:    "list-tag-error",
			Message: err.Error(),
		}, 500)
	}
	data := make([]interface{}, 0)
	for _, item := range list {
		data = append(data, item)
	}
	return model.ListResponse(data)
}


//GetBranches - list git branches inside a repository
func GetBranches(web *router.WebRequest) *model.Container {
	projectName := web.GetQueryParam("project")
	name := web.GetQueryParam("name")
	if projectName == "" || name == "" {
		return model.ErrorResponse(model.MessageItem{
			Code:    "invalid-request",
			Message: "project name is required",
		}, 500)
	}
	bit := bitbucket.NewBitBucketClient()
	list, err := bit.ListBranches(projectName, name)
	if err != nil {
		return model.ErrorResponse(model.MessageItem{
			Code:    "list-tag-error",
			Message: err.Error(),
		}, 500)
	}
	data := make([]interface{}, 0)
	for _, item := range list {
		data = append(data, item)
	}
	return model.ListResponse(data)
}