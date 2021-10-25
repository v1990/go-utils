package placeholder

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestParseLayout(t *testing.T) {
	tests := []struct {
		name    string
		p       PlaceHolder
		content []byte
		want    Layout
		wantErr error
		keys    []string
	}{
		{
			name:    "simple",
			p:       DefaultPlaceholder,
			content: []byte(`I'm ${nick},from ${city}.`),
			want: Layout{
				chunks:   strings2bytes(`I'm `, `nick`, `,from `, `city`, `.`),
				types:    []int{txtChunk, keyChunk, txtChunk, keyChunk, txtChunk},
				capacity: 0,
			},
			keys: []string{"nick", "city"},
		},
		{
			name:    "escape",
			p:       DefaultPlaceholder,
			content: []byte(`I'm \\${nick} \${nick},\\ from ${city}`),
			want: Layout{
				chunks:   strings2bytes(`I'm \`, `nick`, ` ${nick},\\ from `, `city`),
				types:    []int{txtChunk, keyChunk, txtChunk, keyChunk},
				capacity: 0,
			},
			keys: []string{"nick", "city"},
		},
		{
			name:    "format err",
			p:       DefaultPlaceholder,
			content: []byte(`I'm ${nick ,from ${city}`),
			wantErr: newParseErr(4, []byte(`I'm ${nick ,from ${city}`)),
		},
		{
			name:    "format err 2",
			p:       DefaultPlaceholder,
			content: []byte(`I'm ${nick} ,from ${city xxx`),
			wantErr: newParseErr(18, []byte(`I'm ${nick} ,from ${city xxx`)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseLayout(tt.p, tt.content)
			require.Equal(t, tt.wantErr, err, "got err: %s", err)
			require.Equal(t, tt.want.chunks, got.chunks, "got chunks: %v", bytes2strings(got.chunks))
			require.Equal(t, tt.want.types, got.types)
			require.ElementsMatch(t, tt.keys, got.Keys())
		})
	}
}
func strings2bytes(s ...string) [][]byte {
	b := make([][]byte, len(s))
	for i, ss := range s {
		b[i] = []byte(ss)
	}
	return b
}
func bytes2strings(b [][]byte) []string {
	s := make([]string, len(b))
	for i, bb := range b {
		s[i] = "`" + string(bb) + "`,"
	}
	return s
}
