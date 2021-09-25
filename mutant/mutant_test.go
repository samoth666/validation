package mutant

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsMutantFail(t *testing.T) {
	c := require.New(t)

	dna := []string{
		"ATATAT",
		"TGGCGC",
		"TATATA",
		"GCGCGA",
		"ATATAC",
		"GCGCGC",
	}

	c.Equal(false, IsMutant(dna))

	defer func() {
		dna = []string{}
	}()
}

func TestIsMutantOK(t *testing.T) {
	c := require.New(t)

	dna := []string{
		"ATGCGA",
		"CAGTGC",
		"TTATGT",
		"AGAAGG",
		"CCCCTA",
		"TCACTG",
	}

	c.Equal(true, IsMutant(dna))

	defer func() {
		dna = []string{}
	}()
}
