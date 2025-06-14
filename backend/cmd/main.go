package main

import (
	"fmt"
	"net/http"

	"github.com/vnkot/piklnk/configs"
	"github.com/vnkot/piklnk/internal/auth"
	"github.com/vnkot/piklnk/internal/link"
	"github.com/vnkot/piklnk/internal/stat"
	"github.com/vnkot/piklnk/internal/user"
	"github.com/vnkot/piklnk/pkg/db"
	"github.com/vnkot/piklnk/pkg/event"
	"github.com/vnkot/piklnk/pkg/middleware"
)

func main() {
	conf := configs.LoadConfig()

	db := db.NewDb(conf)
	router := http.NewServeMux()
	eventBus := event.NewEventBus()

	linkRepository := link.NewLinkRepository(db)
	userRepository := user.NewUserRepository(db)
	statRepository := stat.NewStatRepository(db)

	authService := auth.NewAuthService(userRepository)
	linkService := link.NewLinkService(&link.LinkServiceDeps{
		LinkRepository: linkRepository,
	})
	statService := stat.NewStatService(&stat.StatServiceDeps{
		EventBus:       eventBus,
		StatRepository: statRepository,
	})

	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config:      conf,
		AuthService: authService,
	})
	link.NewLinkHandler(router, link.LinkHandlerDeps{
		Config:         conf,
		EventBus:       eventBus,
		LinkService:    linkService,
		LinkRepository: linkRepository,
		UserRepository: userRepository,
	})
	stat.NewStatHandler(router, stat.StatHandlerDeps{
		Config:         conf,
		StatRepository: statRepository,
		LinkRepository: linkRepository,
	})

	stack := middleware.Chain(
		middleware.CORS,
		middleware.Logging,
	)

	server := http.Server{
		Addr:    ":8000",
		Handler: stack(router),
	}

	fmt.Println("Server is listening on port 8000")

	go statService.AddClick()
	server.ListenAndServe()
}
