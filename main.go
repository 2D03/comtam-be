package main

import (
	"crypto/tls"
	"github.com/2D03/comtam-be/api"
	"github.com/2D03/comtam-be/conf"
	"github.com/2D03/comtam-be/model"
	"github.com/globalsign/mgo"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	// Echo instance
	e := echo.New()

	setupDB()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//CORs
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	// Routes
	e.GET("/api-info", api.GetAPIInfo)
	e.GET("/menu", api.GetMenu)
	e.GET("/dish", api.GetDish)
	e.PUT("/dish", api.UpdateDish)
	e.DELETE("/dish", api.DeleteDish)
	e.POST("/dish", api.CreateDish)

	// Start server
	e.Logger.Fatal(e.Start(":8000"))
}

func setupDB() {
	println("Connecting db")

	envConfig, err := conf.GetConfigDBMap()
	if err != nil {
		panic(err)
	}

	//db main
	mainDB := &mgo.DialInfo{
		Addrs:   strings.Split(envConfig["dbHost"], ","),
		Timeout: 60 * time.Second,
		//Database: AuthDatabase,
		Username: envConfig["dbUser"],
		Password: envConfig["dbPassword"],
	}

	env := os.Getenv("env")
	if env != "local" {
		tlsConfig := &tls.Config{}
		mainDB.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
			conn, err := tls.Dial("tcp", addr.String(), tlsConfig) // add TLS config
			return conn, err
		}
	}
	mainDBSession, err := mgo.DialWithInfo(mainDB)
	if err != nil {
		panic(err)
	}

	onDBConnected(mainDBSession)
	println("Connected db")
}

func onDBConnected(s *mgo.Session) {
	model.InitDishModel(s, conf.Config.MainDBName)
}
