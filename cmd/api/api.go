package api

import (
	"database/sql"

	"github.com/gorilla/mux"
)

type APIServer struct {
	db *sql.DB
	addr string

}

func NewAPIServer(addr string, db sql.DB){
	return &APIServer{
		addr: addr,
		db: &db,
	}
}

func (s*APIServer) Run() error {
router := mux.NewRouter()
subrouter:= router.PathPrefix("/api/v1").Subrouter()
}