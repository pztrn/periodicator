package gitlab

import (
	"errors"
	"log"
	"net/http"

	"github.com/xanzy/go-gitlab"
)

// Client is a Gitlab's client controlling structure.
type Client struct {
	client *gitlab.Client
	config *Config
}

// NewGitlabClient creates new Gitlab's client controlling structure.
func NewGitlabClient(cfg *Config) *Client {
	// nolint:exhaustruct
	c := &Client{
		config: cfg,
	}
	c.initialize()

	return c
}

// CreateIssue creates issue in designated project (by ID) using passed options.
// Returns error if something went wrong.
func (c *Client) CreateIssue(projectID int, options *gitlab.CreateIssueOptions) error {
	_, _, err := c.client.Issues.CreateIssue(projectID, options)

	// ToDo: fix it!
	// nolint:wrapcheck
	return err
}

// GetClient returns underlying Gitlab's client.
func (c *Client) GetClient() *gitlab.Client {
	return c.client
}

// GetIssuesByTitle returns list of issues that matches passed title in specific
// project.
func (c *Client) GetIssuesByTitle(projectID int, title string) ([]*gitlab.Issue, error) {
	// nolint:exhaustruct
	issues, resp, err := c.client.Issues.ListProjectIssues(projectID, &gitlab.ListProjectIssuesOptions{
		ListOptions: gitlab.ListOptions{
			PerPage: 1000,
			Page:    1,
		},
		Search: &title,
	})
	if err != nil {
		log.Println("Failed to execute HTTP request to Gitlab: " + err.Error())

		// ToDo: fix it!
		// nolint:wrapcheck
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		log.Println("Got status: " + resp.Status)

		// ToDo: fix it!
		// nolint:goerr113
		return nil, errors.New("not HTTP 200 from Gitlab")
	}

	return issues, nil
}

func (c *Client) initialize() {
	client, err := gitlab.NewClient(
		c.config.Token,
		gitlab.WithBaseURL(c.config.BaseURL),
	)
	if err != nil {
		panic("Failed to create Gitlab Client: " + err.Error())
	}

	c.client = client
}
