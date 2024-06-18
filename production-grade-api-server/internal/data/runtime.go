package data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Runtime int32

var ErrInvalidRuntimeFormat = errors.New("invalid runtime format")

func (runtime Runtime) MarshalJSON() ([]byte, error) {
	formattedRuntime := fmt.Sprintf("%d mins", runtime)

	formattedRuntime = strconv.Quote(formattedRuntime)

	return []byte(formattedRuntime), nil
}

func (runtime *Runtime) UnmarshalJSON(b []byte) error {
	/* Sanity check on received value */
	s, err := strconv.Unquote(string(b))
	if err != nil {
		return ErrInvalidRuntimeFormat
	}
	spacedFields := strings.Split(s, " ") /* Expecting the string in format "8 mins" */
	if len(spacedFields) != 2 || spacedFields[1] != "mins" {
		return ErrInvalidRuntimeFormat
	}
	time, err := strconv.Atoi(spacedFields[0])
	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	*runtime = Runtime(time)
	return nil
}
