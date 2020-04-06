package controller

import "github.com/skhatri/api-router-go/router"

func Configure(configurer router.ApiConfigurer) {

	configurer.Get("/api/projects", GetProjects)
	configurer.Get("/api/repositories", GetRepositories)
	configurer.Get("/api/repository", GetRepository)
	configurer.Get("/api/tags", GetTags)
	configurer.Get("/api/branches", GetBranches)
}
