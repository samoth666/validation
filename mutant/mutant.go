package mutant

func isNitrogenBaseEqual(nb0, nb1, nb2, nb3 byte) bool {
	if nb0 == nb1 &&
		nb0 == nb2 &&
		nb0 == nb3 {
		return true
	}

	return false
}

func findMatches(dna []string) (bool, bool, bool) {
	arrayLength := len(dna[0])
	isMatchVertical := false
	isMatchHorizontal := false
	isMatchOblique := false

	for i := 0; i < arrayLength-3; i++ {
		firstSecuence := dna[i]
		secondSecuence := dna[i+1]
		thirdSecuence := dna[i+2]
		quarterSecuence := dna[i+3]

		for j := 0; j < arrayLength; j++ {
			// Verticals matches
			if isNitrogenBaseEqual(firstSecuence[j], secondSecuence[j], thirdSecuence[j], quarterSecuence[j]) {
				isMatchVertical = true
			}

			if j+3 < arrayLength {
				// Oblique matches
				if isNitrogenBaseEqual(firstSecuence[j], secondSecuence[j+1], thirdSecuence[j+2], quarterSecuence[j+3]) {
					isMatchOblique = true
				}

				// Horizontal matches
				if isNitrogenBaseEqual(firstSecuence[j], firstSecuence[j+1], firstSecuence[j+2], firstSecuence[j+3]) {
					isMatchHorizontal = true
				}
				if isNitrogenBaseEqual(quarterSecuence[j], quarterSecuence[j+1], quarterSecuence[j+2], quarterSecuence[j+3]) {
					isMatchHorizontal = true
				}
			}
		}
	}
	return isMatchHorizontal, isMatchVertical, isMatchOblique
}

func IsMutant(dna []string) bool {
	isMatchHorizontal, isMatchVertical, isMatchOblique := findMatches(dna)
	if isMatchVertical && isMatchHorizontal || isMatchVertical && isMatchOblique || isMatchHorizontal && isMatchOblique {
		return true
	}

	return false
}
