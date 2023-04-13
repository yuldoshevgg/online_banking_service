package handlers

import (
	"online_banking_service/api/http"
	config "online_banking_service/config"
	"online_banking_service/grpc/client"
	"online_banking_service/pkg/logger"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	cfg     config.Config
	log     logger.LoggerI
	service client.ServiceManagerI
}

func NewHandler(cfg config.Config, log logger.LoggerI, server client.ServiceManagerI) Handler {
	return Handler{
		cfg:     cfg,
		log:     log,
		service: server,
	}
}

func (h *Handler) handleResponse(c *gin.Context, status http.Status, data interface{}) {
	switch code := status.Code; {
	case code < 300:
		h.log.Info(
			"---Response--->",
			logger.Int("code", status.Code),
			logger.String("status", status.Status),
			logger.Any("description", status.Description),
		)
	case code < 400:
		h.log.Warn(
			"---Response--->",
			logger.Int("code", status.Code),
			logger.String("status", status.Status),
			logger.Any("description", status.Description),
			logger.Any("data", data),
		)
	default:
		h.log.Error(
			"---Response--->",
			logger.Int("code", status.Code),
			logger.String("status", status.Status),
			logger.Any("description", status.Description),
			logger.Any("data", data),
		)
	}

	c.JSON(status.Code, http.Response{
		Status:      status.Status,
		Description: status.Description,
		Data:        data,
	})
}
