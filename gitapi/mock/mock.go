package mock

import (
	"github.com/skhatri/git-reader/gitapi"
	"github.com/skhatri/git-reader/gitapi/model"
)

type gitMock struct {
}

func (git *gitMock) ListProjects() ([]model.Project, error) {
	return []model.Project{}, nil
}
func (git *gitMock) ListRepositories(project string) ([]model.Repository, error) {
	return []model.Repository{}, nil
}
func (git *gitMock) GetRepositoryDetail(project string, slug string) (*model.Repository, error) {
	return &model.Repository{}, nil
}
func (git *gitMock) ListTags(project string, slug string) ([]model.Tag, error) {
	return []model.Tag{}, nil
}
func (git *gitMock) ListBranches(project string, slug string) ([]model.Branch, error) {
	return []model.Branch{}, nil
}

func NewGitMock() gitapi.GitReader {
	return &gitMock{

	}
}
