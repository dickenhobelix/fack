package main

// Importing the required packages
import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/alecthomas/kong"
	"github.com/mergestat/timediff"
)

type CLI struct {
	Since    string `short:"s" name:"since" help:"check if file has been modified in timespan (e.g. 10h12s)"`
	Detailed bool   `short:"d" name:"detailed" help:"make output more verbose. is ignored if --silent is used"`
	Silent   bool   `name:"non-interactive" help:"no output to shell"`
	Filename string `arg:"" name:"file" help:"file to be checked"`
}

func print_last_modified(filename string, modificationTime time.Time, modified_readable string, silent bool, detailed bool) {
	//early exit on silent mode
	if silent {
		os.Exit(0)
	}

	fmt.Print(filename, " was last modified ")

	if detailed {
		fmt.Print(modificationTime, " (", modified_readable, ")\n")
	} else {
		fmt.Print(modified_readable, "\n")
	}
	os.Exit(0)
}

func print_is_current(filename string, duration time.Duration, modificationTime time.Time, modified_readable string, silent bool, detailed bool) {
	isCurrent := modificationTime.Add(duration).After(time.Now())

	//early exit with result in silent mode
	if silent {
		if isCurrent {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}

	if isCurrent {
		fmt.Print(filename, " is current")
	} else {
		print(filename, " is outdated")
	}

	if detailed {
		fmt.Print(" (last modified ", modified_readable, ")\n")
	} else {
		fmt.Print("\n")
	}

	//return success on current files, failure on outdated files
	if isCurrent {
		os.Exit(0)
	} else {
		os.Exit(1)
	}
}

func main() {
	var cli CLI

	args := kong.Parse(&cli,
		kong.Name("fack"),
		kong.Description("File Age Checker - checks file's last-modified date"),
		kong.UsageOnError(),
	)

	args.FatalIfErrorf(args.Error)

	// Get the fileinfo
	fileInfo, err := os.Stat(cli.Filename)

	// Checks for the error
	if err != nil {
		log.Fatal(err)
	}

	modificationTime := fileInfo.ModTime()
	modified_readable := timediff.TimeDiff(modificationTime)

	if cli.Since != "" {
		duration, err := time.ParseDuration(cli.Since)

		if err != nil {
			if cli.Silent {
				os.Exit(-2)
			} else {
				log.Fatal(err)
			}
		}

		print_is_current(fileInfo.Name(), duration, modificationTime, modified_readable, cli.Silent, cli.Detailed)

	} else {
		//no explicit since given, just print last modified
		print_last_modified(fileInfo.Name(), modificationTime, modified_readable, cli.Silent, cli.Detailed)
	}

}
