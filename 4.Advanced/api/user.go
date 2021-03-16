package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	l "github.com/gtldhawalgandhi/go-training/3.Intermediate/logger"
	"github.com/gtldhawalgandhi/go-training/4.Advanced/db"
	errs "github.com/pkg/errors"
)

var (
	errUserExists   = errors.New("user already exists")
	errGettingUser  = errors.New("failed to get users")
	errCreatingUser = errors.New("failed to create user")
	errUpdatingUser = errors.New("failed to update user")
)

func (server *Server) getUsers(ctx *gin.Context) {
	users, err := server.store.GetUsers(ctx)

	if err != nil {
		l.D(err)
		ctx.JSON(http.StatusInternalServerError, errorResponse(errGettingUser))
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func (server *Server) createUser(ctx *gin.Context) {
	var req db.CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	user, err := server.store.CreateUser(ctx, req)
	if err != nil {
		code, ok := db.PGErrorCode(errs.Cause(err))
		l.E("PG error code ", code)

		if !ok {
			ctx.JSON(http.StatusInternalServerError, errorResponse(errCreatingUser))
			return
		}

		ctx.JSON(http.StatusBadRequest, errorResponse(errUserExists))
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (server *Server) updateUser(ctx *gin.Context) {
	var req db.UpdateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	user, err := server.store.UpdateUser(ctx, req)
	if err != nil {

		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}
