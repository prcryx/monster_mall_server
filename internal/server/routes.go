package server

import (
	"github.com/go-chi/chi/v5"
	authCtrl "github.com/prcryx/monster_mall/internal/apis/auth/auth_ctrl"
)

// global routes
func (ap *APIServer) registerGlobalRoutes() {
	ap.setAuthRoute(V1)
}

func (ap *APIServer) setAuthRoute(path string) {
	ac := authCtrl.NewAuthController(ap.db)
	router := chi.NewRouter()
	router.Post(SendOtp, ac.SendOtp)

	ap.rootRouter.Mount(path, router)
}
