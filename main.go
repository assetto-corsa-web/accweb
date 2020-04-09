package main

import (
	"github.com/assetto-corsa-web/accweb/api"
	"github.com/assetto-corsa-web/accweb/cfg"
	serverList "github.com/assetto-corsa-web/accweb/server"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	staticDir               = "public/static"
	staticDirPrefix         = "/static/"
	buildJsFile             = "public/dist/build.js"
	buildJsPrefix           = "/dist/build.js"
	cssFile                 = "public/dist/main.css"
	cssFilePrefix           = "/dist/main.css"
	indexFile               = "public/index.html"
	rootDirPrefix           = "/"
	defaultHttpWriteTimeout = 20
	defaultHttpReadTimeout  = 20
)

var (
	buildJs      []byte
	watchBuildJs = false
)

func configureLog() {
	logrus.Info("Configure logging...")
	level := strings.ToLower(cfg.Get().Loglevel)

	if level == "debug" {
		logrus.SetLevel(logrus.DebugLevel)
	} else if level == "info" {
		logrus.SetLevel(logrus.InfoLevel)
	} else {
		logrus.SetLevel(logrus.WarnLevel)
	}
}

func createConfigDir() {
	if err := os.MkdirAll(cfg.Get().ConfigPath, 0755); err != nil {
		logrus.WithError(err).Fatal("Error creating config directory")
	}
}

func loadBuildJs() {
	logrus.Info("Loading build.js...")
	watchBuildJs = cfg.Get().Dev
	content, err := ioutil.ReadFile(buildJsFile)

	if err != nil {
		logrus.WithField("err", err).Fatal("build.js not found")
	}

	buildJs = content
}

func setupRouter() *mux.Router {
	router := mux.NewRouter()

	// REST endpoints
	router.Handle("/api/token", api.AuthMiddleware(api.TokenHandler, false, false)).Methods("GET")
	router.Handle("/api/server", api.AuthMiddleware(api.SaveServerSettingsHandler, true, false)).Methods("POST")
	router.Handle("/api/server", api.AuthMiddleware(api.CopyServerSetttingsHandler, true, false)).Methods("PUT")
	router.Handle("/api/server", api.AuthMiddleware(api.DeleteServerHandler, true, false)).Methods("DELETE")
	router.Handle("/api/server", api.AuthMiddleware(api.GetServerHandler, false, false)).Methods("GET")
	router.HandleFunc("/api/status", api.GetServerStatusHandler).Methods("GET")
	router.Handle("/api/server/import", api.AuthMiddleware(api.ImportServerHandler, true, false)).Methods("POST")
	router.Handle("/api/instance", api.AuthMiddleware(api.StartInstanceHandler, false, true)).Methods("POST")
	router.Handle("/api/instance", api.AuthMiddleware(api.StopInstanceHandler, false, true)).Methods("DELETE")
	router.Handle("/api/instance", api.AuthMiddleware(api.GetInstanceLogsHandler, false, false)).Methods("GET")
	router.HandleFunc("/api/server/export/{filename}", api.ExportServerHandler).Methods("GET")
	router.HandleFunc("/api/login", api.LoginHandler).Methods("POST")

	// static content
	router.PathPrefix(staticDirPrefix).Handler(http.StripPrefix(staticDirPrefix, http.FileServer(http.Dir(staticDir))))
	router.PathPrefix(buildJsPrefix).Handler(http.StripPrefix(buildJsPrefix, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if watchBuildJs {
			loadBuildJs()
		}

		w.Header().Add("Content-Type", "text/javascript")

		if _, err := w.Write(buildJs); err != nil {
			w.WriteHeader(http.StatusNotFound)
		}
	})))
	router.PathPrefix(cssFilePrefix).Handler(http.StripPrefix(cssFilePrefix, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/css")
		http.ServeFile(w, r, cssFile)
	})))
	router.PathPrefix(rootDirPrefix).Handler(http.StripPrefix(rootDirPrefix, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, indexFile)
	})))

	return router
}

func configureCors(router *mux.Router) http.Handler {
	logrus.Info("Configuring CORS...")

	origins := strings.Split(cfg.Get().CORS.Origins, ",")
	c := cors.New(cors.Options{
		AllowedOrigins:   origins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		Debug:            strings.ToLower(cfg.Get().CORS.Loglevel) == "debug",
	})
	return c.Handler(router)
}

func start(handler http.Handler) {
	logrus.Info("Starting server...")
	server := &http.Server{
		Handler:      handler,
		Addr:         cfg.Get().Webserver.Host,
		WriteTimeout: time.Duration(cfg.Get().Webserver.WriteTimeout) * time.Second,
		ReadTimeout:  time.Duration(cfg.Get().Webserver.ReadTimeout) * time.Second,
	}

	if cfg.Get().Webserver.TLS {
		logrus.Info("TLS enabled")
		logrus.Fatal(server.ListenAndServeTLS(cfg.Get().Webserver.Cert, cfg.Get().Webserver.PrivateKey))
	} else {
		logrus.Fatal(server.ListenAndServe())
	}
}

func main() {
	cfg.Load()
	configureLog()
	createConfigDir()
	api.LoadConfig()
	loadBuildJs()
	serverList.LoadServerList()
	router := setupRouter()
	corsConfig := configureCors(router)
	start(corsConfig)
}
