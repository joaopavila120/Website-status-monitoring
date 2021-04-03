package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const numberOfMonitorings = 3 
const delay               = 5

func main() {
    for {
        fmt.Println("1 - Start website status monitoring")
        fmt.Println("2 - Show log history")
        fmt.Println("3 - Close")
    
        instruction := readInstruction()

        switch instruction {
            case 1: 
                startMonitoring()
            case 2:
                fmt.Println("Opening logs.. ")
                printLogs()
            case 3:
                fmt.Println("Closing...")
                os.Exit(0) //close
            default:
                fmt.Println("Invalid instruction")
                os.Exit(-1) //close throwing error
            }
    }
} 

func readInstruction() int {
    var instruction int 
    fmt.Scanf("%d", &instruction)
    fmt.Println("You choosed option n°", instruction)
    fmt.Println("")

    return instruction
}

func startMonitoring() {
    fmt.Println("Starting monitoring...")
    website :=  readFile()

    for i:=0; i < numberOfMonitorings; i++ {
        for i, website := range website{
            fmt.Println("Connecting to website n°", i, ".Adress: ", website)
            siteTest(website)
        }      
        fmt.Println("")
        fmt.Println("[#######]")
        fmt.Println("Waiting to restart application")
        fmt.Println("[#######]")
        fmt.Println("")
        time.Sleep(delay * time.Second)  
    }        
}

func siteTest(website string) {
    response, err := http.Get(website) //get function has two returns, _ is used to treat errors

    if err != nil {
        fmt.Println("An error occurred:", err)
    }

    if response.StatusCode == 200 { //200 confirms that the website is online
        fmt.Println("The website:", website, "is online")
        registerLogs(website, true)
    }else {
        fmt.Println("The website", website, "is offline")
        fmt.Println("Status Code:", response.StatusCode)
        registerLogs(website, false)
    }
}

func readFile() []string {
    var websites []string

    file, err := os.Open("websites.txt")

    if err != nil {
        fmt.Println("An error occurred while trying to read the file", err)
    }

    reader := bufio.NewReader(file)
    for {
        line, err        := reader.ReadString('\n') // \n is the delimiter byte for broke lines in txt files
        currentLine      := strings.TrimSpace(line) //remove whitespaces from the txt file

        websites = append(websites, currentLine)
        
        if err == io.EOF {
            break
        }
    }

    file.Close()
    return websites
}

func registerLogs(website string, status bool) {
    file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666) //read and write flag, create file flag, file permission 

    if err != nil {
        fmt.Println("An error occurred:", err)
    }

    file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + website + " - online: " + strconv.FormatBool(status) + "\n")

    file.Close()
}

func printLogs() {
    file, err := ioutil.ReadFile("log.txt")

    if err != nil {
        fmt.Println("An error occurred:", err)
    }

    fmt.Println(string(file))
}