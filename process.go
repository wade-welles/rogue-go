package rogue

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type GameProcess struct {
	BinaryName    string
	ProcDirectory string
	Process       *os.Process
	PID           int
	EngineSO      string
	ClientSO      string
	Memory        string
}

func FindGameProcess() *GameProcess {
	gameProcess := &GameProcess{
		BinaryName: "csgo_linux64",
		EngineSO:   "engine_client.so",
		ClientSO:   "client_panorama_client.so",
	}
	cmdlineFiles, _ := filepath.Glob("/proc/*/cmdline")
	for _, file := range cmdlineFiles {
		cmdlineData, _ := ioutil.ReadFile(file)
		if len(cmdlineData) != 0 {
			cmdline := string(cmdlineData)
			directory, binaryName := filepath.Split(cmdline)
			if strings.Contains(binaryName, gameProcess.BinaryName) {
				gameProcess.ProcDirectory = directory
				pathSegments := strings.Split(file, "/")
				gameProcess.PID, _ = strconv.Atoi(pathSegments[2])
				Info("Successfully found ", White("Counter-strike:GO"), LightGray(" process with the PID: "), Yellow(strconv.Itoa(gameProcess.PID)))
			}
		}
	}
	if gameProcess.PID != 0 {
		return gameProcess
	} else {
		Warning("Failed to locate a running instance of ", White("Counter-strike:GO"), LightGray(". Please start the game..."))
		Info("Waiting [", White("15 seconds"), LightGray("] and then searching again, use [")+Yellow("CTRL+C")+LightGray("] to cancel."))
		time.Sleep(15 * time.Second)
		FindGameProcess()
		return nil
	}
}
