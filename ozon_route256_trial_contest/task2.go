package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var (
		initString  string
		stickersCnt int
	)
	fmt.Fscanln(in, &initString)
	fmt.Fscanln(in, &stickersCnt)

	result := []byte(initString)
	for i := 0; i < stickersCnt; i++ {
		var (
			l, r    int
			sticker string
		)
		fmt.Fscanln(in, &l, &r, &sticker)

		for j := l; j <= r; j++ {
			result[j-1] = sticker[j-l]
		}
	}
	fmt.Fprintln(out, string(result))
}
