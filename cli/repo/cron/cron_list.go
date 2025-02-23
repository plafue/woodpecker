// Copyright 2022 Woodpecker Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cron

import (
	"context"
	"html/template"
	"os"

	"github.com/urfave/cli/v3"

	"go.woodpecker-ci.org/woodpecker/v3/cli/common"
	"go.woodpecker-ci.org/woodpecker/v3/cli/internal"
	"go.woodpecker-ci.org/woodpecker/v3/woodpecker-go/woodpecker"
)

var cronListCmd = &cli.Command{
	Name:      "ls",
	Usage:     "list cron jobs",
	ArgsUsage: "[repo-id|repo-full-name]",
	Action:    cronList,
	Flags: []cli.Flag{
		common.RepoFlag,
		common.FormatFlag(tmplCronList, true),
	},
}

func cronList(ctx context.Context, c *cli.Command) error {
	var (
		format           = c.String("format") + "\n"
		repoIDOrFullName = c.String("repository")
	)
	if repoIDOrFullName == "" {
		repoIDOrFullName = c.Args().First()
	}
	client, err := internal.NewClient(ctx, c)
	if err != nil {
		return err
	}
	repoID, err := internal.ParseRepo(client, repoIDOrFullName)
	if err != nil {
		return err
	}
	opt := woodpecker.CronListOptions{}
	list, err := client.CronList(repoID, opt)
	if err != nil {
		return err
	}
	tmpl, err := template.New("_").Parse(format)
	if err != nil {
		return err
	}
	for _, cron := range list {
		if err := tmpl.Execute(os.Stdout, cron); err != nil {
			return err
		}
	}
	return nil
}

// tTemplate for pipeline list information.
var tmplCronList = "\x1b[33m{{ .Name }} \x1b[0m" + `
ID: {{ .ID }}
Branch: {{ .Branch }}
Schedule: {{ .Schedule }}
NextExec: {{ .NextExec }}
`
