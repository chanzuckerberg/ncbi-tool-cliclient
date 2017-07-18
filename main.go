package main

import (
	"github.com/urfave/cli"
	"log"
	"os"
)

// Entry is a file version entry
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

// main sets up and starts the command line client.
func main() {
	log.SetOutput(os.Stderr)
	log.SetOutput(os.Stdout)
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Start client
	app := cli.NewApp()
	app.Name = "NCBI Tool CLI Client"
	app.Usage = "For accessing files and directories on the NCBI tool server."
	app.Commands = setupCommands()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal("Error in running client. ", err)
	}
}

// setupCommands sets up the command line structure.
func setupCommands() []cli.Command {
	// Set up command line flags
	downloadFlag := cli.BoolFlag{
		Name:        "download",
		Destination: &download,
		Usage:       "include to download the files to local disk",
	}
	destFlag := cli.StringFlag{
		Name:        "dest",
		Destination: &dest,
		Usage:       "download destination on local disk",
	}
	inputTimeFlag := cli.StringFlag{
		Name:        "input-time",
		Destination: &inputTime,
		Usage:       "input time for 'at time' requests. Ex: 2017-07-07T00:06:12",
	}
	versionNumFlag := cli.StringFlag{
		Name:        "version-num",
		Destination: &versionNum,
		Usage:       "version number for file requests",
	}
	startDateFlag := cli.StringFlag{
		Name:        "start-date",
		Destination: &startDate,
		Usage:       "start date for directory diff comparisons",
	}
	endDateFlag := cli.StringFlag{
		Name:        "end-date",
		Destination: &endDate,
		Usage:       "end date for directory diff comparisons",
	}
	atTimeFlags := []cli.Flag{
		inputTimeFlag,
		downloadFlag,
		destFlag,
	}

	// Set up commands and sub-commands
	fileCommands := cli.Command{
		Name:   "file",
		Action: fileSimple,
		Usage:  "file actions",
		Flags: []cli.Flag{
			versionNumFlag,
			downloadFlag,
			destFlag,
		},
		Subcommands: []cli.Command{
			{
				Name:   "at-time",
				Action: fileAtTime,
				Usage:  "get a file version at or before a point in time",
				Flags:  atTimeFlags,
			},
			{
				Name:   "history",
				Usage:  "get the version history of a file",
				Action: fileHistory,
			},
		},
	}
	directoryCommands := cli.Command{
		Name:   "directory",
		Action: directorySimple,
		Usage:  "directory actions",
		Flags: []cli.Flag{
			downloadFlag,
			destFlag,
		},
		Subcommands: []cli.Command{
			{
				Name:   "at-time",
				Action: directoryAtTime,
				Usage:  "get a directory state at or before a point in time",
				Flags:  atTimeFlags,
			},
			{
				Name:   "compare",
				Action: directoryCompare,
				Usage:  "compare a directory state across a start and end date",
				Flags: []cli.Flag{
					startDateFlag,
					endDateFlag,
				},
			},
		},
	}
	return []cli.Command{fileCommands, directoryCommands}
}