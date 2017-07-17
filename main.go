package main

import (
	"github.com/urfave/cli"
	"os"
	"log"
)

type Entry struct {
	Path    string
	Version int
	ModTime string
	URL     string
}

// Global variables
var versionNum, inputTime, dest, startDate, endDate string
var download bool
var server = "http://czbiohub-ncbi-tool-server.us-west-2.elasticbeanstalk.com"

func main() {
	log.SetOutput(os.Stderr)
	log.SetOutput(os.Stdout)
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Setup
	app := cli.NewApp()

	app.Commands = []cli.Command{
		{ // File commands
			Name:    "file",
			Action: fileSimple,
			Flags: 	[]cli.Flag{
				cli.StringFlag{
					Name:        "version-num",
					Destination: &versionNum,
				},
				cli.BoolFlag{
					Name:        "download",
					Destination: &download,
				},
				cli.StringFlag{
					Name:        "dest",
					Destination: &dest,
				},
			},
			Subcommands: []cli.Command{
				{
					Name:  "at-time",
					Action: fileAtTime,
					Flags: 	[]cli.Flag{
						cli.StringFlag{
							Name:        "input-time",
							Destination: &inputTime,
						},
						cli.BoolFlag{
							Name:        "download",
							Destination: &download,
						},
						cli.StringFlag{
							Name:        "dest",
							Destination: &dest,
						},
					},
				},
				{
					Name:  "history",
					Action: fileHistory,
				},
			},
		},
		{ // Directory commands
			Name:    "directory",
			Action: directorySimple,
			Flags: 	[]cli.Flag{
				cli.BoolFlag{
					Name:        "download",
					Destination: &download,
				},
				cli.StringFlag{
					Name:        "dest",
					Destination: &dest,
				},
			},
			Subcommands: []cli.Command{
				{
					Name:  "at-time",
					Action: directoryAtTime,
					Flags: 	[]cli.Flag{
						cli.StringFlag{
							Name:        "input-time",
							Destination: &inputTime,
						},
						cli.BoolFlag{
							Name:        "download",
							Destination: &download,
						},
						cli.StringFlag{
							Name:        "dest",
							Destination: &dest,
						},
					},
				},
				{
					Name:  "compare",
					Action: directoryCompare,
					Flags: 	[]cli.Flag{
						cli.StringFlag{
							Name:        "start-date",
							Destination: &startDate,
						},
						cli.StringFlag{
							Name:        "end-date",
							Destination: &endDate,
						},
					},
				},
			},
		},
	}
	app.Run(os.Args)
}