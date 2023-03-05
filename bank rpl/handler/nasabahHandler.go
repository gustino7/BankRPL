package handler

import (
	"bankrpl/common"
	"bankrpl/entity"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type NasabahHandler struct {
	DB *gorm.DB
}

// post nasabah
func (h *NasabahHandler) HandlerRegisterNasabah(ctx *gin.Context) {
	var Nasabah entity.Nasabah

	//take input from json
	err := ctx.ShouldBind(&Nasabah)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			common.Response{
				Status:  false,
				Message: err.Error(),
				Data:    nil,
			})
		return
	}
	//create new in database
	tx := h.DB.Create(&Nasabah)
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
			Message: "Nasabah created successfully",
			Data:    Nasabah,
		})
	//fmt.Println(Nasabah.Alamat)
}

// put nasabah
func (h *NasabahHandler) HandlerUpdateNasabah(ctx *gin.Context) {
	var nasabah entity.Nasabah

	//take input from json
	err := ctx.ShouldBind(&nasabah)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			common.Response{
				Status:  false,
				Message: err.Error(),
				Data:    nil,
			})
		return
	}

	//find an id in db
	var updNasabah entity.Nasabah
	tx := h.DB.First(&updNasabah, nasabah.ID)
	if tx.Error != nil {
		ctx.JSON(http.StatusBadRequest,
			common.Response{
				Status:  false,
				Message: "Can't find ID",
				Data:    nil,
			})
		return
	}

	//update database
	flag := false
	if nasabah.Alamat != "" {
		h.DB.Model(&updNasabah).Update("Alamat", nasabah.Alamat)
		flag = true
	}
	if nasabah.Pekerjaan != "" {
		h.DB.Model(&updNasabah).Update("Pekerjaan", nasabah.Pekerjaan)
		flag = true
	}
	if nasabah.WargaNgr != "" {
		h.DB.Model(&updNasabah).Update("WargaNgr", nasabah.WargaNgr)
		flag = true
	}

	if flag == true {
		ctx.JSON(http.StatusOK,
			common.Response{
				Status:  true,
				Message: "Nasabah updated successfully",
				Data:    nil,
			})
	} else {
		ctx.JSON(http.StatusBadRequest,
			common.Response{
				Status:  false,
				Message: "Nasabah didn't update",
				Data:    nil,
			})
	}
}

// get nasabah by id and show the rekening in it
func (h *NasabahHandler) HandlerGetNasabah(ctx *gin.Context) {
	//take an id from api
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			common.Response{
				Status:  false,
				Message: err.Error(),
				Data:    nil,
			})
		return
	}

	//take data from db nasabah by id and also the rekening
	//var nasabah Nasabah
	var rek entity.Rekening
	//tx := h.DB.Where("id = ?", id).Preload("Rekening_nsbh").Take(&nasabah)
	tx := h.DB.First(&rek, id)
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
			Message: "user fetched successfully",
			Data:    rek,
		})
}

// delete nasabah
func (h *NasabahHandler) HandlerDeleteNasabah(ctx *gin.Context) {
	var nasabah entity.Nasabah

	//take input from json
	err := ctx.ShouldBind(&nasabah)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			common.Response{
				Status:  false,
				Message: err.Error(),
				Data:    nil,
			})
		return
	}

	//delete nasabah in database
	tx := h.DB.Delete(&entity.Nasabah{}, nasabah.ID)
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
			Message: "Nasabah deleted successfully",
			Data:    nil,
		})
}
