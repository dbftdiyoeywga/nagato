package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	msg := `
YUKI.N>みえてる?

ああ

YUKI.N>そっちの時空間とはまだ完全に連結を絶たれていない。それも時間の問題。そうなれば最後。

どうすりゃいい？

YUKI.N>どうにもならない。情報統合思念体は失望している。これで進化の可能性は失われた。

YUKI.N>涼宮ハルヒは何もない所から情報を生み出す力を持っていた。 
YUKI.N>それは情報統合思念体にもない力。この情報創造能力を解析すれば自律進化への糸口がつかめるかもしれないと考えた。
YUKI.N>あなたに賭ける。

何をだよ

YUKI.N>もう一度こちらへ回帰することを我々は望んでいる。涼宮ハルヒは 重要な観察対象。わたしという個体もあなたには戻ってきて欲しいと感じている。

YUKI.N>また図書館に＿

YUKI.N>sleeping beauty_`

	list := strings.Split(msg, "\n")

	// tb := time.NewTicker(2 * time.Second)
	// ts := time.NewTicker(85 * time.Millisecond)
	for _, object := range list {
		if object == "" {
			// select {
			// case <-tb.C:
			// 	continue
			// }
			time.Sleep(1 * time.Second)
		}
		for _, s := range object {
			// select {
			// case <-ts.C:
			// 	fmt.Printf("%s", string(rune(s)))
			// }
			if string(rune(s)) != "。" {
				fmt.Printf("%s", string(rune(s)))
				time.Sleep(85 * time.Millisecond)
			} else {
				fmt.Printf("%s", string(rune(s)))
				time.Sleep(450 * time.Millisecond)

			}
		}
		fmt.Println("")
	}
	fmt.Println("")
	// 	tb.Stop()
	// 	ts.Stop()
}
