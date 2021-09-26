package tasks

import (
	"log"
	"time"

	"github.com/robfig/cron"
	g "github.com/xanzy/go-gitlab"
	"go.dev.pztrn.name/periodicator/internal/gitlab"
)

// BaseTask is a base task structure.
type BaseTask struct {
	client *gitlab.Client

	projectID int
	title     string
	body      string
	tags      []string

	executionStartTimestamp time.Time
	cron                    string
	dueIn                   time.Duration
}

func (b *BaseTask) checkIfOpenedTaskExists(issues []*g.Issue) bool {
	b.log("Checking if opened task already exists...")

	var foundAndNotClosed bool

	for _, issue := range issues {
		if issue.Title == b.title && issue.ClosedAt == nil {
			foundAndNotClosed = true

			break
		}
	}

	return foundAndNotClosed
}

func (b *BaseTask) log(message string) {
	log.Println("Task '" + b.title + "': " + message)
}

// Run executes task procedure.
func (b *BaseTask) Run() {
	// Get similar tasks.
	issues, err := b.client.GetIssuesByTitle(b.projectID, b.title)
	if err != nil {
		b.log("Error while getting issues from Gitlab: " + err.Error())

		return
	}

	// Check if we have opened task. We should not create another task until already
	// created task is closed.

	if b.checkIfOpenedTaskExists(issues) {
		b.log("Found already existing task that isn't closed, won't create new task")

		return
	}

	b.log("No still opened tasks found, checking if we should create new one...")

	// Get latest task creation timestamp from Gitlab.
	lastTaskCreationTS := b.executionStartTimestamp

	for _, issue := range issues {
		if issue.ClosedAt != nil && issue.CreatedAt.After(lastTaskCreationTS) {
			lastTaskCreationTS = *issue.CreatedAt
		}
	}

	b.log("Last task creation timestamp: " + lastTaskCreationTS.String())

	// Set up cron job parser.
	cp := cron.NewParser(cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow)

	schedule, err := cp.Parse(b.cron)
	if err != nil {
		b.log("Failed to parse cron string: " + err.Error())

		return
	}

	// Figure out next task creation and deadline timestamps.
	nextCreationTS := schedule.Next(lastTaskCreationTS)

	// Check if task should be created and create if so.
	if nextCreationTS.Before(time.Now()) {
		// Deadlines should be calculated until first one will appear AFTER today.
		nextDeadlineTSAssumed := nextCreationTS.Add(b.dueIn)

		var nextDeadlineTS g.ISOTime

		for {
			if nextDeadlineTSAssumed.After(time.Now()) {
				nextDeadlineTS = g.ISOTime(nextDeadlineTSAssumed)

				break
			}

			nextDeadlineTSAssumed = nextDeadlineTSAssumed.Add(b.dueIn)
		}

		b.log("Found no opened tasks and task should be created, doing so. Task deadline: " + nextDeadlineTS.String())

		// nolint:exhaustivestruct
		err := b.client.CreateIssue(b.projectID, &g.CreateIssueOptions{
			Title:       &b.title,
			Description: &b.body,
			Labels:      b.tags,
			DueDate:     &nextDeadlineTS,
		})
		if err != nil {
			b.log("Failed to create task: " + err.Error())

			return
		}
	} else {
		b.log("Next task creation in future (" + nextCreationTS.String() + "), skipping for now")

		return
	}
}
