package util

import (
	"errors"
	"fmt"
	"strconv"
)

func ParseInt64(idString string) (int64, error) {
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {

		return 0, errors.New(fmt.Sprintf("%s is invalid reference", idString))
	}

	return id, nil
}
