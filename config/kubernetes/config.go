package kubernetes

import "github.com/upbound/upjet/pkg/config"

// ConfigureCluster configures a Kubernetes Cluster CRD.
func ConfigureCluster(p *config.Provider) {
	p.AddResourceConfigurator("civo_kubernetes_cluster", func(r *config.Resource) {
		r.ShortGroup = ""
		r.Kind = "KubernetesCluster"
		r.References["network_id"] = config.Reference{
			Type: "github.com/upsidr/provider-civo-upjet/apis/civo/v1alpha1.Network",
		}
		r.References["firewall_id"] = config.Reference{
			Type: "github.com/upsidr/provider-civo-upjet/apis/civo/v1alpha1.Firewall",
		}
	})
}
