const fs = require("fs");

if (!fs.existsSync("./build")) {
    fs.mkdirSync("./build");
}

require('esbuild').build({
    entryPoints: ['./src/main.ts'],
    bundle: true,
    outdir: './build',
    format: "cjs",
    platform: 'node',
    external: ["vscode"],
    minify: true,
    target: "node14"
}).catch(() => process.exit(1));

// files/dirs on root level which should be included inside the build
const include = [
    "snippets",
    "syntaxes",
    "CHANGELOG.md",
    "language-configuration.json",
    "LICENSE.txt",
    "package.json",
    "package-lock.json",
    "README.md",
];

function copyToBuild(name) {
    fs.cpSync(`./${name}`, `./build/${name}`, {recursive: true});
}
include.forEach(name => copyToBuild(name));

const pkgjson = JSON.parse(fs.readFileSync("./build/package.json"));
pkgjson.main = "main.js"
fs.writeFileSync("./build/package.json", JSON.stringify(pkgjson))