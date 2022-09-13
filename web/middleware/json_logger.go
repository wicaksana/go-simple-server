package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
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
				Dict("httpRequest", zerolog.Dict().
					Str("requestMethod", ctx.Request.Method).
					Str("requestUrl", ctx.Request.RequestURI).
					Str("remoteIp", ctx.ClientIP()).
					Int("status", ctx.Writer.Status()).
					Str("referer", ctx.Request.Referer()),
				).Send()
		}
	}
}
