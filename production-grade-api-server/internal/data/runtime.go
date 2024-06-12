package data

import (
	"fmt"
	"strconv"
)

type Runtime int32

func (runtime Runtime) MarshalJSON() ([]byte, error) {
	formattedRuntime := fmt.Sprintf("%d mins", runtime)

	formattedRuntime = strconv.Quote(formattedRuntime)

	return []byte(formattedRuntime), nil
}
