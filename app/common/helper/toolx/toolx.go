package toolx

import (
	"strconv"
	"strings"
)

func StringTouIntArray(str string) ([]uint, error) {
	var idsUint []uint
	for _, idStr := range strings.Split(str, ",") {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return nil, err
		}
		idsUint = append(idsUint, uint(id))
	}
	return idsUint, nil
}
