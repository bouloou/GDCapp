package localDatabase

import (
	"regexp"
	"strconv"
)

type Client struct {
	Nom           string
	Concepteur    string
	Poseur        string
	PhaseProjet   int
	NumeroDossier string
}

func ConvertIntToFileNumber(a uint64) string {

	lenStr := 0
	aStr := ""

	if a < 100000000 {
		aStr = strconv.FormatUint(a, 10)
		lenStr = len(aStr)
	}

	zeroStr := ""
	for i := 0; i < 8-lenStr; i++ {
		zeroStr += "0"
	}

	return zeroStr + aStr
}

func IsValidFileNumberFormat(a string) bool {
	var re = regexp.MustCompile(`^[0-9]+$`)
	return len(a) == 8 && re.MatchString(a)
}
