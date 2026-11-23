package main

import (
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"strings"
	"yathahi/source/library"
	packagemanager "yathahi/source/library/package-manager"
)

//--------------------------------------------------------------------------
//	Initiating The Container
//--------------------------------------------------------------------------

func Containersetup() {

	packagemanager.PackageManager("lxc")

	slog.Info("Initializing LXC Deployment...", "container", library.ProjectName)

	// 1. Enable lxc-net Service
	slog.Info("Enabling lxc-net service...")
	out, err := exec.Command("systemctl", "enable", "--now", "lxc-net").CombinedOutput()
	if err != nil {
		slog.Error("Failed to enable lxc-net",
			"error", err,
			"output", string(out),
		)
		os.Exit(1)
	}

	// 2. Create the LXC Container

	_, err = os.Stat("/var/lib/lxc/yathahi/")
	if err != nil {
		slog.Info("Creating Container...", "container", library.ProjectName)
		out, err = exec.Command(
			"lxc-create",
			"-n", library.ProjectName,
			"-t", "download",
			"--",
			"--dist", "debian",
			"--release", "trixie",
			"--arch", "amd64",
		).CombinedOutput()
		if err != nil {
			slog.Error("Failed to create container",
				"container", library.ProjectName,
				"error", err,
				"output", string(out),
			)
			os.Exit(1)
		}
		if len(out) > 0 {
			slog.Info("Container creation output", "output", string(out))
		}

	}

	out, err = exec.Command("sudo", "lxc-info", "-n", "yathahi", "-s").Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fields := strings.Fields(string(out))
	if len(fields) >= 2 && fields[1] == "STOPPED" {
		fmt.Println("Container is in stopped state!")
		// 3. Start the Container
		slog.Info("Starting LXC Container...", "container", library.ProjectName)
		out, err = exec.Command("lxc-start", "-n", library.ProjectName).CombinedOutput()
		if err != nil {
			slog.Error("Failed to start container",
				"container", library.ProjectName,
				"error", err,
				"output", string(out),
			)
			return
		}
	}

	slog.Info("LXC Deployment Complete!", "container", library.ProjectName)
	slog.Info("Configuring Firewall To make internet access inside the lxc Container")
	err = provisionFirewall()
	if err != nil {
		slog.Error("Something went wrong with Firewall", "error", err)
	}
}

//-----------------------------------------------------------------------------
//	Provision Firewall
//-----------------------------------------------------------------------------

func provisionFirewall() error {

	// 1. Check for firewalld
	fwActiveCmd := exec.Command("systemctl", "is-active", "firewalld")
	fwOut, _ := fwActiveCmd.Output()

	ufwActiveCmd := exec.Command("systemctl", "is-active", "ufw")
	ufwOut, _ := ufwActiveCmd.Output()

	if strings.TrimSpace(string(fwOut)) == "active" {
		err := configureFirewalld()
		if err != nil {
			return err
		}

	} else if strings.TrimSpace(string(ufwOut)) == "active" {
		err := configureUfw()
		if err != nil {
			return err
		}
	}

	slog.Warn("Neither firewalld nor ufw is currently active..\nAn active firewall is Recommended.")
	return nil

}

//--------------------------------------------------------------------------
//	Configure Firewalld
//--------------------------------------------------------------------------

func configureFirewalld() error {

	slog.Info("Detected active firewalld. Applying rules...")

	rules := [][]string{
		{"--zone=trusted", "--add-interface=lxcbr0", "--permanent"},
		{"--zone=public", "--add-masquerade", "--permanent"},
		{"--reload"},
	}

	for _, args := range rules {
		cmd := exec.Command("firewall-cmd", args...)
		if out, err := cmd.CombinedOutput(); err != nil {
			return fmt.Errorf("firewall-cmd %s failed: %v (output: %s)\n", strings.Join(args, " "), err, string(out))
		}
	}
	slog.Info("firewalld configuration complete.")
	return nil

}

//--------------------------------------------------------------------------
//	Configure UFW
//--------------------------------------------------------------------------

func configureUfw() error {
	// 2. Check for UFW
	slog.Info("Detected active UFW. Applying rules...")

	rules := [][]string{
		{"allow", "in", "on", "lxcbr0"},
		{"route", "allow", "in", "on", "lxcbr0"},
	}

	for _, args := range rules {
		cmd := exec.Command("ufw", args...)
		if out, err := cmd.CombinedOutput(); err != nil {
			return fmt.Errorf("ufw %s failed: %v (output: %s)\n", strings.Join(args, " "), err, string(out))
		}
	}
	slog.Info("UFW configuration complete.")
	return nil
}
