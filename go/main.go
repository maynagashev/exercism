package main

import (
	"errors"
	"fmt"
)

type Classification uint8

const (
	ClassificationDeficient Classification = iota + 1
	ClassificationAbundant
	ClassificationPerfect
)

var ErrOnlyPositive error = errors.New("expected only positive number")

func main() {
	Classify(1)

	fmt.Printf("def: %T %v \nabund: %T %v \nperfect %T %v",
		ClassificationDeficient, ClassificationDeficient,
		ClassificationAbundant, ClassificationAbundant,
		ClassificationPerfect, ClassificationPerfect,
	)
}

func Classify(n int64) (class Classification, err error) {

	if n < 1 {
		return 0, ErrOnlyPositive
	}

	return
}
