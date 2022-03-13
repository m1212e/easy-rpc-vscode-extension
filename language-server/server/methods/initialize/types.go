package initialize

type traceOption string

const (
	off     traceOption = "off"
	messges traceOption = "messages"
	verbose traceOption = "verbose"
)

type workDoneProgressParams struct {
	/*
		An optional token that a server can use to report work done progress.
	*/
	workDoneToken interface{} `json:"workDoneToken,omitempty"` // string or int
}

type Params struct {
	workDoneProgressParams
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
		name *string `json:"name,omitempty"`
		/*
			The client's version as defined by the client.
		*/
		version *string `json:"version,omitempty"`
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
	Capabilities clientCapabilities `json:"capabilities,omitempty"`

	/*
		User provided initialization options.
	*/
	InitializationOptions interface{} `json:"initializationOptions,omitempty"`

	/*
		The initial trace setting. If omitted trace is disabled ('off').
	*/
	Trace *traceOption `json:"trace,omitempty"`

	/**
	 * The workspace folders configured in the client when the server starts.
	 * This property is only available if the client supports workspace folders.
	 * It can be `null` if the client supports workspace folders but none are
	 * configured.
	 *
	 * @since 3.6.0
	 */
	WorkspaceFolders *[]struct {
		uri  *string `json:"uri,omitempty"`
		name *string `json:"name,omitempty"`
	} `json:"worspaceFolders,omitempty"`
}

type clientCapabilities struct {
	Workspace *struct {
		//TODO
	} `json:"workspace,omitempty"`
}

type Response struct {
	/*
		The capabilities the language server provides.
	*/
	capabilities serverCapabilities `json:"capabilities,omitempty"`

	/*
		Information about the server.

		@since 3.15.0
	*/
	serverInfo *struct {
		/*
			The name of the server as defined by the server.
		*/
		name *string `json:"name,omitempty"`

		/*
			The server's version as defined by the server.
		*/
		version *string `json:"version,omitempty"`
	} `json:"serverInfo,omitempty"`
}

type serverCapabilities struct {
	completionProvider
}