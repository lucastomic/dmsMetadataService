package main

import (
	"github.com/lucastomic/dmsMetadataService/internal/controller"
	"github.com/lucastomic/dmsMetadataService/internal/idgenerator"
	"github.com/lucastomic/dmsMetadataService/internal/logging"
	"github.com/lucastomic/dmsMetadataService/internal/metadataservice"
	"github.com/lucastomic/dmsMetadataService/internal/middleware"
	"github.com/lucastomic/dmsMetadataService/internal/server"
	"github.com/lucastomic/dmsMetadataService/internal/storageserviceurl"
)

func main() {
	apilogger := logging.NewLogrusLogger()
	logicLogger := logging.NewLogrusLogger()
	storageserviceurl := storageserviceurl.New(logicLogger)
	idgenerator := idgenerator.New()
	metadataservice := metadataservice.New(logicLogger, storageserviceurl, &idgenerator)
	controller := controller.NewMetadataController(logicLogger, metadataservice)
	middlewares := []middleware.Middleware{
		middleware.NewLoggingMiddleware(apilogger),
		middleware.NewRequestIDMiddleware(),
	}
	server := server.New(":3002", controller, apilogger, logicLogger, middlewares)
	server.Run()
}
