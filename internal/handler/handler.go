package handler

import (
	"net/http"
	"strconv"
	_ "test-task/docs"
	"test-task/internal/domain"
	"test-task/internal/service"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
)

type Handler struct {
	services  *service.Service
	zapLogger *zap.Logger
}

func NewHandler(services *service.Service, zapLogger *zap.Logger) *Handler {
	return &Handler{
		services:  services,
		zapLogger: zapLogger,
	}
}

func (h *Handler) Routes() (r *gin.Engine) {

	r = gin.Default()
	r.Use(gin.Recovery())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/cars", h.getCars)
	r.POST("/cars", h.createCar)
	r.PUT("/cars/:id", h.updateCar)
	r.DELETE("/cars/:id", h.deleteCar)

	return
}

// @Summary Get cars
// @Description Get all cars
// @Accept  json
// @Produce  json
// @Success 200 {array} domain.Car
// @Failure 400 {json} error
// @Router /cars [get]
func (h *Handler) getCars(c *gin.Context) {

	result, err := h.services.Service.GetData()
	if err != nil {
		h.zapLogger.Error("error in bind JSON", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

// @Summary Create a new car
// @Description Create a new car
// @Accept  json
// @Produce  json
// @Param car body domain.Car true "Car object"
// @Success 201 {json} CarCreatedResponse "car created successfully"
// @Failure 400 {json} error
// @Failure 500 {json} error
// @Router /cars [post]
func (h *Handler) createCar(c *gin.Context) {

	var car domain.Car
	if err := c.BindJSON(&car); err != nil {
		h.zapLogger.Error("error in bind JSON", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.services.Service.SendRequestToExternalAPI(car.RegNums); err != nil {
		h.zapLogger.Error("error in bind JSON", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch data from external API"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "car created successfully"})
}

// @Summary Update a car
// @Description Update a car by ID
// @Accept  json
// @Produce  json
// @Param id path int true "Car ID"
// @Param car body domain.Car true "Car object"
// @Success 200 {json} CarUpdatedResponse "car updated successfully"
// @Failure 400 {json} error
// @Failure 500 {json} error
// @Router /cars/{id} [put]
func (h *Handler) updateCar(c *gin.Context) {

	var car domain.Car
	if err := c.BindJSON(&car); err != nil {
		h.zapLogger.Error("error in bind JSON", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "car updated successfully"})
}

// @Summary Delete a car
// @Description Delete a car by ID
// @Accept  json
// @Produce  json
// @Param id path int true "Car ID"
// @Success 200 {json} DeletedSuccessfully "deleted successfully"
// @Failure 400 {json} error
// @Failure 500 {json} error
// @Router /cars/{id} [delete]
func (h *Handler) deleteCar(c *gin.Context) {

	idStr := c.Param("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)

	if err := h.services.Service.DeleteData(id); err != nil {
		h.zapLogger.Error("error in bind JSON", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "deleted successfully"})
}
