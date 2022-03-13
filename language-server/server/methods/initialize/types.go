package initialize

type TraceOption string
type TextDocumentSyncKind uint8

const (
	Off     TraceOption = "off"
	Messges TraceOption = "messages"
	Verbose TraceOption = "verbose"
)

const (
	Full        TextDocumentSyncKind = 1
	Incremental TextDocumentSyncKind = 2
	None        TextDocumentSyncKind = 0
)

type WorkDoneProgressParams struct {
	/*
		An optional token that a server can use to report work done progress.
	*/
	WorkDoneToken interface{} `json:"workDoneToken,omitempty"` // string or int
}

type Parameters struct {
	WorkDoneProgressParams
	/*
		The process Id of the parent process that started the server. Is null if
		the process has not been started by another process. If the parent
		process is not alive then the server should exit (see exit notification)
		its process.
	*/
	ProcessId *int `json:"processId,omitempty"`
	/*
		Information about the client
		@since 3.15.0
	*/
	ClientInfo *struct {
		/*
			The name of the client as defined by the client.
		*/
		Name *string `json:"name,omitempty"`
		/*
			The client's version as defined by the client.
		*/
		Version *string `json:"version,omitempty"`
	} `json:"clientInfo,omitempty"`

	/*
		The locale the client is currently showing the user interface
		in. This must not necessarily be the locale of the operating
		system.

		Uses IETF language tags as the value's syntax
		(See https:en.wikipedia.orgwikiIETF_language_tag)

		@since 3.16.0
	*/
	Locale *string `json:"locale,omitempty"`

	/*
		The rootPath of the workspace. Is null
		if no folder is open.

		@deprecated in favour of rootUri.
	*/
	RootPath *string `json:"rootPath,omitempty"`

	/*
		The rootUri of the workspace. Is null if no
		folder is open. If both `rootPath` and `rootUri` are set
		`rootUri` wins.

		@deprecated in favour of workspaceFolders.
	*/
	RootUri *string `json:"rootUri,omitempty"`

	/*
		The capabilities provided by the client (editor or tool)
	*/
	Capabilities ClientCapabilities `json:"capabilities,omitempty"`

	/*
		User provided initialization options.
	*/
	InitializationOptions interface{} `json:"initializationOptions,omitempty"`

	/*
		The initial trace setting. If omitted trace is disabled ('off').
	*/
	Trace *TraceOption `json:"trace,omitempty"`

	/**
	 * The workspace folders configured in the client when the server starts.
	 * This property is only available if the client supports workspace folders.
	 * It can be `null` if the client supports workspace folders but none are
	 * configured.
	 *
	 * @since 3.6.0
	 */
	WorkspaceFolders *[]struct {
		Uri  *string `json:"uri,omitempty"`
		Name *string `json:"name,omitempty"`
	} `json:"worspaceFolders,omitempty"`
}

type ClientCapabilities struct {
	Workspace *struct {
		//TODO
	} `json:"workspace,omitempty"`
}

type Response struct {
	/*
		The capabilities the language server provides.
	*/
	Capabilities ServerCapabilities `json:"capabilities,omitempty"`

	/*
		Information about the server.

		@since 3.15.0
	*/
	ServerInfo *struct {
		/*
			The name of the server as defined by the server.
		*/
		Name *string `json:"name,omitempty"`

		/*
			The server's version as defined by the server.
		*/
		Version *string `json:"version,omitempty"`
	} `json:"serverInfo,omitempty"`
}

type ServerCapabilities struct {
	/*
		Defines how text documents are synced. Is either a detailed structure
		defining each notification or for backwards compatibility the
		TextDocumentSyncKind number. If omitted it defaults to
		`TextDocumentSyncKind.None`.
	*/
	TextDocumentSync *TextDocumentSyncOptions `json:"textDocumentSync,omitempty"`

	/*
		The server provides completion support.
	*/
	CompletionProvider *CompletionOptions `json:"completionProvider,omitempty"`

	/*
		The server provides hover support.
	*/
	HoverProvider *bool `json:"hoverProvider,omitempty"`

	/*
		The server provides signature help support.
	*/
	SignatureHelpProvider *SignatureHelpOptions `json:"signatureHelpProvider,omitempty"`

	/*
		The server provides go to declaration support.

		@since 3.14.0
	*/
	DeclarationProvider *bool `json:"declarationProvider,omitempty"`

	/*
		The server provides goto definition support.
	*/
	DefinitionProvider *bool `json:"definitionProvider,omitempty"`

	/*
		The server provides goto type definition support.

		@since 3.6.0
	*/
	TypeDefinitionProvider *bool `json:"typeDefinitionProvider,omitempty"`

	/*
		The server provides goto implementation support.

		@since 3.6.0
	*/
	ImplementationProvider *bool `json:"implementationProvider,omitempty"`

	/*
		The server provides find references support.
	*/
	ReferencesProvider *bool `json:"referencesProvider,omitempty"`

	/*
		The server provides document highlight support.
	*/
	DocumentHighlightProvider *bool `json:"documentHighlightProvider,omitempty"`

	/*
		The server provides document symbol support.
	*/
	DocumentSymbolProvider *bool `json:"documentSymbolProvider,omitempty"`

	/*
		The server provides code actions. The `CodeActionOptions` return type is
		only valid if the client signals code action literal support via the
		property `textDocument.codeAction.codeActionLiteralSupport`.
	*/
	CodeActionProvider *bool `json:"codeActionProvider,omitempty"`

	/*
		The server provides code lens.
	*/
	CodeLensProvider *struct {
		/*
			Code lens has a resolve provider as well.
		*/
		ResolveProvider *bool `json:"resolveProvider,omitempty"`
	} `json:"codeLensProvider,omitempty"`

	//TODO add the rest
}

type TextDocumentSyncOptions struct {
	/*
		Open and close notifications are sent to the server. If omitted open
		close notification should not be sent.
	*/
	OpenClose *bool `json:"openClose,omitempty"`

	Change *TextDocumentSyncKind `json:"change,omitempty"`
}

type CompletionOptions struct {
	WorkDoneProgressParams
	/*
		Most tools trigger completion request automatically without explicitly
		requesting it using a keyboard shortcut (e.g. Ctrl+Space). Typically they
		do so when the user starts to type an identifier. For example if the user
		types `c` in a JavaScript file code complete will automatically pop up
		present `console` besides others as a completion item. Characters that
		make up identifiers don't need to be listed here.

		If code complete should automatically be trigger on characters not being
		valid inside an identifier (for example `.` in JavaScript) list them in
		`triggerCharacters`.
	*/
	TriggerCharacters *[]string `json:"triggerCharacters,omitempty"`

	/*
		The list of all possible characters that commit a completion. This field
		can be used if clients don't support individual commit characters per
		completion item. See client capability
		`completion.completionItem.commitCharactersSupport`.

		If a server provides both `allCommitCharacters` and commit characters on
		an individual completion item the ones on the completion item win.

		@since 3.2.0
	*/
	AllCommitCharacters *[]string `json:"allCommitCharacters,omitempty"`

	/*
		The server provides support to resolve additional
		information for a completion item.
	*/
	ResolveProvider *bool `json:"resolveProvider,omitempty"`

	/*
		The server supports the following `CompletionItem` specific
		capabilities.

		@since 3.17.0 - proposed state
	*/
	CompletionItem *struct {
		/*
			The server has support for completion item label
			details (see also `CompletionItemLabelDetails`) when receiving
			a completion item in a resolve call.

			@since 3.17.0 - proposed state
		*/
		LabelDetailsSupport *bool `json:"labelDetailsSupport,omitempty"`
	} `json:"completionItem,omitempty"`
}

type SignatureHelpOptions struct {
	/*
		The characters that trigger signature help
		automatically.
	*/
	TriggerCharacters *[]string `json:"triggerCharacters,omitempty"`

	/*
		List of characters that re-trigger signature help.

		These trigger characters are only active when signature help is already
		showing. All trigger characters are also counted as re-trigger
		characters.

		@since 3.15.0
	*/
	RetriggerCharacters *[]string `json:"retriggerCharacters,omitempty"`
}
