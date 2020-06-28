package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
	"time"
)

type wan struct {
	name string
}

func Ping() bool {
	pinger := exec.Command("fping", "8.8.8.8")
	err := pinger.Run()
	if err != nil {
		log.Println("ping failed with result", err)
		return false
	} else {
		return true
	}
}

func main() {
	Daemon()
}

func (w *wan) On() {
	log.Println("switching on", w.name)
	cmd := exec.Command("ifup", w.name)
	err := cmd.Run()
	time.Sleep(5 * time.Second)
	if err != nil {
		log.Println(err)
	}
	UpdateDNS()
}

func (w *wan) Off() {
	log.Println("switching off", w.name)
	cmd := exec.Command("ifdown", w.name)
	err := cmd.Run()
	time.Sleep(3 * time.Second)
	if err != nil {
		log.Println(err)
	}
}

func (w *wan) Switch() {
	log.Println("switching from", w.name)
	w.Off()
	switch w.name {
	case "eth1":
		w.name = "eth2"
	case "eth2":
		w.name = "eth1"
	default:
		w.name = "eth1"
	}
	w.On()
	log.Println("switched to", w.name)
}

func InitWan() *wan {
	log.Println("initializing wan")
	wan := wan{name: "eth1"}
	wan.On()
	return &wan
}

func Daemon() {
	log.Println("deamon started")
	wan := InitWan()
	for {
		time.Sleep(5 * time.Second)
		err := Ping()
		if err != true {
			wan.Switch()
		}
	}
}

func DNSConfig() (host, password string) {
	data, err := ioutil.ReadFile("/usr/local/etc/updatedns.conf")
	if err != nil {
		log.Println(err)
		return
	}
	lines := strings.Split(string(data), "\n")
	if len(lines) < 2 {
		return
	}
	host = lines[0]
	password = lines[1]
	return
}

func UpdateDNS() {
	host, password := DNSConfig()
	if host == "" {
		return
	}
	log.Println("updating dns")
	url := fmt.Sprintf("https://dyn.dns.he.net/nic/update?hostname=%s&password=%s", host, password)
	com := exec.Command("curl", "-4", url)
	err := com.Run()
	if err != nil {
		log.Println(err)
	}
}
