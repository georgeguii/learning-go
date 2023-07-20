package main

import (
	"database/sql"
	"fmt"

	"github.com/georgeguii/learning-go/internal/infra/database"
	"github.com/georgeguii/learning-go/internal/usecase"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/testeGo?charset=utf8mb4&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	orderRepository := database.NewOrderRepository(db)

	uc := usecase.NewCalculateFinalPrice(orderRepository)

	input := usecase.OrderInput{
		Id:    "1",
		Price: 10.0,
		Tax:   1.0,
	}

	output, err := uc.Execute(input)
	if err != nil {
		panic(err)
	}
	fmt.Println(output)
}
