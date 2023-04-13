package handlers

import (
	"online_banking_service/api/http"

	us "online_banking_service/genproto/user_service"

	"github.com/gin-gonic/gin"
)

// @ID login
// @Router /login [POST]
// @Summary Login
// @Tags Auth
// @Accept json
// @Produce json
// @Param username query string true "username"
// @Param password query string true "password"
// @Success 200 {object} http.Response{data=string} "Response body"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) Login(c *gin.Context) {

	resp, err := h.service.UserService().Login(
		c.Request.Context(),
		&us.LoginRequest{
			Username: c.Query("username"),
			Password: c.Query("password"),
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// @ID sign_in
// @Router /sign-in [POST]
// @Summary SignIn
// @Tags Auth
// @Accept json
// @Produce json
// @Param username query string true "username"
// @Param password query string true "password"
// @Success 200 {object} http.Response{data=string} "Response body"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) SignIn(c *gin.Context) {

	resp, err := h.service.UserService().SignIn(
		c.Request.Context(),
		&us.SignInRequest{
			Username: c.Query("username"),
			Password: c.Query("password"),
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}
