/*
Copyright 2022 Upbound Inc.
*/

package config

import "github.com/upbound/upjet/pkg/config"

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	// Import requires using a randomly generated ID from provider: nl-2e21sda

	// Vpc
	"ksyun_vpc":                          config.NameAsIdentifier,
	"ksyun_subnet":                       config.NameAsIdentifier,
	"ksyun_security_group":               config.NameAsIdentifier,
	"ksyun_security_group_entry":         config.NameAsIdentifier,
	"ksyun_dnat":                         config.NameAsIdentifier,
	"ksyun_kec_network_interface":        config.NameAsIdentifier,
	"ksyun_nat":                          config.NameAsIdentifier,
	"ksyun_nat_associate":                config.NameAsIdentifier,
	"ksyun_nat_instance_bandwidth_limit": config.NameAsIdentifier,
	"ksyun_network_acl":                  config.NameAsIdentifier,
	"ksyun_network_acl_associate":        config.NameAsIdentifier,
	"ksyun_network_acl_entry":            config.NameAsIdentifier,
	"ksyun_route":                        config.NameAsIdentifier,

	// Instance(KEC)
	"ksyun_instance":                         config.NameAsIdentifier,
	"ksyun_data_guard_group":                 config.NameAsIdentifier,
	"ksyun_kec_network_interface_attachment": config.NameAsIdentifier,
	"ksyun_auto_snapshot_policy":             config.NameAsIdentifier,
	"ksyun_auto_snapshot_volume_association": config.NameAsIdentifier,

	// Eip
	"ksyun_eip":           config.NameAsIdentifier,
	"ksyun_eip_associate": config.NameAsIdentifier,
	"ksyun_bws":           config.NameAsIdentifier,
	"ksyun_bws_associate": config.NameAsIdentifier,

	// Ebs
	"ksyun_snapshot":      config.NameAsIdentifier,
	"ksyun_volume":        config.NameAsIdentifier,
	"ksyun_volume_attach": config.NameAsIdentifier,

	// Alb
	"ksyun_alb":                     config.NameAsIdentifier,
	"ksyun_alb_listener":            config.NameAsIdentifier,
	"ksyun_alb_listener_cert_group": config.NameAsIdentifier,
	"ksyun_alb_rule_group":          config.NameAsIdentifier,

	// Slb
	"ksyun_healthcheck":                config.NameAsIdentifier,
	"ksyun_lb":                         config.NameAsIdentifier,
	"ksyun_lb_acl":                     config.NameAsIdentifier,
	"ksyun_lb_acl_entry":               config.NameAsIdentifier,
	"ksyun_lb_backend_server_group":    config.NameAsIdentifier,
	"ksyun_lb_host_header":             config.NameAsIdentifier,
	"ksyun_lb_listener":                config.NameAsIdentifier,
	"ksyun_lb_listener_associate_acl":  config.NameAsIdentifier,
	"ksyun_lb_listener_server":         config.NameAsIdentifier,
	"ksyun_lb_register_backend_server": config.NameAsIdentifier,
	"ksyun_lb_rule":                    config.NameAsIdentifier,

	// KCM
	"ksyun_certificate": config.NameAsIdentifier,
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
