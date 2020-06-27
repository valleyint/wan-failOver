package main

import (
	"fmt"
	"os/exec"
)

func Ping () bool {
	pinger := exec.Command("ping"  , "-c4" , "-W1500" , "8.8.8.8")
	err := pinger.Run()
	if err != nil {
		fmt.Println("error is" , err)
		return false
	} else {
		return true
	}
}

func main () {
	fmt.Println(Ping())
}
