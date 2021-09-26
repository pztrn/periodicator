package tasks

import (
	"go.dev.pztrn.name/periodicator/internal/gitlab"
)

// Process processes passed tasks.
func Process(client *gitlab.Client, tasks []Config) {
	for _, task := range tasks {
		t := &BaseTask{
			client:                  client,
			projectID:               task.ProjectID,
			title:                   task.Title,
			body:                    task.Body,
			tags:                    task.Tags,
			executionStartTimestamp: task.ExecutionStart.GetTimestamp(),
			cron:                    task.Cron,
			dueIn:                   task.DueIn,
		}

		t.Run()
	}
}
