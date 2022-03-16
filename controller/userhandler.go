package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pawnjiang/go-web/response"
	"github.com/pawnjiang/go-web/util"
)

func Register(ctx *gin.Context) {
	name := ctx.PostForm("name")
	password := ctx.PostForm("password")
	phone := ctx.PostForm("phone")
	if len(phone) != 11 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号码必须是11位")
		return
	}
	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码必须大于6位")
		return
	}
	if len(name) == 0 {
		name = util.RandomString(10)
	}
	if dao.IsPhoneExist(phone) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户已存在")
		return
	}
}
