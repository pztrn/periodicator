package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"go.dev.pztrn.name/periodicator/internal/config"
	"go.dev.pztrn.name/periodicator/internal/gitlab"
	"go.dev.pztrn.name/periodicator/internal/tasks"
)

var (
	showNextCreationTS = flag.Bool("show-next-creation-ts", false, "Show tasks next creation timestamps")
	showVersion        = flag.Bool("version", false, "Show version information and exit")
)

func main() {
	flag.Parse()

	if *showVersion {
		// nolint:forbidigo
		fmt.Println(config.Version)
		os.Exit(0)
	}

	log.Println("Starting periodic tasks creator, version " + config.Version + "...")

	cfg := config.Parse()

	c := gitlab.NewGitlabClient(&cfg.Gitlab)

	if *showNextCreationTS {
		tasks.PrintCreationTSes(c, cfg.Tasks)
		os.Exit(0)
	}

	tasks.Process(c, cfg.Tasks)
}
