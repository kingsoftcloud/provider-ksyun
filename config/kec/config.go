package kec

import "github.com/upbound/upjet/pkg/config"

const shortGroupKec = "kec"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("ksyun_instance", func(r *config.Resource) {

		// we need to override the default group that terrajet generated for
		// this resource, which would be "ksyun"
		r.ShortGroup = shortGroupKec
		r.ExternalName = config.IdentifierFromProvider

		r.References["security_group_id"] = config.Reference{
			Type: "github.com/kingsoftcloud/provider-ksyun/apis/vpc/v1alpha1.SecurityGroup",
		}

		r.References["subnet_id"] = config.Reference{
			Type: "github.com/kingsoftcloud/provider-ksyun/apis/vpc/v1alpha1.Subnet",
		}

		r.References["data_guard_id"] = config.Reference{
			Type: "github.com/kingsoftcloud/provider-ksyun/apis/kec/v1alpha1.DataGuardGroup",
		}

	})
	p.AddResourceConfigurator("ksyun_data_guard_group", func(r *config.Resource) {

		// we need to override the default group that terrajet generated for
		// this resource, which would be "ksyun"
		r.ShortGroup = shortGroupKec

		r.Kind = "DataGuardGroup"

		r.ExternalName = config.IdentifierFromProvider
	})
	p.AddResourceConfigurator("ksyun_kec_network_interface_attachment", func(r *config.Resource) {

		// we need to override the default group that terrajet generated for
		// this resource, which would be "ksyun"
		r.ShortGroup = shortGroupKec

		r.Kind = "KecNetworkInterfaceAttachment"

		r.ExternalName = config.IdentifierFromProvider

		r.References["instance_id"] = config.Reference{
			Type: "github.com/kingsoftcloud/provider-ksyun/apis/kec/v1alpha1.Instance",
		}

	})
}
