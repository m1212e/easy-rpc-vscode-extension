import { window } from "vscode";

// https://github.com/napi-rs/package-template/blob/main/index.js
import {
  mkdirSync,
  existsSync,
  readFileSync,
  createWriteStream,
  rmSync,
  writeFileSync,
} from "fs";
import { https } from "follow-redirects";
import { join } from "path";
const { platform, arch } = process;

const ERPC_VERSION = "0.0.0-alpha.5";

/**
 * Ensures the presence of the correct erpc binary on the system and returns its location as string
 * @param globalStoragePath A path to the global storage where the downloaded binary can be saved
 * @param workspace The workspace the vscode instance is opened in
 * @returns A path to the executeable
 */
export async function getBinary(
  globalStoragePath: string,
  workspace: string
): Promise<string> {
  const binaryName = getBinaryName();
  const erpcStorageDirectory = join(globalStoragePath, "erpc");
  const erpcBinaryLocation = join(erpcStorageDirectory, binaryName);

  // check if the workdir contains a matching executeable, if so use it
  // this is useful for development
  if (existsSync(join(workspace, "easy-rpc"))) {
    window.showInformationMessage("Using workspace easy-rpc binary");
    return join(workspace, "easy-rpc");
  }

  // or on on win
  if (existsSync(join(workspace, "easy-rpc.exe"))) {
    window.showInformationMessage("Using workspace easy-rpc.exe binary");
    return join(workspace, "easy-rpc.exe");
  }

  // or with binary name
  if (existsSync(join(workspace, binaryName))) {
    window.showInformationMessage("Using workspace " + binaryName + " binary");
    return join(workspace, binaryName);
  }

  // check if the correct version is installed
  if (
    existsSync(erpcBinaryLocation) &&
    existsSync(join(erpcStorageDirectory, "version")) &&
    readFileSync(join(erpcStorageDirectory, "version")).toString() ==
      ERPC_VERSION
  ) {
    return erpcBinaryLocation;
  }

  // delete all contents of the storage dir and recreate
  rmSync(erpcStorageDirectory, { recursive: true, force: true });
  mkdirSync(erpcStorageDirectory, { recursive: true });

  // download the matching version
  await new Promise<void>((resolve) => {
    https.get(
      `https://github.com/m1212e/easy-rpc/releases/download/${ERPC_VERSION}/${binaryName}`,
      function (response) {
        if (response.statusCode != 200) {
          throw new Error(
            `Could not get(${response.statusCode}):\nhttps://github.com/m1212e/easy-rpc/releases/download/${ERPC_VERSION}/${binaryName}`
          );
        }
        const file = createWriteStream(erpcBinaryLocation);
        response.pipe(file);

        file.on("finish", () => {
          file.close();
          resolve();
        });
      }
    );
  });

  // svae the downloaded version in a version file
  writeFileSync(join(erpcStorageDirectory, "version"), ERPC_VERSION);

  window.showInformationMessage("A new easy-rpc version has been downloaded");
  return erpcBinaryLocation;
}

/**
 * Generates an OS dependent binary name
 */
function getBinaryName() {
  let filename = "";

  switch (platform) {
    case "win32":
      switch (arch) {
        case "x64":
          filename = "easy-rpc-x86_64-pc-windows-msvc.exe";
          break;
        case "arm64":
          filename = "easy-rpc-aarch64-pc-windows-msvc.exe";
          break;
      }
      break;
    case "darwin":
      switch (arch) {
        case "x64":
          filename = "easy-rpc-x86_64-apple-darwin";
          break;
        case "arm64":
          filename = "easy-rpc-aarch64-apple-darwin";
          break;
      }
      break;
    case "linux":
      switch (arch) {
        case "x64":
          if (isMusl()) {
            filename = "easy-rpc-x86_64-unknown-linux-musl";
          } else {
            filename = "easy-rpc-x86_64-unknown-linux-gnu";
          }
          break;
        case "arm64":
          if (isMusl()) {
            filename = "easy-rpc-aarch64-unknown-linux-musl";
          } else {
            filename = "easy-rpc-aarch64-unknown-linux-gnu";
          }
          break;
      }
  }

  if (filename == "") {
    throw new Error(
      `Unsupported OS: ${platform}, architecture: ${arch}\nPlease open an issue at https://github.com/m1212e/easy-rpc and include this error message, so your platform can be added.`
    );
  }

  return filename;
}

function isMusl() {
  // For Node 10
  if (!process.report || typeof process.report.getReport !== "function") {
    try {
      return readFileSync("/usr/bin/ldd", "utf8").includes("musl");
    } catch (e) {
      return true;
    }
  } else {
    const { glibcVersionRuntime } = process.report.getReport().header;
    return !glibcVersionRuntime;
  }
}
