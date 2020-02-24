package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

// Randomized User Agent
var userAgent = "Mozilla/5.0 (iPhone; CPU iPhone OS 7_0_1 like Mac OS X) AppleWebKit/537.51.1 (KHTML, like Gecko) Version/7.0 Mobile/11A470a Safari/9537.53"

// Path and Files
var pathF = "config"
var extensions = "extensions.txt"
var files = "files.txt"
var folders = "folders.txt"
var foundedFolders []string

func getStatusCode(url string) string {

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	return strconv.Itoa(resp.StatusCode)
}

func scanFiles() {

	for _, fndPTHS := range foundedFolders {
		fmt.Println("\n************* Starting Scan Backups Files. / Founded PATH : " + fndPTHS + " *************\n")

		fileslist, err := os.Open(pathF + "/files.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer fileslist.Close()
		var lastCheck = ""
		fileScan := bufio.NewScanner(fileslist)
		for fileScan.Scan() {

			extensionss, err := os.Open(pathF + "/extensions.txt")
			if err != nil {
				log.Fatal(err)
			}
			defer extensionss.Close()

			scanner := bufio.NewScanner(extensionss)
			for scanner.Scan() {

				var urlE = fndPTHS + "/" + fileScan.Text() + scanner.Text()

				lastCheck = getStatusCode(urlE)

				var chckDrm = urlE + " | Response Code : " + lastCheck

				if lastCheck == "200" || lastCheck == "301" || lastCheck == "302" || lastCheck == "304" || lastCheck == "307" || lastCheck == "403" {
					fmt.Printf("\033[2K\r%s\n", "*Founded :"+chckDrm)
					foundedFolders = append(foundedFolders, urlE)
				} else {
					fmt.Printf("\033[2K\r%s", "Checking : "+chckDrm)
				}

			}

		}

	}

}

func scanPath(filename string, hostname string) string {

	file, err := os.Open(pathF + "/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lastCheck = ""
	scanner := bufio.NewScanner(file)

	fmt.Println("\n************* Starting Scan Backups PATHS *************\n")

	for scanner.Scan() {
		var urlE = hostname + "/" + scanner.Text()
		lastCheck = getStatusCode(urlE)

		var chckDrm = "" + urlE + " | Response Code : " + lastCheck

		if lastCheck == "200" || lastCheck == "301" || lastCheck == "302" || lastCheck == "304" || lastCheck == "307" || lastCheck == "403" {
			fmt.Printf("\033[2K\r%s\n", "* Founded : "+chckDrm)
			foundedFolders = append(foundedFolders, urlE)
		} else {
			fmt.Printf("\033[2K\r%s", "Checking : "+chckDrm)
		}

	}

	fmt.Printf("\033[2K\r%s", "\nPath Scaning Ended.\n")
	scanFiles()
	return lastCheck
}

func main() {

	hostname := flag.String("hostname", "", "Please input hostname")
	flag.Parse()

	fmt.Println(`
      _   _   _   _   _   _   _   _   _   _   _   _  
     / \ / \ / \ / \ / \ / \ / \ / \ / \ / \ / \ / \ 
    ( g | e | t | T | h | e | B | a | c | k | u | p )
     \_/ \_/ \_/ \_/ \_/ \_/ \_/ \_/ \_/ \_/ \_/ \_/ 
						
	Scan backup files and directories.
	Github : github.com/nishitm
	Host : ` + *hostname + `
	`)

	startScan := scanPath(folders, *hostname)
	fmt.Println(startScan)
}
