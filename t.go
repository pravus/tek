package tek

func T[T any](b bool, t, f func() T) T {
	if b {
		return t()
	}
	return f()
}
