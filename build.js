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

function copyToBuild(name) {
    fs.cpSync(`./${name}`, `./build/${name}`, {recursive: true})
}

const include = [
    "snippets",
    "syntaxes",
    "CHANGELOG.md",
    "language-configuration.json",
    "LICENSE.txt",
    "package.json",
    "package-lock.json",
    "README.md",
]

include.forEach(name => copyToBuild(name))