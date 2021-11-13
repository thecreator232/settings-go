package server

import (
	"context"

	"github.com/gorilla/mux"

	"github.com/thecreator232/settings/source/storage"
)

type SettingServer struct {
	Storage storage.Storage
}

func SettingsServerRoutes(ctx context.Context) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/add", AddSettingsHandler)
	r.HandleFunc("/find", FindSettingHandler)
	r.HandleFunc("/update", UpdateSettingHanlder)
	r.HandleFunc("/delete", DeleteSettingHandler)
	return r

}

func (s *SettingServer) InitServer(ctx context.Context) error {

	return nil
}

// func (s *SettingServer) startServer(ctx context.Context, server *Server) error {

// 	return nil

// }
