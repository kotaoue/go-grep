package grep_test

import (
	"reflect"
	"testing"

	"github.com/kotaoue/go-grep/grep"
)

func TestFind(t *testing.T) {
	lines := []string{
		"apple",
		"banana",
		"apricot",
		"cherry",
	}

	tests := []struct {
		name    string
		lines   []string
		pattern string
		want    []string
		wantErr bool
	}{
		{
			name:    "prefix match",
			lines:   lines,
			pattern: "ap",
			want:    []string{"apple", "apricot"},
		},
		{
			name:    "exact match",
			lines:   lines,
			pattern: "banana",
			want:    []string{"banana"},
		},
		{
			name:    "no match",
			lines:   lines,
			pattern: "grape",
			want:    nil,
		},
		{
			name:    "match all",
			lines:   lines,
			pattern: ".",
			want:    lines,
		},
		{
			name:    "empty lines",
			lines:   []string{},
			pattern: "ap",
			want:    nil,
		},
		{
			name:    "invalid pattern",
			lines:   lines,
			pattern: "[invalid",
			wantErr: true,
		},
		{
			name:    "regex anchor",
			lines:   lines,
			pattern: "^a",
			want:    []string{"apple", "apricot"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := grep.Find(tt.lines, tt.pattern)
			if (err != nil) != tt.wantErr {
				t.Errorf("Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Find() = %v, want %v", got, tt.want)
			}
		})
	}
}
