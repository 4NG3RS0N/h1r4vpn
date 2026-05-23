package vpn

import (
	"fmt"
	"os"

	"h1r4vpn/internal/utils"
)

func Status() {
	fmt.Println("[+] Public IP:", utils.PublicIP())

	if _, err := os.Stat("/tmp/h1r4vpn.pid"); err == nil {
		fmt.Println("[+] VPN status: CONNECTED")
	} else {
		fmt.Println("[-] VPN status: DISCONNECTED")
	}
}

func Logs() {
	data, err := os.ReadFile("logs/h1r4vpn.log")
	if err != nil {
		fmt.Println("[-] Could not read log file")
		return
	}

	fmt.Println(string(data))
}