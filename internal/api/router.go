package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

func NewRouter(pool *pgxpool.Pool, log *zap.Logger) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())

	api := r.Group("/api/v1")
	{
		api.GET("/health", HealthHandler(pool))
		api.GET("/users", ListUsersHandler(pool))
		api.POST("/users", CreateUserHandler(pool))
	}

	return r
}
