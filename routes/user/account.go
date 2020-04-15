package user

import (
	"github.com/gin-gonic/gin"
	. "github.com/ilya-sokolov/crypto_kiddies-server/common"
	"github.com/ilya-sokolov/crypto_kiddies-server/model"
	"net/http"
)

func Account(acc *gin.RouterGroup) {
	// TODO: Red_byte Add account-owned routing
}

func CreateAccount(ctx *gin.Context) {
	type request struct {
		NickName string `json:"nickName" validate:"required,gt=1,lt=100"`
		Password string `json:"password" validate:"required,gt=5,lt=100"`
	}
	var r request
	if err := BindAndValidate(ctx, &r); err != nil {
		return
	}
	account, err := model.CreateAccount(r.NickName, r.Password)
	if err != nil {
		ResponseError(ctx, http.StatusInternalServerError, ErrorMessage{Message: err.Error()})
		return
	}
	ResponseSuccess(ctx, http.StatusCreated, AccountResponse{Account: *account, Token: account.Token()})
}

func Authorization(ctx *gin.Context) {
	type request struct {
		AccountId string `json:"accountId" validate:"required,gt=0"`
		Password  string `json:"password" validate:"required,gt=5,lt=100"`
	}
	var r request
	if err := BindAndValidate(ctx, &r); err != nil {
		return
	}
	account, err := model.GetAccount(r.AccountId, r.Password)
	if err != nil {
		ResponseError(ctx, http.StatusInternalServerError, ErrorMessage{Message: err.Error()})
		return
	}
	ResponseSuccess(ctx, http.StatusOK, AccountResponse{Account: *account, Token: account.Token()})
}

type AccountResponse struct {
	model.Account
	Token string `json:"token"`
}
