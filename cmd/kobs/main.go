package main

import (
	"github.com/kobsio/kobs/cmd/kobs/cluster"
	"github.com/kobsio/kobs/cmd/kobs/hub"
	"github.com/kobsio/kobs/cmd/kobs/version"
	"github.com/kobsio/kobs/cmd/kobs/watcher"
	"github.com/kobsio/kobs/pkg/plugins"

	"github.com/kobsio/kobs/pkg/plugins/flux"
	"github.com/kobsio/kobs/pkg/plugins/grafana"
	"github.com/kobsio/kobs/pkg/plugins/harbor"
	"github.com/kobsio/kobs/pkg/plugins/helm"
	"github.com/kobsio/kobs/pkg/plugins/jaeger"
	"github.com/kobsio/kobs/pkg/plugins/opsgenie"
	"github.com/kobsio/kobs/pkg/plugins/prometheus"
	"github.com/kobsio/kobs/pkg/plugins/rss"
	"github.com/kobsio/kobs/pkg/plugins/signalsciences"
	"github.com/kobsio/kobs/pkg/plugins/sonarqube"

	"github.com/alecthomas/kong"
)

var cli struct {
	Hub     hub.Cmd     `cmd:"hub" help:"Start the hub."`
	Watcher watcher.Cmd `cmd:"watcher" help:"Start the watcher."`
	Cluster cluster.Cmd `cmd:"cluster" help:"Start the cluster."`
	Version version.Cmd `cmd:"version" help:"Show version information."`
}

func main() {
	ctx := kong.Parse(&cli)

	registeredPlugins := []plugins.Plugin{
		flux.New(),
		grafana.New(),
		harbor.New(),
		helm.New(),
		jaeger.New(),
		opsgenie.New(),
		prometheus.New(),
		rss.New(),
		signalsciences.New(),
		sonarqube.New(),
	}

	err := ctx.Run(registeredPlugins)
	ctx.FatalIfErrorf(err)
}
