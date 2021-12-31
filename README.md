
# Setup
* Add `gopherjs` binary to path: `export GOPHERJS_GOROOT="$(go1.17.1 env GOROOT)"  # Also add this line to your .profile or equivalent.`

# Core Commands
* build: `GOARCH=wasm GOOS=js go build -o dist/main.wasm main.go`
* start: `go run serve.go`
* format: `go fmt`

# Other Commands
* copy `wasm_exec.js` to `./dist`: `cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" dist`
