package handler

import (
	"bankrpl/common"
	"bankrpl/entity"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RekeningHandler struct {
	DB *gorm.DB
}

// post rekening
func (h *RekeningHandler) HandlerRegisterRekening(ctx *gin.Context) {
	var Rekening entity.Rekening
	//take input from json
	err := ctx.ShouldBind(&Rekening)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			common.Response{
				Status:  false,
				Message: err.Error(),
				Data:    nil,
			})
		return
	}

	//create new rekening and push in database
	tx := h.DB.Create(&Rekening)
	if tx.Error != nil {
		ctx.JSON(http.StatusBadRequest,
			common.Response{
				Status:  false,
				Message: tx.Error.Error(),
				Data:    nil,
			})
		return
	}
	ctx.JSON(http.StatusOK,
		common.Response{
			Status:  true,
			Message: "Rekening created successfully",
			Data:    Rekening,
		})
}

// delet rekening
func (h *RekeningHandler) HandlerDeleteRekening(ctx *gin.Context) {
	var rek entity.Rekening

	//take input from json
	err := ctx.ShouldBind(&rek)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			common.Response{
				Status:  false,
				Message: err.Error(),
				Data:    nil,
			})
		return
	}

	//find id in db
	var updNasabah entity.Rekening
	er := h.DB.First(&updNasabah, rek.ID)
	if er.Error != nil {
		ctx.JSON(http.StatusBadRequest,
			common.Response{
				Status:  false,
				Message: "Can't find ID",
				Data:    nil,
			})
		return
	}

	//delete rekening from db by id
	tx := h.DB.Delete(&entity.Rekening{}, rek.ID)
	if tx.Error != nil {
		ctx.JSON(http.StatusBadRequest,
			common.Response{
				Status:  false,
				Message: tx.Error.Error(),
				Data:    nil,
			})
		fmt.Println("del db err")
		return
	}
	ctx.JSON(http.StatusOK,
		common.Response{
			Status:  true,
			Message: "Rekening deleted successfully",
			Data:    nil,
		})
}
