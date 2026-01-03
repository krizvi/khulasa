# khulasa
Khulasa is a Go project that summarizes long text by extracting key sentences, rewriting into simpler language, or combining both approaches.

## What "khulasa" means
The word "khulasa" means "summary" or "gist" in Arabic and Urdu. It fits the project goal: producing a concise, clear summary from a longer input.
****
## Repository layout
```
.
├── app/               # Example app / CLI entry point
├── khulasa-module/    # Core summarization library module
├── go.work            # Go workspace file
├── go.work.sum        # Checksums for workspace dependencies
├── README.md
└── LICENSE
```

## Go workspace (go.work)
`go.work` is a Go workspace file. It lets the Go toolchain treat multiple local modules as a single workspace during development.

In this repo, the workspace includes:
- `./app`
- `./khulasa-module`

What it provides:
- Local-first module resolution: the `app` module can use the local `khulasa-module` without adding `replace` directives in `go.mod`.
- Multi-module development: edit both modules together with consistent dependency resolution.
- Cleaner commits: no temporary `replace` lines needed for local testing.

`go.work.sum` is the workspace checksum file, similar to `go.sum`, but scoped to workspace mode.
