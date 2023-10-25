package kcm

import "github.com/upbound/upjet/pkg/config"

const shortGroup = "Kcm"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("ksyun_certificate", func(r *config.Resource) {

		// we need to override the default group that terrajet generated for
		// this resource, which would be "ksyun"
		r.ShortGroup = shortGroup
		r.ExternalName = config.IdentifierFromProvider

		r.Kind = "Certificate"

	})

}
