package auth

import (
	"net/http"

	"github.com/vnkot/piklnk/configs"
	"github.com/vnkot/piklnk/pkg/apierr"
	"github.com/vnkot/piklnk/pkg/jwt"
	"github.com/vnkot/piklnk/pkg/req"
	"github.com/vnkot/piklnk/pkg/res"
)

type AuthHandlerDeps struct {
	*AuthService
	*configs.Config
}

type AuthHandler struct {
	*AuthService
	*configs.Config
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config:      deps.Config,
		AuthService: deps.AuthService,
	}

	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[LoginRequest](&w, r)

		if err != nil {
			res.Json(w, apierr.New(http.StatusText(http.StatusBadRequest), http.StatusBadRequest, ""), http.StatusBadRequest)
			return
		}

		userID, err := handler.AuthService.Login(body.Email, body.Password)
		if err != nil {
			res.Json(w, apierr.New(http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized, ""), http.StatusUnauthorized)
			return
		}

		token, err := jwt.NewJWT(handler.Config.Auth.Secret).Create(jwt.JWTData{UserID: *userID})
		if err != nil {
			res.Json(w, apierr.New(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError, ""), http.StatusInternalServerError)
			return
		}

		res.Json(w, LoginResponse{
			Token: token,
		}, http.StatusOK)
	}
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[RegisterRequest](&w, r)

		if err != nil {
			res.Json(w, apierr.New(http.StatusText(http.StatusBadRequest), http.StatusBadRequest, ""), http.StatusBadRequest)
			return
		}

		userID, err := handler.AuthService.Register(body.Email, body.Password, body.Name)
		if err != nil {
			res.Json(w, apierr.New(http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized, ""), http.StatusUnauthorized)
			return
		}

		token, err := jwt.NewJWT(handler.Config.Auth.Secret).Create(jwt.JWTData{UserID: *userID})
		if err != nil {
			res.Json(w, apierr.New(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError, ""), http.StatusInternalServerError)
			return
		}

		res.Json(w, RegisterResponse{
			Token: token,
		}, http.StatusOK)
	}
}
