package vpn

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"h1r4vpn/internal/utils"
)

const pidFile = "/tmp/h1r4vpn.pid"

func validateConfig(cfg string) error {
	if filepath.Ext(cfg) != ".ovpn" {
		return errors.New("invalid config extension")
	}

	info, err := os.Stat(cfg)
	if err != nil {
		return err
	}

	if info.Mode().Perm()&0002 != 0 {
		return errors.New("config is world-writable")
	}

	return nil
}

func Connect(cfg, logFile string) error {
	if err := validateConfig(cfg); err != nil {
		return err
	}

	before := utils.PublicIP()

	fmt.Println("[+] Current IP:", before)
	fmt.Println("[+] Launching OpenVPN...")

	cmd := exec.Command(
		"sudo",
		"openvpn",
		"--config", cfg,
		"--verb", "3",
		"--log-append", logFile,
		"--writepid", pidFile,
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		return err
	}

	fmt.Println("[+] OpenVPN PID:", cmd.Process.Pid)

	time.Sleep(5 * time.Second)

	after := utils.PublicIP()

	fmt.Println("[+] New IP:", after)

	if before == after {
		fmt.Println("[!] Warning: IP unchanged")
	} else {
		fmt.Println("[+] VPN tunnel established")
	}

	return nil
}

func Disconnect() error {
	data, err := os.ReadFile(pidFile)
	if err != nil {
		return errors.New("VPN not running")
	}

	pid := strings.TrimSpace(string(data))

	cmd := exec.Command("sudo", "kill", pid)

	if err := cmd.Run(); err != nil {
		return err
	}

	_ = os.Remove(pidFile)

	fmt.Println("[+] VPN disconnected")

	return nil
}