package coalescing

type Coalescer interface {
	Do(key string, fn func() (interface{}, error)) (v interface{}, err error, shared bool)
}

type None struct {
}

func (_ *None) Do(_ string, fn func() (interface{}, error)) (interface{}, error, bool) {
	v, err := fn()
	return v, err, false
}
