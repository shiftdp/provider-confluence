/*
Copyright 2021 Upbound Inc.
*/

package space

import (
	ujconfig "github.com/upbound/upjet/pkg/config"
)

// Configure configures the null group
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("confluence_space", func(r *ujconfig.Resource) {
		r.ShortGroup = "space"
	})
}
