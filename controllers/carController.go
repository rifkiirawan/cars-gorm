package controllers

import (
	"cars-gorm/models"
	"fmt"
	"net/http"

	"cars-gorm/database"

	"github.com/gin-gonic/gin"
)

var CarDatas = []models.Car{}

// CreateCar godoc
// @Summary Post details for a given id
// @Description Post details of car corresponding to the input Id
// @Tags cars
// @Accept json
// @Produce json
// @Param models.Car body models.Car true "create car"
// @Success 200 {object} models.Car
// @Router /cars [post]
func CreateCar(ctx *gin.Context) {
	var db = database.GetDB()

	var input models.Car
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	carinput := models.Car{Pemilik: input.Pemilik, Merk: input.Merk, Harga: input.Harga, Typecars: input.Typecars}
	db.Create(&carinput)

	ctx.JSON(http.StatusOK, gin.H{"data": carinput})
}

// UpdateCar godoc
// @Summary Update car identified by the given Id
// @Description Update the car corresponding to the input Id
// @Tags cars
// @Accept json
// @Produce json
// @Param id path int true "ID of the car to be updated"
// @Success 200 {object} models.Car
// @Router /cars/{id} [patch]
func UpdateCar(ctx *gin.Context) {
	var db = database.GetDB()

	var car models.Car

	// err := db.First(&car, "Id = ?", ctx.Param("id")).Error
	var input models.Car

	carinput := models.Car{Pemilik: input.Pemilik, Merk: input.Merk, Harga: input.Harga, Typecars: input.Typecars}
	err := db.Model(&car).Where("id = ?", ctx.Param("id")).Updates(carinput).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}

// GetAllCar godoc
// @Summary Get details
// @Description Get details of all car
// @Tags cars
// @Accept json
// @Produce json
// @Success 200 {object} models.Car
// @Router /cars/allcars [get]
func GetAllCar(ctx *gin.Context) {
	var db = database.GetDB()

	var cars []models.Car
	err := db.Find(&cars).Error

	if err != nil {
		fmt.Println("error getting car datas :", err.Error())
	}

	ctx.JSON(http.StatusOK, gin.H{"data": cars})
}

// GetCar godoc
// @Summary Get details for a given id
// @Description Get details of car corresponding to the input id
// @Tags cars
// @Accept json
// @Produce json
// @Param Id path int true "ID of the car"
// @Success 200 {object} models.Car
// @Router /cars/{Id} [get]
func GetCar(ctx *gin.Context) {
	var db = database.GetDB()

	var carOne models.Car

	err := db.First(&carOne, "Id = ?", ctx.Param("id")).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data One": carOne})
}

// DeleteCar godoc
// @Summary Delete car identified by the given id
// @Description Delete the order of corresponding to the input id
// @Tags cars
// @Accept json
// @Produce json
// @Param id path int true "ID of the car to be deleted"
// @Success 204 "No Content"
// @Router /cars/{id} [delete]
func DeleteCar(ctx *gin.Context) {
	var db = database.GetDB()

	var carDelete models.Car

	err := db.First(&carDelete, "Id = ?", ctx.Param("id")).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&carDelete)

	ctx.JSON(http.StatusOK, gin.H{"data": true})
}
