// Package defenv contains methods for extracting environments variables.
//
// There are two types of methods in a package: ordinary and strict.
// If there is no environment variable, or if a parsing error occurs,
// ordinary methods return a default value.
//
// value := defenv.Int("WORKER_NUMBER", 8)
//
// Unlike ordinary methods, strict methods return an error when parsing fails.
//
// value, err := defenv.IntStrict("WORKER_NUMBER", 8)
//
package defenv

import (
	"os"
	"strconv"
	"time"
)

// Bool extracts bool value from environment variable named name
// and returns defaultValue if it is absent or can not be parsed
func Bool(name string, defaultValue bool) bool {
	if strVal, ok := os.LookupEnv(name); ok {
		if res, err := strconv.ParseBool(strVal); err == nil {
			return res
		}
	}

	return defaultValue
}

// BoolStrict extracts bool value from environment variable named name
// and returns defaultValue if it is absent. If the environment variable
// can not be parsed, the method returns an error
func BoolStrict(name string, defaultValue bool) (bool, error) {
	if strVal, ok := os.LookupEnv(name); ok {
		res, err := strconv.ParseBool(strVal)
		if err != nil {
			return false, err
		}

		return res, nil
	}

	return defaultValue, nil
}

// Duration extracts time.Duration value from environment variable named name
// and returns defaultValue if it is absent or can not be parsed
func Duration(name string, defaultValue time.Duration) time.Duration {
	if strVal, ok := os.LookupEnv(name); ok {
		if d, err := time.ParseDuration(strVal); err == nil {
			return d
		}
	}

	return defaultValue
}

// DurationStrict extracts time.Duration value from environment variable named name
// and returns defaultValue if it is absent. If the environment variable
// can not be parsed, the method returns an error
func DurationStrict(name string, defaultValue time.Duration) (time.Duration, error) {
	if strVal, ok := os.LookupEnv(name); ok {
		d, err := time.ParseDuration(strVal)
		if err != nil {
			return 0, err
		}

		return d, nil
	}

	return defaultValue, nil
}

// Float64 extracts float64 value from environment variable named name
// and returns defaultValue if it is absent or can not be parsed
func Float64(name string, defaultValue float64) float64 {
	if strVal, ok := os.LookupEnv(name); ok {
		if f, err := strconv.ParseFloat(strVal, 64); err == nil {
			return f
		}
	}

	return defaultValue
}

// Float64Strict extracts float64 value from environment variable named name
// and returns defaultValue if it is absent. If the environment variable
// can not be parsed, the method returns an error
func Float64Strict(name string, defaultValue float64) (float64, error) {
	if strVal, ok := os.LookupEnv(name); ok {
		f, err := strconv.ParseFloat(strVal, 64)
		if err != nil {
			return 0, err
		}

		return f, nil
	}

	return defaultValue, nil
}

// Int extracts int value from environment variable named name
// and returns defaultValue if it is absent or can not be parsed
func Int(name string, defaultValue int) int {
	if strVal, ok := os.LookupEnv(name); ok {
		if i64, err := strconv.ParseInt(strVal, 10, 0); err == nil {
			return int(i64)
		}
	}

	return defaultValue
}

// IntStrict extracts int value from environment variable named name
// and returns defaultValue if it is absent. If the environment variable
// can not be parsed, the method returns an error
func IntStrict(name string, defaultValue int) (int, error) {
	if strVal, ok := os.LookupEnv(name); ok {
		i64, err := strconv.ParseInt(strVal, 10, 0)
		if err != nil {
			return 0, err
		}

		return int(i64), nil
	}

	return defaultValue, nil
}

// Int64 extracts int64 value from environment variable named name
// and returns defaultValue if it is absent or can not be parsed
func Int64(name string, defaultValue int64) int64 {
	if strVal, ok := os.LookupEnv(name); ok {
		if i64, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return i64
		}
	}

	return defaultValue
}

// Int64Strict extracts int64 value from environment variable named name
// and returns defaultValue if it is absent. If the environment variable
// can not be parsed, the method returns an error
func Int64Strict(name string, defaultValue int64) (int64, error) {
	if strVal, ok := os.LookupEnv(name); ok {
		i64, err := strconv.ParseInt(strVal, 10, 64)
		if err != nil {
			return 0, err
		}

		return i64, nil
	}

	return defaultValue, nil
}

// String extracts string value from environment variable named name
// and returns defaultValue if it is absent or can not be parsed
func String(name, defaultValue string) string {
	if val, ok := os.LookupEnv(name); ok {
		return val
	}
	return defaultValue
}

// Uint extracts uint value from environment variable named name
// and returns defaultValue if it is absent or can not be parsed
func Uint(name string, defaultValue uint) uint {
	if strVal, ok := os.LookupEnv(name); ok {
		if i64, err := strconv.ParseUint(strVal, 10, 0); err == nil {
			return uint(i64)
		} // Bool extracts bool value from environment variable named name
		// and returns defaultValue if it is absent or can not be parsed

	}

	return defaultValue
}

// UintStrict extracts uint value from environment variable named name
// and returns defaultValue if it is absent. If the environment variable
// can not be parsed, the method returns an error
func UintStrict(name string, defaultValue uint) (uint, error) {
	if strVal, ok := os.LookupEnv(name); ok {
		i64, err := strconv.ParseUint(strVal, 10, 0)
		if err != nil {
			return 0, err
		}

		return uint(i64), nil
	}

	return defaultValue, nil
}

// Uint64 extracts uint64 value from environment variable named name
// and returns defaultValue if it is absent or can not be parsed
func Uint64(name string, defaultValue uint64) uint64 {
	if strVal, ok := os.LookupEnv(name); ok {
		if i64, err := strconv.ParseUint(strVal, 10, 64); err == nil {
			return i64
		}
	}

	return defaultValue
}

// Uint64Strict extracts uint64 value from environment variable named name
// and returns defaultValue if it is absent. If the environment variable
// can not be parsed, the method returns an error
func Uint64Strict(name string, defaultValue uint64) (uint64, error) {
	if strVal, ok := os.LookupEnv(name); ok {
		i64, err := strconv.ParseUint(strVal, 10, 64)
		if err != nil {
			return 0, err
		}

		return i64, nil
	}

	return defaultValue, nil
}
