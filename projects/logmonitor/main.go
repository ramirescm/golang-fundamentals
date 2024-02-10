package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const retry = 3
const delay = 1

func main() {
	showIntroduction()
	for {
		showMenu()
		command := readCommand()

		switch command {
		case 1:
			monitoring()
		case 2:
			fmt.Println("Show logs")
			showLogs()
		case 0:
			fmt.Println("Exit")
			os.Exit(0)
		default:
			fmt.Println("Command unknow")
			os.Exit(-1)
		}
	}

}

func showIntroduction() {
	name := "User"
	version := 1.0

	fmt.Println("Hello, Sr. ", name)
	fmt.Println("Program version: ", version)
}

func showMenu() {
	fmt.Println("1- Start monitoring")
	fmt.Println("2- Show logs")
	fmt.Println("0- Exit")
}

func readCommand() int {
	var command int
	fmt.Scan(&command)

	fmt.Println("type of var ", reflect.TypeOf(command))
	fmt.Println("Selected option: ", command)

	return command
}

func monitoring() {
	fmt.Println("Monitoring")

	sites := readUrlFromFile()

	for i := 0; i < retry; i++ {
		for _, site := range sites {
			callSite(site)
		}

		time.Sleep(delay * time.Minute)
		fmt.Println("")
	}
}

func callSite(site string) {
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site: ", site, " - Up!")
		logFileWrite(site, true)
	} else {
		fmt.Println("Site: ", site, "Ops, try again! Status code - ", resp.StatusCode)
		logFileWrite(site, false)
	}
}

func readUrlFromFile() []string {
	var sites []string
	file, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Error ", err)
	}

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		sites = append(sites, line)
		if err == io.EOF {
			break
		}
	}

	file.Close()

	return sites
}

func logFileWrite(site string, status bool) {
	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Error ", err)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " +
		strconv.FormatBool(status) + "\n")

	file.Close()
}

func showLogs() {
	file, err := ioutil.ReadFile("log.txt")
	if err != nil {
		fmt.Println("Error ", err)
	}

	fmt.Println(string(file))
}
