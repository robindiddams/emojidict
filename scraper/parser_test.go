package scraper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNoParser(t *testing.T) {
	assert := assert.New(t)
	l, err := ParseEmojis("99.99", TestFileName, "")
	assert.Error(err)
	assert.Nil(l)
}

func TestParseNone(t *testing.T) {
	assert := assert.New(t)
	l, err := ParseEmojis("13.0", TestFileName, `# ================================================`)
	assert.NoError(err)
	assert.Nil(l)
}

func TestParseClean(t *testing.T) {
	assert := assert.New(t)
	l, err := ParseEmojis("13.0", TestFileName, `1F44E 1F3FC                                ; fully-qualified     # 👎🏼 E1.0 thumbs down: medium-light skin tone`)
	assert.NoError(err)
	assert.EqualValues("thumbs down: medium-light skin tone", l.Name)
	assert.EqualValues("ThumbsDownMediumLightSkinTone", l.ProperName)
	assert.EqualValues([]string{"1f44e", "1f3fc"}, l.CodePoints)
}

func TestParseTestFile13(t *testing.T) {
	assert := assert.New(t)
	l := testFileV13(`1F44E 1F3FC                                ; fully-qualified     # 👎🏼 E1.0 thumbs down: medium-light skin tone`)
	assert.NotNil(l)
	assert.EqualValues("thumbs down: medium-light skin tone", l.Name)
	assert.EqualValues([]string{"1F44E", "1F3FC"}, l.CodePoints)
}
