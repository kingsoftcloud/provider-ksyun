package ebs

import "github.com/upbound/upjet/pkg/config"

const shortGroup = "Ebs"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("ksyun_volume", func(r *config.Resource) {

		// we need to override the default group that terrajet generated for
		// this resource, which would be "ksyun"
		r.ShortGroup = shortGroup
		r.ExternalName = config.IdentifierFromProvider

		r.References["snapshot_id"] = config.Reference{
			Type: "github.com/kingsoftcloud/provider-ksyun/apis/ebs/v1alpha1.Snapshot",
		}
	})

	p.AddResourceConfigurator("ksyun_snapshot", func(r *config.Resource) {

		// we need to override the default group that terrajet generated for
		// this resource, which would be "ksyun"
		r.ShortGroup = shortGroup
		r.ExternalName = config.IdentifierFromProvider

		r.References["volume_id"] = config.Reference{
			Type: "github.com/kingsoftcloud/provider-ksyun/apis/ebs/v1alpha1.Volume",
		}

	})

	p.AddResourceConfigurator("ksyun_volume_attach", func(r *config.Resource) {

		// we need to override the default group that terrajet generated for
		// this resource, which would be "ksyun"
		r.ShortGroup = shortGroup
		r.ExternalName = config.IdentifierFromProvider

		r.Kind = "VolumeAttach"

		r.References["volume_id"] = config.Reference{
			Type: "github.com/kingsoftcloud/provider-ksyun/apis/ebs/v1alpha1.Volume",
		}
	})

}
