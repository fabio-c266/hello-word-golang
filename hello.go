package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

const websitesFileName = "websites.txt"
const logsFileName = "logs.txt"
const requestDelay = 5 * time.Second

func handleConfigAndLogsFiles() {
	websitesFile, _ := os.Open(websitesFileName)

	if websitesFile == nil {
		_, err := os.Create(websitesFileName)

		if err != nil {
			println(err)
			os.Exit(1)
		}
	}

	logsFile, _ := os.Open(logsFileName)

	if logsFile == nil {
		_, err := os.Create(logsFileName)

		if err != nil {
			println(err)
			os.Exit(1)
		}
	}
}

func main() {
	handleConfigAndLogsFiles()

	for {
		showMenu()
		optionInputed := getOption()

		switch optionInputed {
		case 1:
			startMonitoring()
		case 2:
			showLogs()
		case 3:
			println("Ending system...")
			break
		default:
			println("Invalid option. Try again")
		}
	}
}

func showMenu() {
	println("Monitoring System")
	println("1- Start monitoring")
	println("2- Show logs")
	println("3- End system")
}

func startMonitoring() {
	websiteFileInfo, _ := os.Stat(websitesFileName)

	if websiteFileInfo == nil {
		println(websitesFileName, "not founded, try open program again.")
		os.Exit(1)
	}

	if websiteFileInfo.Size() == 0 {
		println("Empty websites file, is required at least 1 website")
		return
	}

	logsFileInfo, _ := os.Stat(logsFileName)

	if logsFileInfo == nil {
		println(logsFileName, "not founded, try open program again.")
		os.Exit(1)
	}

	websitesFile, _ := os.Open(websitesFileName)
	defer websitesFile.Close()

	sc := bufio.NewScanner(websitesFile)
	websites := []string{}

	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		websites = append(websites, line)
	}

	logsFile, _ := os.OpenFile(logsFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer logsFile.Close()

	for {
		for _, website := range websites {
			res, err := http.Get(website)

			if err != nil {
				println("Invali URL:", string(website))
				os.Exit(1)
			}

			var websiteStatus string

			if res.StatusCode == 200 {
				websiteStatus = "ONLINE"
			} else {
				websiteStatus = "OFFLINE"
			}

			logLine := "(" + website + "): " + websiteStatus
			println(logLine)

			currentDateTime := time.Now().Format("02/01/2006 15:04:05")
			logsFile.WriteString("[" + currentDateTime + "]: " + logLine + "\n")
		}

		time.Sleep(requestDelay)
	}
}

func showLogs() {}

func getOption() int {
	var optionInputed int
	fmt.Scan(&optionInputed)

	return optionInputed
}
