package cmd

import (
	"context"
	"time"

	"github.com/src-d/sourced-ce/cmd/sourced/compose"
)

type startCmd struct {
	Command `name:"start" short-description:"Start stopped containers" long-description:"Start stopped containers.\nThe containers must be initialized before with 'init'."`
}

func (c *startCmd) Execute(args []string) error {
	if err := compose.Run(context.Background(), "start"); err != nil {
		return err
	}

	return OpenUI(30 * time.Second)

}

func init() {
	rootCmd.AddCommand(&startCmd{})
}
