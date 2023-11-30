package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/recall704/memory-limit-scheduler-plugin/pkg/plugins"

	"k8s.io/component-base/logs"
	"k8s.io/kubernetes/cmd/kube-scheduler/app"
)

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	command := app.NewSchedulerCommand(
		app.WithPlugin(plugins.Name, plugins.New),
	)

	logs.InitLogs()
	defer logs.FlushLogs()

	if err := command.Execute(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

}
