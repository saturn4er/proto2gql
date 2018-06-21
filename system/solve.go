package main

import (
	"fmt"
)

func main() {
	fmt.Println(systemOddsOld([]float64{1, 2, 3, 4, 5}, 4))
	fmt.Println(systemOddsNew([]float64{1, 2, 3, 4, 5}, 4))

}

type stackItem struct {
	Odds   []float64
	Result float64
}

func (stackItem stackItem) String() string {
	return fmt.Sprint("{Odds:", stackItem.Odds, ",Value:", stackItem.Result, "}")
}

func systemOddsNew(outcomesOdds []float64, dimensionMin int) float64 {
	var stack = make([]stackItem, dimensionMin-1)
	stackPos := -1
	down := true
	tempRes := float64(0)
	odds := outcomesOdds
	for {
		if down {
			tempRes = 0
			stackPos++
			if stackPos+2 == dimensionMin {
				for _, s := range odds[1:] {
					tempRes += s
				}
				tempRes *= odds[0]
				down = false
				fmt.Println("tempRes=", tempRes)
				continue
			}
			stack[stackPos] = stackItem{
				Odds: odds,
			}
			odds = odds[1:]
			fmt.Println(stack[:])
		} else {
			// restoring parent context
			stackPos--
			if stackPos < 0 {
				return stack[0].Result
			}

			fmt.Println(stackPos, "+=", tempRes, "*", stack[stackPos].Odds[0])
			stack[stackPos].Result += tempRes * stack[stackPos].Odds[0]
			length := stackPos + 1
			odds = stack[stackPos].Odds
			if length+len(odds[1:]) >= dimensionMin {
				odds = odds[2:]
				down = true
				continue
			}
			tempRes = stack[stackPos].Result
			continue

			fmt.Println(stackPos)
			fmt.Println(length)
			fmt.Println(odds)
			return 1
			// if stackPos < 0 {
			// 	return tempRes
			// }
			// stack[stackPos].Result += stack[stackPos].TopValue * tempRes
			// if stackPos+1+len(odds) >= dimensionMin {
			// 	stack[stackPos].Odds = stack[stackPos].Odds[1:]
			// 	stack[stackPos].TopValue = stack[stackPos].Odds[0]
			// 	down = true
			// 	odds = stack[stackPos].Odds[1:]
			// 	fmt.Println("down again", stack)
			// 	continue
			// }
			// tempRes = stack[stackPos].Result
			// stackPos--
			// odds = stack[stackPos].Odds
			// stack[stackPos].Result += tempRes * stack[stackPos].TopValue
			// if stackPos+1+len(odds) > dimensionMin {
			// 	stack[stackPos].Odds = stack[stackPos].Odds[1:]
			// 	stack[stackPos].TopValue = stack[stackPos].Odds[0]
			// 	odds = stack[stackPos].Odds[1:]
			// 	fmt.Println("Moving up", stack)
			// 	continue
			// }
			//
			// fmt.Println(stackPos, stack)
			// return stack[0].Result
			// fmt.Println(stack[stackPos].Result, "+=", stack[stackPos].TopValue, "*", tempRes, "=", stack[stackPos].Result+stack[stackPos].TopValue*tempRes)
			// stack[stackPos].Result += stack[stackPos].TopValue * tempRes
			// if stackPos > 0 && len(outcomesOdds)-dimensionMin-stack[stackPos].Odds >= 0 {
			// 	stack[stackPos].Odds++
			// 	stack[stackPos].TopValue = outcomesOdds[stack[stackPos].Odds]
			// 	from = stack[stackPos].Odds + 1
			// 	down = true
			// 	fmt.Println(stackPos, "down again", stack)
			// 	continue
			// }
			// tempRes = stack[stackPos].Result
			// stackPos--
			// if stackPos < 0 {
			// 	return tempRes
			// }
		}
	}
}

func systemOddsOld(odds []float64, dimensionMin int) float64 {
	var calc func(topValue float64, odds []float64, length int) float64
	calc = func(topValue float64, odds []float64, length int) (res float64) {
		if length == 1 {
			for _, o := range odds {
				res += o
			}
			return res
		}
		for i, o := range odds {
			if len(odds[i+1:]) > length-2 && o != 0 {
				res += o * calc(topValue*o, odds[i+1:], length-1)
			}
		}
		return res
	}
	return calc(1, odds, dimensionMin)
}
