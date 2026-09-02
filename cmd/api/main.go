package main

import (
	"fmt"
	"go-messenger/internal/database"
)

func main() {
	fmt.Println("Iniciar servidor.")

	_ = database.NewMySQLConnection()
	fmt.Println("Banco conectado")

}
