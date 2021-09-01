package placeholder

import (
	"bytes"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

func TestPlaceHolder_Replace(t *testing.T) {
	tests := []struct {
		name        string
		placeHolder *PlaceHolder
		content     string
		data        map[string]string
		want        string
		wantErr     bool
		mErr        error
	}{
		{
			name:        "example",
			placeHolder: DefaultPlaceholder,
			content:     "I'm ${nick},from ${city}",
			data: map[string]string{
				"nick": "zhang san",
				"city": "China",
			},
			want: "I'm zhang san,from China",
		},
		{
			// 单转义 与 双转义（非转义）
			name:        "escape",
			placeHolder: DefaultPlaceholder,
			content:     "I'm \\\\${nick} \\${nick},\\\\\\ from ${city}",
			data: map[string]string{
				"nick": "zhang san",
				"city": "China",
			},
			want: "I'm \\zhang san ${nick},\\\\\\ from China",
		},
		{
			name:        "missKey",
			placeHolder: DefaultPlaceholder,
			content:     "I'm ${nick},from ${city}",
			data: map[string]string{
				//"nick": "zhang san",
				"city": "China",
			},
			want:    "I'm ${nick},from China",
			wantErr: true,
		},
		{
			name:        "noPlaceholder",
			placeHolder: DefaultPlaceholder,
			content:     "I'm LiBai,from China",
			data: map[string]string{
				"nick": "zhang san",
				"city": "China",
			},
			want: "I'm LiBai,from China",
		},
		{
			name:        "utf",
			placeHolder: DefaultPlaceholder,
			content:     "😯 我叫${nick},来自${city}",
			data: map[string]string{
				"nick": "李四",
				"city": "⭐️",
			},
			want: "😯 我叫李四,来自⭐️",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := tt.placeHolder
			got, err := p.Replace([]byte(tt.content), tt.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Replace() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(string(got), tt.want) {
				t.Errorf("Replace() \ngot : %s \nwant: %s \nsrc : %s", got, tt.want, tt.content)
			}
		})
	}
}

func TestPlaceHolder_ExtractKeys(t *testing.T) {

	tests := []struct {
		name        string
		placeHolder *PlaceHolder
		content     string
		wantKeys    []string
		wantErr     bool
	}{
		{
			name:        "example",
			placeHolder: DefaultPlaceholder,
			content:     "${k1} \\\\${k2} \\{$k3}",
			wantKeys:    []string{"k1", "k2"},
			wantErr:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := tt.placeHolder
			gotKeys, err := p.ExtractKeys([]byte(tt.content))
			if (err != nil) != tt.wantErr {
				t.Errorf("ExtractKeys() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotKeys, tt.wantKeys) {
				t.Errorf("ExtractKeys() gotKeys = %v, want %v", gotKeys, tt.wantKeys)
			}
		})
	}
}

func TestReplace(t *testing.T) {
	content, err := Replace([]byte("I'm ${nick},from ${city}"), map[string]string{
		"nick": "ZhangSan",
	})
	require.Equal(t, "I'm ZhangSan,from ${city}", string(content))
	require.Error(t, err)

	mke, ok := IsMissKeysErr(err)
	require.True(t, ok)
	require.Equal(t, mke.Keys(), []string{"city"})

	for _, k := range mke.Keys() {
		old := DefaultPlaceholder.MakePlaceholder(k)
		content = bytes.ReplaceAll(content, []byte(old), nil)
	}
	require.Equal(t, "I'm ZhangSan,from ", string(content))

}
