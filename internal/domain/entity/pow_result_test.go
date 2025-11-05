package entity

import (
	"errors"
	"testing"
)

func TestNewPOWResult(t *testing.T) {
	tests := []struct {
		name  string
		valid bool
		err   error
	}{
		{
			name:  "valid result",
			valid: true,
			err:   nil,
		},
		{
			name:  "invalid result",
			valid: false,
			err:   nil,
		},
		{
			name:  "result with error",
			valid: false,
			err:   errors.New("test error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NewPOWResult(tt.valid, tt.err)

			if result == nil {
				t.Fatal("NewPOWResult() returned nil")
			}
			if result.Valid != tt.valid {
				t.Errorf("NewPOWResult() valid = %v, want %v", result.Valid, tt.valid)
			}
			if result.Error != tt.err {
				t.Errorf("NewPOWResult() error = %v, want %v", result.Error, tt.err)
			}
		})
	}
}
