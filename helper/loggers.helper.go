package helper

import (
	"context"
	"reimbursement/constants"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	// elastic "github.com/olivere/elastic/v7"
	log "github.com/sirupsen/logrus"
	// elogrus "gopkg.in/sohlich/elogrus.v7"
)

// InitializeNewLogs ...
func InitializeNewLogs() *log.Logger {
	l := log.New()

	// Connect Elasticsearch
	// client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"))
	// if err != nil {
	// 	log.Panic(err)
	// }
	// hook, err := elogrus.NewAsyncElasticHook(client, "localhost", log.DebugLevel, "skeleton-svc")
	// if err != nil {
	// 	log.Panic(err)
	// }
	// l.Hooks.Add(hook)
	// end connection

	// if GetEnv("ENV_HOST", "PROD") == "PROD" {
	l.SetFormatter(&log.JSONFormatter{
		DataKey:     uuid.NewString(),
		PrettyPrint: false,
	})
	l.SetLevel(log.DebugLevel)
	// } else {
	// 	// The TextFormatter is default, you don't actually have to do this.
	// 	l.SetFormatter(&log.TextFormatter{})
	// 	l.SetLevel(log.InfoLevel)
	// }

	return l
}

// GettingResponseLog ...
func GettingResponseLog(ctxID string, g *gin.Context, req, res interface{}) log.Fields {
	return log.Fields{
		"ID":       ctxID,
		"Method":   g.Request.Method,
		"Path":     g.Request.URL.Path,
		"Header":   g.Request.Header,
		"ClientIP": g.Request.RemoteAddr,
		"Body":     req,
		"Response": res,
	}
}

// GettingDetaultLog ...
func GettingDetaultLog(ctx context.Context, message interface{}) log.Fields {
	return log.Fields{
		"ID":      ctx.Value(constants.SPAN_ID),
		"Message": message,
	}
}
