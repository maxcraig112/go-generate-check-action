# go-generate-check

A reusable GitHub Action that fails a workflow if `go generate` would
produce changes that aren't already committed. It exists so any Go repo
using `go:generate` for code generation (e.g. mock generation, typed
wrapper generation, stringer, etc.) can enforce in CI that the checked-in
generated files are actually up to date with their generators, without
copy-pasting the same shell steps into every workflow.

## Usage

```yaml
- uses: actions/checkout@v7
- uses: actions/setup-go@v7
  with:
    go-version-file: go.mod
- uses: maxcraig112/go-generate-check-action@v1
```

## Inputs

| Name                 | Required | Default   | Description                                                              |
| -------------------- | -------- | --------- | -------------------------------------------------------------------------- |
| `working-directory` | no       | `.`     | Directory to run `go generate` and the git checks in (for monorepos). |
| `args`               | no       | `./...` | Arguments passed to `go generate`, e.g. `./...` or a specific package. |

## Prerequisites

This action does not check out your code or install Go — it assumes both
have already happened:

- The repository has already been checked out (e.g. via
  `actions/checkout`) and the working tree is clean.
- The Go toolchain is already installed on the runner (e.g. via
  `actions/setup-go`).

## How it works

1. Runs `go generate <args>` in `working-directory`.
2. Checks whether that command changed anything in the working tree —
   modified or newly created files — via `git status --porcelain`.
3. If the working tree is dirty, the step fails with a message telling you
   to run `go generate` locally and commit the result, and prints the diff
   and any new untracked files so the failure is actionable directly from
   the Actions log.
4. If the working tree is clean, the step succeeds silently.

## License

Apache-2.0. See [LICENSE](./LICENSE).
