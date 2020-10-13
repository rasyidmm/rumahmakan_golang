package main

import (

	"github.com/rasyidmm/RumahMakan/db"
	"github.com/rasyidmm/RumahMakan/routes"
)

func main()  {
	db.Init()
	e := routes.New()
	e.Logger.Fatal(e.Start(":1234"))
}
