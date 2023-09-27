package app

import (
	"net/http"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/assetto-corsa-web/accweb/internal/pkg/cfg"
	"github.com/assetto-corsa-web/accweb/internal/pkg/server_manager"
	"github.com/assetto-corsa-web/accweb/public"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// @title           ACCWeb API documentation
// @description     ACCweb api documentation
// @termsOfService  http://swagger.io/terms/
// @version         1.19

// @contact.name   ACCWeb project
// @contact.url    https://github.com/assetto-corsa-web/accweb

// @license.name  MIT
// @license.url   https://github.com/assetto-corsa-web/accweb/blob/master/LICENSE

// @host      localhost:8080
// @BasePath  /api

// @securityDefinitions.apikey JWT
// @in header
// @name Authorization

const identityKey = "user_name"

type AccWError struct {
	Error string `json:"error"`
}

func newAccWError(msg string) AccWError {
	return AccWError{Error: msg}
}

type Handler struct {
	sm *server_manager.Service
}

func my(prefix string, fs http.FileSystem) *myFS {
	return &myFS{
		prefix: prefix,
		fs:     fs,
	}
}

type myFS struct {
	prefix string
	fs     http.FileSystem
}

func (f *myFS) Open(name string) (http.File, error) {
	return f.fs.Open(f.prefix + name)
}

func StartServer(config *cfg.Config, sM *server_manager.Service) {
	var r *gin.Engine

	if !config.Dev {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
		r.Use(gin.Recovery())
		_ = r.SetTrustedProxies(nil)
	} else {
		r = gin.Default()
	}

	// setup CORS
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{config.CORS.Origins}
	r.Use(cors.New(corsConfig))

	r.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/")
	})

	// setup routers
	setupRouters(r, sM, config)

	// Starting HTTP Server
	if config.Webserver.TLS {
		if err := r.RunTLS(config.Webserver.Host, config.Webserver.Cert, config.Webserver.PrivateKey); err != nil {
			logrus.WithError(err).Fatal("failed to start http server with TLS")
		}
	} else {
		if err := r.Run(config.Webserver.Host); err != nil {
			logrus.WithError(err).Fatal("failed to start http server")
		}
	}
}

func setupRouters(r *gin.Engine, sM *server_manager.Service, config *cfg.Config) {
	h := Handler{sm: sM}

	if config.Dev {
		basedir := "public"
		r.StaticFile("/", basedir+"/xindex.html")
		r.Static("/static", basedir+"/static")
		r.Static("/dist", basedir+"/dist")
	} else {
		r.GET("/", func(c *gin.Context) {
			c.FileFromFS("xindex.html", http.FS(public.Content))
		})
		r.StaticFS("/static", my("static", http.FS(public.Content)))
		r.StaticFS("/dist", my("dist", http.FS(public.Content)))
	}

	authMW := setupAuthRouters(r, config)

	api := r.Group("/api")
	api.Use(authMW.MiddlewareFunc())

	api.GET("/servers", h.ListServers)
	api.GET("/metadata", h.Metadata)
	api.GET("/instance/:id", h.GetInstance)
	api.GET("/instance/:id/logs", h.GetInstanceLogs)
	api.GET("/instance/:id/live", h.GetInstanceLiveState)
	api.GET("/instance/:id/export", h.ExportInstance)

	// moderator level
	mod := api.Group("")
	{
		mod.Use(ACCWebAuthMiddleware(ACCWebAuthLevel_Mod))
		mod.POST("/servers/stop-all", h.StopAllServers)
		mod.POST("/instance/:id/start", h.StartInstance)
		mod.POST("/instance/:id/stop", h.StopInstance)
	}

	// admin level
	adm := api.Group("")
	{
		adm.Use(ACCWebAuthMiddleware(ACCWebAuthLevel_Adm))
		adm.POST("/instance", h.NewInstance)
		adm.POST("/instance/:id", h.SaveInstance)
		adm.DELETE("/instance/:id", h.DeleteInstance)
		adm.POST("/instance/:id/clone", h.CloneInstance)
	}

}

type LoginPayload struct {
	Password string `json:"password"`
}

type User struct {
	UserName string `json:"user_name"`
	Admin    bool   `json:"admin"`
	Mod      bool   `json:"mod"`
	ReadOnly bool   `json:"read_only"`
}

func setupAuthRouters(r *gin.Engine, config *cfg.Config) *jwt.GinJWTMiddleware {
	// the jwt middleware
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		TokenLookup:      "header:Authorization,query:token",
		Realm:            "test zone",
		SigningAlgorithm: "RS256",
		PrivKeyFile:      config.Auth.PrivateKeyPath,
		PubKeyFile:       config.Auth.PublicKeyPath,
		Timeout:          20 * time.Minute,
		MaxRefresh:       time.Hour,
		IdentityKey:      identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*User); ok {
				return jwt.MapClaims{
					identityKey: v.UserName,
					"admin":     v.Admin,
					"mod":       v.Mod,
					"read_only": v.ReadOnly,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &User{
				UserName: claims[identityKey].(string),
				Admin:    claims["admin"].(bool),
				Mod:      claims["mod"].(bool),
				ReadOnly: claims["read_only"].(bool),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals LoginPayload
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			password := loginVals.Password

			var u *User

			isAdmin := password == config.Auth.AdminPassword
			isMod := password == config.Auth.ModeratorPassword || isAdmin
			isRO := password == config.Auth.ReadOnlyPassword || isMod

			if !isAdmin && !isMod && !isRO {
				return nil, jwt.ErrFailedAuthentication
			}

			u = &User{
				UserName: "username",
				Admin:    isAdmin,
				Mod:      isMod,
				ReadOnly: isRO,
			}
			c.Set(identityKey, u)

			return u, nil
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(*User); ok && (v.Admin || v.Mod || v.ReadOnly) {
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		LoginResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			x, _ := c.Get(identityKey)
			u := x.(*User)

			c.JSON(http.StatusOK, gin.H{
				"code":      http.StatusOK,
				"token":     token,
				"expire":    expire.Format(time.RFC3339),
				"user_name": u.UserName,
				"admin":     u.Admin,
				"mod":       u.Mod,
				"read_only": u.ReadOnly,
			})
		},
	})

	if err != nil {
		logrus.Fatal("JWT Error:" + err.Error())
	}

	// When you use jwt.New(), the function is already automatically called for checking,
	// which means you don't need to call it again.
	errInit := authMiddleware.MiddlewareInit()

	if errInit != nil {
		logrus.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}

	r.POST("/api/login", authMiddleware.LoginHandler)
	r.GET("/api/refresh_token", authMiddleware.RefreshHandler)
	r.GET("/api/logout", authMiddleware.MiddlewareFunc(), authMiddleware.LogoutHandler)
	r.GET("/api/token", authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{})
	})

	return authMiddleware
}

func GetUserFromClaims(c *gin.Context) *User {
	if user, ok := c.Get(identityKey); ok {
		return user.(*User)
	} else {
		claims := jwt.ExtractClaims(c)
		return &User{
			UserName: claims[identityKey].(string),
			Admin:    claims["admin"].(bool),
			Mod:      claims["mod"].(bool),
			ReadOnly: claims["read_only"].(bool),
		}
	}
}
