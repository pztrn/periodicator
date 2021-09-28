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

var showVersion = flag.Bool("version", false, "Show version information and exit")

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

	tasks.Process(c, cfg.Tasks)
}
