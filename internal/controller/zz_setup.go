/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	civofirewall "github.com/upsidr/provider-civo-upjet/internal/controller/civo/civofirewall"
	civokubernetescluster "github.com/upsidr/provider-civo-upjet/internal/controller/civo/civokubernetescluster"
	civonetwork "github.com/upsidr/provider-civo-upjet/internal/controller/civo/civonetwork"
	providerconfig "github.com/upsidr/provider-civo-upjet/internal/controller/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		civofirewall.Setup,
		civokubernetescluster.Setup,
		civonetwork.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
