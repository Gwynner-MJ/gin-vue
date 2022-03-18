package middleware

import (
	"gin-vue/commen"
	"gin-vue/model"
	"gin-vue/response"

	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//获取authorization header
		tokenString := ctx.GetHeader("Authorization")

		//验证token格式
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			response.Response(ctx, http.StatusUnprocessableEntity, 401, nil, "权限不足")
			ctx.Abort()
			return
		}

		tokenString = tokenString[7:]
		token, claims, err := commen.ParsaToken(tokenString)
		if err != nil || !token.Valid {
			response.Response(ctx, http.StatusUnprocessableEntity, 401, nil, "权限不足")
			ctx.Abort()
			return
		}

		//获取token上的UserId
		userId := claims.UserId
		DB := commen.GetDB()
		var user model.User
		DB.First(&user, userId)

		//用户如果不存在
		if user.ID == 0 {
			response.Response(ctx, http.StatusUnprocessableEntity, 401, nil, "权限不足")
			ctx.Abort()
			return
		}

		//用户存在,将user信息写入上下文
		ctx.Set("user", user)

		ctx.Next()

	}
}
