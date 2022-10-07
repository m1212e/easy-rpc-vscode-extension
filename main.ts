import * as path from "path";
import { ExtensionContext, window } from "vscode";
import { LanguageClient, StreamInfo } from "vscode-languageclient/node";
import { ChildProcessWithoutNullStreams, spawn } from "child_process";

let client: LanguageClient;
let p: ChildProcessWithoutNullStreams;

export function activate(context: ExtensionContext) {
  const serverCommand = context.asAbsolutePath(
    path.join("language-server", "easy-rpc.exe")
  );

  client = new LanguageClient(
    "easy-rpc-language-server",
    "easy-rpc-language-server",
    {
      command: serverCommand,
      args: ["-ls"],
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

function createServer(executeablePath: string) {
  p = spawn(executeablePath, ["-ls"], {
    env: { ...process.env, RUST_BACKTRACE: "1" },
  });

  p.stdout.addListener("data", (data: Buffer) => {
    console.log("<-\n" + data.toString() + "\n");
  });
  p.stderr.addListener("data", (data: Buffer) => {
    console.error("<-err\n" + data.toString() + "\n");
  });

  const setDetatchedTrue = () => {
    info.detached = true;
  };

  p.addListener("close", setDetatchedTrue);
  p.addListener("error", setDetatchedTrue);
  p.addListener("disconnect", setDetatchedTrue);
  p.addListener("exit", setDetatchedTrue);

  const info: StreamInfo = {
    writer: p.stdin,
    reader: p.stdout,
  };

  return async () => info;
}
