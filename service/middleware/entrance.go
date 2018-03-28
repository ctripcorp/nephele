package middleware

import (
	"github.com/ctripcorp/nephele/context"
	"github.com/ctripcorp/nephele/log"
	"github.com/gin-gonic/gin"
)

// all panic will be caught here.
func entrance() gin.HandlerFunc {
	return func(httpCtx *gin.Context) {
		ctx := httpCtx.MustGet(context.GlobalName).(*context.Context)
		defer func() {
			err := recover()
			if err != nil {
				//todo
				httpCtx.AbortWithStatus(500)
			}
			log.TraceEndRoot(ctx, err)
		}()

		log.TraceBegin(ctx, "", "middleware", "entrance", "URL", httpCtx.Request.RequestURI)
		httpCtx.Next()
	}
}
