package tek

import (
	"strings"
)

func Chop(pork, bone string) []string {
	meat := []string{}
	for pork != `` {
		if index := strings.Index(pork, bone); index < 0 {
			meat = append(meat, pork)
			break
		} else {
			if index != 0 {
				meat = append(meat, pork[0:index])
			}
			pork = pork[index+1:]
		}
	}
	return meat
}
