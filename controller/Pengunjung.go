package controller

import (
	"bytes"
	"github.com/labstack/echo"
	"github.com/rasyidmm/RumahMakan/models"
	"io/ioutil"
	"net/http"
)

func CreatePengunjung(c echo.Context)error  {
	var bodyBytes []byte
	if c.Request().Body != nil {
		bodyBytes, _ = ioutil.ReadAll(c.Request().Body)
	}

	// Restore the io.ReadCloser to its original state
	c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	// Continue to use the Body, like Binding it to a struct:
	Pengunjung := new(models.Pengunjung)
	if err:= c.Bind(Pengunjung); err !=nil{
		return echo.NewHTTPError(http.StatusBadRequest,err)
	}
	res,err := models.CreatePengunjung(Pengunjung)
	if err != nil{
		return echo.NewHTTPError(http.StatusBadRequest,err)
	}
	return c.JSON(http.StatusCreated,res)
}

func GetAllpengunjung(e echo.Context) error {
	var(
		res models.Response
	)
	result,err := models.GetPengunjungAll()
	if err != nil{
		return echo.NewHTTPError(http.StatusBadRequest,err)
	}
	res.Data = result
	res.Status = true
	res.Message= "Get Data Sukses"
	return e.JSON(http.StatusOK,res)
}

func GetByIdPengujung(e echo.Context)error  {
	id := e.Param("id")
	var res models.Response
	result,err := models.GetByIdPengunjung(id)
	if err != nil{
		return echo.NewHTTPError(http.StatusBadRequest,err)
	}
	res.Message = "Get Data Sukses"
	res.Status = true
	res.Data =result
	return e.JSON(http.StatusOK,res)
}

func UpdatePengunjung(e echo.Context)error  {
	var(
		id = e.Param("id")
		res models.Response
	)

	// Continue to use the Body, like Binding it to a struct:
	order,err := models.GetByIdPengunjung(id)
	if err != nil{
		return echo.NewHTTPError(http.StatusBadRequest,err)
	}
	var bodyBytes []byte
	if e.Request().Body != nil {
		bodyBytes, _ = ioutil.ReadAll(e.Request().Body)
	}
	e.Request().Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	if err:= e.Bind(&order); err !=nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	err = order.UpdatePengunjung()
	if err != nil{
		return echo.NewHTTPError(http.StatusBadRequest,err)
	}
	res.Status=true
	res.Message="update Sukses"
	return  e.JSON(http.StatusOK,res)
}
func DeletePengunjung(e echo.Context)error  {
	var(
		id = e.Param("id")
		res models.Response
	)

	result,err := models.GetByIdPengunjung(id)
	if err != nil{
		return echo.NewHTTPError(http.StatusBadRequest,err)
	}
	err =result.DeletePengunjung()
	res.Status=true
	res.Message="Delete Sukses"
	return e.JSON(http.StatusOK,res)
}