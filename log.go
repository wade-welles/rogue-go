package rogue

import (
	"fmt"
	"log"
)

type category int

const (
	FATAL   = "FATAL"
	ERROR   = "ERROR"
	WARNING = "WARNING"
	INFO    = "INFO"
)

func PrintLog(category string, newline bool, text ...string) {
	content := mergeText(text...)
	if newline {
		fmt.Printf("\n")
	}
	switch category {
	case FATAL:
		log.Println(Red("[ROGUE:FATAL]"), Yellow(content))
	case ERROR:
		log.Println(Red("[ROGUE:ERROR]"), Yellow(content))
	case WARNING:
		log.Println(Yellow("[ROGUE:WARNING]"), LightGray(content))
	case INFO:
		fmt.Println(Blue("[ROGUE:INFO]"), LightGray(content))
	default:
		if category == "" {
			fmt.Println(Blue("[ROGUE]"), LightGray(content))
		} else {
			fmt.Println(Blue("[ROGUE:"+category+"]"), LightGray(content))
		}
	}
}

func mergeText(text ...string) string {
	var content string
	for _, t := range text {
		content += t
	}
	return content
}

func FatalError(err error, text ...string) {
	if err != nil {
		PrintLog(FATAL, false, text...)
	} else {
		content := mergeText(text...)
		PrintLog(FATAL, false, content+": "+err.Error())
	}
}

func Error(err error, text ...string) {
	if err != nil {
		PrintLog(ERROR, false, text...)
	} else {
		content := mergeText(text...)
		PrintLog(ERROR, false, content+": "+err.Error())
	}

}

func Warning(text ...string) {
	PrintLog(WARNING, false, text...)
}

func Info(text ...string) {
	PrintLog(INFO, false, text...)
}

func Log(text ...string) {
	PrintLog("", false, text...)
}
