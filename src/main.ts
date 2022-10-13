import { ExtensionContext, window, workspace } from "vscode";
import { LanguageClient } from "vscode-languageclient/node";
import { ChildProcessWithoutNullStreams } from "child_process";
import { getBinary } from "./downloadBinary";

let client: LanguageClient;
let p: ChildProcessWithoutNullStreams;

export async function activate(context: ExtensionContext) {
  if (!workspace.workspaceFolders) {
    throw new Error("Workspace folders undefined. Can't activate.");
  }

  const workspacePath = workspace.workspaceFolders[0].uri.fsPath.replace(/\s/g, "");

  const binaryPath = await getBinary(context.globalStorageUri.fsPath, workspacePath);

  client = new LanguageClient(
    "easy-rpc-language-server",
    "easy-rpc-language-server",
    {
      command: binaryPath,
      args: ["-ls", "-p", workspacePath],
    },
    {
      outputChannel: window.createOutputChannel("easy-rpc-language-server"),
    }
  );

  client.start();
}

export function deactivate(): Thenable<void> | undefined {
  if (!client) {
    return undefined;
  }
  p.removeAllListeners();
  return client.stop();
}
