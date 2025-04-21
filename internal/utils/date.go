package utils

import (
	"fmt"
	"strconv"
)

func FixDate(dateItem int) string {
	if dateItem < 10 {
		return fmt.Sprintf("0%d", dateItem)
	}

	return strconv.Itoa(dateItem)
}
