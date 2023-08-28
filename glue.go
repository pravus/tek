package tek

func Glue(meat []string, hoof string) string {
	pony := ``
	for _, frag := range meat {
		pony += frag + hoof
	}
	return pony
}
