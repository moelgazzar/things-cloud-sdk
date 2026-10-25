# things cloud sdk

[Things](https://culturedcode.com/things/) comes with a cloud based API, which can
be used to synchronize data between devices.
This is a golang SDK to interact with that API, opening the API so that you
can enhance your Things experience on iOS and Mac.

[![Go](https://github.com/arthursoares/things-cloud-sdk/actions/workflows/go.yml/badge.svg)](https://github.com/arthursoares/things-cloud-sdk/actions/workflows/go.yml)

## Operational Safety

This fork uses the reverse-engineered Things Cloud API for read-only inspection.
Its `things-cli` blocks direct cloud mutations before initializing cloud access.
Use the supported Things URL Scheme, Apple Shortcuts, or AppleScript on a Mac
to create or change tasks. Build the CLI from this reviewed checkout to use its
safety gate. Low-level SDK mutation methods remain research code and are not an
operational write interface.

## Getting Started

### Installation

```bash
go get github.com/arthursoares/things-cloud-sdk
```

### Quick Start

**1. Set up your credentials:**

```bash
export THINGS_USERNAME='your@email.com'
export THINGS_PASSWORD='yourpassword'
```

**2. Create a simple Go program:**

```go
package main

import (
    "fmt"
    "os"
    things "github.com/arthursoares/things-cloud-sdk"
)

func main() {
    client := things.New(
        things.APIEndpoint,
        os.Getenv("THINGS_USERNAME"),
        os.Getenv("THINGS_PASSWORD"),
    )

    // Verify credentials
    resp, err := client.Verify()
    if err != nil {
        panic(err)
    }
    fmt.Printf("✓ Connected: %s\n", resp.Email)

    // Inspect existing histories without creating or changing records.
    histories, err := client.Histories()
    if err != nil {
        panic(err)
    }
    for _, history := range histories {
        fmt.Printf("History: %s\n", history.ID)
    }
}
```

**3. Run it:**

```bash
go run main.go
```

### CLI Quick Start

Build and use the read-only command-line tool from this checkout:

```bash
go build -o things-cli ./cmd/things-cli/
./things-cli list --today
./things-cli projects
```

## Features

- **Verify Credentials** — validate account access
- **History Inspection** - list and read existing histories
- **Item Inspection** - read task, area, tag and checklist history records
- **Task Types** — tasks, projects, and headings (action groups within projects)
- **Structured Notes** — full-text and delta patch support for task notes
- **Recurring Tasks** — neverending, end on date, end after N times
- **Tombstone Deletion** — explicit deletion records via `Tombstone2` entities
- **Device Registration** — register app instances for APNS push notifications
- **Alarm/Reminders** — alarm time offset support on tasks
- **State Aggregation** — in-memory state built from history items, with queries for projects, headings, subtasks, areas, tags, and checklist items
- **Persistent Sync Engine** — SQLite-backed incremental sync with semantic change detection

## CLI

`things-cli` provides read-only Things Cloud inspection. Direct cloud write and
batch commands fail closed; use supported Mac interfaces for task mutations.

### Setup

```bash
export THINGS_USERNAME='your@email.com'
export THINGS_PASSWORD='yourpassword'
go build -o things-cli ./cmd/things-cli/
```

### Commands

```bash
# Read
things-cli list [--today] [--inbox] [--area NAME] [--project NAME]
things-cli show <uuid>
things-cli areas
things-cli projects
things-cli tags

```

The disabled commands are `create`, `create-area`, `create-tag`, `add-checklist`,
`edit`, `complete`, `trash`, `purge`, `move-to-today`, and `batch`.

### Examples

```bash
./things-cli list --inbox
./things-cli projects
./things-cli tags
```

## Advanced SDK Usage

### Working with Histories and Items

Use existing histories and their item streams for read-only inspection and state
aggregation. The quick start above lists existing histories; the sync engine
below builds local views from received records. Historical mutation examples and
low-level write methods are research material and do not establish a supported
cloud-write route. Use the supported Mac interfaces for every task mutation.

## Persistent Sync Engine

The `sync` package provides a SQLite-backed sync engine that tracks "what changed since last sync" — perfect for building agents, automations, or dashboards that react to Things changes.

```go
package main

import (
    "fmt"
    "os"
    things "github.com/arthursoares/things-cloud-sdk"
    "github.com/arthursoares/things-cloud-sdk/sync"
)

func main() {
    client := things.New(
        things.APIEndpoint,
        os.Getenv("THINGS_USERNAME"),
        os.Getenv("THINGS_PASSWORD"),
    )

    // Open persistent sync database
    syncer, _ := sync.Open("things.db", client)
    defer syncer.Close()

    // Fetch changes since last sync
    changes, _ := syncer.Sync()

    for _, c := range changes {
        switch v := c.(type) {
        case sync.TaskCreated:
            fmt.Printf("New task: %s\n", v.Task.Title)
        case sync.TaskCompleted:
            fmt.Printf("Completed: %s\n", v.Task.Title)
        case sync.TaskMovedToToday:
            fmt.Printf("Moved to Today: %s\n", v.Task.Title)
        }
    }

    // Query current state
    state := syncer.State()
    inbox, _ := state.TasksInInbox(sync.QueryOpts{})
    projects, _ := state.AllProjects(sync.QueryOpts{})
}
```

### Semantic Change Types

The sync engine detects 40+ semantic change types:

| Category | Changes |
|----------|---------|
| **Task Lifecycle** | `TaskCreated`, `TaskCompleted`, `TaskUncompleted`, `TaskTrashed`, `TaskDeleted` |
| **Task Movement** | `TaskMovedToInbox`, `TaskMovedToToday`, `TaskMovedToAnytime`, `TaskMovedToSomeday`, `TaskMovedToUpcoming` |
| **Task Organization** | `TaskMovedToProject`, `TaskMovedToArea`, `TaskMovedUnderHeading`, `TaskTagsChanged` |
| **Task Details** | `TaskTitleChanged`, `TaskNoteChanged`, `TaskDeadlineSet`, `TaskDeadlineRemoved` |
| **Projects** | `ProjectCreated`, `ProjectCompleted`, `ProjectTrashed`, `ProjectDeleted` |
| **Areas & Tags** | `AreaCreated`, `AreaDeleted`, `TagCreated`, `TagDeleted` |
| **Checklists** | `ChecklistItemCreated`, `ChecklistItemCompleted`, `ChecklistItemDeleted` |

### State Queries

```go
state := syncer.State()

// Query by location
inbox, _ := state.TasksInInbox(sync.QueryOpts{})
today, _ := state.TasksInToday(sync.QueryOpts{})
anytime, _ := state.TasksInAnytime(sync.QueryOpts{})

// Query by container
tasks, _ := state.TasksInProject(projectUUID, sync.QueryOpts{})
tasks, _ := state.TasksInArea(areaUUID, sync.QueryOpts{})

// List all
projects, _ := state.AllProjects(sync.QueryOpts{})
areas, _ := state.AllAreas(sync.QueryOpts{})
tags, _ := state.AllTags(sync.QueryOpts{})
```

### Change Log Queries

```go
// Changes in last hour
changes, _ := syncer.ChangesSince(time.Now().Add(-1 * time.Hour))

// Changes for a specific task
changes, _ := syncer.ChangesForEntity(taskUUID)

// Changes since server index
changes, _ := syncer.ChangesSinceIndex(150)
```

## Wire Format Notes

Key findings from reverse engineering the Things Cloud sync protocol:

- **UUIDs must be Base58-encoded** (Bitcoin alphabet: `123456789ABCDEFGH...`). Standard UUID strings or other encodings will crash Things.app during sync.
- **`md` (modification date) must be `null` on creates.** Set timestamps only on updates.
- **Schedule field (`st`)**: `0` = Inbox, `1` = Anytime/Today (with `sr`/`tir` dates = Today), `2` = Someday/Upcoming (with dates = Upcoming).
- **Status field (`ss`)**: `0` = Pending, `2` = Canceled, `3` = Completed. Don't confuse with `st` (schedule)!
- **Headings (`tp=2`) must have `st=1`** (anytime). `st=0` (inbox) crashes Things.app.
- **Tasks in projects, headings, or areas** should default to `st=1` (anytime) — they've been triaged out of inbox.
- **Kind strings**: `Task7`, `Task6`, `Tag4`, `ChecklistItem3`, `Area3`, `Tombstone2`

See `docs/client-side-bugs.md` for the full investigation and crash analysis.

## Architecture

The SDK models all changes as immutable Items (events). A History is a sync stream identified by a UUID. The client pushes/pulls Items through Histories, inspired by [operational transformations and Git's internals](https://www.swift.org/blog/how-swifts-server-support-powers-things-cloud/).

## TODO

- [ ] Repeat after completion
- [x] Persistent state storage (see `sync` package)

## Note

As there is no official API documentation available all requests need to be reverse engineered,
which takes some time. Feel free to contribute and improve & extend this implementation.
