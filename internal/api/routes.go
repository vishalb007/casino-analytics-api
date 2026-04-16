package api

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine, h *Handler) {
	api := r.Group("/", AuthMiddleware())

	api.GET("/gross_gaming_rev", h.GetGGR)
}