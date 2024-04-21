package helper

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net/http"
	"os"
	"reimbursement/constants"
	"reimbursement/helper/models"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/kelseyhightower/envconfig"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

// serverfileca ...
var (
	// readtimeout               string
	// readheadertimeout         string
	// writetimeout              string
	// idletimeout               string
	// maxheaderbyte             string
	serverfileca string
	// serverfileprivatekey      string
	// serverfilepubkey          string
	servertls12client string
	// consuladdres              string
	// consulregid               string
	// consulregname             string
	// consulregserver           string
	// consulregport             string
	// consulhealtCheckHttp      string
	// consulhealtcheckInterval  string
	// consulhealtcheckTimeout   string
	// consulonof                string
	// kafkaBrokerUrl            string
	// kafkaClient               string
	// kafkaProducerTimeout      string
	// kafkaProducerDialTimeout  string
	// kafkaProducerReadTimeout  stringz
	// kafkaProducerWriteTimeout string
	// kafkaProducerMaxmsgbyte   string
)

// GetServerConfig ...
func GetServerConfig() *models.ServerConfig {
	var serverCfg models.ServerConfig
	err := envconfig.Process(constants.SERVERPREFIX, &serverCfg)
	fmt.Println("Error Config Consul : ", err)
	return &serverCfg
}

// GetServerTlsConfig ...
func GetServerTlsConfig() *tls.Config {
	if servertls12client == "ON" {

		caCert, err := os.ReadFile(serverfileca)
		if err != nil {
			fmt.Println("Error : ", err)

		}
		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(caCert)

		tlsConfig := &tls.Config{
			ClientCAs:  caCertPool,
			ClientAuth: tls.RequireAndVerifyClientCert,
		}
		tlsConfig.BuildNameToCertificate()
		return tlsConfig
	}
	return &tls.Config{InsecureSkipVerify: true}

}

// GinServerUp ...
func GinServerUp(listenAddr string, router *gin.Engine) error {

	cfg := *GetServerConfig()
	fmt.Println("[TLS.1.2]:", cfg.Servertls12client)
	srv := &http.Server{
		Addr:              listenAddr,
		Handler:           router,
		TLSConfig:         GetServerTlsConfig(),
		ReadTimeout:       cfg.ReadTimeout,
		ReadHeaderTimeout: cfg.ReadHeaderTimeout,
		WriteTimeout:      cfg.WriteTimeout,
		IdleTimeout:       cfg.IdleTimeout,
		MaxHeaderBytes:    cfg.MaxHeaderBytes,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
	}

	if cfg.Servertls12client == "ON" {
		return srv.ListenAndServeTLS(cfg.Serverfilepubkey, cfg.Serverfileprivatekey)
	}
	return srv.ListenAndServe()
}

// OpenTracer - middleware that addes opentracing
func OpenTracer(operationPrefix []byte) gin.HandlerFunc {
	if operationPrefix == nil {
		operationPrefix = []byte("api-request-")
	}
	return func(c *gin.Context) {
		// all before request is handled
		var span opentracing.Span
		if cspan, ok := c.Get("tracing-context"); ok {
			span = StartSpanWithParent(cspan.(opentracing.Span).Context(), string(operationPrefix)+c.Request.RequestURI, c.Request.Method, c.Request.URL.Path)

		} else {
			span = StartSpanWithHeader(&c.Request.Header, string(operationPrefix)+c.Request.RequestURI, c.Request.Method, c.Request.URL.Path)
		}
		defer span.Finish()            // after all the other defers are completed.. finish the span
		c.Set("tracing-context", span) // add the span to the context so it can be used for the duration of the request.
		c.Next()

		span.SetTag(string(ext.HTTPStatusCode), c.Writer.Status())
	}
}

// StartSpanWithParent ...
func StartSpanWithParent(parent opentracing.SpanContext, operationName, method, path string) opentracing.Span {
	options := []opentracing.StartSpanOption{
		opentracing.Tag{Key: ext.SpanKindRPCServer.Key, Value: ext.SpanKindRPCServer.Value},
		opentracing.Tag{Key: string(ext.HTTPMethod), Value: method},
		opentracing.Tag{Key: string(ext.HTTPUrl), Value: path},
		opentracing.Tag{Key: "current-goroutines", Value: runtime.NumGoroutine()},
	}

	if parent != nil {
		options = append(options, opentracing.ChildOf(parent))
	}

	return opentracing.StartSpan(operationName, options...)
}

// StartSpanWithHeader ...
func StartSpanWithHeader(header *http.Header, operationName, method, path string) opentracing.Span {
	var wireContext opentracing.SpanContext
	if header != nil {
		wireContext, _ = opentracing.GlobalTracer().Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(*header))
	}
	span := StartSpanWithParent(wireContext, operationName, method, path)
	span.SetTag("current-goroutines", runtime.NumGoroutine())
	return span
	// return StartSpanWithParent(wireContext, operationName, method, path)
}
