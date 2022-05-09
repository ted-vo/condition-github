package main

import (
	githubCondition "github.com/ted-vo/condition-github/pkg/condition"
	"github.com/ted-vo/semantic-release/v3/pkg/condition"
	"github.com/ted-vo/semantic-release/v3/pkg/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		CICondition: func() condition.CICondition {
			return &githubCondition.GitHubActions{}
		},
	})
}
