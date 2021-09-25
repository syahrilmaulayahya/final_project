package main

import (
	"final_project/app/routes"
	_productUseCase "final_project/business/products"
	_userUseCase "final_project/business/users"
	_productController "final_project/controllers/products"
	_userController "final_project/controllers/users"
	_productDB "final_project/drivers/databases/products"
	_userDB "final_project/drivers/databases/users"
	_mysqlDriver "final_project/drivers/mysql"
	"log"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func init() {
	viper.SetConfigFile(`app/config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func dbMigrate(db *gorm.DB) {
	db.AutoMigrate(&_userDB.User{}, &_productDB.Product{})
}

func main() {
	configDB := _mysqlDriver.ConfigDB{
		DB_Username: viper.GetString(`database.user`),
		DB_Password: viper.GetString(`database.pass`),
		DB_Host:     viper.GetString(`database.host`),
		DB_Port:     viper.GetString(`database.port`),
		DB_Database: viper.GetString(`database.name`),
	}

	Conn := configDB.InitialDB()

	dbMigrate(Conn)
	e := echo.New()
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	userRepository := _userDB.NewMysqlRepository(Conn)
	userUseCase := _userUseCase.NewUserUseCase(userRepository, timeoutContext)
	userController := _userController.NewUserController(userUseCase)
	productRepository := _productDB.NewMysqlRepository(Conn)
	productUseCase := _productUseCase.NewProductUseCase(productRepository, timeoutContext)
	productController := _productController.NewProductController(productUseCase)
	routesInit := routes.ControllerList{
		UserController:    *userController,
		ProductController: *productController,
	}

	routesInit.RouteRegister(e)
	log.Fatal(e.Start(viper.GetString("server.address")))
}
