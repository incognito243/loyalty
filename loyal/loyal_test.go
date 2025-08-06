package loyal

import (
	"strings"
	"testing"
)

func TestFindLoyalCustomers(t *testing.T) {
	day1 := `
2025-08-05T10:00:00Z,home,custA
2025-08-05T10:05:00Z,about,custB
2025-08-05T11:00:00Z,home,custC
2025-08-05T11:30:00Z,contact,custA
`
	day2 := `
2025-08-06T09:00:00Z,home,custA
2025-08-06T09:15:00Z,home,custB
2025-08-06T10:00:00Z,pricing,custA
2025-08-06T10:30:00Z,about,custD
`

	got, err := FindLoyalCustomers(strings.NewReader(day1), strings.NewReader(day2))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := []string{"custA", "custB"}
	if len(got) != len(want) {
		t.Fatalf("got %v, want %v", got, want)
	}
	for i := range got {
		if got[i] != want[i] {
			t.Errorf("index %d: got %q, want %q", i, got[i], want[i])
		}
	}
}

func TestNoLoyal(t *testing.T) {
	d1 := `
t1,p1,custX
t2,p2,custX
`
	d2 := `
t3,p3,custY
`
	got, err := FindLoyalCustomers(strings.NewReader(d1), strings.NewReader(d2))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got) != 0 {
		t.Errorf("got %v, want empty", got)
	}
}

func TestInvalidLog(t *testing.T) {
	bad := "not,a,valid,line"
	_, err := FindLoyalCustomers(strings.NewReader(bad), strings.NewReader(""))
	if err == nil {
		t.Fatal("expected parse error, got nil")
	}
}
