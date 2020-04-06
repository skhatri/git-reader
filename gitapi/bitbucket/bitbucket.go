package bitbucket

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/skhatri/git-reader/gitapi"
	"github.com/skhatri/git-reader/gitapi/model"
	"github.com/skhatri/git-reader/httpclient"
	"log"
	"os"
	"sync"
)

type gitApi struct {
	token  string
	client httpclient.HttpClient
}

var instance *gitApi
var instanceLock = sync.Mutex{}

func NewBitBucketClient() gitapi.GitReader {
	if instance != nil {
		return instance
	}
	initialize()
	return instance
}

func initialize() {
	instanceLock.Lock()
	var token = ""
	if t := os.Getenv("BITBUCKET_TOKEN"); t == "" {
		token = os.Getenv("GIT_READER")
	} else {
		token = t
	}
	if instance == nil {
		instance = &gitApi{
			token:  token,
			client: httpclient.NewHttpClient(2000),
		}
	}
	instanceLock.Unlock()
}

func callAndTransformResult(git *gitApi, url string, transformer func(map[string]interface{}) []interface{}) (items []interface{}, generalError error) {
	var headers = map[string]string{
		"Accept":        "application/json",
		"Authorization": fmt.Sprintf("Bearer %s", git.token),
	}
	response, err := git.client.DoGet(headers, url, httpclient.DefaultHttpOptions)

	if err != nil {
		return nil, err
	}

	if response.Status >= 400 {
		return nil, errors.New(fmt.Sprintf("downstream server returned error status=%d, body=%s",
			response.Status, string(response.Data)))
	}

	defer func() {
		if r := recover(); r != nil {
			log.Println("Error in transformation", r)
			items = nil
			generalError = errors.New(fmt.Sprintf("error in transformation: %v", r))
		}
	}()
	var genericMap = map[string]interface{}{}
	buff := bytes.NewBuffer(response.Data)
	json.NewDecoder(buff).Decode(&genericMap)
	items = transformer(genericMap)
	return items, generalError
}

func (git *gitApi) ListProjects() ([]model.Project, error) {
	url := "http://localhost:7990/rest/api/1.0/projects"
	items, generalError := callAndTransformResult(git, url, transformProjectResult)
	if generalError != nil {
		return nil, generalError
	}
	projects := make([]model.Project, 0)
	for _, item := range items {
		projects = append(projects, item.(model.Project))
	}
	return projects, generalError
}

func transformProjectResult(genericMap map[string]interface{}) []interface{} {
	projects := make([]interface{}, 0)
	values := genericMap["values"].([]interface{})
	for _, item := range values {
		v := item.(map[string]interface{})
		projects = append(projects, model.Project{
			Key:         v["key"].(string),
			Name:        v["name"].(string),
			Description: v["description"].(string),
			Id:          fmt.Sprintf("%.f", v["id"].(float64)),
		})
	}
	return projects
}

func (git *gitApi) ListRepositories(projectName string) ([]model.Repository, error) {
	url := fmt.Sprintf("http://localhost:7990/rest/api/1.0/projects/%s/repos", projectName)

	items, generalError := callAndTransformResult(git, url, transformRepositoryResult)
	if generalError != nil {
		return nil, generalError
	}
	repositories := make([]model.Repository, 0)
	for _, item := range items {
		repositories = append(repositories, item.(model.Repository))
	}
	return repositories, generalError

}

func transformRepositoryResult(genericMap map[string]interface{}) []interface{} {
	repositories := make([]interface{}, 0)
	values := genericMap["values"].([]interface{})
	for _, item := range values {
		v := item.(map[string]interface{})
		repositories = append(repositories, model.Repository{
			Slug:        v["slug"].(string),
			Name:        v["name"].(string),
			Description: v["description"].(string),
			Id:          fmt.Sprintf("%.f", v["id"].(float64)),
		})
	}
	return repositories
}

func (git *gitApi) GetRepositoryDetail(projectName string, slug string) (*model.Repository, error) {
	url := fmt.Sprintf("http://localhost:7990/rest/api/1.0/projects/%s/repos/%s", projectName, slug)
	items, generalError := callAndTransformResult(git, url, transformRepositoryDetailResult)
	if generalError != nil {
		return nil, generalError
	}
	var repository model.Repository
	if len(items) == 1 {
		repository = items[0].(model.Repository)
	}
	return &repository, generalError
}

func transformRepositoryDetailResult(v map[string]interface{}) []interface{} {
	repositories := make([]interface{}, 0)
	repositories = append(repositories, model.Repository{
		Slug:        v["slug"].(string),
		Name:        v["name"].(string),
		Description: v["description"].(string),
		Id:          fmt.Sprintf("%.f", v["id"].(float64)),
	})
	return repositories
}

func (git *gitApi) ListTags(projectName string, slug string) ([]model.Tag, error) {
	url := fmt.Sprintf("http://localhost:7990/rest/api/1.0/projects/%s/repos/%s/tags", projectName, slug)
	items, generalError := callAndTransformResult(git, url, transformTagResult)
	if generalError != nil {
		return nil, generalError
	}
	tags := make([]model.Tag, 0)
	for _, item := range items {
		tags = append(tags, item.(model.Tag))
	}
	return tags, generalError
}

func transformTagResult(genericMap map[string]interface{}) []interface{} {
	repositories := make([]interface{}, 0)
	values := genericMap["values"].([]interface{})
	for _, item := range values {
		v := item.(map[string]interface{})
		var hash = ""
		if v["hash"] != nil {
			hash = v["hash"].(string)
		}
		repositories = append(repositories, model.Tag{
			Type:         v["type"].(string),
			LatestCommit: v["latestCommit"].(string),
			Display:      v["displayId"].(string),
			Hash:         hash,
			Id:           v["id"].(string),
		})
	}
	return repositories
}

func (git *gitApi) ListBranches(projectName string, slug string) ([]model.Branch, error) {
	url := fmt.Sprintf("http://localhost:7990/rest/api/1.0/projects/%s/repos/%s/branches", projectName, slug)
	items, generalError := callAndTransformResult(git, url, transformBranchResult)
	if generalError != nil {
		return nil, generalError
	}
	tags := make([]model.Branch, 0)
	for _, item := range items {
		tags = append(tags, item.(model.Branch))
	}
	return tags, generalError
}

func transformBranchResult(genericMap map[string]interface{}) []interface{} {
	repositories := make([]interface{}, 0)
	values := genericMap["values"].([]interface{})
	for _, item := range values {
		v := item.(map[string]interface{})
		repositories = append(repositories, model.Branch{
			Type:         v["type"].(string),
			Display:      v["displayId"].(string),
			IsDefault:    v["isDefault"].(bool),
			LatestCommit: v["latestCommit"].(string),
			Id:           v["id"].(string),
		})
	}
	return repositories
}
