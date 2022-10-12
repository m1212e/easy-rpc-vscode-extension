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

fs.copyFileSync("./package.json", "./build/package.json");