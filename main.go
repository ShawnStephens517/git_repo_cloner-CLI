package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func directory_Verify() {
	fmt.Println(`
	Checking if projects directory is on your desktop
	If it is not present,
	Magic Spells will make it appear...
	It may be sorcery!
	`)
	if _, err := os.Stat("~/Desktop/thunderchicken"); os.IsNotExist(err) {
		os.Mkdir("~/Desktop/thunderchicken", 0755)
	} else {
		fmt.Println("Projects folder alread exists")
	}

}

func main() {
	projectName :=
		`
--------------------------------------------------------------------------------------



██████  ██ ████████      ██████  ███████ ██████   ██████           ██████ ██       ██████  ███    ██ ███████ ██████  
██       ██    ██         ██   ██ ██      ██   ██ ██    ██         ██      ██      ██    ██ ████   ██ ██      ██   ██ 
██   ███ ██    ██         ██████  █████   ██████  ██    ██         ██      ██      ██    ██ ██ ██  ██ █████   ██████  
██    ██ ██    ██         ██   ██ ██      ██      ██    ██         ██      ██      ██    ██ ██  ██ ██ ██      ██   ██ 
 ██████  ██    ██ ███████ ██   ██ ███████ ██       ██████  ███████  ██████ ███████  ██████  ██   ████ ███████ ██   ██ 
                                                                                                                      
                                                                                                                      
A CLI for cloning all Github or Gitlab projects to your local machine.

Created by Shawn Stephens
                                                                                                                                                 
                                                                                                                                                 


--------------------------------------------------------------------------------------
`
	fmt.Println(projectName)

	directory_Verify()

	var domain string
	fmt.Println("Enter the GitLab Instance url.")
	fmt.Scanln(&domain)

	var username string
	fmt.Println("Enter your Username now")
	fmt.Scanln(&username)

	userInfo, err := http.Get(domain + "/api/v4/users?username=" + username)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(userInfo.Body)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData))
}
