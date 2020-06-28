package main

import (
	"fmt"
	"os/exec"
)

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
