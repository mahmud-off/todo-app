package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"

	todo "github.com/mahmud-off/todo-app/pkg"
	"github.com/mahmud-off/todo-app/pkg/handler"
	"github.com/mahmud-off/todo-app/pkg/repository"
	"github.com/mahmud-off/todo-app/pkg/service"
	"github.com/spf13/viper"
)

// @title			CRUD-REST TodoList
// @version		1.0
// @description	This is a sample Web todo list.
// @termsOfService	http://swagger.io/terms/
// @contact.name	API Support
// @contact.url	https://t.me/vmkdWW
// @contact.email	o.mahmudowvadim2020@ya.ru
// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html
// @host			localhost:8080
// @BasePath		/swagger/api/v1
// @securityDefinitions.basic  BasicAuth
func main() {

	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := InitConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variable: %s", err.Error())
		return
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		log.Fatal("error initialize db: " + err.Error())
		return
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error shutting down: %s", err.Error())
	}
	if err := db.Close(); err != nil {
		logrus.Errorf("error closing database: %s", err.Error())
	}

}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()

}
