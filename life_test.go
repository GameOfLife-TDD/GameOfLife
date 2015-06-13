package lifeofgame

import (
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func TestRegxWorks(t *testing.T) {
	match(t, `.1\n..`, "01\n01")
}

func match(t *testing.T, reg, str string) {
	pattern := regexp.MustCompile(reg)
	if !pattern.MatchString(str) {
		t.Error(reg + " not matched \n" + str)
	}
}

func TestEmpty(t *testing.T) {
	assert.Equal(t, "0", Next("0"))
}

func TestLiveToDead(t *testing.T) {
	assert.Equal(t, "0", Next("1"))
	assert.Equal(t, "00", Next("10"))
	match(t, `.0\n..\n..`, Next("01\n00\n11"))
}

func TestKeepLive(t *testing.T) {
	match(t, `.1\n..`, Next("01\n11"))
	match(t, `.1\n..`, Next("01\n11\n11"))
}

func TestBorn(t *testing.T) {
	match(t, `1.\n..`, Next("01\n11"))
}

func TestDeadCrowd(t *testing.T) {
	match(t, `.0.\n...`, Next("111\n101"))
}

