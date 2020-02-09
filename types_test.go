package emojidict

import (
	"fmt"
	"testing"
)

// 1F468 200D 2764 FE0F 200D 1F468             ; RGI_Emoji_ZWJ_Sequence  ; couple with heart: man, man                                    # E2.0   [1] (👨‍❤️‍👨)

func TestEmojiDefinition(t *testing.T) {
	complexemoji := "👨🏼‍🤝‍👨🏻" // 1F468 1F3FC 200D 1F91D 200D 1F468 1F3FB

	for _, r := range complexemoji {
		fmt.Printf("0x%x, ", r)
	}
	fmt.Printf("\n")
	// 1f468 1f3fc 200d 1f91d 200d 1f468 1f3fb
	points := []rune{0x1f468, 0x1f3fc, 0x200d, 0x1f91d, 0x200d, 0x1f468, 0x1f3fb}

	rebuilt := string(points)

	fmt.Println(len(complexemoji), complexemoji)
	fmt.Println(len(rebuilt), rebuilt)

}
