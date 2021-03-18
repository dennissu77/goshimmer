package ledgerstate

import (
	"sync"

	"github.com/iotaledger/goshimmer/plugins/webapi"
	"github.com/iotaledger/hive.go/node"
)

// region Plugin ///////////////////////////////////////////////////////////////////////////////////////////////////////

var (
	// plugin holds the instance of the plugin.
	plugin *node.Plugin

	// pluginOnce is used to ensure that the plugin is a singleton.
	pluginOnce sync.Once
)

// Plugin returns the Plugin.
func Plugin() *node.Plugin {
	pluginOnce.Do(func() {
		plugin = node.NewPlugin("WebAPI ledgerstate Endpoint", node.Enabled, configure)
	})

	return plugin
}

func configure(*node.Plugin) {
	webapi.Server().GET("ledgerstate/branches/:branchID", GetBranchEndPoint)
	webapi.Server().GET("ledgerstate/branches/:branchID/children", GetBranchChildrenEndPoint)
}

// endregion ///////////////////////////////////////////////////////////////////////////////////////////////////////////

// region ErrorResponse ////////////////////////////////////////////////////////////////////////////////////////////////

// ErrorResponse is the response that is returned when an error occurred in any of the endpoints.
type ErrorResponse struct {
	Error string `json:"error"`
}

// NewErrorResponse returns an ErrorResponse from the given error.
func NewErrorResponse(err error) ErrorResponse {
	return ErrorResponse{
		Error: err.Error(),
	}
}

// endregion ///////////////////////////////////////////////////////////////////////////////////////////////////////////