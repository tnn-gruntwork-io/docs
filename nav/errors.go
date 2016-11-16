package nav

import "fmt"

// Custom errors

type WrongNumberOfCaptureGroupsFoundInPathRegEx struct {
	inputPath string
	regExName string
	regEx     string
}

func (err WrongNumberOfCaptureGroupsFoundInPathRegEx) Error() string {
	return fmt.Sprintf("The wrong number of capture groups was found. This may be because the path did not match the RegEx.\ninputPath = %s\nregExName = %s\nregEx = %s\n", err.inputPath, err.regExName, err.regEx)
}
