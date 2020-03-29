package chapter06

import (
	"fmt"
	"log/syslog"
	"os"
	"path/filepath"
)

func Start03() {
	args := os.Args
	programName := filepath.Base(args[0])
	//sysLog, err := syslog.New(syslog.LOG_INFO|syslog.LOG_LOCAL7, programName)
	sysLog, err := syslog.New(syslog.LOG_ALERT|syslog.LOG_LOCAL7, programName)
	if err != nil {
		fmt.Println("error :", err)
		os.Exit(1)
	}

	sysLog.Crit("Crit: Loggingin go!")
	fmt.Fprintf(sysLog, "logging stuff!!!")
}
