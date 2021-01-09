package service

import (
	"errors"
	"regexp"
	"time"
)

//FormatTaskTimestamp converts a string date into a valid time field in GO.
func FormatTaskTimestamp(taskString string) (time.Time, error) {

	if taskString == "" {
		return time.Now(), nil
	}
	t, err := time.Parse(time.RFC3339, taskString)

	if err != nil {
		t, err := time.Parse("2020-01-01", taskString)
		if err != nil {
			return t, err
		}

		return t, nil

	}

	return t, nil
}

//ValidateStringField checks if a string match a given pattern
func ValidateStringField(taskLabel string, regex string, errormessage string) error {

	resp, _ := regexp.MatchString(regex, taskLabel)

	if resp == false {
		return errors.New(errormessage)
	}

	return nil
}
