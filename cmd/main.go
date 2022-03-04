package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/lib/pq"
	"github.com/nevskyw/todo-app"
	"github.com/nevskyw/todo-app/pkg/handler"
	"github.com/nevskyw/todo-app/pkg/repository"
	"github.com/nevskyw/todo-app/pkg/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// initConfig - инициализируем конфигурационные файлы
func initConfig() error {
	viper.AddConfigPath("configs") // - указываем имя нашей дериктории
	viper.SetConfigName("config")  // - указываем имя нашего файла
	return viper.ReadInConfig()
}

func main() {
	// формат JSON для наших логов
	logrus.SetFormatter(new(logrus.JSONFormatter))

	// initConfig - выполнение функции
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	// Load - переменная окружения, в которой передаем значение пороля
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	// База данных - инициализация
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	// Инициализируем сервер в горутине
	srv := new(todo.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil { // handlers.InitRoutes() - возрвщвет указатель типа gin.Engine
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Print("TodoApp Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit // строка для чтения из канала, которая блокирует выполнение главной горутины мэйн

	logrus.Print("TodoApp Shutting Down")

	// метод отсановки сервера
	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}
	// закрытие всех соединений с базой данных
	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}
}
