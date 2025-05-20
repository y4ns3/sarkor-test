package main

import (
	"log"
	"net"

	"github.com/joho/godotenv"
	"github.com/y4ns3/sarkor-test/config"
	"github.com/y4ns3/sarkor-test/db"
	"github.com/y4ns3/sarkor-test/internal/repository"
	"github.com/y4ns3/sarkor-test/internal/rest"
	"github.com/y4ns3/sarkor-test/internal/rest/handler"
	"github.com/y4ns3/sarkor-test/internal/usecase/product"
)

func main() {
	if err := execute(); err != nil {
		log.Fatal(err)
	}
}
func execute() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	cfg, err := config.NewConfig()
	if err != nil {
		return err
	}
	db, err := db.NewDB(cfg.Db.Url)
	if err != nil {
		return err
	}
	defer db.Close()

	productRepository := repository.NewProductRepository(db)
	productService := product.NewService(productRepository)
	productHandler := handler.NewHandler(productService)
	if err != nil {
		return err
	}

	server := rest.NewServer(productHandler)

	url := net.JoinHostPort(cfg.Server.Host, cfg.Server.Port)

	return server.Run(url)
}
