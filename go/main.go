package main

import (
	"fmt"
	"os/exec"
)

const (
	gen_cert_script = "./gen-cert.sh"
)

func main() {
	err := exec.Command("chmod", "777", gen_cert_script).Run()

	if err != nil {
		fmt.Printf("Failed to chmod: %s", err)
	}

	err = exec.Command("sh", "-c", gen_cert_script).Run()
	if err != nil {
		fmt.Printf("Failed to run: %s", err)
	}

	out, err := exec.Command("ls").Output()
	fmt.Printf("%s", out)
}

