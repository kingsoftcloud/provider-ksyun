package vpc

import "github.com/upbound/upjet/pkg/config"

const shortGroupVpc = "VPC"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("ksyun_vpc", func(r *config.Resource) {

		// we need to override the default group that terrajet generated for
		// this resource, which would be "ksyun"
		r.ShortGroup = shortGroupVpc
		r.ExternalName = config.IdentifierFromProvider
	})
	p.AddResourceConfigurator("ksyun_subnet", func(r *config.Resource) {

		// we need to override the default group that terrajet generated for
		// this resource, which would be "ksyun"
		r.ShortGroup = shortGroupVpc
		r.ExternalName = config.IdentifierFromProvider

		r.References["vpc_id"] = config.Reference{
			Type: "github.com/kingsoftcloud/provider-ksyun/apis/vpc/v1alpha1.VPC",
		}
	})

	p.AddResourceConfigurator("ksyun_security_group", func(r *config.Resource) {

		// we need to override the default group that terrajet generated for
		// this resource, which would be "ksyun"
		r.ShortGroup = shortGroupVpc
		r.ExternalName = config.IdentifierFromProvider

		r.Kind = "SecurityGroup"

		r.References["vpc_id"] = config.Reference{
			Type: "github.com/kingsoftcloud/provider-ksyun/apis/vpc/v1alpha1.VPC",
		}
	})

	p.AddResourceConfigurator("ksyun_security_group_entry", func(r *config.Resource) {

		// we need to override the default group that terrajet generated for
		// this resource, which would be "ksyun"
		r.ShortGroup = shortGroupVpc
		r.ExternalName = config.IdentifierFromProvider

		r.Kind = "SecurityGroupEntry"

		r.References["security_group_id"] = config.Reference{
			Type: "github.com/kingsoftcloud/provider-ksyun/apis/vpc/v1alpha1.SecurityGroup",
		}
	})
}
