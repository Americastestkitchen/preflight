package ruby

import (
	"bytes"
	"os/exec"
)

func CheckSyntax(filename string) (string, error) {
	var out bytes.Buffer

	cmd := exec.Command("ruby", "-c", filename)
	cmd.Stderr = &out
	err := cmd.Run()
	return out.String(), err
}
