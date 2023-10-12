/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	bws "github.com/kingsoftcloud/provider-ksyun/internal/controller/eip/bws"
	bwsassociate "github.com/kingsoftcloud/provider-ksyun/internal/controller/eip/bwsassociate"
	eip "github.com/kingsoftcloud/provider-ksyun/internal/controller/eip/eip"
	eipassociate "github.com/kingsoftcloud/provider-ksyun/internal/controller/eip/eipassociate"
	dataguardgroup "github.com/kingsoftcloud/provider-ksyun/internal/controller/kec/dataguardgroup"
	instance "github.com/kingsoftcloud/provider-ksyun/internal/controller/kec/instance"
	kecnetworkinterfaceattachment "github.com/kingsoftcloud/provider-ksyun/internal/controller/kec/kecnetworkinterfaceattachment"
	providerconfig "github.com/kingsoftcloud/provider-ksyun/internal/controller/providerconfig"
	securitygroup "github.com/kingsoftcloud/provider-ksyun/internal/controller/vpc/securitygroup"
	securitygroupentry "github.com/kingsoftcloud/provider-ksyun/internal/controller/vpc/securitygroupentry"
	subnet "github.com/kingsoftcloud/provider-ksyun/internal/controller/vpc/subnet"
	vpc "github.com/kingsoftcloud/provider-ksyun/internal/controller/vpc/vpc"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		bws.Setup,
		bwsassociate.Setup,
		eip.Setup,
		eipassociate.Setup,
		dataguardgroup.Setup,
		instance.Setup,
		kecnetworkinterfaceattachment.Setup,
		providerconfig.Setup,
		securitygroup.Setup,
		securitygroupentry.Setup,
		subnet.Setup,
		vpc.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
