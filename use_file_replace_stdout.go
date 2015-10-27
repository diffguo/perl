package main

import (
	//"bufio"
	"fmt"
	"os"
	"os/exec"
	//"syscall"
)

func main() {
	cmd := exec.Command("ls", "-l")

	logFile, _ := os.OpenFile("/Users/jason/WorkSpace/WorkSpace_go/src/perl/out.txt", os.O_WRONLY|os.O_CREATE|os.O_SYNC, 0755)

	cmd.Stdout = logFile
	cmd.Stderr = logFile

	err := cmd.Start()
	if err != nil {
		fmt.Println(err)
	}

	cmd.Wait()
}
