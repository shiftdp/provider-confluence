/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	content "github.com/shiftdp/provider-confluence/internal/controller/content/content"
	providerconfig "github.com/shiftdp/provider-confluence/internal/controller/providerconfig"
	space "github.com/shiftdp/provider-confluence/internal/controller/space/space"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		content.Setup,
		providerconfig.Setup,
		space.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
