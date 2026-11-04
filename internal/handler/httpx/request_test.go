package httpx

import (
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
)

func TestParseID(t *testing.T) {
	want := uuid.New()
	req := httptest.NewRequest("PUT", "/api/v1/courses/"+want.String(), nil)
	req.SetPathValue("id", want.String())
	rec := httptest.NewRecorder()

	got, ok := ParseID(rec, req)
	if !ok {
		t.Fatal("ParseID returned false for a valid UUID")
	}
	if got != want {
		t.Fatalf("ParseID returned %s, want %s", got, want)
	}
}

func TestParseIDRejectsInvalidUUID(t *testing.T) {
	req := httptest.NewRequest("PUT", "/api/v1/courses/not-a-uuid", nil)
	req.SetPathValue("id", "not-a-uuid")
	rec := httptest.NewRecorder()

	if _, ok := ParseID(rec, req); ok {
		t.Fatal("ParseID returned true for an invalid UUID")
	}
	if rec.Code != 400 {
		t.Fatalf("ParseID returned status %d, want 400", rec.Code)
	}
}
