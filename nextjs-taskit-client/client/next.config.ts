import type { NextConfig } from "next";

const basePath = process.env.NEXT_PUBLIC_BASE_PATH;
if (basePath == undefined) {
  throw new Error("NEXT_PUBLIC_BASE_PATH must be set")
}

const config: NextConfig = {
  basePath: basePath,
  assetPrefix: basePath,
  output: 'standalone'
};

export default config;
