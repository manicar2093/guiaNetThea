package utils

import (
	"io"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var templateUtil TemplateUtils

func setUp() {
	templateUtil = NewTemplateUtils()
}

func TestTemplateUtils(t *testing.T) {
	setUp()

	w := httptest.NewRecorder()

	e := templateUtil.RenderTemplateToResponseWriter("../../templates/on_dev.html", w, "nil")

	assert.Nil(t, e, "No debió presentar ningun error.")

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)

	if len(body) <= 0 {
		t.Fatal("No hay respuesta html")
	}
}

func TestTemplateUtilsTemplateNotExists(t *testing.T) {
	setUp()

	w := httptest.NewRecorder()

	e := templateUtil.RenderTemplateToResponseWriter("x", w, nil)

	assert.NotNil(t, e, "Se debio regresar error")
	assert.Equal(t, "template was not found", e.Error(), "El mensaje de error no es el correcto")

}

func TestTemplateUtilsUnexpectedError(t *testing.T) {
	t.Skip("This is not necesary by the moment")
	setUp()

	w := httptest.NewRecorder()

	e := templateUtil.RenderTemplateToResponseWriter("../../templates/on_dev.html", w, 12)

	assert.NotNil(t, e, "Se debio regresar error")
	assert.Equal(t, "error al ejecutar el template", e.Error(), "El mensaje de error no es el correcto")

}
