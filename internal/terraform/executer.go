package terraform

import "os/exec"

func Init(path string) error {
	cmd := exec.Command("terraform", "init")
	cmd.Dir = path
	return cmd.Run()
}

func Apply(path string) error {
	cmd := exec.Command("terraform", "apply", "-auto-approve")
	cmd.Dir = path
	return cmd.Run()
}

func Destroy(path string) error {
	cmd := exec.Command("terraform", "destroy", "-auto-approve")
	cmd.Dir = path
	return cmd.Run()
}

func Output(path string) (string, error) {
	cmd := exec.Command("terraform", "output", "-raw", "public_ip")
	cmd.Dir = path
	out, err := cmd.Output()
	return string(out), err
}
