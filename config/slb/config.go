package slb

import "github.com/upbound/upjet/pkg/config"

const shortGroup = "Slb"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("ksyun_lb", func(r *config.Resource) {

		// we need to override the default group that terrajet generated for
		// this resource, which would be "ksyun"
		r.ShortGroup = shortGroup
		r.ExternalName = config.IdentifierFromProvider
		r.Kind = "LB"

		r.References["vpc_id"] = config.Reference{
			Type: "github.com/kingsoftcloud/provider-ksyun/apis/vpc/v1alpha1.VPC",
		}
	})

	p.AddResourceConfigurator("ksyun_healthcheck", func(r *config.Resource) {

		// we need to override the default group that terrajet generated for
		// this resource, which would be "ksyun"
		r.ShortGroup = shortGroup
		r.ExternalName = config.IdentifierFromProvider

		r.Kind = "Healthcheck"

		r.References["listener_id"] = config.Reference{
			Type: "github.com/kingsoftcloud/provider-ksyun/apis/slb/v1alpha1.Alb",
		}

	})

	p.AddResourceConfigurator("ksyun_lb_acl", func(r *config.Resource) {

		// we need to override the default group that terrajet generated for
		// this resource, which would be "ksyun"
		r.ShortGroup = shortGroup
		r.ExternalName = config.IdentifierFromProvider

		r.Kind = "LbAcl"

	})
	p.AddResourceConfigurator("ksyun_lb_acl_entry", func(r *config.Resource) {

		// we need to override the default group that terrajet generated for
		// this resource, which would be "ksyun"
		r.ShortGroup = shortGroup
		r.ExternalName = config.IdentifierFromProvider

		r.Kind = "LbAclEntry"

		r.References["load_balancer_acl_id"] = config.Reference{
			Type: "github.com/kingsoftcloud/provider-ksyun/apis/slb/v1alpha1.LbAcl",
		}
	})
	p.AddResourceConfigurator("ksyun_lb_backend_server_group", func(r *config.Resource) {

		// we need to override the default group that terrajet generated for
		// this resource, which would be "ksyun"
		r.ShortGroup = shortGroup
		r.ExternalName = config.IdentifierFromProvider

		r.Kind = "LbBackendServerGroup"

		r.References["vpc_id"] = config.Reference{
			Type: "github.com/kingsoftcloud/provider-ksyun/apis/vpc/v1alpha1.VPC",
		}
	})
	p.AddResourceConfigurator("ksyun_lb_host_header", func(r *config.Resource) {

		// we need to override the default group that terrajet generated for
		// this resource, which would be "ksyun"
		r.ShortGroup = shortGroup
		r.ExternalName = config.IdentifierFromProvider

		r.Kind = "LbHostHeader"

		r.References["listener_id"] = config.Reference{
			Type: "github.com/kingsoftcloud/provider-ksyun/apis/slb/v1alpha1.LbListener",
		}

	})

	p.AddResourceConfigurator("ksyun_lb_listener", func(r *config.Resource) {

		// we need to override the default group that terrajet generated for
		// this resource, which would be "ksyun"
		r.ShortGroup = shortGroup
		r.ExternalName = config.IdentifierFromProvider

		r.Kind = "LbListener"

		r.References["load_balancer_id"] = config.Reference{
			Type: "github.com/kingsoftcloud/provider-ksyun/apis/slb/v1alpha1.LB",
		}
		r.References["load_balancer_acl_id"] = config.Reference{
			Type: "github.com/kingsoftcloud/provider-ksyun/apis/slb/v1alpha1.LbAcl",
		}

	})
	p.AddResourceConfigurator("ksyun_lb_listener_associate_acl", func(r *config.Resource) {

		// we need to override the default group that terrajet generated for
		// this resource, which would be "ksyun"
		r.ShortGroup = shortGroup
		r.ExternalName = config.IdentifierFromProvider

		r.Kind = "LbListenerAssociateAcl"

		r.References["listener_id"] = config.Reference{
			Type: "github.com/kingsoftcloud/provider-ksyun/apis/slb/v1alpha1.LbListener",
		}
		r.References["load_balancer_acl_id"] = config.Reference{
			Type: "github.com/kingsoftcloud/provider-ksyun/apis/slb/v1alpha1.LbAcl",
		}

	})
	p.AddResourceConfigurator("ksyun_lb_listener_server", func(r *config.Resource) {

		// we need to override the default group that terrajet generated for
		// this resource, which would be "ksyun"
		r.ShortGroup = shortGroup
		r.ExternalName = config.IdentifierFromProvider

		r.Kind = "LbListenerServer"

		r.References["listener_id"] = config.Reference{
			Type: "github.com/kingsoftcloud/provider-ksyun/apis/slb/v1alpha1.LbListener",
		}

	})
	p.AddResourceConfigurator("ksyun_lb_register_backend_server", func(r *config.Resource) {

		// we need to override the default group that terrajet generated for
		// this resource, which would be "ksyun"
		r.ShortGroup = shortGroup
		r.ExternalName = config.IdentifierFromProvider

		r.Kind = "LbRegisterBackendServer"

		r.References["backend_server_group_id"] = config.Reference{
			Type: "github.com/kingsoftcloud/provider-ksyun/apis/slb/v1alpha1.LbBackendServerGroup",
		}

	})

	p.AddResourceConfigurator("ksyun_lb_rule", func(r *config.Resource) {

		// we need to override the default group that terrajet generated for
		// this resource, which would be "ksyun"
		r.ShortGroup = shortGroup
		r.ExternalName = config.IdentifierFromProvider

		r.Kind = "LbRule"

		r.References["host_header_id"] = config.Reference{
			Type: "github.com/kingsoftcloud/provider-ksyun/apis/slb/v1alpha1.LbHostHeader",
		}

		r.References["backend_server_group_id"] = config.Reference{
			Type: "github.com/kingsoftcloud/provider-ksyun/apis/slb/v1alpha1.LbBackendServerGroup",
		}

	})

}
