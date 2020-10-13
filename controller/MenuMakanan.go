package controller

import (
	"bytes"
	"github.com/labstack/echo"
	"github.com/rasyidmm/RumahMakan/models"
	"io/ioutil"
	"net/http"
)

func CreateMenuMakanan(c echo.Context)error  {
	var bodyBytes []byte
	if c.Request().Body != nil {
		bodyBytes, _ = ioutil.ReadAll(c.Request().Body)
	}

	// Restore the io.ReadCloser to its original state
	c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	// Continue to use the Body, like Binding it to a struct:
	order := new(models.MenuMakanan)
	if err:= c.Bind(order); err !=nil{
		return echo.NewHTTPError(http.StatusBadRequest,err)
	}
	res,err := models.Create(order)
	if err != nil{
		return echo.NewHTTPError(http.StatusBadRequest,err)
	}
	return c.JSON(http.StatusCreated,res)
}
func GetMenuMakananAll(c echo.Context) error {
	var res models.Response
	result,err := models.GetMenuMakananAll()
	if err != nil{
		return echo.NewHTTPError(http.StatusBadRequest,err)
	}
	res.Data = result
	res.Status = true
	res.Message= "Get Data Sukses"
	return c.JSON(http.StatusOK,res)
}

func GetByIdMenuMakanan(c echo.Context)error  {
	id := c.Param("id")
	var res models.Response
	result,err := models.GetByIdMenuMakanan(id)
	if err != nil{
		return echo.NewHTTPError(http.StatusBadRequest,err)
	}
	res.Message = "Get Data Sukses"
	res.Status = true
	res.Data =result
	return c.JSON(http.StatusOK,res)
}
func UpdateMenuMakanan(c echo.Context)error  {
	var(
		id = c.Param("id")
		res models.Response
	)

	// Continue to use the Body, like Binding it to a struct:
	order,err := models.GetByIdMenuMakanan(id)
	if err!= nil{
		return echo.NewHTTPError(http.StatusBadRequest,err)
	}
	var bodyBytes []byte
	if c.Request().Body != nil {
		bodyBytes, _ = ioutil.ReadAll(c.Request().Body)
	}
	c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	if err:= c.Bind(&order); err !=nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	err = order.UpdateMenuMakanan()
	if err != nil{
		return echo.NewHTTPError(http.StatusBadRequest,err)
	}
	res.Status=true
	res.Message="update Sukses"
	return  c.JSON(http.StatusOK,res)
}

func DeleteMenuMakanan(c echo.Context)error  {
	var(
		id = c.Param("id")
		res models.Response
	)

	result,err := models.GetByIdMenuMakanan(id)
	if err != nil{
		return echo.NewHTTPError(http.StatusBadRequest,err)
	}
	err =result.DeleteMenuMakanan()
	res.Status=true
	res.Message="Delete Sukses"
	return  c.JSON(http.StatusOK,res)
}