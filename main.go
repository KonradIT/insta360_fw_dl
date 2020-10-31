package main

import (
	"errors"
	"fmt"
	"os"

	icmd "github.com/konradit/insta360_fw_dl/cmd"
	"github.com/konradit/insta360_fw_dl/pkg/insta360"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "camera",
	Short: "Camera to use",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("Specify camera")
		}
		cam, err := insta360.CameraGet(args[0])
		if err != nil {
			return err
		}
		return icmd.RunDownloader(cam)
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
