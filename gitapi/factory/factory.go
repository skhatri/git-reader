package factory

import (
	"github.com/skhatri/git-reader/gitapi"
	"github.com/skhatri/git-reader/gitapi/bitbucket"
	"github.com/skhatri/git-reader/gitapi/mock"
	"os"
	"sync"
)

//NewGitReader creates a git reader for given provider type
func NewGitReader(typeOverride *string) gitapi.GitReader {
	var value = os.Getenv("GIT_PROVIDER")
	if typeOverride != nil {
		value = *typeOverride
	}
	var reader gitapi.GitReader
	switch value {
	case "mock":
		reader = mock.NewGitMock()
	default:
		reader = bitbucket.NewBitBucketClient()
	}
	return reader
}

var defaultGitClient gitapi.GitReader
var clientMux = sync.Mutex{}

func GetGitClient() gitapi.GitReader {
	if defaultGitClient != nil {
		return defaultGitClient

	}
	clientMux.Lock()
	if defaultGitClient == nil {
		defaultGitClient = NewGitReader(nil)
	}
	clientMux.Unlock()
	return defaultGitClient
}
