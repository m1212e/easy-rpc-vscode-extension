package textDocumentDidOpen

type Parameters struct {
	/*
		The document that was opened.
	*/
	TextDocument *TextDocumentItem `json:"textDocument,omitempty"`
}

type TextDocumentItem struct {
	/*
		The text document's URI.
	*/
	Uri *string `json:"uri,omitempty"`

	/*
		The text document's language identifier.
	*/
	LanguageId *string `json:"languageId,omitempty"`

	/*
		The version number of this document (it will increase after each
		change, including undo/redo).
	*/
	Version *int `json:"version,omitempty"`

	/*
		The content of the opened text document.
	*/
	Text *string `json:"text,omitempty"`
}
