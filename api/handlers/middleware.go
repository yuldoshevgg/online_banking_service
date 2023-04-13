package handlers

import (
	"strings"

	"online_banking_service/api/http"
	config "online_banking_service/config"
	"online_banking_service/pkg/security"

	"github.com/gin-gonic/gin"
)

func (h *Handler) AuthMiddleware(c *gin.Context) {
	if ok := h.hasAccess(c); !ok {
		c.Abort()
		return
	}

	c.Next()
}

func (h *Handler) hasAccess(c *gin.Context) bool {

	bearerToken := c.GetHeader("Authorization")
	strArr := strings.Split(bearerToken, " ")
	if len(strArr) != 2 || strArr[0] != "Bearer" {
		h.handleResponse(c, http.Forbidden, "token error: wrong format")
		return false
	}

	token := strArr[1]

	claims, err := security.ExtractClaims(token, h.cfg.SecretKey)
	if err != nil {
		h.handleResponse(c, http.Forbidden, "Invalid token")
		return false
	}

	roleID := claims["role_id"].(string)              // Example auth parts
	clientTypeID := claims["client_type_id"].(string) // Example auth paths

	return roleID == config.RoleID && clientTypeID == config.ClientTypeID
}
