package service

import (
	"fmt"
	"net/http"
	"playground/database"
	"playground/models"
	"playground/request"
	"playground/token"
	"playground/util"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	var (
		req  request.LoginRequest
		user models.User
		err  error
		jwt  string
		resp request.LoginResponse
	)
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusInternalServerError, util.GetError(err))
		return
	}
	if user, err = database.GetUserByUsername(req.Username); err != nil {
		ctx.JSON(http.StatusInternalServerError, util.GetError(err))
		return
	}
	if !util.CheckPasswordHash(req.Password, user.Password) {
		ctx.JSON(http.StatusInternalServerError, util.GetError(fmt.Errorf("password is incorrect")))
		return
	}
	if jwt, err = token.CreateToken(user.Username); err != nil {
		ctx.JSON(http.StatusInternalServerError, util.GetError(err))
		return
	}
	resp.Username = user.Username
	resp.Token = jwt
	ctx.JSON(http.StatusOK, resp)
}

func CreateUser(ctx *gin.Context) {
	var (
		req      request.CreateUserRequest
		user     models.User
		err      error
		password string
	)
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusInternalServerError, util.GetError(err))
		return
	}
	if password, err = util.HashPassword(req.Password); err != nil {
		ctx.JSON(http.StatusInternalServerError, util.GetError(err))
		return
	}
	req.Password = password
	if user, err = database.CreateUser(req); err != nil {
		ctx.JSON(http.StatusInternalServerError, util.GetError(err))
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func GetUser(ctx *gin.Context) {
	var (
		req  request.GetUserRequest
		user models.User
		err  error
	)
	if err = ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusInternalServerError, util.GetError(err))
		return
	}
	if user, err = database.GetUser(req); err != nil {
		ctx.JSON(http.StatusInternalServerError, util.GetError(err))
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func GetUserLists(ctx *gin.Context) {
	var (
		req   request.GetUserListsRequest
		users []models.User
		err   error
	)
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusInternalServerError, util.GetError(err))
		return
	}
	if users, err = database.GetUserLists(req); err != nil {
		ctx.JSON(http.StatusInternalServerError, util.GetError(err))
		return
	}
	ctx.JSON(http.StatusOK, users)
}
