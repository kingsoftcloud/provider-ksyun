package eip

import "github.com/upbound/upjet/pkg/config"

const shortGroup = "Eip"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("ksyun_eip", func(r *config.Resource) {

		// we need to override the default group that terrajet generated for
		// this resource, which would be "ksyun"
		r.ShortGroup = shortGroup
		r.ExternalName = config.IdentifierFromProvider
	})

	p.AddResourceConfigurator("ksyun_eip_associate", func(r *config.Resource) {

		// we need to override the default group that terrajet generated for
		// this resource, which would be "ksyun"
		r.ShortGroup = shortGroup
		r.ExternalName = config.IdentifierFromProvider

		r.Kind = "EipAssociate"

		r.References["allocation_id"] = config.Reference{
			Type: "github.com/kingsoftcloud/provider-ksyun/apis/eip/v1alpha1.EIP",
		}

	})

	p.AddResourceConfigurator("ksyun_bws", func(r *config.Resource) {

		// we need to override the default group that terrajet generated for
		// this resource, which would be "ksyun"
		r.ShortGroup = shortGroup
		r.ExternalName = config.IdentifierFromProvider
	})
	p.AddResourceConfigurator("ksyun_bws_associate", func(r *config.Resource) {

		// we need to override the default group that terrajet generated for
		// this resource, which would be "ksyun"
		r.ShortGroup = shortGroup
		r.ExternalName = config.IdentifierFromProvider

		r.Kind = "BwsAssociate"

		r.References["allocation_id"] = config.Reference{
			Type: "github.com/kingsoftcloud/provider-ksyun/apis/eip/v1alpha1.EIP",
		}
		r.References["band_width_share_id"] = config.Reference{
			Type: "github.com/kingsoftcloud/provider-ksyun/apis/eip/v1alpha1.Bws",
		}
	})

}
