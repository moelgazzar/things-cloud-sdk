# Things Cloud SDK Codex Notes

## Read Order
1. `CLAUDE.md`
2. `README.md`
3. `cmd/README.md`

## What This Repo Is
This is a Go SDK for the Things Cloud API, including CLI tools and a persistent sync engine.

## Non-Negotiables
- Preserve existing behavior described in `CLAUDE.md`
- Never fake API knowledge beyond what is documented in the repo
- Run focused Go tests for touched packages when making code changes

## Codex Role
Use Codex as a repo-aware coding assistant layered on top of the existing Claude instructions. Keep changes narrow and testable.

## Key Paths
- `CLAUDE.md`
- `README.md`
- `cmd/README.md`
- `sync/`
- `state/memory/`
- `syncutil/`
- `tapes/`
