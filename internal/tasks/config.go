package tasks

import "time"

// Config is a task's configuration as should be defined in configuration file.
// nolint:tagliatelle
type Config struct {
	ExecutionStart TaskStartTime `yaml:"execution_start"`
	Title          string        `yaml:"title"`
	Body           string        `yaml:"body"`
	Cron           string        `yaml:"cron"`
	Tags           []string      `yaml:"tags"`
	ProjectID      int           `yaml:"project_id"`
	DueIn          time.Duration `yaml:"due_in"`
}

// TaskStartTime holds task's start time for next creation timestamp calculation.
type TaskStartTime struct {
	ts time.Time
}

func (tts *TaskStartTime) GetTimestamp() time.Time {
	return tts.ts
}

func (tts *TaskStartTime) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var timeData string

	if err := unmarshal(&timeData); err != nil {
		return err
	}

	timeField, err := time.Parse("2006-01-02 15:04:05", timeData)
	if err != nil {
		// ToDo: fix it!
		// nolint:wrapcheck
		return err
	}

	tts.ts = timeField

	return nil
}
