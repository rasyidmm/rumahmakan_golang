package routes

import (
	"github.com/labstack/echo"
	"github.com/rasyidmm/RumahMakan/controller"
	"net/http"
)

func New() *echo.Echo {
	e:=echo.New()
	e.GET("/", func(context echo.Context) error {
		return context.String(http.StatusOK,"heelo")
	})
	e.POST("/jenis-makanan",controller.CreateJenisMakanan)
	e.GET("/jenis-makanan",controller.FindJenisMakananAl)
	e.GET("/jenis-makanan/:id",controller.FindJenisMakananById)
	e.PUT("/jenis-makanan/:id",controller.UpdateJenisMakanana)
	e.DELETE("/jenis-makanan/:id",controller.DeleteJenisMakanan)

	e.POST("/menu-makanan",controller.CreateMenuMakanan)
	e.GET("/menu-makanan",controller.GetMenuMakananAll)
	e.GET("/menu-makanan/:id",controller.GetByIdMenuMakanan)
	e.PUT("/menu-makanan/:id",controller.UpdateMenuMakanan)
	e.DELETE("/menu-makanan/:id",controller.DeleteMenuMakanan)

	e.POST("/pengunjung",controller.CreatePengunjung)
	e.GET("/pengunjung",controller.GetAllpengunjung)
	e.GET("/pengunjung/:id",controller.GetByIdPengujung)
	e.PUT("/pengunjung/:id",controller.UpdatePengunjung)
	e.DELETE("/pengunjung/:id",controller.DeletePengunjung)

	return e
}