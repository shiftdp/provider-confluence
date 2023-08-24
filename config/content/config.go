/*
Copyright 2021 Upbound Inc.
*/

package content

import (
	ujconfig "github.com/upbound/upjet/pkg/config"
)

// Configure configures the null group
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("confluence_content", func(r *ujconfig.Resource) {
		r.ShortGroup = "content"
	})
}
