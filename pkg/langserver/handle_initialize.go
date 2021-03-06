package langserver

import (
	"context"
	"encoding/json"
	"path/filepath"

	"github.com/sourcegraph/jsonrpc2"
)

func (h *langHandler) handleInitialize(ctx context.Context, conn *jsonrpc2.Conn, req *jsonrpc2.Request) (result interface{}, err error) {
	if req.Params == nil {
		return nil, &jsonrpc2.Error{Code: jsonrpc2.CodeInvalidParams}
	}
	h.conn = conn

	var params InitializeParams
	if err := json.Unmarshal(*req.Params, &params); err != nil {
		return nil, err
	}

	// Only try to parse the workspace root if its not null. Otherwise initialize will fail
	if params.RootURI != "" {
		rootPath, err := fromURI(params.RootURI)
		if err != nil {
			return nil, err
		}
		h.rootPath = filepath.Clean(rootPath)
		h.addFolder(rootPath)
	}

	var completion = &CompletionProvider{
		TriggerCharacters: []string{"*"},
	}

	return InitializeResult{
		Capabilities: ServerCapabilities{
			TextDocumentSync:           TDSKFull,
			DocumentFormattingProvider: true,
			DocumentSymbolProvider:     true,
			DefinitionProvider:         true,
			ReferencesProvider:         true,
			CompletionProvider:         completion,
			HoverProvider:              true,
			CodeActionProvider:         true,
			Workspace: &ServerCapabilitiesWorkspace{
				WorkspaceFolders: WorkspaceFoldersServerCapabilities{
					Supported:           true,
					ChangeNotifications: true,
				},
			},
		},
	}, nil
}
