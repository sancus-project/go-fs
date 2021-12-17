package main

import (
	"errors"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"

	"go.sancus.dev/fs/dirfs"
	"go.sancus.dev/fs/fuse"
)

var mountCmd = &cobra.Command{
	Use: "mount [flags] <back> <front>",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			return errors.New("Not enough arguments")
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {

		// dirfs
		fsys, err := dirfs.New(args[0])
		if err != nil {
			return err
		}
		defer fsys.Close()

		// fuse daemon
		srv, err := fuse.New(fsys, args[1])
		if err != nil {
			return err
		}
		defer srv.Close()

		// watch signals
		go func() {
			sig := make(chan os.Signal, 1)
			signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)

			for signum := range sig {
				switch signum {
				case syscall.SIGHUP:
					// reload
					if err := srv.Reload(); err != nil {
						log.Println("Failed to reload:", err)
					}
				case syscall.SIGINT, syscall.SIGTERM:
					// terminate
					log.Println("Terminating...")
					if err := srv.Unmount(); err != nil {
						log.Println("Failed to terminate:", err)
						continue
					}

					return
				}
			}
		}()

		return srv.Serve()
	},
}

func init() {
	rootCmd.AddCommand(mountCmd)
}
