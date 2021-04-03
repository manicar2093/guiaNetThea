package services

import (
	"fmt"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/manicar2093/guianetThea/app/dao"
)

var (
	letterA     = rune(65)
	sheetName   = "Reporte"
	fileHeaders = []string{
		"Usuario",
		"Correo Electrónico",
		"Rol",
		"Tiempo de sesión",
		"Logueado desde",
		"Página visitada",
		"Tipo de Logout",
		"Fecha incio sesión",
		"Fecha termino sesión",
	}
)

type LoginRegistryService interface {
	// CreateLoginRegistryXLS crea el archivo xls con los datos que se encuentren entre las fechas indicadas.
	// La fecha inicial no debe ser mayoy a la final.
	// El nombre que se le da el archivo es el intervalo de tiempo que se solicitó con el prefix 'login_registry'
	CreateLoginRegistryXLS(init, final time.Time) (*excelize.File, error)
}

type LoginRegistryServiceImpl struct {
	loginRegistryDao dao.LoginRegistryDao
}

func NewLoginRegistryService(loginRegistryDao dao.LoginRegistryDao) LoginRegistryService {
	return &LoginRegistryServiceImpl{loginRegistryDao}
}

func (l LoginRegistryServiceImpl) CreateLoginRegistryXLS(init, final time.Time) (res *excelize.File, e error) {
	data, e := l.loginRegistryDao.LogRegistrySearch(init, final)
	if e != nil {
		return
	}
	res = l.createNewFile()
	l.setFileHeaders(res, letterA)
	letter := letterA
	counter := 2
	for _, v := range data {
		res.SetCellValue(sheetName, fmt.Sprintf("%s%d", string(letter), counter), v.Name)
		letter++
		res.SetCellValue(sheetName, fmt.Sprintf("%s%d", string(letter), counter), v.Email)
		letter++
		res.SetCellValue(sheetName, fmt.Sprintf("%s%d", string(letter), counter), v.Rol)
		letter++
		res.SetCellValue(sheetName, fmt.Sprintf("%s%d", string(letter), counter), v.Time)
		letter++
		res.SetCellValue(sheetName, fmt.Sprintf("%s%d", string(letter), counter), v.Host)
		letter++
		res.SetCellValue(sheetName, fmt.Sprintf("%s%d", string(letter), counter), v.Page)
		letter++
		res.SetCellValue(sheetName, fmt.Sprintf("%s%d", string(letter), counter), v.TypeLogOut)
		letter++
		res.SetCellValue(sheetName, fmt.Sprintf("%s%d", string(letter), counter), init)
		letter++
		res.SetCellValue(sheetName, fmt.Sprintf("%s%d", string(letter), counter), final)
		counter++
		letter = letterA
	}
	return
}

func (l LoginRegistryServiceImpl) setFileHeaders(file *excelize.File, column rune) {
	row := 1
	for _, v := range fileHeaders {
		file.SetCellValue(sheetName, fmt.Sprintf("%s%d", string(column), row), v)
		column++
	}
}

func (l LoginRegistryServiceImpl) createNewFile() *excelize.File {
	file := excelize.NewFile()
	file.NewSheet(sheetName)
	file.DeleteSheet("Sheet1")
	return file
}
