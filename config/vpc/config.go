package vpc

import "github.com/upbound/upjet/pkg/config"

const shortGroup = "VPC"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("ksyun_vpc", func(r *config.Resource) {

		// we need to override the default group that terrajet generated for
		// this resource, which would be "ksyun"
		r.ShortGroup = shortGroup
		r.ExternalName = config.IdentifierFromProvider
	})
	p.AddResourceConfigurator("ksyun_subnet", func(r *config.Resource) {

		// we need to override the default group that terrajet generated for
		// this resource, which would be "ksyun"
		r.ShortGroup = shortGroup
		r.ExternalName = config.IdentifierFromProvider

		r.References["vpc_id"] = config.Reference{
			Type: "github.com/kingsoftcloud/provider-ksyun/apis/vpc/v1alpha1.VPC",
		}
	})

	p.AddResourceConfigurator("ksyun_security_group", func(r *config.Resource) {

		// we need to override the default group that terrajet generated for
		// this resource, which would be "ksyun"
		r.ShortGroup = shortGroup
		r.ExternalName = config.IdentifierFromProvider

		r.Kind = "SecurityGroup"

		r.References["vpc_id"] = config.Reference{
			Type: "github.com/kingsoftcloud/provider-ksyun/apis/vpc/v1alpha1.VPC",
		}
	})

	p.AddResourceConfigurator("ksyun_security_group_entry", func(r *config.Resource) {

		// we need to override the default group that terrajet generated for
		// this resource, which would be "ksyun"
		r.ShortGroup = shortGroup
		r.ExternalName = config.IdentifierFromProvider

		r.Kind = "SecurityGroupEntry"

		r.References["security_group_id"] = config.Reference{
			Type: "github.com/kingsoftcloud/provider-ksyun/apis/vpc/v1alpha1.SecurityGroup",
		}
	})

	p.AddResourceConfigurator("ksyun_route", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.ExternalName = config.IdentifierFromProvider
		r.Kind = "Route"

		r.References["vpc_id"] = config.Reference{
			Type: "github.com/kingsoftcloud/provider-ksyun/apis/vpc/v1alpha1.VPC",
		}
	})

	p.AddResourceConfigurator("ksyun_network_acl", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.ExternalName = config.IdentifierFromProvider
		r.Kind = "NetworkAcl"

		r.References["vpc_id"] = config.Reference{
			Type: "github.com/kingsoftcloud/provider-ksyun/apis/vpc/v1alpha1.VPC",
		}
	})
	p.AddResourceConfigurator("ksyun_network_acl_associate", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.ExternalName = config.IdentifierFromProvider
		r.Kind = "NetworkAclAssociate"

		r.References["vpc_id"] = config.Reference{
			Type: "github.com/kingsoftcloud/provider-ksyun/apis/vpc/v1alpha1.VPC",
		}
		r.References["network_acl_id"] = config.Reference{
			Type: "github.com/kingsoftcloud/provider-ksyun/apis/vpc/v1alpha1.NetworkAcl",
		}

	})
	p.AddResourceConfigurator("ksyun_network_acl_entry", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.ExternalName = config.IdentifierFromProvider
		r.Kind = "NetworkAclEntry"

		r.References["network_acl_id"] = config.Reference{
			Type: "github.com/kingsoftcloud/provider-ksyun/apis/vpc/v1alpha1.NetworkAcl",
		}
	})
	p.AddResourceConfigurator("ksyun_nat", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.ExternalName = config.IdentifierFromProvider
		r.Kind = "Nat"

		r.References["vpc_id"] = config.Reference{
			Type: "github.com/kingsoftcloud/provider-ksyun/apis/vpc/v1alpha1.VPC",
		}
	})
	p.AddResourceConfigurator("ksyun_nat_associate", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.ExternalName = config.IdentifierFromProvider
		r.Kind = "NatAssociate"

		r.References["nat_id"] = config.Reference{
			Type: "github.com/kingsoftcloud/provider-ksyun/apis/vpc/v1alpha1.Nat",
		}
		r.References["subnet_id"] = config.Reference{
			Type: "github.com/kingsoftcloud/provider-ksyun/apis/vpc/v1alpha1.Subnet",
		}
	})
	p.AddResourceConfigurator("ksyun_nat_instance_bandwidth_limit", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.ExternalName = config.IdentifierFromProvider
		r.Kind = "NatInstanceBandwidthLimit"

		r.References["nat_id"] = config.Reference{
			Type: "github.com/kingsoftcloud/provider-ksyun/apis/vpc/v1alpha1.Nat",
		}

	})
	p.AddResourceConfigurator("ksyun_kec_network_interface", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.ExternalName = config.IdentifierFromProvider
		r.Kind = "KecNetworkInterface"

		r.References["security_group_ids"] = config.Reference{
			Type: "github.com/kingsoftcloud/provider-ksyun/apis/vpc/v1alpha1.SecurityGroup",
		}

		r.References["subnet_id"] = config.Reference{
			Type: "github.com/kingsoftcloud/provider-ksyun/apis/vpc/v1alpha1.Subnet",
		}

	})

	p.AddResourceConfigurator("ksyun_dnat", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.ExternalName = config.IdentifierFromProvider
		r.Kind = "Dnat"

		r.References["nat_id"] = config.Reference{
			Type: "github.com/kingsoftcloud/provider-ksyun/apis/vpc/v1alpha1.Nat",
		}

	})
}
