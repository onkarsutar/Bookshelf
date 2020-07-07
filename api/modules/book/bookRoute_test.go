package book

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestInit(test *testing.T) {
	e := echo.New()
	o := e.Group("/o")
	r := e.Group("/r")

	Init(o, r)
}

func Test_createBookRoute(t *testing.T) {
	e := echo.New()
	bookJSON := `{ "title": "Yugandhar",    "author": "Shivaji Sawant",    "isbn": "9788184984897"}`
	req, _ := http.NewRequest(echo.POST, "/o/book/create", strings.NewReader(bookJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	if assert.NoError(t, createBookRoute(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func Test_readBookRoute(t *testing.T) {
	e := echo.New()
	req, _ := http.NewRequest(echo.GET, "/o/book/read/a14b1dc1-e08e-4a2f-9a8e-d55e5a3ae1ba", strings.NewReader(""))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	if assert.NoError(t, updateBookRoute(c)) {
		assert.Equal(t, http.StatusExpectationFailed, rec.Code)
	}
}


func Test_updateBookRoute(t *testing.T) {
	e := echo.New()
	bookJSON := `{ "title": "Yugandhar (Marathi)",    "author": "Shivaji Sawant",    "isbn": "9788184984897"}`

	req, _ := http.NewRequest(echo.POST, "/o/book/update", strings.NewReader(bookJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	if assert.NoError(t, updateBookRoute(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}


func Test_deleteBookRoute(t *testing.T) {
	e := echo.New()
	bookJSON := `{ "id": "a14b1dc1-e08e-4a2f-9a8e-d55e5a3ae1ba"}`
	req, _ := http.NewRequest(echo.DELETE, "/o/book/delete",strings.NewReader(bookJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	if assert.NoError(t, deleteBookRoute(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
