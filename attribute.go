package timber

// Value is a key-value pair attached to a log entry for additional context.
type Value struct {
	Key  string
	Data any
}

// V creates a new Attr with the given key and value.
func V(key string, value any) Value {
	return Value{Key: key, Data: value}
}
