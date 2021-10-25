package placeholder

var (
	DefaultPlaceholder = NewPlaceHolder("${", "}")
)

func Replace(content []byte, data map[string]string) ([]byte, error) {
	return DefaultPlaceholder.Replace(content, data)
}

// PlaceHolder 占位符替换器
type PlaceHolder struct {
	left, right string
}

// NewPlaceHolder ...
func NewPlaceHolder(left string, right string) PlaceHolder {
	return PlaceHolder{left: left, right: right}
}

// ParseLayout ...
func (p PlaceHolder) ParseLayout(content []byte) (Layout, error) {
	return ParseLayout(p, content)
}

// ExtractKeys 提取文本中使用通配符的key
func (p PlaceHolder) ExtractKeys(content []byte) (keys []string, err error) {
	layout, err := p.ParseLayout(content)
	if err != nil {
		return nil, err
	}
	return layout.Keys(), nil
}

// Replace content
func (p PlaceHolder) Replace(content []byte, data map[string]string) ([]byte, error) {
	layout, err := p.ParseLayout(content)
	if err != nil {
		return nil, err
	}
	return layout.Execute(data)
}
