// // configuring the server

package http

// import (
// 	cfg "media_api/config"
// 	film_handling "media_api/internal/adapter/delivery/http/handler"
// 	usecase "media_api/internal/usecase"
// 	"net/http"

// 	"github.com/rs/cors"
// 	"golang.org/x/exp/slog"
// )

// func ConfigureHttpServer(
// 	http_config *cfg.HTTP,
// 	usecase *usecase.RentalUseCase,
// ) *http.Server {

// 	slog.Info("Configuring the HTTP Server\n")

// 	filmHandler := film_handling.New(usecase)

// 	apiHandler := NewHttpHandler(filmHandler)

// 	configuredHandler := cors.New(
// 		cors.Options{
// 			AllowedOrigins:   []string{"*"},
// 			AllowOriginFunc:  func(origin string) bool { return true },
// 			AllowCredentials: true,
// 			ExposedHeaders:   []string{"Content-Length", "ETag", "Link", "X-RateLimit-Limit", "X-RateLimit-Remaining"},
// 			Debug:            true,
// 		},
// 	).Handler(apiHandler)

// 	mux := http.NewServeMux()
// 	mux.Handle("/films", configuredHandler)

// 	return &http.Server{
// 		Addr:    http_config.Host + ":" + http_config.Port,
// 		Handler: mux,
// 	}
// }
