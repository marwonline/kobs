package cluster

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/kobsio/kobs/pkg/cluster/api"
	"github.com/kobsio/kobs/pkg/cluster/kubernetes"
	clusterPlugins "github.com/kobsio/kobs/pkg/cluster/plugins"
	"github.com/kobsio/kobs/pkg/instrument/log"
	"github.com/kobsio/kobs/pkg/instrument/metrics"
	"github.com/kobsio/kobs/pkg/instrument/tracer"
	"github.com/kobsio/kobs/pkg/plugins"
	"github.com/kobsio/kobs/pkg/plugins/plugin"
	"github.com/kobsio/kobs/pkg/utils/config"
	"go.uber.org/zap"
)

type Cmd struct {
	Config string `env:"KOBS_CONFIG" default:"config.yaml" help:"The path to the configuration file for the cluster."`

	Log     log.Config     `json:"log" embed:"" prefix:"log." envprefix:"KOBS_LOG_"`
	Tracer  tracer.Config  `json:"tracer" embed:"" prefix:"tracer." envprefix:"KOBS_TRACER_"`
	Metrics metrics.Config `json:"metrics" embed:"" prefix:"metrics." envprefix:"KOBS_METRICS_"`

	Cluster struct {
		Kubernetes kubernetes.Config `json:"kubernetes" embed:"" prefix:"kubernetes." envprefix:"KUBERNETES_"`
		API        api.Config        `json:"api" embed:"" prefix:"api." envprefix:"API_"`
		Plugins    []plugin.Instance `json:"plugins" kong:"-"`
	} `json:"cluster" embed:"" prefix:"cluster." envprefix:"KOBS_CLUSTER_"`
}

func (r *Cmd) Run(plugins []plugins.Plugin) error {
	cfg, err := config.Load(r.Config, *r)
	if err != nil {
		log.Error(nil, "Could not load configuration", zap.Error(err))
		return err
	}

	logger, err := log.Setup(cfg.Log)
	if err != nil {
		return err
	}
	defer logger.Sync()

	tracerClient, err := tracer.Setup(cfg.Tracer)
	if err != nil {
		return err
	}
	if tracerClient != nil {
		defer tracerClient.Shutdown()
	}

	metricsServer := metrics.New(cfg.Metrics)
	go metricsServer.Start()
	defer metricsServer.Stop()

	kubernetesClient, err := kubernetes.NewClient(cfg.Cluster.Kubernetes)
	if err != nil {
		log.Error(nil, "Could not create Kubernetes client", zap.Error(err))
		return err
	}

	pluginsClient, err := clusterPlugins.NewClient(plugins, cfg.Cluster.Plugins, kubernetesClient)
	if err != nil {
		log.Error(nil, "Could not create plugins client", zap.Error(err))
		return err
	}

	apiServer, err := api.New(cfg.Cluster.API, kubernetesClient, pluginsClient)
	if err != nil {
		log.Error(nil, "Could not create client server", zap.Error(err))
		return err
	}
	go apiServer.Start()

	// All components should be terminated gracefully. For that we are listen for the SIGINT and SIGTERM signals and try
	// to gracefully shutdown the started kobs components. This ensures that established connections or tasks are not
	// interrupted.
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGTERM)

	log.Debug(nil, "Start listining for SIGINT and SIGTERM signal")
	<-done
	log.Info(nil, "Shutdown kobs client...")

	apiServer.Stop()

	log.Info(nil, "Shutdown is done")

	return nil
}
