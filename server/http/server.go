package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/JeanLouiseFinch/calendar/config"
	"github.com/JeanLouiseFinch/calendar/internal/adapters/nosql"
	"github.com/JeanLouiseFinch/calendar/internal/domain/entities"
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
	req.ParseForm()
	resp.Header().Set("Content-Type", "application/json")
	newEvent := EventHTTP{
		Description: req.Form.Get("description"),
		EndString:   req.Form.Get("end"),
		Owner:       req.Form.Get("owner"),
		StartString: req.Form.Get("start"),
		Title:       req.Form.Get("title"),
	}
	start, err := toTimeFromString(newEvent.StartString)
	if err != nil {
		h.Logger.Error("Parse time", zap.Error(err))
		resp.WriteHeader(http.StatusBadRequest)
		return
	}
	newEvent.Start = start
	end, err := toTimeFromString(newEvent.EndString)
	if err != nil {
		h.Logger.Error("Parse time", zap.Error(err))
		resp.WriteHeader(http.StatusBadRequest)
		return
	}
	newEvent.End = end
	val, err := h.Storage.AddEvent(context.Background(), &entities.Event{
		Owner:       newEvent.Owner,
		Title:       newEvent.Title,
		Description: newEvent.Description,
		Start:       start,
		End:         end,
	})
	if err != nil {
		he := HTTPError{
			Error: err,
		}
		err = json.NewEncoder(resp).Encode(he)
		if err != nil {
			resp.WriteHeader(http.StatusInternalServerError)
			return
		}
		resp.WriteHeader(http.StatusOK)
		return
	}
	newEvent.ID = val
	err = json.NewEncoder(resp).Encode(newEvent)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}
	h.Logger.Info("Data:", zap.Int("ID", int(newEvent.ID)), zap.String("owner", newEvent.Owner), zap.String("title", newEvent.Title), zap.String("description", newEvent.Description), zap.Any("start", start), zap.Any("end", end))
	resp.WriteHeader(http.StatusOK)
	return
}
func (h *MyHandler) Delete(resp http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	resp.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(req.Form.Get("id"))
	if err != nil {
		h.Logger.Error("Parse id", zap.Error(err))
		resp.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.Storage.DeleteEvent(context.Background(), uint(id))
	if err != nil {
		he := HTTPError{
			Error: err,
		}
		err = json.NewEncoder(resp).Encode(he)
		if err != nil {
			resp.WriteHeader(http.StatusInternalServerError)
			return
		}
		resp.WriteHeader(http.StatusOK)
		return
	}
	newEvent := EventHTTP{
		ID: uint(id),
	}
	err = json.NewEncoder(resp).Encode(newEvent)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}
	h.Logger.Info("Delete:", zap.Int("ID", int(newEvent.ID)))
	resp.WriteHeader(http.StatusOK)
	return
}
func (h *MyHandler) Update(resp http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	resp.Header().Set("Content-Type", "application/json")
	newEvent := EventHTTP{
		Description: req.Form.Get("description"),
		EndString:   req.Form.Get("end"),
		Owner:       req.Form.Get("owner"),
		StartString: req.Form.Get("start"),
		Title:       req.Form.Get("title"),
	}
	id, err := strconv.Atoi(req.Form.Get("id"))
	if err != nil {
		h.Logger.Error("Parse id", zap.Error(err))
		resp.WriteHeader(http.StatusBadRequest)
		return
	}
	newEvent.ID = uint(id)
	start, err := toTimeFromString(newEvent.StartString)
	if err != nil {
		h.Logger.Error("Parse time", zap.Error(err))
		resp.WriteHeader(http.StatusBadRequest)
		return
	}
	newEvent.Start = start
	end, err := toTimeFromString(newEvent.EndString)
	if err != nil {
		h.Logger.Error("Parse time", zap.Error(err))
		resp.WriteHeader(http.StatusBadRequest)
		return
	}
	newEvent.End = end
	err = h.Storage.EditEvent(context.Background(), newEvent.ID, &entities.Event{
		Owner:       newEvent.Owner,
		Title:       newEvent.Title,
		Description: newEvent.Description,
		Start:       start,
		End:         end,
	})
	if err != nil {
		he := HTTPError{
			Error: err,
		}
		err = json.NewEncoder(resp).Encode(he)
		if err != nil {
			resp.WriteHeader(http.StatusInternalServerError)
			return
		}
		resp.WriteHeader(http.StatusOK)
		return
	}

	err = json.NewEncoder(resp).Encode(newEvent)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}
	h.Logger.Info("Update data:", zap.Int("ID", int(newEvent.ID)), zap.String("owner", newEvent.Owner), zap.String("title", newEvent.Title), zap.String("description", newEvent.Description), zap.Any("start", start), zap.Any("end", end))
	resp.WriteHeader(http.StatusOK)
	return
}
func (h *MyHandler) EventsForDay(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")

	events, err := h.Storage.GetEventsByTimeRange(context.Background(), entities.TimeRangeDay)
	if err != nil {
		he := HTTPError{
			Error: err,
		}
		err = json.NewEncoder(resp).Encode(he)
		if err != nil {
			resp.WriteHeader(http.StatusInternalServerError)
			return
		}
		resp.WriteHeader(http.StatusOK)
		return
	}
	err = json.NewEncoder(resp).Encode(events)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}
	for key := range events {
		h.Logger.Info("Data for day:", zap.String("owner", events[key].Owner), zap.String("title", events[key].Title), zap.String("description", events[key].Description), zap.Any("start", events[key].Start), zap.Any("end", events[key].End))

	}
	resp.WriteHeader(http.StatusOK)
	return
}
func (h *MyHandler) EventsForWeek(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")

	events, err := h.Storage.GetEventsByTimeRange(context.Background(), entities.TimeRangeWeek)
	if err != nil {
		he := HTTPError{
			Error: err,
		}
		err = json.NewEncoder(resp).Encode(he)
		if err != nil {
			resp.WriteHeader(http.StatusInternalServerError)
			return
		}
		resp.WriteHeader(http.StatusOK)
		return
	}
	err = json.NewEncoder(resp).Encode(events)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}
	for key := range events {
		h.Logger.Info("Data for week:", zap.String("owner", events[key].Owner), zap.String("title", events[key].Title), zap.String("description", events[key].Description), zap.Any("start", events[key].Start), zap.Any("end", events[key].End))

	}
	resp.WriteHeader(http.StatusOK)
	return
}
func (h *MyHandler) EventsForMonth(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")

	events, err := h.Storage.GetEventsByTimeRange(context.Background(), entities.TimeRangeMounth)
	if err != nil {
		he := HTTPError{
			Error: err,
		}
		err = json.NewEncoder(resp).Encode(he)
		if err != nil {
			resp.WriteHeader(http.StatusInternalServerError)
			return
		}
		resp.WriteHeader(http.StatusOK)
		return
	}
	err = json.NewEncoder(resp).Encode(events)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}
	for key := range events {
		h.Logger.Info("Data for mounth:", zap.String("owner", events[key].Owner), zap.String("title", events[key].Title), zap.String("description", events[key].Description), zap.Any("start", events[key].Start), zap.Any("end", events[key].End))

	}
	resp.WriteHeader(http.StatusOK)
	return
}
func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {}
