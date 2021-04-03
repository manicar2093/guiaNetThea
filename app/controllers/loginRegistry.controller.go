package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/manicar2093/guianetThea/app/models"
	"github.com/manicar2093/guianetThea/app/services"
	"github.com/manicar2093/guianetThea/app/utils"
)

type LoginRegistryController interface {
	LoginRegistryInform(w http.ResponseWriter, r *http.Request)
}

type LoginRegistryControllerImpl struct {
	loginRegistryService services.LoginRegistryService
	validatorService     services.ValidatorService
}

func NewLoginRegistryController(loginRegistryService services.LoginRegistryService, validatorService services.ValidatorService) LoginRegistryController {
	return &LoginRegistryControllerImpl{loginRegistryService, validatorService}
}

func (l LoginRegistryControllerImpl) LoginRegistryInform(w http.ResponseWriter, r *http.Request) {
	var data models.LoginRegistryData
	if e := json.NewDecoder(r.Body).Decode(&data); e != nil {
		utils.Error.Printf("Error al decodificar los datos del body. Detalles: \n\t%v", e)
		utils.JSON(map[string]interface{}{"message": "No se logró obtener la información necesaria. Valida la documentación"}, http.StatusInternalServerError, w)
		return
	}

	err, valid := l.validatorService.Validate(data)
	if !valid {
		utils.JSON(map[string]interface{}{
			"message": "La información es incorrecta. Favor de revisar la documentación",
			"errores": err,
		}, http.StatusBadRequest, w)
		return
	}

	file, e := l.loginRegistryService.CreateLoginRegistryXLS(data.InitDate, data.FinalDate)
	if e != nil {
		utils.Error.Printf("Error crear el documento con los datos de login. Detalles: \n\t%v", e)
		utils.JSON(map[string]interface{}{"message": "Hubo un error inesperado al crear el documento. Favor de contactar a soporte"}, http.StatusInternalServerError, w)
		return
	}
	dateFormat := "20060102"
	fileName := fmt.Sprintf("reporte_%s_%s.xlsx", data.InitDate.Format(dateFormat), data.FinalDate.Format(dateFormat))

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", fmt.Sprintf(`attachment;filename="%s"`, fileName))
	w.Header().Set("File-Name", fileName)
	w.Header().Set("Content-Transfer-Encoding", "binary")
	w.Header().Set("Expires", "0")

	e = file.Write(w)
	if e != nil {
		utils.Error.Printf("Error al escribir archivo en respuesta. Detalles: \n\t%v", e)
		utils.JSON(map[string]interface{}{"message": "Hubo un error inesperado al crear el documento. Favor de contactar a soporte"}, http.StatusInternalServerError, w)
		return
	}
}
