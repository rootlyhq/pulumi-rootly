// Package rootlyshim exposes the upstream Plugin Framework provider to the Pulumi bridge.
// Its module path permits importing the upstream internal provider without copying it.
package rootlyshim

import (
	"github.com/hashicorp/terraform-plugin-framework/provider"
	rootly "github.com/rootlyhq/terraform-provider-rootly/v5/internal/provider"
)

// New creates the upstream Plugin Framework provider.
func New(version string) provider.Provider {
	return rootly.New(version)()
}
