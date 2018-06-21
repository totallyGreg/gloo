package translator

import (
	envoyapi "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	"github.com/hashicorp/go-multierror"
	"github.com/solo-io/gloo/pkg/api/types/v1"
)

// aggregate user config errors inside this map
// the top level Translate function should convert these to reports
// to append errors, use multierror.Append
type configErrors map[v1.ConfigObject]error

func (e configErrors) addError(obj v1.ConfigObject, err error) {
	if err == nil {
		return
	}
	e[obj] = multierror.Append(e[obj], err)
}

// the set of resources returned by one iteration for a single v1.Listener
// the top level Translate function should aggregate these into a finished snapshot
type listenerResources struct {
	clusters     []*envoyapi.Cluster
	routeConfig  *envoyapi.RouteConfiguration
	listener     *envoyapi.Listener
	configErrors configErrors
}
