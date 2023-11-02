package fun

// Apply creates a function that chains all functions given from left to right
// Apply(a, b, c) will return a function that calls a(b(c()))
// See test for an example
func Apply[T any, F func(T) T](funcs ...F) F {

	return func(t T) T {

		res := t

		for i := len(funcs) - 1; i >= 0; i-- {
			res = funcs[i](res)
		}

		return res
	}

}
