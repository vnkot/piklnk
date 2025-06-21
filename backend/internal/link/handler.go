package link

import (
	"net/http"
	"strconv"

	"github.com/vnkot/piklnk/configs"
	"github.com/vnkot/piklnk/internal/auth/repository"
	"github.com/vnkot/piklnk/pkg/apierr"
	"github.com/vnkot/piklnk/pkg/event"
	"github.com/vnkot/piklnk/pkg/jwt"
	"github.com/vnkot/piklnk/pkg/middleware"
	"github.com/vnkot/piklnk/pkg/req"
	"github.com/vnkot/piklnk/pkg/res"
	"gorm.io/gorm"
)

type LinkHandlerDeps struct {
	EventBus       *event.EventBus
	Config         *configs.Config
	LinkService    *LinkService
	LinkRepository *LinkRepository
	UserRepository *repository.UserRepository
}

type LinkHandler struct {
	EventBus       *event.EventBus
	LinkService    *LinkService
	LinkRepository *LinkRepository
	UserRepository *repository.UserRepository
}

func NewLinkHandler(router *http.ServeMux, deps LinkHandlerDeps) {
	handler := &LinkHandler{
		EventBus:       deps.EventBus,
		LinkService:    deps.LinkService,
		LinkRepository: deps.LinkRepository,
		UserRepository: deps.UserRepository,
	}

	router.HandleFunc("GET /{hash}", handler.GoTo())
	router.Handle("POST /link/create", middleware.IsMaybeAuthed(handler.Create(), deps.Config.Auth))
	router.Handle("GET /link/list", middleware.IsAuthed(handler.GetAll(), deps.Config.Auth))
	router.Handle("PATCH /link/update/{id}", middleware.IsAuthed(handler.Update(), deps.Config.Auth))
	router.Handle("DELETE /link/delete/{id}", middleware.IsAuthed(handler.Delete(), deps.Config.Auth))
}

func (handler *LinkHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var userID *uint

		jwtData, ok := r.Context().Value(middleware.ContextJWTDataKey).(*jwt.JWTData)
		if ok && jwtData != nil {
			userID = &jwtData.UserID
		}

		body, err := req.HandleBody[LinkCreateRequest](&w, r)
		if err != nil {
			res.Json(w, apierr.New(http.StatusText(http.StatusBadRequest), http.StatusBadRequest, ""), http.StatusBadRequest)
			return
		}

		link, err := handler.LinkService.Create(body.Url, userID)
		if err != nil {
			res.Json(w, apierr.New(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError, err.Error()), http.StatusInternalServerError)
			return
		}

		res.Json(w, link, http.StatusCreated)
	}
}

func (handler *LinkHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		jwtData, ok := r.Context().Value(middleware.ContextJWTDataKey).(*jwt.JWTData)

		if !ok || jwtData == nil {
			res.Json(w, apierr.New(http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized, ""), http.StatusUnauthorized)
			return
		}

		body, err := req.HandleBody[LinkUpdateRequest](&w, r)
		if err != nil {
			res.Json(w, apierr.New(http.StatusText(http.StatusBadRequest), http.StatusBadRequest, ""), http.StatusBadRequest)
			return
		}

		id, err := strconv.ParseUint(r.PathValue("id"), 10, 32)

		if err != nil {
			res.Json(w, apierr.New(http.StatusText(http.StatusBadRequest), http.StatusBadRequest, ""), http.StatusBadRequest)
			return
		}

		err = handler.LinkRepository.Update(&Link{
			Model: gorm.Model{ID: uint(id)},
			Url:   body.Url,
			Hash:  body.Hash,
		}, jwtData.UserID)
		if err != nil {
			res.Json(w, apierr.New(http.StatusText(http.StatusBadRequest), http.StatusBadRequest, err.Error()), http.StatusBadRequest)
			return
		}

		res.Json(w, http.StatusText(http.StatusOK), http.StatusOK)
	}
}

func (handler *LinkHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		jwtData, ok := r.Context().Value(middleware.ContextJWTDataKey).(*jwt.JWTData)

		if !ok || jwtData == nil {
			res.Json(w, apierr.New(http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized, ""), http.StatusUnauthorized)
			return
		}

		id, err := strconv.ParseUint(r.PathValue("id"), 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = handler.LinkRepository.Delete(uint(id), jwtData.UserID)
		if err != nil {
			res.Json(w, apierr.New(http.StatusText(http.StatusForbidden), http.StatusForbidden, err.Error()), http.StatusForbidden)
			return
		}

		res.Json(w, http.StatusText(http.StatusOK), http.StatusOK)
	}
}

func (handler *LinkHandler) GoTo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hash := r.PathValue("hash")

		link, err := handler.LinkRepository.GetByHash(hash)
		if err != nil {
			res.Json(w, apierr.New(http.StatusText(http.StatusNotFound), http.StatusNotFound, err.Error()), http.StatusNotFound)
			return
		}

		handler.EventBus.Publish(event.Event{
			Data: link.ID,
			Type: event.EventLinkClick,
		})

		http.Redirect(w, r, link.Url, http.StatusTemporaryRedirect)
	}
}

func (handler *LinkHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		jwtData, ok := r.Context().Value(middleware.ContextJWTDataKey).(*jwt.JWTData)

		if !ok || jwtData == nil {
			res.Json(w, apierr.New(http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized, ""), http.StatusUnauthorized)
			return
		}

		queryParams, err := req.HandleQueryParams[LinkGetAllParamsRequest](&w, r)
		if err != nil {
			res.Json(w, apierr.New(http.StatusText(http.StatusBadRequest), http.StatusBadRequest, err.Error()), http.StatusBadRequest)
			return
		}

		res.Json(w, LinkGetAllResponse{
			Count: handler.LinkRepository.GetCount(jwtData.UserID),
			Links: handler.LinkRepository.GetAll(queryParams.Limit, queryParams.Offset, jwtData.UserID),
		}, http.StatusOK)
	}
}
