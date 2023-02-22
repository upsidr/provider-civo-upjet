/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	firewall "github.com/upsidr/provider-civo-upjet/internal/controller/civo/firewall"
	kubernetescluster "github.com/upsidr/provider-civo-upjet/internal/controller/civo/kubernetescluster"
	network "github.com/upsidr/provider-civo-upjet/internal/controller/civo/network"
	providerconfig "github.com/upsidr/provider-civo-upjet/internal/controller/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		firewall.Setup,
		kubernetescluster.Setup,
		network.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
