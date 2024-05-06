package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Bank struct {
	rtod, rtoe, dtor, dtoe, etor, etod float64
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscanln(in, &t)
	for i := 0; i < t; i++ {
		var bankA, bankB, bankC Bank
		bankA = readBankInfo(in)
		bankB = readBankInfo(in)
		bankC = readBankInfo(in)
		results := []float64{
			bankA.rtod,
			bankB.rtod,
			bankC.rtod,

			bankA.rtoe * bankB.etod,
			bankB.rtoe * bankA.etod,
			bankA.rtoe * bankC.etod,
			bankC.rtoe * bankA.etod,
			bankB.rtoe * bankC.etod,
			bankC.rtoe * bankB.etod,

			bankA.rtoe * bankB.etor * bankC.rtod,
			bankB.rtoe * bankA.etor * bankC.rtod,
			bankA.rtoe * bankC.etor * bankB.rtod,
			bankC.rtoe * bankA.etor * bankB.rtod,
			bankB.rtoe * bankC.etor * bankA.rtod,
			bankC.rtoe * bankB.etor * bankA.rtod,

			bankA.rtod * bankB.dtor * bankC.rtod,
			bankB.rtod * bankA.dtor * bankC.rtod,
			bankA.rtod * bankC.dtor * bankB.rtod,
			bankC.rtod * bankA.dtor * bankB.rtod,
			bankB.rtod * bankC.dtor * bankA.rtod,
			bankC.rtod * bankB.dtor * bankA.rtod,

			bankA.rtod * bankB.dtoe * bankC.etod,
			bankB.rtod * bankA.dtoe * bankC.etod,
			bankA.rtod * bankC.dtoe * bankB.etod,
			bankC.rtod * bankA.dtoe * bankB.etod,
			bankB.rtod * bankC.dtoe * bankA.etod,
			bankC.rtod * bankB.dtoe * bankA.etod,
		}
		maxExchange := 0.0
		for _, result := range results {
			maxExchange = math.Max(maxExchange, result)
		}
		fmt.Fprintln(out, maxExchange)
	}
}

func readBankInfo(in *bufio.Reader) Bank {
	var l11, l12, l21, l22, l31, l32, l41, l42, l51, l52, l61, l62 float64
	fmt.Fscanln(in, &l11, &l12)
	fmt.Fscanln(in, &l21, &l22)
	fmt.Fscanln(in, &l31, &l32)
	fmt.Fscanln(in, &l41, &l42)
	fmt.Fscanln(in, &l51, &l52)
	fmt.Fscanln(in, &l61, &l62)
	return Bank{l12 / l11, l22 / l21, l32 / l31, l42 / l41, l52 / l51, l62 / l61}
}
