# Copilot instructions for App-for-Supply-Department

Short, actionable guidance to help AI coding agents get productive quickly.

## Big picture
- Multi-module Go monorepo (see [go.work](go.work#L1-L5)). Modules: `inventory`, `order`, `payment`, `shared`.
- Contract-first architecture: Protobufs live under `shared/proto` and generated code is committed to `shared/pkg/proto` and `shared/pkg/openapi`.
  - Proto generation configured via `shared/proto/buf.gen.yaml` (uses local `bin/` plugins).
  - Example generated code: [shared/pkg/proto/payment/v1/payment.pb.go](shared/pkg/proto/payment/v1/payment.pb.go#L1-L10).

## How services integrate
- Services expose gRPC; an HTTP gateway is generated via `protoc-gen-grpc-gateway` and OpenAPI via `protoc-gen-openapiv2`.
- Validation rules use `protoc-gen-validate` (see `proto_deps/validate/validate.proto`).
- Generated OpenAPI + server stubs live in `shared/pkg/openapi` (see `shared/pkg/openapi/order/v1/oas_handlers_gen.go`).

## Developer workflows (single commands)
- Install project tooling and formatters: `task install-formatters`.
- Install lint tool: `task install-golangci-lint` then `task lint` to run checks.
- Install buf and proto plugins: `task install-buf` and `task proto:install-plugins`.
- Generate protos and OpenAPI: `task proto:gen` (runs `buf generate` in `shared/proto`).
- Generate OpenAPI client/handlers: `task ogen:gen` (uses `shared/api/order/v1/order.openapi.v1.yaml`).

Examples (run from repo root):
```
task proto:install-plugins
task proto:gen
task ogen:gen
task format
task lint
```

## Important repo conventions
- All CLI/protoc tools are installed into the repo-local `bin/` directory by Taskfile tasks — avoid relying on global installs. See `Taskfile.yml` for versions and paths.
- Generated code is checked into `shared/pkg/*`. Do not hand-edit generated files; update source `.proto` / `.yaml` and re-run generation.
- Go modules are composed via `go.work`; use `go` commands from the repo root when possible (the workspace references each module).

## Code patterns to follow (concrete pointers)
- Look for generated gRPC+gateway glue in `shared/pkg/proto/*/*_grpc.pb.go` and gateway files like `*_gw.go` to understand HTTP <> gRPC mapping.
- Validation appears in generated `*_pb.validate.go` files (example: [shared/pkg/proto/payment/v1/payment.pb.validate.go](shared/pkg/proto/payment/v1/payment.pb.validate.go#L1-L10)).
- OpenAPI server/client code is under `shared/pkg/openapi/*` and contains `oas_handlers_gen.go`, `oas_server_gen.go`, etc.

## When editing or adding features
- If you change a `.proto` file, run `task proto:gen` and check `shared/pkg/proto` and `shared/api` outputs.
- If you update OpenAPI spec under `shared/api`, run `task ogen:gen` to regenerate `shared/pkg/openapi`.
- Run `task format` before commits to match repo import grouping and formatting conventions.

## Quick troubleshooting tips
- If generated code doesn't appear, ensure `bin/` plugins exist (Taskfile tasks install them) and run `task proto:gen` from repo root.
- To rebuild a single module: `cd order && go build ./...` or from repo root use `go build ./order/...`.

## Files to inspect for deeper context
- Taskfile workflow and versions: [Taskfile.yml](Taskfile.yml#L1-L10)
- Buf proto generation config: [shared/proto/buf.gen.yaml](shared/proto/buf.gen.yaml#L1-L20)
- go modules workspace: [go.work](go.work#L1-L5)

If anything here is unclear or you want more detail in any section (e.g., step-by-step proto regeneration, recommended local dev run loops, or examples of typical PR changes), tell me which part to expand.
