package main

import (
	"log"

	"go.dev.pztrn.name/periodicator/internal/config"
	"go.dev.pztrn.name/periodicator/internal/gitlab"
	"go.dev.pztrn.name/periodicator/internal/tasks"
)

func main() {
	log.Println("Starting periodic tasks creator...")

	cfg := config.Parse()

	c := gitlab.NewGitlabClient(&cfg.Gitlab)

	tasks.Process(c, cfg.Tasks)
}
