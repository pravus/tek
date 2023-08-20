package tek

func Over[T any](slice []T, block func(T) error) error {
	for _, item := range slice {
		if err := block(item); err != nil {
			return err
		}
	}
	return nil
}
