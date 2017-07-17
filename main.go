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

// main sets up the command line client structure.
func main() {
	log.SetOutput(os.Stderr)
	log.SetOutput(os.Stdout)
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Set up command line flags
	downloadFlag := cli.BoolFlag{
		Name:        "download",
		Destination: &download,
	}
	destFlag := cli.StringFlag{
		Name:        "dest",
		Destination: &dest,
	}
	inputTimeFlag := cli.StringFlag{
		Name:        "input-time",
		Destination: &inputTime,
	}
	versionNumFlag := cli.StringFlag{
		Name:        "version-num",
		Destination: &versionNum,
	}
	startDateFlag := cli.StringFlag{
		Name:        "start-date",
		Destination: &startDate,
	}
	endDateFlag := cli.StringFlag{
		Name:        "end-date",
		Destination: &endDate,
	}
	atTimeFlags := []cli.Flag{
		inputTimeFlag,
		downloadFlag,
		destFlag,
	}

	fileCommands := cli.Command{
		Name:   "file",
		Action: fileSimple,
		Flags: []cli.Flag{
			versionNumFlag,
			downloadFlag,
			destFlag,
		},
		Subcommands: []cli.Command{
			{
				Name:   "at-time",
				Action: fileAtTime,
				Flags:  atTimeFlags,
			},
			{
				Name:   "history",
				Action: fileHistory,
			},
		},
	}

	directoryCommands := cli.Command{
		Name:   "directory",
		Action: directorySimple,
		Flags: []cli.Flag{
			downloadFlag,
			destFlag,
		},
		Subcommands: []cli.Command{
			{
				Name:   "at-time",
				Action: directoryAtTime,
				Flags:  atTimeFlags,
			},
			{
				Name:   "compare",
				Action: directoryCompare,
				Flags: []cli.Flag{
					startDateFlag,
					endDateFlag,
				},
			},
		},
	}

	app := cli.NewApp()
	app.Name = "NCBI Tool CLI Client"
	app.Usage = "For accessing the NCBI tool server"
	app.Commands = []cli.Command{fileCommands, directoryCommands}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal("Error in running client. ", err)
	}
}
