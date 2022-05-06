package langserver

import (
	"context"
	"encoding/json"

	"github.com/sourcegraph/jsonrpc2"
)

func (h *langHandler) handleWorkspaceDidChangeConfiguration(ctx context.Context, conn *jsonrpc2.Conn, req *jsonrpc2.Request) (result interface{}, err error) {
	if req.Params == nil {
		return nil, &jsonrpc2.Error{Code: jsonrpc2.CodeInvalidParams}
	}

	var params DidChangeConfigurationParams
	if err := json.Unmarshal(*req.Params, &params); err != nil {
		return nil, err
	}

	return h.didChangeConfiguration(&params.Settings)
}

func (h *langHandler) didChangeConfiguration(config *Config) (interface{}, error) {
	if config.LogLevel > 0 {
		h.loglevel = config.LogLevel
	}
	return nil, nil
}
