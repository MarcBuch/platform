package main

import (
	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
	"cuelang.org/go/cue/parser"
	"github.com/charmbracelet/log"
)

type ClusterConfig struct {
	Name               string
	Kubernetes_version string
	Owner              string
	Default_pool       PoolConfig
	Network            NetworkConfig
	Features           FeatureConfig
}

type PoolConfig struct {
	Min_count int
	Max_count int
	Sku       string
}

type NetworkConfig struct {
	Name string
	Cidr string
}

type FeatureConfig struct {
	External_secrets bool
	Grafana          bool
	Snyk             bool
	Argo             bool
	Flux             bool
}

type CueReader struct {
	values cue.Value
}

// ReadFile reads a CUE file and returns a cue.Value
func (reader CueReader) ReadFile(filepath string) (cue.Value, error) {
	ctx := cuecontext.New()

	f, err := parser.ParseFile(filepath, nil)
	if err != nil {
		log.Error(err.Error())
		return cue.Value{}, err
	}

	values := ctx.BuildFile(f)
	err = values.Validate()
	if err != nil {
		log.Error(err.Error())
		return cue.Value{}, err
	}
	return values, nil
}

// GetValue looksup a path and returns a specific value
func (reader CueReader) GetValue(path string) cue.Value {
	return reader.values.LookupPath(cue.ParsePath(path))
}

func mapConfig(reader CueReader) ClusterConfig {
	var clusterConfig ClusterConfig

	clusterConfig.Name, _ = reader.GetValue("cluster.name").String()
	clusterConfig.Kubernetes_version, _ = reader.GetValue("cluster.kubernetes_version").String()
	clusterConfig.Owner, _ = reader.GetValue("cluster.owner").String()
	minCount, _ := reader.GetValue("cluster.default_pool.min_count").Int64()
	clusterConfig.Default_pool.Min_count = int(minCount)
	maxCount, _ := reader.GetValue("cluster.default_pool.max_count").Int64()
	clusterConfig.Default_pool.Max_count = int(maxCount)
	clusterConfig.Default_pool.Sku, _ = reader.GetValue("cluster.default_pool.sku").String()
	clusterConfig.Network.Name, _ = reader.GetValue("network.name").String()
	clusterConfig.Network.Cidr, _ = reader.GetValue("network.cidr").String()
	clusterConfig.Features.External_secrets, _ = reader.GetValue("features.external_secrets").Bool()
	clusterConfig.Features.Grafana, _ = reader.GetValue("features.grafana").Bool()
	clusterConfig.Features.Snyk, _ = reader.GetValue("features.snyk").Bool()
	clusterConfig.Features.Argo, _ = reader.GetValue("features.argo").Bool()
	clusterConfig.Features.Flux, _ = reader.GetValue("features.flux").Bool()

	return clusterConfig
}
