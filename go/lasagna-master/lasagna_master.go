package lasagna

// PreparationTime returns the total preparation time of all the layers
func PreparationTime(layers []string, timePerLayer int) int {
	if timePerLayer == 0 {
		timePerLayer = 2
	}
	return timePerLayer * len(layers)
}

func Quantities(layers []string) (noodles int, sauce float64) {
	for _, v := range layers {
		if v == "noodles" {
			noodles += 50
		}
		if v == "sauce" {
			sauce += 0.2
		}
	}
	return
}

func AddSecretIngredient(layers []string) []string {
	return append(layers, "secret ingredient")
}

func ScaleRecipe(quantities []float64, portionSize float64) []float64 {
	res := make([]float64, len(quantities))
	for i, v := range quantities {
		res[i] = v * portionSize
	}
	return res
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
