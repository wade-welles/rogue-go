package rogue

import (
	"os"
	"os/exec"
	"path/filepath"
)

func RunAsRoot() {
	// Root is required because we are directly interacting with process memory
	if os.Geteuid() != 0 {
		Info("The " + White("Rogue:GO") + LightGray(" hack engine requires root access to interact with game memory, re-running the process with sudo...\n"))
		executable, _ := os.Executable()
		directory, binary := filepath.Split(executable)
		os.Chdir(directory)
		cmd := exec.Command("/bin/sh", "-c", "sudo ./"+binary)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
		err := cmd.Run()
		if err != nil {
			FatalError(err, "Failed to re-run the executable with sudo: ")
			os.Exit(1)
		} else {
			// Exit because we just relaunched as root
			os.Exit(0)
		}
	}
}
