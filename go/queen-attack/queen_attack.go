package queenattack

import (
	"errors"
	"math"
)

func CanQueenAttack(whitePosition, blackPosition string) (bool, error) {
	// Validate input positions
	if whitePosition == blackPosition {
		return false, errors.New("queens cannot occupy the same position")
	}
	if len(whitePosition) != 2 || len(blackPosition) != 2 {
		return false, errors.New("invalid position")
	}

	// Convert positions to coordinates
	whiteFile, whiteRank := int(whitePosition[0]-'a'), int(whitePosition[1]-'1')
	blackFile, blackRank := int(blackPosition[0]-'a'), int(blackPosition[1]-'1')

	// Validate coordinates are on the board
	if whiteFile < 0 || whiteFile > 7 || whiteRank < 0 || whiteRank > 7 || blackFile < 0 || blackFile > 7 || blackRank < 0 || blackRank > 7 {
		return false, errors.New("position off board")
	}

	// Check for attack possibilities
	sameRank := whiteRank == blackRank
	sameFile := whiteFile == blackFile
	// Angle between the two points is 45 degrees if the difference in the x and y coordinates is the same
	sameDiagonal := math.Abs(float64(whiteFile-blackFile)) == math.Abs(float64(whiteRank-blackRank))

	return sameRank || sameFile || sameDiagonal, nil
}
