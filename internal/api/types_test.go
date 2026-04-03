package api

import (
	"encoding/json"
	"testing"
)

func TestFlexInt64_UnmarshalString(t *testing.T) {
	input := `{"expire_time":"1710460800000","is_whole_day":true}`
	var r InlineReminder
	if err := json.Unmarshal([]byte(input), &r); err != nil {
		t.Fatalf("unmarshal string expire_time: %v", err)
	}
	if int64(r.ExpireTime) != 1710460800000 {
		t.Errorf("expected 1710460800000, got %d", r.ExpireTime)
	}
}

func TestFlexInt64_UnmarshalNumber(t *testing.T) {
	input := `{"expire_time":1710460800000,"is_whole_day":true}`
	var r InlineReminder
	if err := json.Unmarshal([]byte(input), &r); err != nil {
		t.Fatalf("unmarshal numeric expire_time: %v", err)
	}
	if int64(r.ExpireTime) != 1710460800000 {
		t.Errorf("expected 1710460800000, got %d", r.ExpireTime)
	}
}
