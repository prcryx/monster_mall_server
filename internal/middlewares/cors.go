package middlewares

import (
	"net/http"

	"github.com/go-chi/cors"
	"github.com/prcryx/monster_mall/internal/common/constants"
)

func Cors() func(http.Handler) http.Handler {
    return cors.Handler(
		cors.Options{
			AllowedOrigins:   constants.AllowedOrigins(),
			AllowedHeaders:   constants.AllowedHeaders(),
			AllowedMethods:   constants.AllowedMethods(),
			AllowCredentials: constants.AllowCredentials,
			MaxAge:           constants.MaxAge,
			ExposedHeaders:   constants.ExposedHeaders(),
		},
	);
}