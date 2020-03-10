package tools

func Merge(s map[string]interface{}, t map[string] interface{}) map[string] interface{} {
	for k,v := range t {
		s[k] = v
	}
	return s
}

func Concat(s map[int]interface{}, t map[int]interface{}) map[int]interface{} {
	for k,v := range t {
		s[len(s) + k] = v
	}
	return s
}
