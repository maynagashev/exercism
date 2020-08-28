package perfect

import "errors"

type Classification uint8

const (
	ClassificationDeficient Classification = iota + 1
	ClassificationAbundant
	ClassificationPerfect
)

var ErrOnlyPositive error = errors.New("expected only positive number")

func Classify(n int64) (class Classification, err error) {

	if n<1 {
		return 0, ErrOnlyPositive
	}

	return ClassificationPerfect, nil
}
