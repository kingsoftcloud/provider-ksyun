/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	"github.com/kingsoftcloud/provider-ksyun/config/ebs"
	"github.com/kingsoftcloud/provider-ksyun/config/eip"
	"github.com/kingsoftcloud/provider-ksyun/config/kcm"
	"github.com/kingsoftcloud/provider-ksyun/config/kec"
	"github.com/kingsoftcloud/provider-ksyun/config/slb"
	"github.com/kingsoftcloud/provider-ksyun/config/vpc"
	ujconfig "github.com/upbound/upjet/pkg/config"
)

const (
	resourcePrefix = "ksyun"
	modulePath     = "github.com/kingsoftcloud/provider-ksyun"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("kingsoftcloud.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		vpc.Configure,
		kec.Configure,
		eip.Configure,
		ebs.Configure,
		slb.Configure,
		kcm.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
