package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/prcryx/monster_mall/internal/middlewares"
	"gorm.io/gorm"
)

type APIServer struct {
	port       uint16
	serverAddr string
	db         *gorm.DB
	rootRouter *chi.Mux
}

func NewAPIServer(serverAddr string, port uint16, db *gorm.DB) *APIServer {
	return &APIServer{
		serverAddr: serverAddr,
		port:       port,
		db:         db,
		rootRouter: chi.NewRouter(),
	}
}

func (ap *APIServer) setupRoutes() {
	// ap.rootRouter.Use(middlewares.Cors())
	ap.rootRouter.Use(middlewares.Cors())
	ap.registerGlobalRoutes()
	// ap.registerGuardedRoute()
}

func (ap *APIServer) Run() {
	// defer ap.cleanUp()
	ap.setupRoutes()
	log.Printf("Starting server at Port :%v", ap.port)
	err := http.ListenAndServe(fmt.Sprintf("%v:%v", ap.serverAddr, ap.port), ap.rootRouter)
	if err != nil {
		log.Fatal("Error occurred", err)
		return
	}
}

//clean ups
