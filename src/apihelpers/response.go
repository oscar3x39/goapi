package apihelpers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RespondData struct {
	Code int `json:"code"`
	Status int `json:"status"`
	Data interface {} `json:"data"`
}

func JSON (w *gin.Context, status int, payload interface {}) {
	fmt.Println("status", status)

	var res RespondData
	res.Status = status
	res.Data = payload

	w.JSON(http.StatusOK, res)
}