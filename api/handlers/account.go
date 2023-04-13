package handlers

import (
	"strconv"

	"online_banking_service/api/http"
	as "online_banking_service/genproto/account_service"

	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// @ID create_account
// @Router /create-account [POST]
// @Summary Create online bank account
// @Tags Accounts
// @Accept json
// @Produce json
// @Param user_id query string true "user_id"
// @Param account_number query string true "account_number"
// @Success 200 {object} http.Response{data=string} "Response body"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateAccount(c *gin.Context) {

	resp, err := h.service.AccountService().CreateAccount(
		c.Request.Context(),
		&as.CreateAccountRequest{
			UserId:        c.Query("user_id"),
			AccountNumber: c.Query("account_number"),
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// @Security ApiKeyAuth
// @ID pay_for_account
// @Router /pay-for-account [PUT]
// @Summary Do payment for an account
// @Tags Accounts
// @Accept json
// @Produce json
// @Param account_number query string true "account_number"
// @Param balance query string true "balance"
// @Success 200 {object} http.Response{data=string} "Response body"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) PayForAccount(c *gin.Context) {

	balance, err := strconv.Atoi(c.Query("balance"))
	if err != nil {
		h.handleResponse(c, http.InvalidArgument, err.Error())
		return
	}

	resp, err := h.service.AccountService().PayForAccount(
		c.Request.Context(),
		&as.PayForAccountRequest{
			AccountNumber: c.Query("account_number"),
			Balance:       float64(balance),
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// @Security ApiKeyAuth
// @ID withdraw_from_account
// @Router /withdraw-from-account [PUT]
// @Summary Withdraw money from a account
// @Tags Accounts
// @Accept json
// @Produce json
// @Param account_number query string true "account_number"
// @Param balance query string true "balance"
// @Success 200 {object} http.Response{data=string} "Response body"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) WithdrawFromAccount(c *gin.Context) {

	balance, err := strconv.Atoi(c.Query("balance"))
	if err != nil {
		h.handleResponse(c, http.InvalidArgument, err.Error())
		return
	}

	resp, err := h.service.AccountService().WithdrawFromAccount(
		c.Request.Context(),
		&as.WithdrawFromAccountRequest{
			AccountNumber: c.Query("account_number"),
			Balance:       float64(balance),
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// @Security ApiKeyAuth
// @ID get_accounts
// @Router /accounts [GET]
// @Summary Get all user accounts by your username
// @Tags Accounts
// @Accept json
// @Produce json
// @Param username query string true "username"
// @Success 200 {object} http.Response{data=string} "Response body"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetAccounts(c *gin.Context) {

	resp, err := h.service.AccountService().GetAccounts(
		c.Request.Context(),
		&as.GetAccountsRequest{
			Username: c.Query("username"),
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// @Security ApiKeyAuth
// @ID transfer_balance
// @Router /transfer-balance [POST]
// @Summary Transfer balance between accounts
// @Tags Accounts
// @Accept json
// @Produce json
// @Param sender query string true "sender"
// @Param recipient query string true "recipient"
// @Param balance query string true "balance"
// @Success 200 {object} http.Response{data=string} "Response body"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) TransferBalance(c *gin.Context) {

	balance, err := strconv.Atoi(c.Query("balance"))
	if err != nil {
		h.handleResponse(c, http.InvalidArgument, err.Error())
		return
	}

	resp, err := h.service.AccountService().TransferBalance(
		c.Request.Context(),
		&as.TransferBalanceRequest{
			Sender:    c.Query("sender"),
			Recipient: c.Query("recipient"),
			Balance:   float64(balance),
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}
