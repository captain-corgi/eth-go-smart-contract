package main

import (
	"github.com/captain-corgi/eth-go-smart-contract/api"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"net/http"
)

func main() {
	client, err := ethclient.Dial("http://127.0.0.1:7545")
	if err != nil {
		panic(err)
	}
	conn, err := api.NewApi(common.HexToAddress("0x1DD1b49e045786B80020268419e1f985dB0A4a00"), client)
	if err != nil {
		panic(err)
	}
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/greet/:message", func(c echo.Context) error {
		message := c.Param("message")
		reply, err := conn.Greet(&bind.CallOpts{}, message)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, reply)
	})
	e.GET("/hello", func(c echo.Context) error {
		reply, err := conn.Hello(&bind.CallOpts{})
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, reply) // Hello World
	})

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
