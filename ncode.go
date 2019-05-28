package narou

import (
	"math"
	"strconv"
	"strings"
)

func Ncode2Serial(ncode string) (int, error) {
	serial, err := strconv.Atoi(ncode[1:5])
	if err != nil {
		return 0, err
	}

	for i := 1; i <= len(ncode[5:]); i++ {
		serial += (9999 * int(math.Pow(26, float64(i-1))) * int(strings.ToUpper(ncode)[len(ncode)-i]-"A"[0]))
	}

	return serial, nil
}
