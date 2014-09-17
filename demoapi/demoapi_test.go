package demoapi

import (
	"fmt"
	"github.com/agoravoting/agora-http-go/middleware"
	stest "github.com/agoravoting/agora-http-go/server/testing"
	"net/http"
	"testing"
)

const (
	newEvent = `{
	"name": "foo election",
	"auth_method": "sms-code",
	"auth_method_config": {
		"probando": "lo que sea"
	}
}`
	secret = "somesecret"
)

var (
	SharedSecret = "somesecret"
	Config       = `{
	"Debug": true,
	"DbMaxIddleConnections": 5,
	"DbConnectString": "user=agora_http_go password=agora_http_go dbname=agora_http_go sslmode=disable",

	"SharedSecret": "somesecret",
	"Admins": ["test@example.com"],
	"ActiveModules": [
		"github.com/agoravoting/agora-http-go/demoapi"
	],
	"RavenDSN": ""
}`
)

func TestEventApi(t *testing.T) {
	ts := stest.New(t, Config)
	defer ts.TearDown()
	auth_admin := map[string]string{"Authorization": middleware.AuthHeader("superuser", SharedSecret)}

	// do a post and get it back
	newEvent := ts.RequestJson("POST", "/api/v1/event/", http.StatusAccepted, auth_admin, newEvent).(map[string]interface{})
	apiPath := fmt.Sprintf("/api/v1/event/%.0f", newEvent["id"])
	ts.Request("GET", apiPath, http.StatusOK, auth_admin, "")
	// deleting the event
	ts.Request("DELETE", apiPath, http.StatusOK, auth_admin, "")
	ts.Request("GET", apiPath, http.StatusNotFound, auth_admin, "")
	// 	fmt.Printf("req-out = %s\n", ret)

}
