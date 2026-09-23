package refind

import "testing"

func TestQuoteRefindValue(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "plain command line",
			input: "quiet rw root=UUID=test",
			want:  `"quiet rw root=UUID=test"`,
		},
		{
			name:  "double quote",
			input: `foo="bar"`,
			want:  `"foo=\"bar\""`,
		},
		{
			name:  "backslash",
			input: `foo\bar`,
			want:  `"foo\\bar"`,
		},
		{
			name:  "mixed escape and quote",
			input: `A "quote" and \path`,
			want:  `"A \"quote\" and \\path"`,
		},
		{
			name:  "unknown escape preserved",
			input: `foo=\x1b`,
			want:  `"foo=\\x1b"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := quoteRefindValue(tt.input); got != tt.want {
				t.Fatalf("quoteRefindValue() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestUnquoteRefindValue(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "plain value",
			input: `"quiet rw root=UUID=test"`,
			want:  "quiet rw root=UUID=test",
		},
		{
			name:  "escaped double quote",
			input: `"foo=\"bar\""`,
			want:  `foo="bar"`,
		},
		{
			name:  "escaped backslash",
			input: `"foo\\bar"`,
			want:  `foo\bar`,
		},
		{
			name:  "unknown escape preserved",
			input: `"foo=\\x1b"`,
			want:  `foo=\x1b`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := unquoteRefindValue(tt.input); got != tt.want {
				t.Fatalf("unquoteRefindValue() = %q, want %q", got, tt.want)
			}
		})
	}
}
