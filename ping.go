package main

import (
	"fmt"
	"os/exec"
)

func Ping () bool {
	pinger := exec.Command("fping" , "8.8.8.8")
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

func On (portNo string) error {
	cmd := exec.Command("ifup" , fmt.Sprintf("etc%v" , portNo))
	err := cmd.Run()
	if err != nil {
		return err
	} else {
		return nil
	}
}

func Off (portNo string) error {
	cmd := exec.Command("ifdown" , fmt.Sprintf("etc%v" , portNo))
	err := cmd.Run()
	if err != nil {
		return err
	} else {
		return nil
	}
}

func Reboot () error {
	cmd := exec.Command("reboot")
	err := cmd.Run()
	if err != nil {
		return err
	} else {
		return nil
	}
}

