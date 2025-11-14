package stat

import (
	"api-project/configs"
	"api-project/pkg/middleware"
	"api-project/pkg/request"
	"api-project/pkg/response"
	"net/http"
	"time"
)

const (
	GroupByDay   = "day"
	GroupByMonth = "month"
)

type StatHandlerDeps struct {
	Config     *configs.AuthConfig
	Repository *StatRepository
}
type StatHandler struct {
	Repository *StatRepository
}

func NewHandler(router *http.ServeMux, deps StatHandlerDeps) {
	handler := &StatHandler{
		Repository: deps.Repository,
	}
	router.Handle("GET /stat", middleware.IsAuthed(handler.getStat(), deps.Config))
}
func (handler *StatHandler) getStat() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fromStr, err := request.PrepareParam[string](&w, r, "query", "from", true)
		if err != nil {
			return
		}
		from, err := time.Parse("2006-01-02", fromStr)
		if err != nil {
			response.BadRequestJson(w, "Invalid from param")
			return
		}
		toStr, err := request.PrepareParam[string](&w, r, "query", "to", true)
		if err != nil {
			return
		}
		to, err := time.Parse("2006-01-02", toStr)
		if err != nil {
			response.BadRequestJson(w, "Invalid to param")
			return
		}
		by, err := request.PrepareParam[string](&w, r, "query", "by", true)
		if err != nil {
			return
		}
		if !(by == GroupByDay || by == GroupByMonth) {
			response.BadRequestJson(w, "Invalid by param")
			return
		}
		stats := handler.Repository.GetStats(by, from, to)
		response.Json(w, stats, http.StatusOK)
	}
}
