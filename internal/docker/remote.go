package docker

import (
	"fmt"
	"os/exec"
)

func RunRemoteContainer(ip, keyPath, name string, port int, image string) error {
	cmd := exec.Command(
		"ssh",
		"-i", keyPath,
		"-o", "StrictHostKeyChecking=no",
		fmt.Sprintf("ec2-user@%s", ip),
		fmt.Sprintf(
			"docker run -d --name %s -p %d:%d --restart always %s",
			name, port, port, image,
		),
	)

	return cmd.Run()
}
