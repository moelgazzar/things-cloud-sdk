package main

import (
	"encoding/json"
	"testing"

	thingscloud "github.com/arthursoares/things-cloud-sdk"
)

func TestNewTaskWriteEnvelopeUsesTask7(t *testing.T) {
	t.Parallel()

	for _, action := range []int{int(thingscloud.ItemActionCreated), int(thingscloud.ItemActionModified)} {
		action := action
		t.Run(thingscloud.ItemAction(action).String(), func(t *testing.T) {
			t.Parallel()
			envelope := newTaskWriteEnvelope("task-id", action, map[string]any{"tt": "test"})
			wire, err := json.Marshal(envelope)
			if err != nil {
				t.Fatalf("marshal task envelope: %v", err)
			}

			var decoded struct {
				Action  int             `json:"t"`
				Kind    string          `json:"e"`
				Payload json.RawMessage `json:"p"`
			}
			if err := json.Unmarshal(wire, &decoded); err != nil {
				t.Fatalf("unmarshal task envelope: %v", err)
			}
			if decoded.Action != action {
				t.Fatalf("action=%d, want %d", decoded.Action, action)
			}
			if decoded.Kind != string(thingscloud.ItemKindTask7) {
				t.Fatalf("kind=%q, want %q", decoded.Kind, thingscloud.ItemKindTask7)
			}
			if len(decoded.Payload) == 0 {
				t.Fatal("task envelope payload is empty")
			}
		})
	}
}
