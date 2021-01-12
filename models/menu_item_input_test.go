package models

import (
	"reflect"
	"testing"
	"time"
)

func TestMenuItemInput_availableAt(t *testing.T) {
	midday := time.Date(3000, 5, 1, 12, 0, 0, 0, time.UTC)
	two := time.Date(3000, 5, 1, 14, 0, 0, 0, time.UTC)
	eleven := time.Date(3000, 5, 1, 23, 0, 0, 0, time.UTC)
	midnight := time.Date(3000, 5, 2, 0, 0, 0, 0, time.UTC)
	type fields struct {
		ID     string
		Status ItemStatus
	}
	type args struct {
		now time.Time
	}
	tests := []struct {
		name            string
		fields          fields
		args            args
		wantAvailableAt *time.Time
	}{
		{
			name: "add two hours if unavailable and two hours before midnight",
			fields: fields{
				ID:     "1",
				Status: ItemStatusUnavailable,
			},
			args: args{
				midday,
			},
			wantAvailableAt: &two,
		},
		{
			name: "midnight if less than two hours before midnight",
			fields: fields{
				ID:     "2",
				Status: ItemStatusUnavailable,
			},
			args: args{
				eleven,
			},
			wantAvailableAt: &midnight,
		},
		{
			name: "nil if status is hidden",
			fields: fields{
				ID:     "3",
				Status: ItemStatusHidden,
			},
			args: args{
				eleven,
			},
			wantAvailableAt: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MenuItemInput{
				ID:     tt.fields.ID,
				Status: tt.fields.Status,
			}
			if gotAvailableAt := m.availableAt(tt.args.now); !reflect.DeepEqual(gotAvailableAt, tt.wantAvailableAt) {
				t.Errorf("availableAt() = %v, want %v", gotAvailableAt, tt.wantAvailableAt)
			}
		})
	}
}
