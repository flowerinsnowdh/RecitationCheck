package util

import "math/rand/v2"

func FisherYates[T any](slice []T) {
    for i, _ := range slice {
        var targetIndex int = rand.IntN(len(slice))
        slice[i], slice[targetIndex] = slice[targetIndex], slice[i]
    }
}
