import * as esbuild from "esbuild";
import { sassPlugin } from "esbuild-sass-plugin";
import { copy } from "esbuild-plugin-copy";
import fs from "node:fs";

const watch = process.argv.includes('--watch');

/** @type {import('esbuild').BuildOptions} */
const config = {
  entryPoints: ["resources/js/app.js", "resources/css/style.scss"],
  bundle: true,
  minify: true,
  logLevel: "debug",
  metafile: true,
  sourcemap: true,
  legalComments: "linked",
  allowOverwrite: true,
  outbase: "resources",
  outdir: "./assets/public/",
  target: ["es2020", "chrome58", "edge16", "firefox57", "safari11"],
  plugins: [
    sassPlugin({
      basedir: "resources/css",
      loadPaths: ["node_modules", "resources/css"],
    }),
    copy({
      resolveFrom: "cwd",
      assets: {
        from: ["./node_modules/htmx.org/dist/htmx.min.js"],
        to: ["./assets/public/js/vendor"],
      },
      watch: true,
    }),
    copy({
      resolveFrom: "cwd",
      assets: {
        from: ["./node_modules/htmx.org/dist/ext/loading-states.js"],
        to: ["./assets/public/js/vendor"],
      },
      watch: true,
    }),
  ],
}

async function build() {
  if (watch) {
    const ctx = await esbuild.context(config);
    await ctx.watch();
    return;
  }

  const result = await esbuild.build(config);
  fs.writeFileSync("meta.json", JSON.stringify(result.metafile));
}

build();
