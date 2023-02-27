package firewall

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("civo_firewall", func(r *config.Resource) {
		r.ShortGroup = ""
		r.Kind = "CivoFirewall"
		r.References["network_id"] = config.Reference{
			Type: "github.com/upsidr/provider-civo-upjet/apis/civo/v1alpha1.CivoNetwork",
		}
	})
}
