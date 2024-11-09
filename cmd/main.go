package main

import (
	"fmt"
	handler "golang_template/internal/api"

	"log"
	"net/http"
	"os"

	_ "github.com/a-h/templ"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
		AllowCredentials: true,
	}))

	handler.Router(e)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))))

}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
}
