/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	listener "github.com/kingsoftcloud/provider-ksyun/internal/controller/alb/listener"
	listenercertgroup "github.com/kingsoftcloud/provider-ksyun/internal/controller/alb/listenercertgroup"
	rulegroup "github.com/kingsoftcloud/provider-ksyun/internal/controller/alb/rulegroup"
	snapshot "github.com/kingsoftcloud/provider-ksyun/internal/controller/ebs/snapshot"
	volume "github.com/kingsoftcloud/provider-ksyun/internal/controller/ebs/volume"
	volumeattach "github.com/kingsoftcloud/provider-ksyun/internal/controller/ebs/volumeattach"
	bws "github.com/kingsoftcloud/provider-ksyun/internal/controller/eip/bws"
	bwsassociate "github.com/kingsoftcloud/provider-ksyun/internal/controller/eip/bwsassociate"
	eip "github.com/kingsoftcloud/provider-ksyun/internal/controller/eip/eip"
	eipassociate "github.com/kingsoftcloud/provider-ksyun/internal/controller/eip/eipassociate"
	certificate "github.com/kingsoftcloud/provider-ksyun/internal/controller/kcm/certificate"
	autosnapshotpolicy "github.com/kingsoftcloud/provider-ksyun/internal/controller/kec/autosnapshotpolicy"
	autosnapshotvolumeassociation "github.com/kingsoftcloud/provider-ksyun/internal/controller/kec/autosnapshotvolumeassociation"
	dataguardgroup "github.com/kingsoftcloud/provider-ksyun/internal/controller/kec/dataguardgroup"
	instance "github.com/kingsoftcloud/provider-ksyun/internal/controller/kec/instance"
	kecnetworkinterfaceattachment "github.com/kingsoftcloud/provider-ksyun/internal/controller/kec/kecnetworkinterfaceattachment"
	alb "github.com/kingsoftcloud/provider-ksyun/internal/controller/ksyun/alb"
	providerconfig "github.com/kingsoftcloud/provider-ksyun/internal/controller/providerconfig"
	healthcheck "github.com/kingsoftcloud/provider-ksyun/internal/controller/slb/healthcheck"
	lb "github.com/kingsoftcloud/provider-ksyun/internal/controller/slb/lb"
	lbacl "github.com/kingsoftcloud/provider-ksyun/internal/controller/slb/lbacl"
	lbaclentry "github.com/kingsoftcloud/provider-ksyun/internal/controller/slb/lbaclentry"
	lbbackendservergroup "github.com/kingsoftcloud/provider-ksyun/internal/controller/slb/lbbackendservergroup"
	lbhostheader "github.com/kingsoftcloud/provider-ksyun/internal/controller/slb/lbhostheader"
	lblistener "github.com/kingsoftcloud/provider-ksyun/internal/controller/slb/lblistener"
	lblistenerassociateacl "github.com/kingsoftcloud/provider-ksyun/internal/controller/slb/lblistenerassociateacl"
	lblistenerserver "github.com/kingsoftcloud/provider-ksyun/internal/controller/slb/lblistenerserver"
	lbregisterbackendserver "github.com/kingsoftcloud/provider-ksyun/internal/controller/slb/lbregisterbackendserver"
	lbrule "github.com/kingsoftcloud/provider-ksyun/internal/controller/slb/lbrule"
	dnat "github.com/kingsoftcloud/provider-ksyun/internal/controller/vpc/dnat"
	kecnetworkinterface "github.com/kingsoftcloud/provider-ksyun/internal/controller/vpc/kecnetworkinterface"
	nat "github.com/kingsoftcloud/provider-ksyun/internal/controller/vpc/nat"
	natassociate "github.com/kingsoftcloud/provider-ksyun/internal/controller/vpc/natassociate"
	natinstancebandwidthlimit "github.com/kingsoftcloud/provider-ksyun/internal/controller/vpc/natinstancebandwidthlimit"
	networkacl "github.com/kingsoftcloud/provider-ksyun/internal/controller/vpc/networkacl"
	networkaclassociate "github.com/kingsoftcloud/provider-ksyun/internal/controller/vpc/networkaclassociate"
	networkaclentry "github.com/kingsoftcloud/provider-ksyun/internal/controller/vpc/networkaclentry"
	route "github.com/kingsoftcloud/provider-ksyun/internal/controller/vpc/route"
	securitygroup "github.com/kingsoftcloud/provider-ksyun/internal/controller/vpc/securitygroup"
	securitygroupentry "github.com/kingsoftcloud/provider-ksyun/internal/controller/vpc/securitygroupentry"
	subnet "github.com/kingsoftcloud/provider-ksyun/internal/controller/vpc/subnet"
	vpc "github.com/kingsoftcloud/provider-ksyun/internal/controller/vpc/vpc"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		listener.Setup,
		listenercertgroup.Setup,
		rulegroup.Setup,
		snapshot.Setup,
		volume.Setup,
		volumeattach.Setup,
		bws.Setup,
		bwsassociate.Setup,
		eip.Setup,
		eipassociate.Setup,
		certificate.Setup,
		autosnapshotpolicy.Setup,
		autosnapshotvolumeassociation.Setup,
		dataguardgroup.Setup,
		instance.Setup,
		kecnetworkinterfaceattachment.Setup,
		alb.Setup,
		providerconfig.Setup,
		healthcheck.Setup,
		lb.Setup,
		lbacl.Setup,
		lbaclentry.Setup,
		lbbackendservergroup.Setup,
		lbhostheader.Setup,
		lblistener.Setup,
		lblistenerassociateacl.Setup,
		lblistenerserver.Setup,
		lbregisterbackendserver.Setup,
		lbrule.Setup,
		dnat.Setup,
		kecnetworkinterface.Setup,
		nat.Setup,
		natassociate.Setup,
		natinstancebandwidthlimit.Setup,
		networkacl.Setup,
		networkaclassociate.Setup,
		networkaclentry.Setup,
		route.Setup,
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
