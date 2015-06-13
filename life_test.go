package lifeofgame

import (
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func match(t *testing.T, reg, str string) {
	pattern := regexp.MustCompile(reg)
	if !pattern.MatchString(str) {
		t.Fatal("\n" + str + "\nnot matched")
	}
}

func TestEmpty(t *testing.T) {
	assert.Equal(t, "0", Next("0"))
}

func TestLiveToDead(t *testing.T) {

	assert.Equal(t, "0", Next("1"))
}

func TestLiveToDead2(t *testing.T) {
	assert.Equal(t, "00", Next("10"))
}

func TestRegxWorks(t *testing.T) {
	match(t, `.1\n..`, "01\n01")
}

func TestKeepLive(t *testing.T) {
	match(t, `.1\n..`, Next("01\n11"))
}

func TestBorn(t *testing.T) {
	match(t, `1.\n..`, Next("01\n11"))
}

func TestDeadCrowd(t *testing.T) {
	match(t, `.0.\n...`, Next("111\n101"))
}

func Test2Cells_keep_dead(t *testing.T) {
	match(t, `10.\n...`, Next("111\n101"))
}
