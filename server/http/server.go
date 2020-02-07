package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/JeanLouiseFinch/calendar/config"
	"github.com/JeanLouiseFinch/calendar/nosql"
	"go.uber.org/zap"
)

type MyHandler struct {
	Logger  *zap.Logger
	Config  *config.Config
	Storage *nosql.Storage
}

func RunServer(cfg *config.Config, log *zap.Logger) error {
	handler := &MyHandler{
		Logger:  log,
		Config:  cfg,
		Storage: nosql.NewStorage(),
	}

	mux := http.NewServeMux()
	mux.Handle("/", handler)
	mux.HandleFunc("/add", handler.postOnly(handler.Add))
	mux.HandleFunc("/delete", handler.postOnly(handler.Delete))
	mux.HandleFunc("/update", handler.postOnly(handler.Update))
	mux.HandleFunc("/events_for_day", handler.getOnly(handler.EventsForDay))
	mux.HandleFunc("/events_for_week", handler.getOnly(handler.EventsForWeek))
	mux.HandleFunc("/events_for_month", handler.getOnly(handler.EventsForMonth))
	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", handler.Config.IP, handler.Config.Port),
		Handler: mux,
	}
	handler.Logger.Info("Runnig server...", zap.String("host", handler.Config.IP), zap.Int("port", handler.Config.Port))
	handler.Logger.Fatal("ups", zap.Error(server.ListenAndServe()))
	return nil
}
func (h *MyHandler) postOnly(hh http.HandlerFunc) http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPost {
			h.Logger.Info("Get zapros on only post page")
			resp.WriteHeader(http.StatusInternalServerError)
			return
		}
		hh(resp, req)
	}
}
func (h *MyHandler) getOnly(hh http.HandlerFunc) http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodGet {
			h.Logger.Info("Post zapros on only get page")
			resp.WriteHeader(http.StatusInternalServerError)
			return
		}
		hh(resp, req)
	}
}
func (h *MyHandler) Add(resp http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var data EventHTTP
	err := decoder.Decode(&data)
	if err != nil {
		h.Logger.Error("parsing input data error", zap.Error(err))
	} else {
		h.Logger.Info("Data:", zap.String("owner", data.Owner), zap.String("title", data.Title), zap.String("description", data.Description), zap.String("start", data.Start), zap.String("end", data.End))
	}
	return
}
func (h *MyHandler) Delete(resp http.ResponseWriter, req *http.Request) {
	return
}
func (h *MyHandler) Update(resp http.ResponseWriter, req *http.Request) {
	return
}
func (h *MyHandler) EventsForDay(resp http.ResponseWriter, req *http.Request) {
	return
}
func (h *MyHandler) EventsForWeek(resp http.ResponseWriter, req *http.Request) {
	return
}
func (h *MyHandler) EventsForMonth(resp http.ResponseWriter, req *http.Request) {
	return
}
func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {}
