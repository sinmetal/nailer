package backend

import (
	"context"
	"crypto/sha256"
	"fmt"
	"net/http"

	"github.com/favclip/ucon"
	"github.com/favclip/ucon/swagger"
)

func setupDBAPI(swPlugin *swagger.Plugin) {
	api := &DBAPI{}
	tag := swPlugin.AddTag(&swagger.Tag{Name: "DB", Description: "DB API list"})
	var hInfo *swagger.HandlerInfo

	hInfo = swagger.NewHandlerInfo(api.Post)
	ucon.Handle(http.MethodPost, "/api/1/db", hInfo)
	hInfo.Description, hInfo.Tags = "post to database schema", []string{tag.Name}
}

// DBAPI is Spanner Database に関するAPI
type DBAPI struct{}

// DBAPIPostRequest is DB Schema 登録 API Request
type DBAPIPostRequest struct {
	Schema []byte `json:"schema"`
}

// DBAPIPostResponse is DB Schema 登録 API Response
type DBAPIPostResponse struct {
	Checksum string `json:"checksum"`
}

// Post is DB Schemaを登録する
func (api *DBAPI) Post(ctx context.Context, form *DBAPIPostRequest) (*DBAPIPostResponse, error) {
	return &DBAPIPostResponse{
		checksum(form.Schema),
	}, nil
}

func checksum(bytes []byte) string {
	return fmt.Sprintf("%x", sha256.Sum256(bytes))
}
