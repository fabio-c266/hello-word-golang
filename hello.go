package main

import "fmt"

func startMonitoring() {
	println("Monitoring started")
}

func showLogs() {
	println("Showing Logs")
}

func main() {
	println("Monitoring System")
	println("1- Start monitoring")
	println("2- Show logs")
	println("3- End system")

	endSystem := false

	for !endSystem {
		var optionInputed int
		fmt.Scan(&optionInputed)

		switch optionInputed {
		case 1:
			startMonitoring()
		case 2:
			showLogs()
		case 3:
			println("Ending system...")
			println("System ended")
			endSystem = true
		default:
			println("Invalid option. Try again")
		}
	}
}
