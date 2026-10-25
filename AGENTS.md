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


## Shared Cleanliness And Task Closeout

Apply the [global working agreements](/home/ubuntu/.codex/AGENTS.md), including **Keep Things Organized**, and the complete [workspace resource lifecycle](</home/ubuntu/projects/Agent Workspace/.codex/maintenance/resource-lifecycle.md>) alongside this repository's existing instructions. These cover project-owned storage, essential originals, disposable scratch/builds/screenshots, task-owned tabs/test fixtures, and Git closeout.

Clean only eligible task-owned resources after authority and dependency checks. Preserve active or unpublished work, skills, history, browser logins, and service/recovery paths. Use existing project lifecycle and release gates; verify exact removals and publication. Record holds with owner, reason, dependencies, and next trigger. Missing visibility remains a hold; these instructions do not create a background deletion service.
