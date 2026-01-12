package docker

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func RunRemoteContainer(ip, keyPath, name string, hostport int, containerPort int, image string) error {
	cmd := exec.Command(
		"ssh",
		"-i", keyPath,
		"-o", "StrictHostKeyChecking=no",
		fmt.Sprintf("ec2-user@%s", ip),
		fmt.Sprintf(
			"docker rm -f %s || true && docker run -d --name %s -p %d:%d --restart always %s",
			name,
			name,
			hostport,
			containerPort,
			image,
		),
	)

	return cmd.Run()
}

func runSSH(ip, keyPath, command string) error {
	cmd := exec.Command(
		"ssh",
		"-i", keyPath,
		"-o", "StrictHostKeyChecking=no",
		fmt.Sprintf("ec2-user@%s", ip),
		command,
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func CloneRepo(ip, keyPath, repo string) error {
	return runSSH(ip, keyPath,
		fmt.Sprintf("rm -rf app && git clone %s app", repo),
	)
}

func BuildImage(ip, keyPath, image string) error {
	return runSSH(ip, keyPath,
		fmt.Sprintf("cd app && docker build -t %s .", image),
	)
}

func RunContainer(ip, keyPath, name string, port int, image string) error {
	return runSSH(ip, keyPath,
		fmt.Sprintf(
			"docker rm -f %s || true && docker run -d --name %s -p %d:%d --restart always %s",
			name, name, port, port, image,
		),
	)
}

func RemoveContainer(ip, keyPath, name string) error {
	return runSSH(ip, keyPath,
		fmt.Sprintf("docker rm -f %s || true", name),
	)
}

func HealthCheck(ip, keyPath string, port int) error {
	var lastErr error

	for i := 0; i < 5; i++ {
		err := runSSH(ip, keyPath,
			fmt.Sprintf("curl -sf http://localhost:%d", port),
		)
		if err == nil {
			return nil
		}

		lastErr = err
		time.Sleep(2 * time.Second)
	}

	return fmt.Errorf("health check failed after retries: %w", lastErr)
}
