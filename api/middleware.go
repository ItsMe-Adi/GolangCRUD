package api

import (
	"fmt"
	"net/http"
	"playground/token"
	"playground/util"
	"strings"

	"github.com/gin-gonic/gin"
)

func authMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := ctx.GetHeader("authorization")
		if len(header) == 0 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, util.GetError(fmt.Errorf("unauthorized")))
			return
		}
		fields := strings.Fields(header)
		if len(fields) < 2 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, util.GetError(fmt.Errorf("unauthorized")))
			return 
		}
		accessToken := fields[1]
		if !token.ValidateToken(accessToken) {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, util.GetError(fmt.Errorf("unauthorized")))
			return
		}
		ctx.Next()
	}
}
