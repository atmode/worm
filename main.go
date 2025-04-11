package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	a := []int{1, 2, 3, 1}
	fmt.Println(a)
	fmt.Println(containsDuplicate(a))

	arr1 := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	fmt.Println("missing number :", missingNumber(arr1))

	arr2 := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Println("Missing numbers:", findDisappearedNumbers(arr2))

	arr3 := []int{2, 11, 7, 15}
	target := 9
	result := twoSum(arr3, target)
	fmt.Println("Indices:", result)

	fmt.Println(smallerNumbersThanCurrent([]int{8, 1, 2, 2, 3}))
	fmt.Println(smallerNumbersThanCurrent2([]int{8, 1, 2, 2, 3}))

	points1 := [][]int{{1, 1}, {3, 4}, {-1, 0}}
	result1 := minTimeToVisitAllPoints(points1)
	fmt.Println("Minimum time for points1:", result1)

	result2 := minTimeToVisitAllPoints(points1)
	fmt.Println("Minimum time for points1:", result2)

	points2 := [][]int{{3, 2}, {-2, 2}}
	fmt.Println("Minimum time for points2 ", minTimeToVisitAllPoints2(points2))

	//island of numbers
	grid := [][]byte{
		[]byte("11000"),
		[]byte("11000"),
		[]byte("00100"),
		[]byte("00011"),
	}
	fmt.Println("Number of islands:", numIsLands(grid))

}
func containsDuplicate(arr []int) bool {
	seen := make(map[int]bool)
	for _, value := range arr {
		if _, exists := seen[value]; exists {
			return true
		}
		seen[value] = true
	}
	return false

}

func missingNumber(nums []int) int {
	n := len(nums)
	total := (n * (n + 1)) / 2
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return total - sum
}

func findDisappearedNumbers(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		index := abs(nums[i]) - 1
		if nums[index] > 0 {
			nums[index] = -nums[index] // mark as seen by making it negative
		}
	}
	var result []int
	for i, nums := range nums {
		if nums > 0 {
			result = append(result, i+1)
		}
	}

	return result
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		complement := target - num // that would add up with nu;m to reach the target.
		if j, found := m[complement]; found {
			return []int{j, i}
		}
		//if not found, store the current number and its index in the map.
		m[num] = i
	}
	// if not solution is found return nil
	return nil

}

func smallerNumbersThanCurrent(nums []int) []int {
	sortedNums := make([]int, len(nums))
	copy(sortedNums, nums)
	sort.Ints(sortedNums)

	numToSmallerCount := make(map[int]int)
	for i, num := range sortedNums {
		if _, ok := numToSmallerCount[num]; !ok {
			numToSmallerCount[num] = i
		}
	}
	result := make([]int, len(nums))
	for i, num := range nums {
		result[i] = numToSmallerCount[num]
	}
	return result
}
func smallerNumbersThanCurrent2(nums []int) (ans []int) {
	arr := make([]int, len(nums))
	copy(arr, nums)
	sort.Ints(arr)

	for i, x := range nums {
		nums[i] = sort.SearchInts(arr, x)
	}
	return nums
}

func minTimeToVisitAllPoints(points [][]int) (ans int) {
	for i, p := range points[1:] {
		dx := abs(p[0] - points[i][0])
		dy := abs(p[1] - points[i][1])
		ans += max(dx, dy)
	}
	return
}

func minTimeToVisitAllPoints2(points [][]int) int {
	totalTime := 0
	for i := 1; i < len(points); i++ {
		xDist := int(math.Abs(float64(points[i][0] - points[i-1][0])))
		yDist := int(math.Abs(float64(points[i][1] - points[i-1][1])))
		totalTime += xDist + yDist
	}
	return totalTime

}

func numIsLands(grid [][]byte) int {

	if len(grid) == 0 {
		return 0
	}

	rows := len(grid)
	cols := len(grid[0])
	count := 0
	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r < 0 || r >= rows || c < 0 || c >= cols || grid[r][c] != '1' {
			return
		}
		grid[r][c] = '0'
		dfs(r+1, c) // down
		dfs(r-1, c) // up
		dfs(r, c+1) // right
		dfs(r, c-1) // left

	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			// if we find an unvisited land cell, perform DFS
			if grid[i][j] == '1' {
				dfs(i, j)
				count++
			}
		}
	}

	return count
}
