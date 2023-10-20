package route

import (
	"bytes"
	"github.com/gin-contrib/requestid"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	_ "github.com/leslieleung/gin-application-template/docs" // docs is generated by Swag CLI, you have to import it.
	"github.com/leslieleung/gin-application-template/internal/controller"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"time"
)

func RegisterRoute() *gin.Engine {
	r := gin.Default()

	registerMiddlewares(r)

	r.GET("/ping", controller.PingPong)
	// Swagger API docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return r
}

func registerMiddlewares(r *gin.Engine) {
	r.Use(requestid.New())
	logger, _ := zap.NewProduction()
	r.Use(ginzap.GinzapWithConfig(logger, &ginzap.Config{
		TimeFormat: time.RFC3339Nano,
		UTC:        true,
		SkipPaths:  []string{},
		Context: ginzap.Fn(func(c *gin.Context) []zapcore.Field {
			fields := []zapcore.Field{}
			fields = append(fields, zap.String("request_id", requestid.Get(c)))

			// log request body
			var body []byte
			var buf bytes.Buffer
			tee := io.TeeReader(c.Request.Body, &buf)
			body, _ = io.ReadAll(tee)
			c.Request.Body = io.NopCloser(&buf)
			fields = append(fields, zap.String("body", string(body)))
			return fields
		}),
	}))
}
