package main

import (
	_middleware "api-mini-project/app/middlewares"
	_routes "api-mini-project/app/routes"
	_categoryUsecase "api-mini-project/businesses/categories"
	_productUsecase "api-mini-project/businesses/products"
	_userUsecase "api-mini-project/businesses/users"
	_categoryController "api-mini-project/controllers/categories"
	_productController "api-mini-project/controllers/products"
	_userController "api-mini-project/controllers/users"
	_driverFactory "api-mini-project/drivers"
	_dbDriver "api-mini-project/drivers/mysql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"api-mini-project/util"

	"context"

	"github.com/labstack/echo/v4"
)

type operation func(ctx context.Context) error

func main() {
	configDB := _dbDriver.ConfigDB{
		DB_USERNAME: util.GetConfig("DB_USERNAME"),
		DB_PASSWORD: util.GetConfig("DB_PASSWORD"),
		DB_NAME:     util.GetConfig("DB_NAME"),
		DB_HOST:     util.GetConfig("DB_HOST"),
		DB_PORT:     util.GetConfig("DB_PORT"),
	}

	db := configDB.InitDB()

	_dbDriver.InitMigrate(db)

	configJWT := _middleware.ConfigJWT{
		SecretJWT:       util.GetConfig("JWT_SECRET_KEY"),
		ExpiresDuration: 1,
	}

	configLogger := _middleware.ConfigLogger{
		Format: "[${time_rfc3339}] ${status} ${method} ${host} ${path} ${latency_human}" + "\n",
	}

	server := echo.New()

	categoryRepo := _driverFactory.NewCategoryRepository(db)
	categoryUsecase := _categoryUsecase.NewCategoryUsecase(categoryRepo)
	categoryCtrl := _categoryController.NewCategoryController(categoryUsecase)

	productRepo := _driverFactory.NewProductRepository(db)
	productUsecase := _productUsecase.NewProductUsecase(productRepo)
	productCtrl := _productController.NewProductController(productUsecase)

	userRepo := _driverFactory.NewUserRepository(db)
	userUsecase := _userUsecase.NewUserUsecase(userRepo, &configJWT)
	userCtrl := _userController.NewAuthController(userUsecase)

	routesInit := _routes.ControllerList{
		LoggerMiddleware:   configLogger.Init(),
		JWTMiddleware:      configJWT.Init(),
		CategoryController: *categoryCtrl,
		ProductController:  *productCtrl,
		AuthController:     *userCtrl,
	}

	routesInit.Route(server)

	go func() {
		if err := server.Start(":8080"); err != nil && err != http.ErrServerClosed {
			server.Logger.Fatal("Shutting Down The Server")
		}
	}()

	wait := gracefulShutdown(context.Background(), 2*time.Second, map[string]operation{
		"database": func(ctx context.Context) error {
			return _dbDriver.CloseDB(db)
		},
		"http-server": func(ctx context.Context) error {
			return server.Shutdown(context.Background())
		},
	})

	<-wait
}

func gracefulShutdown(ctx context.Context, timeout time.Duration, ops map[string]operation) <-chan struct{} {
	wait := make(chan struct{})
	go func() {
		s := make(chan os.Signal, 1)

		signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
		<-s

		log.Println("shutting down")

		timeoutFunc := time.AfterFunc(timeout, func() {
			log.Printf("timeout %d ms has been elapsed, force exit", timeout.Milliseconds())
			os.Exit(0)
		})

		defer timeoutFunc.Stop()

		var wg sync.WaitGroup

		for key, op := range ops {
			wg.Add(1)
			innerOp := op
			innerKey := key
			go func() {
				defer wg.Done()

				log.Printf("cleaning up: %s", innerKey)
				if err := innerOp(ctx); err != nil {
					log.Printf("%s: clean up failed: %s", innerKey, err.Error())
					return
				}

				log.Printf("%s was shutdown gracefully", innerKey)
			}()
		}

		wg.Wait()

		close(wait)
	}()

	return wait
}
