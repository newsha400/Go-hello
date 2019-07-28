package change

import (
	"fmt"
	"log"
	"sort"
)

//GiveChange is used for calling Change from main package for testing purposes
func GiveChange() {
	var coins = []int{2, 5, 10, 20, 50}
	var target = 21
	Change(coins, target)
}

//Change takes a set of coins and a target ammount, and returns a set of minimum number of coins to make change
func Change(coins []int, target int) (expectedChange []int, err error) {
	if target < 0 {
		err = fmt.Errorf("no change available for negative value %v", target)
		log.Println(err)
		return
	}
	if target < coins[0] && target != 0 {
		err = fmt.Errorf("target change %v is smaller than the smallest coin %v", target, coins[0])
		log.Println(err)
		return
	}
	if target == 0 {
		expectedChange = make([]int, 0)
		log.Println(expectedChange)
		return
	}
	var table [][]int
	var targetArray = make([]int, target+1)
	dst := make([]int, len(targetArray))
	copy(dst, targetArray)
	table = append(table, dst)
	for _, coin := range coins {
		for tempTarget := 0; tempTarget <= target; tempTarget++ {
			if tempTarget >= coin {
				if targetArray[tempTarget-coin] == 0 && tempTarget%coin != 0 {
				} else if (targetArray[tempTarget-coin]+1) < targetArray[tempTarget] || targetArray[tempTarget] == 0 {
					targetArray[tempTarget] = targetArray[tempTarget-coin] + 1
				}
			} else if tempTarget == 0 {
				targetArray[tempTarget] = 0
			}

		}
		dst := make([]int, len(targetArray))
		copy(dst, targetArray)
		table = append(table, dst)
	}
	numberOfCoins := table[len(coins)][target]
	changeIsAvailable := false
	for i := 0; i < len(table); i++ {
		if table[i][target] != 0 {
			changeIsAvailable = true
			break
		}
	}
	if !changeIsAvailable {
		err = fmt.Errorf("no combination of coins %v can make the change %v", coins, target)
		log.Println(err)
		return
	}
	expectedChange = make([]int, numberOfCoins)
	row := len(coins)
	col := target
	i := 0
	for i < len(expectedChange) {
		if table[row][col] < table[row-1][col] || table[row-1][col] == 0 {
			expectedChange[i] = coins[row-1]
			i++
			col = col - coins[row-1]
		} else if table[row][col] == table[row-1][col] {
			row = row - 1
		}
	}
	sort.Ints(expectedChange)
	log.Println(expectedChange)
	return expectedChange, err
}

//[ ,  ,   ,   ,   ] [0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21]
//[ ,  ,   ,   ,   ] [0 0 0 0 0 0 0 0 0 0  0  0  0  0  0  0  0  0  0  0  0  0]
//[2,  ,   ,   ,   ] [0 0 1 0 2 0 3 0 4 0  5  0  6  0  7  0  8  0  9  0  10 0]
//[2, 5,   ,   ,   ] [0 0 1 0 2 1 3 2 4 3  2  4  3  5  4  3  5  4  6  5  4  6]
//[2, 5, 10,   ,   ] [0 0 1 0 2 1 3 2 4 3  1  4  2  5  3  2  4  3  5  4  2  5]
//[2, 5, 10, 20,   ] [0 0 1 0 2 1 3 2 4 3  1  4  2  5  3  2  4  3  5  4  1  5]
//[2, 5, 10, 20, 50] [0 0 1 0 2 1 3 2 4 3  1  4  2  5  3  2  4  3  5  4  1  5]
