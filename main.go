package main

import (
	"context"
	"fmt"
	"github.com/assetto-corsa-web/accweb/auth"
	"github.com/assetto-corsa-web/accweb/config"
	"github.com/assetto-corsa-web/accweb/pages"
	"github.com/emvi/logbuch"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"
)

const (
	logTimeFormat   = "2006-01-02_15:04:05"
	shutdownTimeout = time.Second * 30
	dataDir         = "data"
	staticDir       = "static"
)

func getRouter() *mux.Router {
	router := mux.NewRouter()
	router.Handle("/", auth.Middleware(pages.Overview))
	router.HandleFunc("/login", pages.Login)
	router.HandleFunc("/logout", pages.Logout)
	router.Handle("/user", auth.Middleware(pages.User))
	router.Handle("/user/new", auth.Middleware(pages.NewUser))
	router.Handle("/user/edit", auth.Middleware(pages.EditUser))
	router.Handle("/user/delete", auth.Middleware(pages.DeleteUser))
	router.Handle("/server", auth.Middleware(pages.Server))
	router.Handle("/logs", auth.Middleware(pages.Logs))
	router.HandleFunc("/status", pages.Status)
	staticDirPrefix := fmt.Sprintf("/%s/", staticDir)
	router.PathPrefix(staticDirPrefix).Handler(http.StripPrefix(staticDirPrefix, http.FileServer(http.Dir(staticDir))))
	router.NotFoundHandler = http.HandlerFunc(pages.NotFound)
	return router
}

func configureLogging() {
	logbuch.SetFormatter(logbuch.NewFieldFormatter(logTimeFormat, "\t\t"))
	level := strings.ToLower(config.Get().Loglevel)
	logbuch.Info("Configuring logging...", logbuch.Fields{"level": level})

	if level == "debug" {
		logbuch.SetLevel(logbuch.LevelDebug)
	} else {
		logbuch.SetLevel(logbuch.LevelInfo)
	}
}

func configureCors(router *mux.Router) http.Handler {
	logbuch.Info("Configuring CORS...")
	origins := strings.Split(config.Get().CORS.Origins, ",")
	c := cors.New(cors.Options{
		AllowedOrigins:   origins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		Debug:            strings.ToLower(config.Get().CORS.Loglevel) == "debug",
	})
	return c.Handler(router)
}

func createDataDirectory() {
	if _, err := os.Stat(dataDir); err != nil {
		if err := os.Mkdir(dataDir, 0744); err != nil {
			logbuch.Fatal("Error creating data directory. Make sure the data directory can be created", logbuch.Fields{"err": err, "path": dataDir})
		}
	}
}

func startServer(handler http.Handler) {
	logbuch.Info("Starting server...")
	server := &http.Server{
		Handler:      handler,
		Addr:         config.Get().Server.Host,
		WriteTimeout: time.Duration(config.Get().Server.WriteTimeout) * time.Second,
		ReadTimeout:  time.Duration(config.Get().Server.ReadTimeout) * time.Second,
	}

	go func() {
		sigint := make(chan os.Signal)
		signal.Notify(sigint, os.Interrupt)
		<-sigint
		logbuch.Info("Shutting down server...")
		ctx, _ := context.WithTimeout(context.Background(), shutdownTimeout)

		if err := server.Shutdown(ctx); err != nil {
			logbuch.Fatal("Error shutting down server gracefully", logbuch.Fields{"err": err})
		}
	}()

	if config.Get().Server.TLS {
		logbuch.Info("TLS enabled")

		if err := server.ListenAndServeTLS(config.Get().Server.Cert, config.Get().Server.PrivateKey); err != http.ErrServerClosed {
			logbuch.Fatal("Error starting server with TLS enabled", logbuch.Fields{"err": err})
		}
	} else {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			logbuch.Fatal("Error starting server with TLS disabled", logbuch.Fields{"err": err})
		}
	}
}

func main() {
	config.Load()
	configureLogging()
	createDataDirectory()
	auth.LoadConfig()
	pages.LoadTemplate()
	router := getRouter()
	configureCors(router)
	startServer(router)
}
