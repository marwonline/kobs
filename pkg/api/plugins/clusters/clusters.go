package clusters

import (
	"context"
	"fmt"
	"os"
	"sort"
	"time"

	applicationProto "github.com/kobsio/kobs/pkg/api/plugins/application/proto"
	"github.com/kobsio/kobs/pkg/api/plugins/clusters/cluster"
	"github.com/kobsio/kobs/pkg/api/plugins/clusters/proto"
	clustersProto "github.com/kobsio/kobs/pkg/api/plugins/clusters/proto"
	"github.com/kobsio/kobs/pkg/api/plugins/clusters/provider"

	"github.com/sirupsen/logrus"
	flag "github.com/spf13/pflag"
)

var (
	log                     = logrus.WithFields(logrus.Fields{"package": "clusters"})
	cacheDurationNamespaces string
	cacheDurationTopology   string
)

// init is used to define all command-line flags for the clusters package. Currently this is only the cache duration,
// which is used to cache the namespaces for a cluster.
func init() {
	defaultCacheDurationNamespaces := "5m"
	if os.Getenv("KOBS_CLUSTERS_CACHE_DURATION_NAMESPACES") != "" {
		defaultCacheDurationNamespaces = os.Getenv("KOBS_CLUSTERS_CACHE_DURATION_NAMESPACES")
	}

	defaultCacheDurationTopology := "60m"
	if os.Getenv("KOBS_CLUSTERS_CACHE_DURATION_TOPOLOGY") != "" {
		defaultCacheDurationTopology = os.Getenv("KOBS_CLUSTERS_CACHE_DURATION_TOPOLOGY")
	}

	flag.StringVar(&cacheDurationNamespaces, "clusters.cache-duration.namespaces", defaultCacheDurationNamespaces, "The duration, for how long requests to get the list of namespaces should be cached.")
	flag.StringVar(&cacheDurationTopology, "clusters.cache-duration.topology", defaultCacheDurationTopology, "The duration, for how long the topology data should be cached.")
}

// Config is the configuration required to load all clusters.
type Config struct {
	Providers []provider.Config `yaml:"providers"`
}

// Clusters contains all fields and methods to interact with the configured Kubernetes clusters. It must implement the
// Clusters service from the protocol buffers definition.
type Clusters struct {
	clustersProto.UnimplementedClustersServer
	clusters []*cluster.Cluster
	edges    []*clustersProto.Edge
	nodes    []*clustersProto.Node
}

func (c *Clusters) getCluster(name string) *cluster.Cluster {
	for _, cl := range c.clusters {
		if cl.GetName() == name {
			return cl
		}
	}

	return nil
}

// GetClusters returns all loaded Kubernetes clusters.
// We are not returning the complete cluster structure. Instead we are returning just the names of the clusters. We are
// also sorting the clusters alphabetically, to improve the user experience in the frontend.
// NOTE: Maybe we can also save the cluster names slice, since the name of a cluster couldn't change during runtime.
func (c *Clusters) GetClusters(ctx context.Context, getClustersRequest *clustersProto.GetClustersRequest) (*clustersProto.GetClustersResponse, error) {
	log.Tracef("GetClusters")

	var clusters []string

	for _, cluster := range c.clusters {
		clusters = append(clusters, cluster.GetName())
	}

	sort.Slice(clusters, func(i, j int) bool {
		return clusters[i] < clusters[j]
	})

	log.WithFields(logrus.Fields{"clusters": clusters}).Tracef("GetClusters")

	return &clustersProto.GetClustersResponse{
		Clusters: clusters,
	}, nil
}

// GetNamespaces returns all namespaces for the given clusters.
// As we did it for the clusters, we are also just returning the names of all namespaces. After we retrieved all
// namespaces we have to depulicate them, so that our frontend logic can handle them properly. We are also sorting the
// namespaces alphabetically.
func (c *Clusters) GetNamespaces(ctx context.Context, getNamespacesRequest *clustersProto.GetNamespacesRequest) (*clustersProto.GetNamespacesResponse, error) {
	log.WithFields(logrus.Fields{"clusters": getNamespacesRequest.Clusters}).Tracef("GetNamespaces")

	var namespaces []string

	for _, clusterName := range getNamespacesRequest.Clusters {
		cluster := c.getCluster(clusterName)
		if cluster == nil {
			return nil, fmt.Errorf("invalid cluster name")
		}

		clusterNamespaces, err := cluster.GetNamespaces(ctx)
		if err != nil {
			return nil, err
		}

		if clusterNamespaces != nil {
			namespaces = append(namespaces, clusterNamespaces...)
		}

	}

	keys := make(map[string]bool)
	uniqueNamespaces := []string{}
	for _, namespace := range namespaces {
		if _, value := keys[namespace]; !value {
			keys[namespace] = true
			uniqueNamespaces = append(uniqueNamespaces, namespace)
		}
	}

	sort.Slice(uniqueNamespaces, func(i, j int) bool {
		return uniqueNamespaces[i] < uniqueNamespaces[j]
	})

	log.WithFields(logrus.Fields{"namespaces": uniqueNamespaces}).Tracef("GetNamespaces")

	return &clustersProto.GetNamespacesResponse{
		Namespaces: uniqueNamespaces,
	}, nil
}

// GetCRDs returns all CRDs for all clusters.
// Instead of only returning the CRDs for a list of specified clusters, we return all CRDs, so that we only have to call
// this function once from the React app. The CRDs form all loaded clusters are merged and then deduplicated.
func (c *Clusters) GetCRDs(ctx context.Context, getCRDsRequest *clustersProto.GetCRDsRequest) (*clustersProto.GetCRDsResponse, error) {
	log.Tracef("GetCRDs")
	var crds []*clustersProto.CRD

	for _, cluster := range c.clusters {
		crds = append(crds, cluster.GetCRDs()...)
	}

	keys := make(map[string]bool)
	uniqueCRDs := []*clustersProto.CRD{}
	for _, crd := range crds {
		if _, value := keys[crd.Resource+"."+crd.Path]; !value {
			keys[crd.Resource+"."+crd.Path] = true
			uniqueCRDs = append(uniqueCRDs, crd)
		}
	}

	log.WithFields(logrus.Fields{"count": len(uniqueCRDs)}).Tracef("GetCRDs")

	return &clustersProto.GetCRDsResponse{
		Crds: uniqueCRDs,
	}, nil
}

// GetResources returns a list of resources for the given clusters and namespaces.
// To generate this list, we loop over every cluster and namespace and try to get the resources for this. A resource is
// identified by it's Kubernetes API path and name.
func (c *Clusters) GetResources(ctx context.Context, getResourcesRequest *clustersProto.GetResourcesRequest) (*clustersProto.GetResourcesResponse, error) {
	log.WithFields(logrus.Fields{"clusters": getResourcesRequest.Clusters, "namespaces": getResourcesRequest.Namespaces, "resource": getResourcesRequest.Resource, "path": getResourcesRequest.Path, "paramName": getResourcesRequest.ParamName, "param": getResourcesRequest.Param}).Tracef("GetResources")

	var resources []*proto.Resources

	for _, clusterName := range getResourcesRequest.Clusters {
		cluster := c.getCluster(clusterName)
		if cluster == nil {
			return nil, fmt.Errorf("invalid cluster name")
		}

		if getResourcesRequest.Namespaces == nil {
			list, err := cluster.GetResources(ctx, "", getResourcesRequest.Path, getResourcesRequest.Resource, getResourcesRequest.ParamName, getResourcesRequest.Param)
			if err != nil {
				return nil, err
			}

			resources = append(resources, &clustersProto.Resources{
				Cluster:      clusterName,
				Namespace:    "",
				ResourceList: list,
			})
		} else {
			for _, namespace := range getResourcesRequest.Namespaces {
				list, err := cluster.GetResources(ctx, namespace, getResourcesRequest.Path, getResourcesRequest.Resource, getResourcesRequest.ParamName, getResourcesRequest.Param)
				if err != nil {
					return nil, err
				}

				resources = append(resources, &clustersProto.Resources{
					Cluster:      clusterName,
					Namespace:    namespace,
					ResourceList: list,
				})
			}
		}

	}

	log.WithFields(logrus.Fields{"count": len(resources)}).Tracef("GetResources")

	return &clustersProto.GetResourcesResponse{
		Resources: resources,
	}, nil
}

// GetLogs returns the log line for the given pod, which is identified by the cluster, namespace and name.
func (c *Clusters) GetLogs(ctx context.Context, getLogsRequest *clustersProto.GetLogsRequest) (*clustersProto.GetLogsResponse, error) {
	log.WithFields(logrus.Fields{"cluster": getLogsRequest.Cluster, "namespace": getLogsRequest.Namespace, "pod": getLogsRequest.Name, "container": getLogsRequest.Container, "regex": getLogsRequest.Regex, "since": getLogsRequest.Since, "previous": getLogsRequest.Previous}).Tracef("GetLogs")

	cluster := c.getCluster(getLogsRequest.Cluster)
	if cluster == nil {
		return nil, fmt.Errorf("invalid cluster name")
	}

	logs, err := cluster.GetLogs(ctx, getLogsRequest.Namespace, getLogsRequest.Name, getLogsRequest.Container, getLogsRequest.Regex, getLogsRequest.Since, getLogsRequest.Previous)
	if err != nil {
		return nil, err
	}

	log.WithFields(logrus.Fields{"count": len(logs)}).Tracef("GetLogs")

	return &clustersProto.GetLogsResponse{
		Logs: logs,
	}, nil
}

// GetApplications returns a list of applications for the given clusters and namespaces.
// To generate this list, we loop over every cluster and namespace and try to get the applications for this.
func (c *Clusters) GetApplications(ctx context.Context, getApplicationsRequest *clustersProto.GetApplicationsRequest) (*clustersProto.GetApplicationsResponse, error) {
	log.WithFields(logrus.Fields{"clusters": getApplicationsRequest.Clusters, "namespaces": getApplicationsRequest.Namespaces}).Tracef("GetApplications")

	var applications []*applicationProto.Application

	for _, clusterName := range getApplicationsRequest.Clusters {
		cluster := c.getCluster(clusterName)
		if cluster == nil {
			return nil, fmt.Errorf("invalid cluster name")
		}

		for _, namespace := range getApplicationsRequest.Namespaces {
			list, err := cluster.GetApplications(ctx, namespace)
			if err != nil {
				return nil, err
			}

			applications = append(applications, list...)
		}
	}

	log.WithFields(logrus.Fields{"count": len(applications)}).Tracef("GetApplications")

	return &clustersProto.GetApplicationsResponse{
		Applications: applications,
	}, nil
}

// GetApplication returns a single application with the given name in the given cluster and namespace. If there isn't,
// such an application an error is returned.
func (c *Clusters) GetApplication(ctx context.Context, getApplicationRequest *clustersProto.GetApplicationRequest) (*clustersProto.GetApplicationResponse, error) {
	log.WithFields(logrus.Fields{"cluster": getApplicationRequest.Cluster, "namespace": getApplicationRequest.Namespace, "name": getApplicationRequest.Name}).Tracef("GetApplication")

	cluster := c.getCluster(getApplicationRequest.Cluster)
	if cluster == nil {
		return nil, fmt.Errorf("invalid cluster name")
	}

	application, err := cluster.GetApplication(ctx, getApplicationRequest.Namespace, getApplicationRequest.Name)
	if err != nil {
		return nil, err
	}

	log.WithFields(logrus.Fields{"application": application}).Tracef("GetApplication")

	return &clustersProto.GetApplicationResponse{
		Application: application,
	}, nil
}

// GetApplicationsTopology returns the topology for the given list of clusters and namespaces. We add an additional node
// for each cluster and namespace. These nodes are used to group the applications by the cluster and namespace.
func (c *Clusters) GetApplicationsTopology(ctx context.Context, getApplicationsTopologyRequest *clustersProto.GetApplicationsTopologyRequest) (*clustersProto.GetApplicationsTopologyResponse, error) {
	var edges []*clustersProto.Edge
	var nodes []*clustersProto.Node

	for _, clusterName := range getApplicationsTopologyRequest.Clusters {
		nodes = append(nodes, &clustersProto.Node{
			Id:        clusterName,
			Label:     clusterName,
			Type:      "cluster",
			Parent:    "",
			Cluster:   clusterName,
			Namespace: "",
			Name:      "",
		})

		for _, namespace := range getApplicationsTopologyRequest.Namespaces {
			nodes = append(nodes, &clustersProto.Node{
				Id:        clusterName + "-" + namespace,
				Label:     namespace,
				Type:      "namespace",
				Parent:    clusterName,
				Cluster:   clusterName,
				Namespace: namespace,
				Name:      "",
			})

			for _, edge := range c.edges {
				if (edge.SourceCluster == clusterName && edge.SourceNamespace == namespace) || (edge.TargetCluster == clusterName && edge.TargetNamespace == namespace) {
					edges = appendEdgeIfMissing(edges, edge)
				}
			}
		}
	}

	for _, edge := range edges {
		for _, node := range c.nodes {
			if node.Id == edge.Source || node.Id == edge.Target {
				nodes = appendNodeIfMissing(nodes, node)
			}
		}
	}

	return &clustersProto.GetApplicationsTopologyResponse{
		Edges: edges,
		Nodes: nodes,
	}, nil
}

// Load loads all clusters for the given configuration.
// The clusters can be retrieved from different providers. Currently we are supporting incluster configuration and
// kubeconfig files. In the future it is planning to directly support GKE, EKS, AKS, etc.
func Load(config Config) (*Clusters, error) {
	var clusters []*cluster.Cluster

	for _, p := range config.Providers {
		providerClusters, err := provider.GetClusters(&p)
		if err != nil {
			return nil, err
		}

		if providerClusters != nil {
			clusters = append(clusters, providerClusters...)
		}
	}

	d, err := time.ParseDuration(cacheDurationNamespaces)
	if err != nil {
		return nil, err
	}

	for _, c := range clusters {
		c.SetOptions(d)
	}

	cs := &Clusters{
		clusters: clusters,
	}

	go cs.generateTopology()

	return cs, nil
}