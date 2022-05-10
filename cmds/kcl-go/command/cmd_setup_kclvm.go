// Copyright 2021 The KCL Authors. All rights reserved.

package command

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"

	"kusionstack.io/kclvm-go/scripts"
)

var cmdSetupKclvmFlags = []cli.Flag{
	&cli.BoolFlag{
		Name:  "all",
		Usage: "setup kclvm for all platform",
	},
	&cli.StringFlag{
		Name:  "triple",
		Usage: "set kclvm triple",
		Value: scripts.DefaultKclvmTriple,
	},
	&cli.StringFlag{
		Name:  "outdir",
		Usage: "set kclvm output dir",
		Value: "_build",
	},
	&cli.StringSliceFlag{
		Name:  "mirrors",
		Usage: "set mirror list",
	},
}

func NewSetpupKclvmCmd() *cli.Command {
	return &cli.Command{
		Hidden: false,
		Name:   "setup-kclvm",
		Usage:  "setup kclvm command",
		Flags:  cmdSetupKclvmFlags,
		Action: func(c *cli.Context) error {
			// go run ./cmds/kcl-go/ setup-kclvm --triple=kclvm-ubuntu

			if c.Args().Len() > 0 {
				fmt.Println("ERR: invalid arguments:", c.Args().Slice())
				cli.ShowCommandHelpAndExit(c, "setup-kclvm", 0)
			}

			all := c.Bool("all")
			triple := c.String("triple")
			outdir := c.String("outdir")
			mirrors := c.StringSlice("mirrors")

			if len(mirrors) != 0 {
				scripts.KclvmDownloadUrlBase_mirrors = append(scripts.KclvmDownloadUrlBase_mirrors, mirrors...)
			}

			if all {
				err := scripts.SetupKclvmAll(outdir)
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				return nil
			}

			if triple == "" || outdir == "" {
				cli.ShowCommandHelpAndExit(c, "setup-kclvm", 0)
			}

			if triple != scripts.DefaultKclvmTriple {
				if outdir == "" || outdir == "_"+scripts.DefaultKclvmTriple+"-root_" {
					outdir = "_" + triple + "-root_"
				}
			}

			scripts.DefaultKclvmTriple = triple

			err := scripts.SetupKclvm(outdir)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			return nil
		},
	}
}
