package boot

import (
	"sync"

	durationdata "github.com/IacopoMelani/Go-Starter-Project/pkg/models/duration_data"

	"github.com/IacopoMelani/musicgang/controllers"

	"github.com/IacopoMelani/musicgang/config"
	"github.com/IacopoMelani/musicgang/routes"

	"github.com/labstack/echo/middleware"

	"github.com/labstack/echo"
)

var e *echo.Echo

func initEchoRoutes(e *echo.Echo) {

	routes.InitGetRoutes(e)
	routes.InitPostRoutes(e)
	routes.InitStaticRoutes(e)
}

// InitServer - Si occupa di lanciare l'applicazione con tutte le dovute operazioni iniziali
func InitServer() {

	var wg sync.WaitGroup

	wg.Add(4)

	go func() {
		defer wg.Done()
		config.GetInstance()
	}()

	go func() {
		defer wg.Done()
		durationdata.InitDurationData()
	}()

	go func() {
		defer wg.Done()
		e = echo.New()
		e.Use(middleware.Recover())
		e.Use(middleware.Logger())
		initEchoRoutes(e)

	}()

	go func() {
		defer wg.Done()
		controllers.InitCustomHandler()
	}()

	wg.Wait()

	config := config.GetInstance()

	e.Logger.Fatal(e.Start(config.AppPort))
}
