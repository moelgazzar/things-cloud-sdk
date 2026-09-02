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
- Treat the reverse-engineered Things Cloud API as read-only. Direct cloud writes are disabled because unsupported records can crash Things and corrupt synced libraries.
- Route task creation and changes through Cultured Code's supported Mac interfaces: the Things URL Scheme, Apple Shortcuts, or AppleScript.
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
