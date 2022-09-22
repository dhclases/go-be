package main

import (
	configDb "crudgolang/config"
	deliveryHttp "crudgolang/delivery"
	pbRepo "crudgolang/repository"
	pbService "crudgolang/service"
	"net/http"

	"os"

	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	configDb.InitConnMySQLDB()
}

func main() {
	// TODO: Read from .env
	port := "8001"

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if gin.IsDebugging() {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	log.Logger = log.Output(
		zerolog.ConsoleWriter{
			Out:     os.Stderr,
			NoColor: false,
		},
	)

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	// Add a logger middleware, which:
	//   - Logs all requests, like a combined access and error log.
	//   - Logs to stdout.
	r.Use(logger.SetLogger())

	r.Use(gin.Recovery())

	db, err := configDb.GetMySQLDB()
	if err != nil {
		log.Error().Msg("Failed Connect to database =" + err.Error())
		return
	}

	defer db.Close()

	phoneRepo := pbRepo.NewPhoneBookRepo(db)
	phoneSrv := pbService.NewPhoneBookService(phoneRepo)

	deliveryHttp.NewPhoneBookHTTPHandler(r, phoneSrv)
	log.Info().Msg("Service Running at port : " + port)

	// TODO: Evaluate
	//api for login user
	// r.POST("api/login", authMiddleware.LoginHandler)

	if errHTTP := http.ListenAndServe(":"+port, r); errHTTP != nil {
		log.Error().Msg(errHTTP.Error())
	}
}
