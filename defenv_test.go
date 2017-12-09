package defenv

import (
	"errors"
	"fmt"
	"os"
	"testing"
	"time"
)

func TestBool(t *testing.T) {
	for _, tc := range []struct {
		name         string
		setEnv       bool
		envValue     string
		defaultValue bool
		expRes       bool
	}{
		{
			name:         `true then environment value is "true"`,
			setEnv:       true,
			envValue:     "true",
			defaultValue: false,
			expRes:       true,
		},
		{
			name:         `true then environment value is "1"`,
			setEnv:       true,
			envValue:     "1",
			defaultValue: false,
			expRes:       true,
		},
		{
			name:         `use default value then environment value is ""`,
			setEnv:       true,
			envValue:     "",
			defaultValue: true,
			expRes:       true,
		},
		{
			name:         `use default value then environment value is "bad"`,
			setEnv:       true,
			envValue:     "bad",
			defaultValue: true,
			expRes:       true,
		},
		{
			name:         `false then environment value as "F"`,
			setEnv:       true,
			envValue:     "F",
			defaultValue: true,
			expRes:       false,
		},
		{
			name:         `use default value then environment value is not set`,
			setEnv:       false,
			defaultValue: true,
			expRes:       true,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if err := os.Unsetenv("VALUE"); err != nil {
					t.Errorf("coudn't unset VALUE: %s", err)
				}
			}()

			if tc.setEnv {
				if err := os.Setenv("VALUE", tc.envValue); err != nil {
					t.Fatal(err)
				}
			}

			res := Bool("VALUE", tc.defaultValue)
			if res != tc.expRes {
				t.Errorf("expected value: %t, got: %t", tc.expRes, res)
			}
		})
	}
}

func TestBoolStrict(t *testing.T) {
	for _, tc := range []struct {
		name         string
		setEnv       bool
		envValue     string
		defaultValue bool
		expRes       bool
		expErr       error
	}{
		{
			name:         `true then environment value is "true"`,
			setEnv:       true,
			envValue:     "true",
			defaultValue: false,
			expRes:       true,
		},
		{
			name:         `true then environment value is "1"`,
			setEnv:       true,
			envValue:     "1",
			defaultValue: false,
			expRes:       true,
		},
		{
			name:         `fail then environment value is ""`,
			setEnv:       true,
			envValue:     "",
			defaultValue: true,
			expRes:       false,
			expErr:       errors.New(`strconv.ParseBool: parsing "": invalid syntax`),
		},
		{
			name:         `fail then environment value is "bad"`,
			setEnv:       true,
			envValue:     "bad",
			defaultValue: true,
			expRes:       false,
			expErr:       errors.New(`strconv.ParseBool: parsing "bad": invalid syntax`),
		},
		{
			name:         `false then environment value is "F"`,
			setEnv:       true,
			envValue:     "F",
			defaultValue: true,
			expRes:       false,
		},
		{
			name:         "use default value then environment value is not set",
			setEnv:       false,
			defaultValue: true,
			expRes:       true,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if err := os.Unsetenv("VALUE"); err != nil {
					t.Errorf("coudn't unset VALUE: %s", err)
				}
			}()

			if tc.setEnv {
				if err := os.Setenv("VALUE", tc.envValue); err != nil {
					t.Fatal(err)
				}
			}

			res, err := BoolStrict("VALUE", tc.defaultValue)
			if fmt.Sprint(err) != fmt.Sprint(tc.expErr) {
				t.Errorf("expected error: %v, got: %v", tc.expErr, err)
			}
			if res != tc.expRes {
				t.Errorf("expected value: %t, got: %t", tc.expRes, res)
			}
		})
	}
}

func TestDuration(t *testing.T) {
	for _, tc := range []struct {
		name         string
		setEnv       bool
		envValue     string
		defaultValue time.Duration
		expRes       time.Duration
	}{
		{
			name:         `200 milliseconds then environment value is "200ms"`,
			setEnv:       true,
			envValue:     "200ms",
			defaultValue: 3 * time.Second,
			expRes:       200 * time.Millisecond,
		},
		{
			name:         `-2 hours then environment value is "200ms"`,
			setEnv:       true,
			envValue:     "-2h",
			defaultValue: 3 * time.Second,
			expRes:       -2 * time.Hour,
		},
		{
			name:         `2 hours 5 minutes and 20 seconds then environment value is "2h5m20s"`,
			setEnv:       true,
			envValue:     "2h5m20s",
			defaultValue: 3 * time.Second,
			expRes:       2*time.Hour + 5*time.Minute + 20*time.Second,
		},
		{
			name:         `use default value then environment value is "30"`,
			setEnv:       true,
			envValue:     "30",
			defaultValue: 3 * time.Second,
			expRes:       3 * time.Second,
		},
		{
			name:         `use default value then environment value is ""`,
			setEnv:       true,
			envValue:     "",
			defaultValue: 3 * time.Second,
			expRes:       3 * time.Second,
		},
		{
			name:         `use default value then then environment value is "bad"`,
			setEnv:       true,
			envValue:     "bad",
			defaultValue: 3 * time.Second,
			expRes:       3 * time.Second,
		},
		{
			name:         `use default value then environment value is not set`,
			setEnv:       false,
			defaultValue: 3 * time.Second,
			expRes:       3 * time.Second,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if err := os.Unsetenv("VALUE"); err != nil {
					t.Errorf("coudn't unset VALUE: %s", err)
				}
			}()

			if tc.setEnv {
				if err := os.Setenv("VALUE", tc.envValue); err != nil {
					t.Fatal(err)
				}
			}

			res := Duration("VALUE", tc.defaultValue)
			if res != tc.expRes {
				t.Errorf("expected value: %v, got: %v", tc.expRes, res)
			}
		})
	}
}

func TestDurationStrict(t *testing.T) {
	for _, tc := range []struct {
		name         string
		setEnv       bool
		envValue     string
		defaultValue time.Duration
		expRes       time.Duration
		expErr       error
	}{
		{
			name:         `200 milliseconds then environment value is "200ms"`,
			setEnv:       true,
			envValue:     "200ms",
			defaultValue: 3 * time.Second,
			expRes:       200 * time.Millisecond,
		},
		{
			name:         `-2 hours then environment value is "200ms"`,
			setEnv:       true,
			envValue:     "-2h",
			defaultValue: 3 * time.Second,
			expRes:       -2 * time.Hour,
		},
		{
			name:         `2 hours 5 minutes and 20 seconds then environment value is "2h5m20s"`,
			setEnv:       true,
			envValue:     "2h5m20s",
			defaultValue: 3 * time.Second,
			expRes:       2*time.Hour + 5*time.Minute + 20*time.Second,
		},
		{
			name:         `fail then environment value is "30"`,
			setEnv:       true,
			envValue:     "30",
			defaultValue: 3 * time.Second,
			expErr:       errors.New("time: missing unit in duration 30"),
		},
		{
			name:         `fail then environment value is ""`,
			setEnv:       true,
			envValue:     "",
			defaultValue: 3 * time.Second,
			expErr:       errors.New("time: invalid duration "),
		},
		{
			name:         `fail then environment is "bad"`,
			setEnv:       true,
			envValue:     "bad",
			defaultValue: 3 * time.Second,
			expErr:       errors.New("time: invalid duration bad"),
		},
		{
			name:         `use default value then environment value is not set`,
			setEnv:       false,
			defaultValue: 3 * time.Second,
			expRes:       3 * time.Second,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if err := os.Unsetenv("VALUE"); err != nil {
					t.Errorf("coudn't unset VALUE: %s", err)
				}
			}()

			if tc.setEnv {
				if err := os.Setenv("VALUE", tc.envValue); err != nil {
					t.Fatal(err)
				}
			}

			res, err := DurationStrict("VALUE", tc.defaultValue)
			if fmt.Sprint(err) != fmt.Sprint(tc.expErr) {
				t.Errorf("expected error: %v, got: %v", tc.expErr, err)
			}
			if res != tc.expRes {
				t.Errorf("expected value: %v, got: %v", tc.expRes, res)
			}
		})
	}
}

func TestFloat64(t *testing.T) {
	for _, tc := range []struct {
		name         string
		setEnv       bool
		envValue     string
		defaultValue float64
		expRes       float64
	}{
		{
			name:         `3.14 then environment value is "3.14"`,
			setEnv:       true,
			envValue:     "3.14",
			defaultValue: 1.2,
			expRes:       3.14,
		},
		{
			name:         `-321.123 then environment value is "-321.123"`,
			setEnv:       true,
			envValue:     "-321.123",
			defaultValue: 1.2,
			expRes:       -321.123,
		},
		{
			name:         `30 then environment value is "30"`,
			setEnv:       true,
			envValue:     "30",
			defaultValue: 1.2,
			expRes:       30,
		},
		{
			name:         `0 then environment value is "0"`,
			setEnv:       true,
			envValue:     "0",
			defaultValue: 1.2,
			expRes:       0,
		},
		{
			name:         `use default value then environment is ""`,
			setEnv:       true,
			envValue:     "",
			defaultValue: 1.2,
			expRes:       1.2,
		},
		{
			name:         `use default value then environment is "bad"`,
			setEnv:       true,
			envValue:     "bad",
			defaultValue: 1.2,
			expRes:       1.2,
		},
		{
			name:         `use default value then environment value is not set`,
			setEnv:       false,
			defaultValue: 1.2,
			expRes:       1.2,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if err := os.Unsetenv("VALUE"); err != nil {
					t.Errorf("coudn't unset VALUE: %s", err)
				}
			}()

			if tc.setEnv {
				if err := os.Setenv("VALUE", tc.envValue); err != nil {
					t.Fatal(err)
				}
			}

			res := Float64("VALUE", tc.defaultValue)
			if res != tc.expRes {
				t.Errorf("expected value: %f, got: %f", tc.expRes, res)
			}
		})
	}
}

func TestFloat64Strict(t *testing.T) {
	for _, tc := range []struct {
		name         string
		setEnv       bool
		envValue     string
		defaultValue float64
		expRes       float64
		expErr       error
	}{
		{
			name:         `3.14 then environment value is "3.14"`,
			setEnv:       true,
			envValue:     "3.14",
			defaultValue: 1.2,
			expRes:       3.14,
		},
		{
			name:         `-321.123 then environment value is "-321.123"`,
			setEnv:       true,
			envValue:     "-321.123",
			defaultValue: 1.2,
			expRes:       -321.123,
		},
		{
			name:         `30 then environment value is "30"`,
			setEnv:       true,
			envValue:     "30",
			defaultValue: 1.2,
			expRes:       30,
		},
		{
			name:         `0 then environment value is "0"`,
			setEnv:       true,
			envValue:     "0",
			defaultValue: 1.2,
			expRes:       0,
		},
		{
			name:         `fail then environment value is ""`,
			setEnv:       true,
			envValue:     "",
			defaultValue: 1.2,
			expErr:       errors.New(`strconv.ParseFloat: parsing "": invalid syntax`),
		},
		{
			name:         `fail then environment value is "bad"`,
			setEnv:       true,
			envValue:     "bad",
			defaultValue: 1.2,
			expErr:       errors.New(`strconv.ParseFloat: parsing "bad": invalid syntax`),
		},
		{
			name:         `use default value then environment value is not set`,
			setEnv:       false,
			defaultValue: 1.2,
			expRes:       1.2,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if err := os.Unsetenv("VALUE"); err != nil {
					t.Errorf("coudn't unset VALUE: %s", err)
				}
			}()

			if tc.setEnv {
				if err := os.Setenv("VALUE", tc.envValue); err != nil {
					t.Fatal(err)
				}
			}

			res, err := Float64Strict("VALUE", tc.defaultValue)
			if fmt.Sprint(err) != fmt.Sprint(tc.expErr) {
				t.Errorf("expected error: %v, got: %v", tc.expErr, err)
			}
			if res != tc.expRes {
				t.Errorf("expected value: %f, got: %f", tc.expRes, res)
			}
		})
	}
}

func TestInt(t *testing.T) {
	for _, tc := range []struct {
		name         string
		setEnv       bool
		envValue     string
		defaultValue int
		expRes       int
	}{
		{
			name:         `123 then environment value is "123"`,
			setEnv:       true,
			envValue:     "123",
			defaultValue: 321,
			expRes:       123,
		},
		{
			name:         `-1 then environment value is "-1"`,
			setEnv:       true,
			envValue:     "-1",
			defaultValue: 321,
			expRes:       -1,
		},
		{
			name:         `use default value then environment value is "3.1"`,
			setEnv:       true,
			envValue:     "3.1",
			defaultValue: 321,
			expRes:       321,
		},
		{
			name:         `0 then environment value is "0"`,
			setEnv:       true,
			envValue:     "0",
			defaultValue: 321,
			expRes:       0,
		},
		{
			name:         `use default value then environment value is ""`,
			setEnv:       true,
			envValue:     "",
			defaultValue: 321,
			expRes:       321,
		},
		{
			name:         `use default value then environment value is "bad"`,
			setEnv:       true,
			envValue:     "bad",
			defaultValue: 321,
			expRes:       321,
		},
		{
			name:         `use default value then environment value is more then int max value`,
			setEnv:       true,
			envValue:     "12345678901234567890",
			defaultValue: 321,
			expRes:       321,
		},
		{
			name:         "use default value then environment value is not set",
			setEnv:       false,
			defaultValue: 321,
			expRes:       321,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if err := os.Unsetenv("VALUE"); err != nil {
					t.Errorf("coudn't unset VALUE: %s", err)
				}
			}()

			if tc.setEnv {
				if err := os.Setenv("VALUE", tc.envValue); err != nil {
					t.Fatal(err)
				}
			}

			res := Int("VALUE", tc.defaultValue)
			if res != tc.expRes {
				t.Errorf("expected value: %d, got: %d", tc.expRes, res)
			}
		})
	}
}

func TestIntStrict(t *testing.T) {
	for _, tc := range []struct {
		name         string
		setEnv       bool
		envValue     string
		defaultValue int
		expRes       int
		expErr       error
	}{
		{
			name:         `123 from environment as "123"`,
			setEnv:       true,
			envValue:     "123",
			defaultValue: 321,
			expRes:       123,
		},
		{
			name:         `-1 then environment value is "-1"`,
			setEnv:       true,
			envValue:     "-1",
			defaultValue: 321,
			expRes:       -1,
		},
		{
			name:         `fail then environment value is "3.1"`,
			setEnv:       true,
			envValue:     "3.1",
			defaultValue: 321,
			expErr:       errors.New(`strconv.ParseInt: parsing "3.1": invalid syntax`),
		},
		{
			name:         `0 then environment value is "0"`,
			setEnv:       true,
			envValue:     "0",
			defaultValue: 321,
			expRes:       0,
		},
		{
			name:         `fail then environment value is ""`,
			setEnv:       true,
			envValue:     "",
			defaultValue: 321,
			expErr:       errors.New(`strconv.ParseInt: parsing "": invalid syntax`),
		},
		{
			name:         `fail then environment value is "bad"`,
			setEnv:       true,
			envValue:     "bad",
			defaultValue: 321,
			expErr:       errors.New(`strconv.ParseInt: parsing "bad": invalid syntax`),
		},
		{
			name:         `fail then environment value is more then then int max value`,
			setEnv:       true,
			envValue:     "12345678901234567890",
			defaultValue: 321,
			expErr:       errors.New(`strconv.ParseInt: parsing "12345678901234567890": value out of range`),
		},
		{
			name:         `use default value then environment value is not set`,
			setEnv:       false,
			defaultValue: 321,
			expRes:       321,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if err := os.Unsetenv("VALUE"); err != nil {
					t.Errorf("coudn't unset VALUE: %s", err)
				}
			}()

			if tc.setEnv {
				if err := os.Setenv("VALUE", tc.envValue); err != nil {
					t.Fatal(err)
				}
			}

			res, err := IntStrict("VALUE", tc.defaultValue)
			if fmt.Sprint(err) != fmt.Sprint(tc.expErr) {
				t.Errorf("expected error: %v, got: %v", tc.expErr, err)
			}
			if res != tc.expRes {
				t.Errorf("expected value: %d, got: %d", tc.expRes, res)
			}
		})
	}
}

func TestInt64(t *testing.T) {
	for _, tc := range []struct {
		name         string
		setEnv       bool
		envValue     string
		defaultValue int64
		expRes       int64
	}{
		{
			name:         `123 then environment value is "123"`,
			setEnv:       true,
			envValue:     "123",
			defaultValue: 321,
			expRes:       123,
		},
		{
			name:         `-1 then environment value is "-1"`,
			setEnv:       true,
			envValue:     "-1",
			defaultValue: 321,
			expRes:       -1,
		},
		{
			name:         `use default value then environment value is "3.1"`,
			setEnv:       true,
			envValue:     "3.1",
			defaultValue: 321,
			expRes:       321,
		},
		{
			name:         `0 then environment value is "0"`,
			setEnv:       true,
			envValue:     "0",
			defaultValue: 321,
			expRes:       0,
		},
		{
			name:         `use default value then environment value is ""`,
			setEnv:       true,
			envValue:     "",
			defaultValue: 321,
			expRes:       321,
		},
		{
			name:         `use default value then environment value is "bad"`,
			setEnv:       true,
			envValue:     "bad",
			defaultValue: 321,
			expRes:       321,
		},
		{
			name:         `use default value then environment value is more then int max value`,
			setEnv:       true,
			envValue:     "12345678901234567890",
			defaultValue: 321,
			expRes:       321,
		},
		{
			name:         "use default value then environment value is not set",
			setEnv:       false,
			defaultValue: 321,
			expRes:       321,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if err := os.Unsetenv("VALUE"); err != nil {
					t.Errorf("coudn't unset VALUE: %s", err)
				}
			}()

			if tc.setEnv {
				if err := os.Setenv("VALUE", tc.envValue); err != nil {
					t.Fatal(err)
				}
			}

			res := Int64("VALUE", tc.defaultValue)
			if res != tc.expRes {
				t.Errorf("expected value: %d, got: %d", tc.expRes, res)
			}
		})
	}
}

func TestInt64Strict(t *testing.T) {
	for _, tc := range []struct {
		name         string
		setEnv       bool
		envValue     string
		defaultValue int64
		expRes       int64
		expErr       error
	}{
		{
			name:         `123 from environment as "123"`,
			setEnv:       true,
			envValue:     "123",
			defaultValue: 321,
			expRes:       123,
		},
		{
			name:         `-1 then environment value is "-1"`,
			setEnv:       true,
			envValue:     "-1",
			defaultValue: 321,
			expRes:       -1,
		},
		{
			name:         `fail then environment value is "3.1"`,
			setEnv:       true,
			envValue:     "3.1",
			defaultValue: 321,
			expErr:       errors.New(`strconv.ParseInt: parsing "3.1": invalid syntax`),
		},
		{
			name:         `0 then environment value is "0"`,
			setEnv:       true,
			envValue:     "0",
			defaultValue: 321,
			expRes:       0,
		},
		{
			name:         `fail then environment value is ""`,
			setEnv:       true,
			envValue:     "",
			defaultValue: 321,
			expErr:       errors.New(`strconv.ParseInt: parsing "": invalid syntax`),
		},
		{
			name:         `fail then environment value is "bad"`,
			setEnv:       true,
			envValue:     "bad",
			defaultValue: 321,
			expErr:       errors.New(`strconv.ParseInt: parsing "bad": invalid syntax`),
		},
		{
			name:         `fail then environment value is more then then int max value`,
			setEnv:       true,
			envValue:     "12345678901234567890",
			defaultValue: 321,
			expErr:       errors.New(`strconv.ParseInt: parsing "12345678901234567890": value out of range`),
		},
		{
			name:         `use default value then environment value is not set`,
			setEnv:       false,
			defaultValue: 321,
			expRes:       321,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if err := os.Unsetenv("VALUE"); err != nil {
					t.Errorf("coudn't unset VALUE: %s", err)
				}
			}()

			if tc.setEnv {
				if err := os.Setenv("VALUE", tc.envValue); err != nil {
					t.Fatal(err)
				}
			}

			res, err := Int64Strict("VALUE", tc.defaultValue)
			if fmt.Sprint(err) != fmt.Sprint(tc.expErr) {
				t.Errorf("expected error: %v, got: %v", tc.expErr, err)
			}
			if res != tc.expRes {
				t.Errorf("expected value: %d, got: %d", tc.expRes, res)
			}
		})
	}
}

func TestString(t *testing.T) {
	for _, tc := range []struct {
		name         string
		setEnv       bool
		envValue     string
		defaultValue string
		expRes       string
	}{
		{
			name:         `"test" then environment value is "test"`,
			setEnv:       true,
			envValue:     "test",
			defaultValue: "default",
			expRes:       "test",
		},
		{
			name:         `"" then environment value is ""`,
			setEnv:       true,
			envValue:     "",
			defaultValue: "default",
			expRes:       "",
		},
		{
			name:         "use default value then environment value is not set",
			setEnv:       false,
			defaultValue: "default",
			expRes:       "default",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if err := os.Unsetenv("VALUE"); err != nil {
					t.Errorf("coudn't unset VALUE: %s", err)
				}
			}()

			if tc.setEnv {
				if err := os.Setenv("VALUE", tc.envValue); err != nil {
					t.Fatal(err)
				}
			}

			res := String("VALUE", tc.defaultValue)
			if res != tc.expRes {
				t.Errorf("expected value: %s, got: %s", tc.expRes, res)
			}
		})
	}
}

func TestUint(t *testing.T) {
	for _, tc := range []struct {
		name         string
		setEnv       bool
		envValue     string
		defaultValue uint
		expRes       uint
	}{
		{
			name:         `123 then environment value is "123"`,
			setEnv:       true,
			envValue:     "123",
			defaultValue: 321,
			expRes:       123,
		},
		{
			name:         `use default value then environment value is "-1"`,
			setEnv:       true,
			envValue:     "-1",
			defaultValue: 321,
			expRes:       321,
		},
		{
			name:         `use default value then environment value is "3.1"`,
			setEnv:       true,
			envValue:     "3.1",
			defaultValue: 321,
			expRes:       321,
		},
		{
			name:         `0 then environment value is "0"`,
			setEnv:       true,
			envValue:     "0",
			defaultValue: 321,
			expRes:       0,
		},
		{
			name:         `use default value then environment value is ""`,
			setEnv:       true,
			envValue:     "",
			defaultValue: 321,
			expRes:       321,
		},
		{
			name:         `use default value then environment value is "bad"`,
			setEnv:       true,
			envValue:     "bad",
			defaultValue: 321,
			expRes:       321,
		},
		{
			name:         `12345678901234567890 then environment value is "12345678901234567890"`,
			setEnv:       true,
			envValue:     "12345678901234567890",
			defaultValue: 321,
			expRes:       12345678901234567890,
		},
		{
			name:         `use default value then environment value is more then then uint max value`,
			setEnv:       true,
			envValue:     "123456789012345678901",
			defaultValue: 321,
			expRes:       321,
		},
		{
			name:         "use default value then environment value is not set",
			setEnv:       false,
			defaultValue: 321,
			expRes:       321,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if err := os.Unsetenv("VALUE"); err != nil {
					t.Errorf("coudn't unset VALUE: %s", err)
				}
			}()

			if tc.setEnv {
				if err := os.Setenv("VALUE", tc.envValue); err != nil {
					t.Fatal(err)
				}
			}

			res := Uint("VALUE", tc.defaultValue)
			if res != tc.expRes {
				t.Errorf("expected value: %d, got: %d", tc.expRes, res)
			}
		})
	}
}

func TestUintStrict(t *testing.T) {
	for _, tc := range []struct {
		name         string
		setEnv       bool
		envValue     string
		defaultValue uint
		expRes       uint
		expErr       error
	}{
		{
			name:         `123 from environment as "123"`,
			setEnv:       true,
			envValue:     "123",
			defaultValue: 321,
			expRes:       123,
		},
		{
			name:         `fail then environment value is "-1"`,
			setEnv:       true,
			envValue:     "-1",
			defaultValue: 321,
			expErr:       errors.New(`strconv.ParseUint: parsing "-1": invalid syntax`),
		},
		{
			name:         `fail then environment value is "3.1"`,
			setEnv:       true,
			envValue:     "3.1",
			defaultValue: 321,
			expErr:       errors.New(`strconv.ParseUint: parsing "3.1": invalid syntax`),
		},
		{
			name:         `0 then environment value is "0"`,
			setEnv:       true,
			envValue:     "0",
			defaultValue: 321,
			expRes:       0,
		},
		{
			name:         `fail then environment value is ""`,
			setEnv:       true,
			envValue:     "",
			defaultValue: 321,
			expErr:       errors.New(`strconv.ParseUint: parsing "": invalid syntax`),
		},
		{
			name:         `fail then environment value is "bad"`,
			setEnv:       true,
			envValue:     "bad",
			defaultValue: 321,
			expErr:       errors.New(`strconv.ParseUint: parsing "bad": invalid syntax`),
		},
		{
			name:         `12345678901234567890 then environment value is "12345678901234567890"`,
			setEnv:       true,
			envValue:     "12345678901234567890",
			defaultValue: 321,
			expRes:       12345678901234567890,
		},
		{
			name:         `fail then environment value is more then then uint max value`,
			setEnv:       true,
			envValue:     "123456789012345678901",
			defaultValue: 321,
			expErr:       errors.New(`strconv.ParseUint: parsing "123456789012345678901": value out of range`),
		},
		{
			name:         `use default value then environment value is not set`,
			setEnv:       false,
			defaultValue: 321,
			expRes:       321,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if err := os.Unsetenv("VALUE"); err != nil {
					t.Errorf("coudn't unset VALUE: %s", err)
				}
			}()

			if tc.setEnv {
				if err := os.Setenv("VALUE", tc.envValue); err != nil {
					t.Fatal(err)
				}
			}

			res, err := UintStrict("VALUE", tc.defaultValue)
			if fmt.Sprint(err) != fmt.Sprint(tc.expErr) {
				t.Errorf("expected error: %v, got: %v", tc.expErr, err)
			}
			if res != tc.expRes {
				t.Errorf("expected value: %d, got: %d", tc.expRes, res)
			}
		})
	}
}

func TestUint64(t *testing.T) {
	for _, tc := range []struct {
		name         string
		setEnv       bool
		envValue     string
		defaultValue uint64
		expRes       uint64
	}{
		{
			name:         `123 then environment value is "123"`,
			setEnv:       true,
			envValue:     "123",
			defaultValue: 321,
			expRes:       123,
		},
		{
			name:         `use default value then environment value is "-1"`,
			setEnv:       true,
			envValue:     "-1",
			defaultValue: 321,
			expRes:       321,
		},
		{
			name:         `use default value then environment value is "3.1"`,
			setEnv:       true,
			envValue:     "3.1",
			defaultValue: 321,
			expRes:       321,
		},
		{
			name:         `0 then environment value is "0"`,
			setEnv:       true,
			envValue:     "0",
			defaultValue: 321,
			expRes:       0,
		},
		{
			name:         `use default value then environment value is ""`,
			setEnv:       true,
			envValue:     "",
			defaultValue: 321,
			expRes:       321,
		},
		{
			name:         `use default value then environment value is "bad"`,
			setEnv:       true,
			envValue:     "bad",
			defaultValue: 321,
			expRes:       321,
		},
		{
			name:         `12345678901234567890 then environment value is "12345678901234567890"`,
			setEnv:       true,
			envValue:     "12345678901234567890",
			defaultValue: 321,
			expRes:       12345678901234567890,
		},
		{
			name:         `use default value then environment value is more then then uint max value`,
			setEnv:       true,
			envValue:     "123456789012345678901",
			defaultValue: 321,
			expRes:       321,
		},
		{
			name:         "use default value then environment value is not set",
			setEnv:       false,
			defaultValue: 321,
			expRes:       321,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if err := os.Unsetenv("VALUE"); err != nil {
					t.Errorf("coudn't unset VALUE: %s", err)
				}
			}()

			if tc.setEnv {
				if err := os.Setenv("VALUE", tc.envValue); err != nil {
					t.Fatal(err)
				}
			}

			res := Uint64("VALUE", tc.defaultValue)
			if res != tc.expRes {
				t.Errorf("expected value: %d, got: %d", tc.expRes, res)
			}
		})
	}
}

func TestUint64Strict(t *testing.T) {
	for _, tc := range []struct {
		name         string
		setEnv       bool
		envValue     string
		defaultValue uint64
		expRes       uint64
		expErr       error
	}{
		{
			name:         `123 from environment as "123"`,
			setEnv:       true,
			envValue:     "123",
			defaultValue: 321,
			expRes:       123,
		},
		{
			name:         `fail then environment value is "-1"`,
			setEnv:       true,
			envValue:     "-1",
			defaultValue: 321,
			expErr:       errors.New(`strconv.ParseUint: parsing "-1": invalid syntax`),
		},
		{
			name:         `fail then environment value is "3.1"`,
			setEnv:       true,
			envValue:     "3.1",
			defaultValue: 321,
			expErr:       errors.New(`strconv.ParseUint: parsing "3.1": invalid syntax`),
		},
		{
			name:         `0 then environment value is "0"`,
			setEnv:       true,
			envValue:     "0",
			defaultValue: 321,
			expRes:       0,
		},
		{
			name:         `fail then environment value is ""`,
			setEnv:       true,
			envValue:     "",
			defaultValue: 321,
			expErr:       errors.New(`strconv.ParseUint: parsing "": invalid syntax`),
		},
		{
			name:         `fail then environment value is "bad"`,
			setEnv:       true,
			envValue:     "bad",
			defaultValue: 321,
			expErr:       errors.New(`strconv.ParseUint: parsing "bad": invalid syntax`),
		},
		{
			name:         `12345678901234567890 then environment value is "12345678901234567890"`,
			setEnv:       true,
			envValue:     "12345678901234567890",
			defaultValue: 321,
			expRes:       12345678901234567890,
		},
		{
			name:         `fail then environment value is more then then uint max value`,
			setEnv:       true,
			envValue:     "123456789012345678901",
			defaultValue: 321,
			expErr:       errors.New(`strconv.ParseUint: parsing "123456789012345678901": value out of range`),
		},
		{
			name:         `use default value then environment value is not set`,
			setEnv:       false,
			defaultValue: 321,
			expRes:       321,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if err := os.Unsetenv("VALUE"); err != nil {
					t.Errorf("coudn't unset VALUE: %s", err)
				}
			}()

			if tc.setEnv {
				if err := os.Setenv("VALUE", tc.envValue); err != nil {
					t.Fatal(err)
				}
			}

			res, err := Uint64Strict("VALUE", tc.defaultValue)
			if fmt.Sprint(err) != fmt.Sprint(tc.expErr) {
				t.Errorf("expected error: %v, got: %v", tc.expErr, err)
			}
			if res != tc.expRes {
				t.Errorf("expected value: %d, got: %d", tc.expRes, res)
			}
		})
	}
}
