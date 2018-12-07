package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"syscall"
	//	"strings"
)

const PGDATA = "PGDATA"
const BORN_SLAVE = "BORN_SLAVE"

type Environment struct {
	dataDir  string
	isMaster bool
}

func dirIsEmpty(dirName string) (bool, error) {
	f, err := os.Open(dirName)
	if err != nil {
		return false, err
	}
	defer f.Close()

	_, err = f.Readdirnames(1) // Or f.Readdir(1)
	if err == io.EOF {
		return true, nil
	}
	return false, err // Either not empty or error, suits both cases
}

func runProcess(proc string, args []string) {
	binary, lookErr := exec.LookPath(proc)
	if lookErr != nil {
		panic(lookErr)
	}
	env := os.Environ()
	args = append([]string{proc}, args...)
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}

func startPgOnMaintenancePort(environment Environment) {

}

func initDB() {
	runProcess("/docker-entrypoint.sh", os.Args[1:])
}

func main() {

	enironment := Environment{
		dataDir:  os.Getenv(PGDATA),
		isMaster: os.Getenv(BORN_SLAVE) != "true",
	}
	fmt.Printf("%#v\n", enironment)
	empty, err := dirIsEmpty(enironment.dataDir)
	if err != nil {
		panic(err)
	}
	if empty {
		initDB()
		startPgOnMaintenancePort(enironment)
	} else {
		fmt.Printf("Fuck")
	}
}
