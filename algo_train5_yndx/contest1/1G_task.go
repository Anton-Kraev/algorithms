package main

import (
	"fmt"
	"math"
)

func Main1G() int {
	var myCnt, hp, enemyNew int
	fmt.Scan(&myCnt, &hp, &enemyNew)
	return Solve1G(myCnt, hp, enemyNew)
}

func Solve1G(myCnt, hp, enemyNew int) int {
	hp -= myCnt
	if hp <= 0 {
		return 1
	}

	fibLimit := 1.618
	rounds := 1
	enemyCnt := enemyNew
	bestRounds := -1

	for ; (hp > 0 || enemyCnt > 0) && myCnt > 0; rounds++ {
		stepsToFinish := hp/myCnt + 1
		least := myCnt - hp%myCnt
		if hp%myCnt == 0 {
			stepsToFinish--
			least -= myCnt
		}
		enemyCntAfterFinish := enemyCnt + (stepsToFinish-1)*enemyNew
		damageBeforeFinish := (enemyCnt+enemyCntAfterFinish)*stepsToFinish/2 - least
		enemyCntAfterFinish -= least
		if int(float64(myCnt-damageBeforeFinish)*fibLimit) >= enemyCntAfterFinish {
			if bestRounds == -1 {
				bestRounds = rounds + stepsToFinish + lastSteps(myCnt-damageBeforeFinish, enemyCntAfterFinish)
			} else {
				bestRounds = int(math.Min(float64(bestRounds), float64(rounds+stepsToFinish+lastSteps(myCnt-damageBeforeFinish, enemyCntAfterFinish))))
			}
		}

		dmg := int(math.Min(float64(myCnt), float64(enemyCnt)))
		enemyCnt -= dmg
		hp -= myCnt - dmg
		myCnt -= enemyCnt

		if hp > 0 {
			enemyCnt += enemyNew
		}
		if enemyNew == myCnt && enemyCnt == myCnt && hp > 0 {
			break
		}
	}

	return bestRounds
}

func lastSteps(myCnt, enemyCnt int) int {
	rounds := 0
	for ; enemyCnt > 0 && myCnt > 0; rounds++ {
		enemyCnt -= myCnt
		myCnt -= enemyCnt
	}
	return rounds
}
