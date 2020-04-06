package gitapi

import "github.com/skhatri/git-reader/gitapi/model"

type GitReader interface {
	ListProjects() ([]model.Project, error)
	ListRepositories(project string) ([]model.Repository, error)
	GetRepositoryDetail(project string, slug string) (*model.Repository, error)
	ListTags(project string, slug string) ([]model.Tag, error)
	ListBranches(project string, slug string) ([]model.Branch, error)
}
