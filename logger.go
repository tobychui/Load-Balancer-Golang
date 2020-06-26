package main

import (
  "log"
  "time"
)

// HH/mm/SS DD/MM/YYYY
const formatString string = "15:04:05 01/02/2006"

func LogWarn(msg string) {
  log.Println("[Warning] " + msg)
}

func LogErr(msg string) {
  log.Println("[ERROR] " + msg)
}

func LogErrAndCrash(msg string) {
  panic(msg)
}

func getFormattedTime() string {
  return time.Now().Format(formatString)
}
