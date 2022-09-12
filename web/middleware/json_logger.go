package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func JSONLogMIddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// process request
		ctx.Next()

		if ctx.Writer.Status() >= 500 {
			log.Error().Msg(ctx.Errors.String())
		} else {
			log.Info().
				Str("client_ip", ctx.ClientIP()).
				Str("method", ctx.Request.Method).
				Str("path", ctx.Request.RequestURI).
				Int("status", ctx.Writer.Status()).
				Str("user_id", ctx.Request.URL.User.Username()).
				Str("referrer", ctx.Request.Referer()).
				Str("request_id", ctx.Writer.Header().Get("Request-Id")).
				Send()
		}

	}
}
