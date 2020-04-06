package controller

import (
	"fmt"
	"github.com/skhatri/api-router-go/router"
	"testing"
)

type MockApiConfigurer struct {
	m map[string]router.HandlerFunc
}

func (mc *MockApiConfigurer) Get(uri string, hf router.HandlerFunc) {

	mc.m[makeKey("get", uri)] = hf
}

func makeKey(method string, uri string) string {
	return fmt.Sprintf("%s%s", method, uri)
}

func (mc *MockApiConfigurer) Post(uri string, hf router.HandlerFunc) {
	mc.m[makeKey("post", uri)] = hf
}

func (mc *MockApiConfigurer) Method(method string, uri string, hf router.HandlerFunc) {
	mc.m[makeKey(method,uri)] = hf
}

func TestRegistersApis(t *testing.T) {

}
