package main

import (
	"context"
	"os"
	"sync"

	"todo_graphql/config"
	"todo_graphql/instance"
	"todo_graphql/runner"

	"github.com/urfave/cli"
)

func main() {
	config.Load()
	instance.Init()
	defer instance.Destroy()


	clientApp := cli.NewApp()
	clientApp.Name = "todo-list"
	clientApp.Version = "0.0.1"
	clientApp.Commands = []cli.Command{
		{
			Name:  "start",
			Usage: "Start the service",
			Action: func(c *cli.Context) error {
				ctx := context.Background()

				var wg sync.WaitGroup
				wg.Add(1)

				go runner.NewAPI().Go(ctx, &wg)

				wg.Wait()
				return nil
			},
		},
	}
	if err := clientApp.Run(os.Args); err != nil {
		panic(err)
	}
}
