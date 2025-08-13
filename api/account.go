package api

import (
	"net/http"
	"strconv"

	db "github.com/devsirose/simplebank/db/sqlc"
	"github.com/devsirose/simplebank/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CreateAccountRequest struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,oneof=USD EUR VND "`
}

type GetAccountRequest struct {
	ID int64 `json:"owner" uri:"id" binding:"required,min=1"`
}

func (s *Server) CreateAccount(ctx *gin.Context) {
	var req CreateAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	acc, err := s.store.CreateAccount(ctx, db.CreateAccountParams{
		Owner:    req.Owner,
		Currency: req.Currency,
		Balance:  0,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusCreated, acc)
	return
}

func (s *Server) GetAccountBy(ctx *gin.Context) {
	var req GetAccountRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		logger.Log.Info("bind uri error or uri not found", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	acc, err := s.store.GetAccount(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, notFoundResponse(strconv.FormatInt(req.ID, 10)))
		return
	}
	ctx.JSON(http.StatusOK, acc)
}
