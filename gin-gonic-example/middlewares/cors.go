package middlewares

import "github.com/rs/cors"

func CORS() *cors.Cors {
return cors.New(cors.Options{
	AllowedOrigins: []string{"*"},
	AllowedMethods: []string{"*"},
	AllowedHeaders: []string{"*"},
})
}