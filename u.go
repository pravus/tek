package tek

func U[T any](b []byte, t T, u func([]byte, any) error) (*T, error) {
	if err := u(b, &t); err != nil {
		return nil, err
	}
	return &t, nil
}
