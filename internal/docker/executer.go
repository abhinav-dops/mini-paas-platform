package docker

import (
	"fmt"
	"os/exec"
)

func RunContainer(name string, port int, image string) error {
	cmd := exec.Command(
		"docker", "run", "-d",
		"--name", name,
		"-p", fmt.Sprintf("%d:%d", port, port),
		"--restart", "always",
		image,
	)
	return cmd.Run()
}
