package control

import (
	"os"
	"os/exec"

	"github.com/codegangsta/cli"
	"github.com/rancher/os/log"
)

func udevSettleAction(c *cli.Context) {
	if err := UdevSettle(); err != nil {
		log.Fatal(err)
	}
}

func UdevSettle() error {
	cmd := exec.Command("udevd", "--daemon")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}

	cmd = exec.Command("udevadm", "trigger", "--action=add")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}

	cmd = exec.Command("udevadm", "settle")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
