package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/manicar2093/guianetThea/app/models"
	"github.com/stretchr/testify/assert"
)

func TestLoginRegistryInform(t *testing.T) {

	setUp()
	requestData := models.LoginRegistryData{
		InitDate:  time.Date(2021, time.March, 10, 00, 00, 00, 00, time.UTC),
		FinalDate: time.Date(2021, time.March, 19, 00, 00, 00, 00, time.UTC),
	}

	file := excelize.NewFile()
	w, r := httptest.NewRecorder(), httptest.NewRequest(http.MethodPost, "/admin/login_registry/create", serialize(t, &requestData))

	loginRegistryService.On("CreateLoginRegistryXLS", requestData.InitDate, requestData.FinalDate).Return(file, nil)
	validatorService.On("Validate", requestData).Return([]models.ErrorValidationDetail{}, true)

	controller := NewLoginRegistryController(loginRegistryService, validatorService)

	router.HandleFunc("/admin/login_registry/create", controller.LoginRegistryInform)

	router.ServeHTTP(w, r)

	loginRegistryService.AssertExpectations(t)
	validatorService.AssertExpectations(t)

	assert.Equal(t, "application/octet-stream", w.Header().Get("Content-Type"), "El content-type no es correcto")
	assert.Equal(t, `attachment;filename="reporte_20210310_20210319.xlsx"`, w.Header().Get("Content-Disposition"))
	assert.Equal(t, "reporte_20210310_20210319.xlsx", w.Header().Get("File-Name"), "El filename no coincide")
	assert.Equal(t, "binary", w.Header().Get("Content-Transfer-Encoding"), "El encoding de transferencia no es correcto")
	assert.Equal(t, "0", w.Header().Get("Expires"), "El encoding de transferencia no es correcto")

}
