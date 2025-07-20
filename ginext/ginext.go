package ginext

import (
	"bytes"
	"encoding/json"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type Engine struct {
	*gin.Engine
}

type responseBody struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

type Context = gin.Context

func New() *Engine {
	return &Engine{gin.New()}
}

func (e *Engine) GET(relativePath string, handlers ...gin.HandlerFunc) {
	e.Engine.GET(relativePath, handlers...)
}

func (e *Engine) Run(addr ...string) error {
	return e.Engine.Run(addr...)
}

func (e Engine) WithLogging(logger *zerolog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		logLevel := logger.GetLevel()

		if logLevel == zerolog.Disabled {
			c.Next()
			return
		}

		start := time.Now()

		var writer *responseBody
		if logLevel <= zerolog.DebugLevel {
			writer = &responseBody{
				ResponseWriter: c.Writer,
				body:           bytes.NewBufferString(""),
			}
			c.Writer = writer
		}

		c.Next()

		logCtx := logger.With().Logger()

		// logCtx := logger.With().
		// 	Str("Request method", c.Request.Method).
		// 	Str("Request path", c.Request.URL.Path).
		// 	Str("Request query", c.Request.URL.RawQuery).
		// 	Str("Request ip", c.ClientIP()).
		// 	Str("Request Content-Type", c.ContentType()).
		// 	Str("Request user-agent", c.Request.UserAgent()).
		// 	Int("Response status", c.Writer.Status()).
		// 	Int("Response size", c.Writer.Size())

		fields := map[string]any{
			"Request method":       c.Request.Method,
			"Request path":         c.Request.URL.Path,
			"Request query":        c.Request.URL.RawQuery,
			"Request ip":           c.ClientIP(),
			"Request Content-Type": c.ContentType(),
			"Request user-agent":   c.Request.UserAgent(),
			"Response status":      c.Writer.Status(),
			"Response size":        c.Writer.Size(),
		}

		if logLevel <= zerolog.DebugLevel && writer != nil {
			if body := writer.body.String(); body != "" {
				if strings.Contains(writer.Header().Get("Content-Type"), "application/json") {
					var pretty bytes.Buffer
					if err := json.Indent(&pretty, []byte(body), "", "  "); err != nil {
						// logCtx = logCtx.RawJSON("Response body", []byte(body))
						fields["Response body"] = pretty.String()
					} else {
						// logCtx = logCtx.RawJSON("Response body", []byte(pretty.Bytes()))
						fields["Response body"] = body
					}
				} else {
					// var truncateBody string
					maxLenBody := 1024
					if len(body) > maxLenBody {
						// truncateBody = body[:maxLenBody] + "...[truncated]"
						fields["Response body"] = body[:maxLenBody] + "...[truncated]"
					} else {
						// truncateBody = body
						// logCtx = logCtx.Str("Response body", truncateBody)
						fields["Response body"] = body
					}
				}
			}
		}

		// logCtx = logCtx.Dur("latency", time.Since(start))
		fields["Latency"] = time.Since(start)

		// log := logCtx.Logger()

		msg := "Request handled"

		if len(c.Errors) > 0 {
			errors := make([]error, len(c.Errors))
			for i, e := range c.Errors {
				errors[i] = e.Err
			}
			logCtx.Error().Errs("errors", errors).Msg(msg)
		} else {
			switch {
			case logLevel <= zerolog.DebugLevel:
				logCtx.Debug().Fields(fields).Msg(msg)
			case logLevel <= zerolog.InfoLevel:
				logCtx.Info().Fields(fields).Msg(msg)
			}
		}
	}
}

func (w *responseBody) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
