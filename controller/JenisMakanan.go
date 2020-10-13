package controller

import (
	"bytes"
	"github.com/labstack/echo"
	"github.com/rasyidmm/RumahMakan/models"
	"io/ioutil"
	"net/http"
)

func CreateJenisMakanan(c echo.Context)error  {
	var bodyBytes []byte
	if c.Request().Body != nil {
		bodyBytes, _ = ioutil.ReadAll(c.Request().Body)
	}

	// Restore the io.ReadCloser to its original state
	c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	// Continue to use the Body, like Binding it to a struct:
	order := new(models.JenisMakanan)
	if err:= c.Bind(order); err !=nil{
		return echo.NewHTTPError(http.StatusBadRequest,err)
	}
	res,err := models.CreateJenisMakanan(order)
	if err != nil{
		return echo.NewHTTPError(http.StatusBadRequest,err)
	}
	return c.JSON(http.StatusCreated,res)
}
func FindJenisMakananAl(c echo.Context) error {
	result, err := models.FindJenisMakananAl()
	if err != nil{
		return echo.NewHTTPError(http.StatusInternalServerError, "get Data Error")
	}
	return c.JSON(http.StatusOK,result)

}
func UpdateJenisMakanana(c echo.Context)(error)  {
	id := c.Param("id")
	defer c.Request().Body.Close()

	jenismakanan,err := models.FindJenisMakananByIdNative(id)
	if err!= nil{
		return echo.NewHTTPError(http.StatusBadRequest,err)
	}
	var bodyBytes []byte
	if c.Request().Body != nil {
		bodyBytes, _ = ioutil.ReadAll(c.Request().Body)
	}
	c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	if err:= c.Bind(&jenismakanan); err !=nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	jenismakanan.UpdateJenisMakanan()
	return c.JSON(http.StatusOK,jenismakanan)
}

func DeleteJenisMakanan(c echo.Context) error {
	id := c.Param("id")
	jenismakanan,err := models.FindJenisMakananByIdNative(id)
	if err!= nil{
		return echo.NewHTTPError(http.StatusBadRequest,err)
	}
	err = jenismakanan.DeleteJenisMakanan()
	if err!= nil{
		return echo.NewHTTPError(http.StatusBadRequest,err)
	}
	return c.JSON(http.StatusOK,jenismakanan)
}

func FindJenisMakananById(c echo.Context)error  {
	id:=c.Param("id")
	jenismakanan,err := models.FindJenisMakananByIdNative(id)
	if err!= nil{
		return echo.NewHTTPError(http.StatusBadRequest,err)
	}
	return c.JSON(http.StatusOK,jenismakanan)
}