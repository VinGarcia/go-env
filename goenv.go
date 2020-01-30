package env

import (
	"fmt"
	"os"
	"strconv"
)

// GetStr parses the given environment variable as an string
// in case it is missing and there is no default variable,
// it reports the problem and exits the program.
func GetStr(name string, defaultValue ...string) string {
	value, ok := os.LookupEnv(name)
	if !ok && len(defaultValue) == 0 {
		fmt.Printf("ERROR: missing %s environment variable!\n", name)
		os.Exit(1)
	}

	if !ok {
		value = defaultValue[0]
	}

	return value
}

// GetInt parses the given environment variable as an int
// in case of error with no default variable, it reports the problem
// and exits the program.
func GetInt(name string, defaultValue ...int) int {
	value, ok := os.LookupEnv(name)
	intValue, err := strconv.Atoi(value)

	if (!ok || err != nil) && len(defaultValue) == 0 {
		fmt.Printf("ERROR: missing or invalid %s environment variable!\n", name)
		os.Exit(1)
	}

	if !ok || err != nil {
		intValue = defaultValue[0]
	}

	return intValue
}

// GetBool parses the given environment variable as a bool
// in case of error with no default variable, it reports the problem
// and exits the program.
func GetBool(name string, defaultValue ...bool) bool {
	value, ok := os.LookupEnv(name)
	boolValue, err := strconv.ParseBool(value)

	if (!ok || err != nil) && len(defaultValue) == 0 {
		fmt.Printf("ERROR: missing or invalid %s environment variable!\n", name)
		os.Exit(1)
	}

	if !ok || err != nil {
		boolValue = defaultValue[0]
	}

	return boolValue
}
