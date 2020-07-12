package cmd

import (
	"github.com/perun-network/perun-eth-demo/cmd/demo"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"time"
)

var startGanacheCmd = &cobra.Command{
	Use:   "ganache",
	Short: "Start ganache-cli and deploy contracts",
	Run:   startGanache,
}

func init() {
	rootCmd.AddCommand(startGanacheCmd)
}

func startGanache(_ *cobra.Command, _ []string) {
	cmd := exec.Command("ganache-cli", "--account=\"0x7d51a817ee07c3f28581c47a5072142193337fdca4d7911e58c5af2d03895d1a,100000000000000000000000\"", "--account=\"0x6aeeb7f09e757baa9d3935a042c3d0d46a2eda19e9b676283dce4eaf32e29dc9,100000000000000000000000\"", "-h", "0.0.0.0")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	select {
	case <-time.After(10 * time.Second):
		err := demo.DeployContracts()
		if err != nil {
			log.Fatal(err)
		}
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		select {
		case <-c:
			err := cmd.Process.Kill()
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
