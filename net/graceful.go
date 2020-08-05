package net

import (
	"os"
	"os/exec"
)

func Parent(){
	file := netListener.File() // this returns a Dup()
	path := "/path/to/executable"
	args := []string{
		"-graceful"}

	cmd := exec.Command(path, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.ExtraFiles = []*os.File{file}

	err := cmd.Start()
	if err != nil {
		log.Fatalf("gracefulRestart: Failed to launch, error: %v", err)
	}
}