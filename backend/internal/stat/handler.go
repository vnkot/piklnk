package stat

import (
	"net/http"
	"strconv"
	"time"

	"github.com/vnkot/piklnk/configs"
	"github.com/vnkot/piklnk/internal/link"
	"github.com/vnkot/piklnk/pkg/apierr"
	"github.com/vnkot/piklnk/pkg/jwt"
	"github.com/vnkot/piklnk/pkg/middleware"
	"github.com/vnkot/piklnk/pkg/req"
	"github.com/vnkot/piklnk/pkg/res"
)

type StatHandler struct {
	StatRepository *StatRepository
	Config         *configs.Config
	LinkRepository *link.LinkRepository
}

type StatHandlerDeps struct {
	StatRepository *StatRepository
	Config         *configs.Config
	LinkRepository *link.LinkRepository
}

func NewStatHandler(router *http.ServeMux, deps StatHandlerDeps) {
	handler := &StatHandler{
		StatRepository: deps.StatRepository,
		LinkRepository: deps.LinkRepository,
	}

	router.Handle("GET /stat/group/{id}", middleware.IsAuthed(handler.GetGroupStat(), deps.Config.Auth))
}

func (handler *StatHandler) GetGroupStat() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		jwtData, ok := r.Context().Value(middleware.ContextJWTDataKey).(*jwt.JWTData)

		if !ok || jwtData == nil {
			res.Json(w, apierr.New(http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized, ""), http.StatusUnauthorized)
			return
		}

		linkID, err := strconv.ParseUint(r.PathValue("id"), 10, 32)

		if err != nil {
			res.Json(w, apierr.New(http.StatusText(http.StatusBadRequest), http.StatusBadRequest, err.Error()), http.StatusBadRequest)
			return
		}

		link, err := handler.LinkRepository.GetById(uint(linkID))

		if err != nil {
			res.Json(w, apierr.New(http.StatusText(http.StatusBadRequest), http.StatusBadRequest, err.Error()), http.StatusBadRequest)
			return
		}

		if jwtData.UserID != *link.UserId {
			res.Json(w, apierr.New(http.StatusText(http.StatusForbidden), http.StatusForbidden, ""), http.StatusForbidden)
			return
		}

		queryParams, err := req.HandleQueryParams[GetGroupStatParamsRequest](&w, r)
		if err != nil {
			res.Json(w, apierr.New(http.StatusText(http.StatusBadRequest), http.StatusBadRequest, err.Error()), http.StatusBadRequest)
			return
		}

		dateFrom, errFrom := time.Parse("2006-01-02", queryParams.From)
		dateTo, errTo := time.Parse("2006-01-02", queryParams.To)
		if errFrom != nil || errTo != nil {
			res.Json(w, apierr.New(http.StatusText(http.StatusBadRequest), http.StatusBadRequest, "invalid date format, use YYYY-MM-DD"), http.StatusBadRequest)
			return
		}

		validPeriods := map[string]bool{"day": true, "month": true}
		if !validPeriods[queryParams.By] {
			res.Json(w, apierr.New(http.StatusText(http.StatusBadRequest), http.StatusBadRequest, "invalid period: use 'day' or 'month'"), http.StatusBadRequest)
			return
		}

		stats, err := handler.StatRepository.GetGroupStat(uint(linkID), queryParams.By, dateFrom, dateTo)

		if err != nil {
			res.Json(w, apierr.New(http.StatusText(http.StatusBadRequest), http.StatusBadRequest, err.Error()), http.StatusBadRequest)
			return
		}

		res.Json(w, stats, http.StatusOK)
	}
}
