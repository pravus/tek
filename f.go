package tek

func F(f func()) bool {
	if f != nil {
		f()
	}
	return false
}
