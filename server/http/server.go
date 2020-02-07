package main

import (
	"fmt"
	"net/http"

	"github.com/JeanLouiseFinch/calendar/config"
	"go.uber.org/zap"
)

type MyHandler struct {
	Logger *zap.Logger
	Config *config.Config
}

func (h *MyHandler) Add(resp http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}
	return
}
func (h *MyHandler) Delete(resp http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}
	return
}
func (h *MyHandler) Update(resp http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}
	return
}
func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {}
func main() {
	handler := &MyHandler{}
	cfg, err := config.GetConfig("../../config/config.yaml")
	if err != nil {
		panic(err)
	}
	handler.Config = cfg
	mux := http.NewServeMux()
	mux.Handle("/", handler)
	mux.HandleFunc("add", handler.Add)
	mux.HandleFunc("delete", handler.Delete)
	mux.HandleFunc("update", handler.Update)
	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", handler.Config.IP, handler.Config.Port),
		Handler: mux,
	}
	handler.Logger.Fatal("ups", zap.Error(server.ListenAndServe()))
}
