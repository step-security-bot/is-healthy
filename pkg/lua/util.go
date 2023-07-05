package lua

import (
	"github.com/gobwas/glob"
)

func Match(pattern, text string, separators ...rune) bool {
	compiledGlob, err := glob.Compile(pattern, separators...)
	if err != nil {
		return false
	}
	return compiledGlob.Match(text)
}
