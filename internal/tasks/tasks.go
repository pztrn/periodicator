package tasks

import (
	"go.dev.pztrn.name/periodicator/internal/gitlab"
)

// PrintCreationTSes prints tasks creation timestamps.
func PrintCreationTSes(client *gitlab.Client, tasks []Config) {
	for _, task := range tasks {
		taskData := &BaseTask{
			client:                  client,
			projectID:               task.ProjectID,
			title:                   task.Title,
			body:                    task.Body,
			tags:                    task.Tags,
			executionStartTimestamp: task.ExecutionStart.GetTimestamp(),
			cron:                    task.Cron,
			dueIn:                   task.DueIn,
		}

		// Get similar tasks.
		// ToDo: refactor?
		issues, err := taskData.getIssues()
		if err != nil {
			panic("Error while getting issues from Gitlab: " + err.Error())
		}

		taskData.log(taskData.getNextCreationTimestamp(taskData.getLastCreationTimestamp(issues)).String())
	}
}

// Process processes passed tasks.
func Process(client *gitlab.Client, tasks []Config) {
	for _, task := range tasks {
		taskData := &BaseTask{
			client:                  client,
			projectID:               task.ProjectID,
			title:                   task.Title,
			body:                    task.Body,
			tags:                    task.Tags,
			executionStartTimestamp: task.ExecutionStart.GetTimestamp(),
			cron:                    task.Cron,
			dueIn:                   task.DueIn,
		}

		taskData.Run()
	}
}
