package tek

func Mepr[A, Z any](slice []A, block func(A) (Z, error)) ([]Z, error) {
	mepr := make([]Z, len(slice))
	for i := range slice {
		if value, err := block(slice[i]); err != nil {
			return nil, err
		} else {
			mepr[i] = value
		}
	}
	return mepr, nil
}
