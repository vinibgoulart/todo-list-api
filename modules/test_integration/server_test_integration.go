package test_integration

import (
	"net/http/httptest"

	"github.com/vinibgoulart/todo-list/server"
)

func RunTestServer() *httptest.Server {
	return httptest.NewServer(server.SetupServer())
}
