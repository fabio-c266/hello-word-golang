package main

import "fmt"

func main() {
	fmt.Println("Sistema de monitoramento")
	fmt.Println("1- Iniciar monitoramento")
	fmt.Println("2- Exibir logs")
	fmt.Println("3- Sair do programa")

	var optionInputed int
	fmt.Scan(&optionInputed)

	fmt.Println(optionInputed)
}
