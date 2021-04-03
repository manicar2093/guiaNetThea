package services

import (
	"fmt"
	"testing"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/manicar2093/guianetThea/app/entities"
	"github.com/stretchr/testify/assert"
)

func TestCreateLoginRegistryXLS(t *testing.T) {
	initDate := time.Date(2021, time.March, 10, 00, 00, 00, 00, time.UTC)
	finalDate := time.Date(2021, time.March, 19, 00, 00, 00, 00, time.UTC)
	dataMock := []entities.LoginRegistry{
		{Name: "name1", Email: "correo1", Rol: "rol1", Time: "time1", Host: "host1", TypeLogOut: "type1", Page: "page1"},
		{Name: "name2", Email: "correo2", Rol: "rol2", Time: "time2", Host: "host2", TypeLogOut: "type2", Page: "page2"},
		{Name: "name3", Email: "correo3", Rol: "rol3", Time: "time3", Host: "host3", TypeLogOut: "type3", Page: "page3"},
		{Name: "name4", Email: "correo4", Rol: "rol4", Time: "time4", Host: "host4", TypeLogOut: "type4", Page: "page4"},
		{Name: "name5", Email: "correo5", Rol: "rol5", Time: "time5", Host: "host5", TypeLogOut: "type5", Page: "page5"},
		{Name: "name6", Email: "correo6", Rol: "rol6", Time: "time6", Host: "host6", TypeLogOut: "type6", Page: "page6"},
	}

	logRegistryDao.On("LogRegistrySearch", initDate, finalDate).Return(dataMock, nil)

	service := NewLoginRegistryService(logRegistryDao)

	xls, e := service.CreateLoginRegistryXLS(initDate, finalDate)

	logRegistryDao.AssertExpectations(t)

	assert.Nil(t, e, "No debió regresar error")
	assert.IsType(t, &excelize.File{}, xls, "No se recibió el tipo de dato requerido")
	assert.Equal(t, 1, xls.SheetCount, "No se crearon las hojas indicadas")

	e = xls.SaveAs(fmt.Sprintf("%s_%s.xlsx", "./test", time.Now().Format("20060102")))
	if e != nil {
		t.Log(e)
		t.Fatal("Error al guardar el archivo :S")
	}

}
