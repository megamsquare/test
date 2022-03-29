package main

import "fmt"

func main() {
  fmt.Println(twoOldestAges([]int{1, 2, 10, 8}))
  fmt.Println(twoOldestAges([]int{1, 5, 87, 45, 8, 8}))
}

func twoOldestAges(ageArray []int) [2]int {
  var oldestAge int
  var secondOldestAge int
  for _, v := range ageArray {
    if v > oldestAge {
      secondOldestAge = oldestAge
      oldestAge = v
    } else if v > secondOldestAge {
      secondOldestAge = v
    }
  }
  result := [2]int{oldestAge, secondOldestAge}
  return result

}