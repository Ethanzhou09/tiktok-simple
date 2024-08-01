package middleware

// import (
// 	"context"
// 	"github.com/cloudwego/hertz/pkg/app"
// 	"tiktok-simple/pkg/jwt"
// 	"tiktok-simple/idl/kitex_gen/user"
// )

// func JWT() app.HandlerFunc {
// 	return func(c context.Context, ctx *app.RequestContext) {
// 		var code uint32
// 		req := api.TaskRequest{}
// 		code = 200
// 		token := string(ctx.GetHeader("token"))
// 		if token == "" {
// 			code = 404
// 			ctx.JSON(500, map[string]interface{}{
// 				"code": code,
// 				"msg":  "鉴权失败",
// 			})
// 			ctx.Abort()
// 		}
// 		claims, err := jwt.ParseToken(token)
// 		if err != nil {
// 			code = 401
// 			ctx.JSON(500, map[string]interface{}{
// 				"code": code,
// 				"msg":  "鉴权失败",
// 			})
// 			ctx.Abort()
// 		}
// 		if err := ctx.Bind(&req); err != nil {
// 			ctx.JSON(500, "CreateTaskHandler-ShouldBindJSON")
// 			ctx.Abort()
// 		}
// 		ctx_id := req.Uid
// 		if claims.Id != uint(ctx_id) {
// 			ctx.JSON(500, map[string]interface{}{
// 				"code": code,
// 				"msg":  "请求用户信息不匹配",
// 			})
// 			ctx.Abort()
// 		}
// 		ctx.Next(c)
// 	}
// }