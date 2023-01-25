package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	cmd := exec.Command("./child")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	fmt.Println("manager: Running interrupt child")
	cmd.Start()
	go func() {
		time.Sleep(1 * time.Second)
		cmd.Process.Signal(os.Interrupt)
	}()
	cmd.Wait()
	fmt.Println("manager: child interrupted")

	fmt.Print("\n\n\n")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	cmd = exec.CommandContext(ctx, "./child")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	fmt.Println("manager: Running kill child")
	cmd.Start()
	cmd.Wait()
	fmt.Println("manager: child killed")
}
