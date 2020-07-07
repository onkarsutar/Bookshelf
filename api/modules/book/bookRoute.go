package book

import (
	validationhelper "GoStudy/WebFrameworks/Echo/api/helper/validation"
	"GoStudy/WebFrameworks/Echo/api/model"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

// Init book init
func Init(o, r *echo.Group) {
	o.POST("/book/create", createBookRoute)
	o.GET("/book/read/:id", readBookRoute)
	o.POST("/book/update", updateBookRoute)
	o.DELETE("/book/delete", deleteBookRoute)
}

func createBookRoute(c echo.Context) error {
	bookObj := model.Book{}
	err := c.Bind(&bookObj)
	if err != nil {
		log.Error("createBookRoute Bind Error : ", err)
		return c.JSON(http.StatusExpectationFailed, "BIND_ERROR")
	}

	errors := validationhelper.Validate(bookObj)
	if errors != nil {
		return c.JSON(http.StatusExpectationFailed, "REQUIRED_FIELD_VALIDATION_FAILED")
	}
	id,err := createBookService(bookObj)
	if err != nil {
		return c.JSON(http.StatusExpectationFailed,err.Error())
	}
	return  c.JSON(http.StatusOK, "Book Created : "+ id)
}

func readBookRoute(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusExpectationFailed, "REQUIRED_FIELD_VALIDATION_FAILED")
	}
	bookObj, err := readBookService(id)
	if err != nil {
		return c.JSON(http.StatusExpectationFailed, err.Error())
	}

	return  c.JSON(http.StatusOK, bookObj)
}
func updateBookRoute(c echo.Context) error {
	bookObj := model.Book{}

	err := c.Bind(&bookObj)
	if err != nil {
		log.Error("createBookRoute Bind Error : ", err)
		return c.JSON(http.StatusExpectationFailed, "BIND_ERROR")
	}
	errors := validationhelper.Validate(bookObj)
	if errors != nil {
		return c.JSON(http.StatusExpectationFailed, "REQUIRED_FIELD_VALIDATION_FAILED")
	}

	id,err := updateBookService(bookObj)
	if err != nil {
		return c.JSON(http.StatusExpectationFailed, err.Error())
	}
	return  c.JSON(http.StatusOK, "Book Updated : "+id)
}

func deleteBookRoute(c echo.Context) error {
	bookObj := model.Book{}

	err := c.Bind(&bookObj)
	if err != nil {
		log.Error("createBookRoute Bind Error : ", err)
		return c.JSON(http.StatusExpectationFailed, "BIND_ERROR")
	}
	if bookObj.ID== "" {
		return c.JSON(http.StatusExpectationFailed, "REQUIRED_FIELD_VALIDATION_FAILED")
	}
	err = deleteBookService(bookObj.ID)
	if err != nil {
		return c.JSON(http.StatusExpectationFailed, err.Error())
	}

	return  c.JSON(http.StatusOK, "Book Deleted")
}
