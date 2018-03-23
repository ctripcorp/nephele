package middleware

import (
	_ "github.com/ctripcorp/nephele/context"
	"github.com/gin-gonic/gin"
)

// all panic will be caught here.
func recovery() gin.HandlerFunc {
	return func(httpCtx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				//todo
				//ctx := httpCtx.MustGet(context.GlobalName).(*context.Context)
				httpCtx.AbortWithStatus(500)
			}
		}()

		httpCtx.Next()
	}
}
