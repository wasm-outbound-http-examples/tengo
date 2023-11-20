# Use TengoHTTP to send HTTP(s) requests from inside WASM

## Instructions for this devcontainer

Tested with Go 1.21.4, Bun 1.0.13, Deno 1.38.2, Tengo 2.10.0, and TengoHTTP 0.5.0.

### Preparation

1. Open this repo in devcontainer, e.g. using Github Codespaces.
   Type or copy/paste following commands to devcontainer's terminal.

### Building

1. `cd` into the folder of this example:

```sh
cd browser-and-deno-tengohttp
```

2. Install TengoHTTP as library:
```sh
go get github.com/damiva/TengoHTTP
```

This also installs Tengo as dependency.

3. Compile the example:

```sh
GOOS=js GOARCH=wasm go build -o main.wasm main.go
```

4. Copy the glue JS from Golang distribution to example's folder:

```sh
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
```

### Test with browser

1. Run simple HTTP server to temporarily publish project to Web:

```sh
python3 -m http.server
```

Codespace will show you "Open in Browser" button. Just click that button or
obtain web address from "Forwarded Ports" tab.

2. As `index.html` and a **11M**-sized `main.wasm` are loaded into browser, refer to browser developer console
   to see the results.

### Test with Node.js

Impossible yet due to https://github.com/golang/go/issues/59605.

### Test with Bun

1. Install Bun:

```sh
curl -fsSL https://bun.sh/install | bash
```

2. Run with Bun:

```sh
~/.bun/bin/bun bun.js
```

### Test with Deno

1. Install Deno:

```sh
curl -fsSL https://deno.land/x/install/install.sh | sh
```

2. Run with Deno:

```sh
~/.deno/bin/deno run --allow-read --allow-net deno.js
```

### Finish

Perform your own experiments if desired.
