package main

import (
	"fmt"
	"os/exec"
	"os"
)

func main() {
	err := exec.Command("openssl", "genrsa", "-out", "ca.key", "2048").Run()
	fmt.Println("err: ", err)
	out, err := exec.Command("ls").Output()
	fmt.Printf("%s", out)
	os.Remove("./ca.key")
}

