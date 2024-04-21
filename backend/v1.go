package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jacobmaizel/aura/backend/api"
)

func AddV1Routes(rg *gin.RouterGroup) {

	v1Routes := rg.Group("/v1")

	api.AddPatientRoutes(v1Routes)
	api.AddMedicationRoutes(v1Routes)
	api.AddTemperatureRoutes(v1Routes)
}
