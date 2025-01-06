package main

import (
	"fmt"
	"log"
	"split_it_backend/app/routes"
	_middleware "split_it_backend/app/middleware"
	_userUseCase "split_it_backend/businesses/users"

	_userController "split_it_backend/controllers/users"

	postgres_driver "split_it_backend/drivers/database/postgres"
	"split_it_backend/drivers/database/users"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)
func init(){
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".","_"))
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}

}
func dbMigrate(db *gorm.DB){
	db.AutoMigrate(
		&users.User{},
	)
}

func main() {
	configDB := postgres_driver.ConfigDB{
		DB_Host: viper.GetString(`db.host`),
		DB_User: viper.GetString(`db.user`),
		DB_Password : viper.GetString(`db.password`),
		DB_Name : viper.GetString(`db.dbname`),
		DB_Port : viper.GetInt(`db.port`),
		DB_SSLMode : viper.GetString(`db.sslmode`),
		DB_TimeZone : viper.GetString(`db.timezone`),
	}
	db := configDB.InitialDB()
	dbMigrate(db)

	timeoutContext := time.Duration(viper.GetInt("server.context_timeout")) * time.Second
	
	e := echo.New()
	e.Use(middleware.CORS())
	jwt := _middleware.ConfigJWT{
		SecretKey : viper.GetString(`jwt.secret`),
		ExpiredTime : viper.GetInt(`jwt.expired`),
	}

	
	userRepo := users.NewUserRepository(db)
	userUseCase := _userUseCase.NewUseCase(userRepo,timeoutContext,jwt)
	userController := _userController.NewUserController(userUseCase)

	routesInit := routes.ControllerList{
		UserController: *userController,
		JWTConfig	: jwt.Init(),
	}

	routesInit.RouteRegister(e)
	port := viper.GetInt("server.port")
	log.Fatal(e.Start(fmt.Sprintf(":%d", port)))

}