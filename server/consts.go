package server

const (
	// esm.sh build version
	VERSION = 115
	// esm.sh stable build version, used for UI libraries like react, to make sure the runtime is single copy
	// change this carefully
	STABLE_VERSION = 112
)

const (
	nodejsMinVersion = 16
	nodejsLatestLTS  = "18.15.0"
	nodeTypesVersion = "18.15.11"
	denoStdVersion   = "0.177.0"
)

// fix some package versions
var fixedPkgVersions = map[string]string{
	"@types/react@17": "17.0.57",
	"@types/react@18": "18.0.34",
	"isomorphic-ws@4": "5.0.0",
	"resolve@1.22":    "1.22.2", // 1.22.3 will read package.json from disk
}

// css packages
var cssPackages = map[string]string{
	"normalize.css": "normalize.css",
	"@unocss/reset": "tailwind.css",
	"reset-css":     "reset.css",
}

// stable build for UI libraries like react, to make sure the runtime is single copy
var stableBuild = map[string]bool{
	"react":    true,
	"preact":   true,
	"solid-js": true,
	"vue":      true,
}

// reserved packages, for `deno` target use `npm:package` to import (skip build)
var reservedPackages = []string{
	"fsevent",
	"default-gateway",
	"@achingbrain/ssdp",
}

// allowlist for require mode when parsing cjs exports fails
var requireModeAllowList = []string{
	"@babel/types",
	"domhandler",
	"he",
	"jsbn",
	"netmask",
	"xml2js",
	"keycode",
	"lru_map",
	"lz-string",
	"maplibre-gl",
	"postcss-selector-parser",
	"react-draggable",
	"resolve",
	"safe-buffer",
	"seedrandom",
	"stream-browserify",
	"stream-http",
	"typescript",
	"vscode-oniguruma",
	"web-streams-ponyfill",
}

var denoNextUnspportedNodeModules = map[string]bool{
	"inspector": true,
}
