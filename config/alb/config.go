package alb

import "github.com/upbound/upjet/pkg/config"

const shortGroup = "Alb"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("ksyun_alb", func(r *config.Resource) {

		// we need to override the default group that terrajet generated for
		// this resource, which would be "ksyun"
		r.ShortGroup = shortGroup
		r.ExternalName = config.IdentifierFromProvider

		r.References["vpc_id"] = config.Reference{
			Type: "github.com/kingsoftcloud/provider-ksyun/apis/vpc/v1alpha1.VPC",
		}
	})

	p.AddResourceConfigurator("ksyun_alb_listener", func(r *config.Resource) {

		// we need to override the default group that terrajet generated for
		// this resource, which would be "ksyun"
		r.ShortGroup = shortGroup
		r.ExternalName = config.IdentifierFromProvider

		r.Kind = "AlbListener"

		r.References["alb_id"] = config.Reference{
			Type: "github.com/kingsoftcloud/provider-ksyun/apis/ebs/v1alpha1.Alb",
		}

	})

	p.AddResourceConfigurator("ksyun_alb_listener_cert_group", func(r *config.Resource) {

		// we need to override the default group that terrajet generated for
		// this resource, which would be "ksyun"
		r.ShortGroup = shortGroup
		r.ExternalName = config.IdentifierFromProvider

		r.Kind = "AlbListenerCertGroup"

		r.References["alb_listener_id"] = config.Reference{
			Type: "github.com/kingsoftcloud/provider-ksyun/apis/ebs/v1alpha1.AlbListener",
		}
	})
	p.AddResourceConfigurator("ksyun_alb_rule_group", func(r *config.Resource) {

		// we need to override the default group that terrajet generated for
		// this resource, which would be "ksyun"
		r.ShortGroup = shortGroup
		r.ExternalName = config.IdentifierFromProvider

		r.Kind = "AlbRuleGroup"

		r.References["alb_listener_id"] = config.Reference{
			Type: "github.com/kingsoftcloud/provider-ksyun/apis/ebs/v1alpha1.AlbListener",
		}
		r.References["backend_server_group_id"] = config.Reference{
			Type: "github.com/kingsoftcloud/provider-ksyun/apis/ebs/v1alpha1.LbBackendServerGroup",
		}
	})

}
