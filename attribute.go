package timber

// Attr is a key-value pair attached to a log entry for additional context.
type Attr struct {
	Key   string
	Value any
}

// A creates a new Attr with the given key and value.
func A(key string, value any) Attr {
	return Attr{Key: key, Value: value}
}
