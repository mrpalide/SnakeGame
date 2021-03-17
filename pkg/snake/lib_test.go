package snake

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var initGameFlagTests = []struct {
	in  string
	out GameState
}{
	{"a", GameState{}},
	{"a,1", GameState{}},
	{"1,a", GameState{}},
	{"12", GameState{}},
	{"1,2,3", GameState{}},
	{"1,0", GameState{}},
	{"0,1", GameState{}},
	{"0,0", GameState{}},
	{"1,1", GameState{}},
	{"1,2", GameState{}},
}

func TestInitGame(t *testing.T) {
	assert := assert.New(t)
	// Error in Game Initialization
	for _, test := range initGameFlagTests {
		game, _ := initGame(test.in)
		assert.Equal(game, test.out)
	}

	// Correct Game Initialization
	_, err := initGame("2,2")
	assert.Nil(err)
}
