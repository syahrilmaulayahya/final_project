package main

import (
	_middleware "final_project/app/middleware"
	"final_project/app/routes"
	_adminUseCase "final_project/business/admins"
	_productUseCase "final_project/business/products"
	_transactionUseCase "final_project/business/transactions"
	_userUseCase "final_project/business/users"
	_adminController "final_project/controllers/admins"
	_productController "final_project/controllers/products"
	_transactionController "final_project/controllers/transactions"
	_userController "final_project/controllers/users"
	_adminDB "final_project/drivers/databases/admins"
	_productDB "final_project/drivers/databases/products"
	_transactionDB "final_project/drivers/databases/transactions"
	_userDB "final_project/drivers/databases/users"
	_mysqlDriver "final_project/drivers/mysql"
	"log"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func dbMigrate(db *gorm.DB) {
	db.AutoMigrate(&_userDB.User{}, &_productDB.Product{}, &_userDB.Review_Rating{}, &_productDB.Product_description{}, &_productDB.Product_type{},
		&_productDB.Size{}, &_transactionDB.Shopping_Cart{}, &_transactionDB.Payment_Method{}, &_transactionDB.Shipment{}, &_transactionDB.Transaction{},
		&_transactionDB.Transaction_Detail{}, _adminDB.Admin{})
}

func main() {
	configDB := _mysqlDriver.ConfigDB{
		DB_Username: viper.GetString(`database.user`),
		DB_Password: viper.GetString(`database.pass`),
		DB_Host:     viper.GetString(`database.host`),
		DB_Port:     viper.GetString(`database.port`),
		DB_Database: viper.GetString(`database.name`),
	}
	configJWT := _middleware.ConfigJWT{
		SecretJWT:       viper.GetString(`jwt.secret`),
		ExpiresDuration: viper.GetInt(`jwt.expired`),
	}
	Conn := configDB.InitialDB()

	dbMigrate(Conn)
	e := echo.New()
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	userRepository := _userDB.NewMysqlRepository(Conn)
	userUseCase := _userUseCase.NewUserUseCase(userRepository, timeoutContext, configJWT)
	userController := _userController.NewUserController(userUseCase)

	productRepository := _productDB.NewMysqlRepository(Conn)
	productUseCase := _productUseCase.NewProductUseCase(productRepository, timeoutContext)
	productController := _productController.NewProductController(productUseCase)

	transactionRepository := _transactionDB.NewMysqlRepository(Conn)
	transactionUseCase := _transactionUseCase.NewTransactionUseCase(transactionRepository, timeoutContext, configJWT)
	transactionController := _transactionController.NewTransactionController(transactionUseCase)

	adminRepository := _adminDB.NewMysqlRepository(Conn)
	adminUseCase := _adminUseCase.NewAdminUseCase(adminRepository, timeoutContext)
	adminController := _adminController.NewAdminController(adminUseCase)
	routesInit := routes.ControllerList{
		UserController:        *userController,
		ProductController:     *productController,
		TransactionController: *transactionController,
		AdminController:       *adminController,
		JWTMiddleware:         configJWT.Init(),
	}

	routesInit.RouteRegister(e)
	log.Fatal(e.Start(viper.GetString("server.address")))
}
