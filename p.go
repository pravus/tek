package tek

func P(p any, f func(any, string, string) ([]byte, error)) string {
	if b, err := f(p, ``, `  `); err != nil {
		return err.Error()
	} else {
		return string(b)
	}
}
