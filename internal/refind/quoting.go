package refind

import "strings"

// escapeRefindQuotedString escapes characters that are special inside a
// double-quoted rEFInd value.
func escapeRefindQuotedString(value string) string {
	var b strings.Builder
	b.Grow(len(value))

	for _, r := range value {
		switch r {
		case '\\':
			b.WriteString(`\\`)
		case '"':
			b.WriteString(`\"`)
		case '\n', '\r':
			continue
		default:
			b.WriteRune(r)
		}
	}

	return b.String()
}

// quoteRefindValue returns one double-quoted rEFInd value.
func quoteRefindValue(value string) string {
	return "\"" + escapeRefindQuotedString(value) + "\""
}

// unquoteRefindValue removes one pair of surrounding quotes and decodes only the
// escapes that quoteRefindValue emits. Unknown escapes are preserved.
func unquoteRefindValue(value string) string {
	value = strings.TrimSpace(value)
	if len(value) < 2 || value[0] != '"' || value[len(value)-1] != '"' {
		return value
	}

	value = value[1 : len(value)-1]
	var b strings.Builder
	b.Grow(len(value))

	for i := 0; i < len(value); i++ {
		if value[i] != '\\' || i+1 >= len(value) {
			b.WriteByte(value[i])
			continue
		}

		next := value[i+1]
		if next == '\\' || next == '"' {
			b.WriteByte(next)
			i++
			continue
		}

		b.WriteByte('\\')
	}

	return b.String()
}
