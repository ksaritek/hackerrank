/*
You are given an unordered array consisting of consecutive integers  [1, 2, 3, ..., n] without any duplicates. You are allowed to swap any two elements. Find the minimum number of swaps required to sort the array in ascending order.

# Example

Perform the following steps:

i   arr                         swap (indices)
0   [7, 1, 3, 2, 4, 5, 6]   swap (0,3)
1   [2, 1, 3, 7, 4, 5, 6]   swap (0,1)
2   [1, 2, 3, 7, 4, 5, 6]   swap (3,4)
3   [1, 2, 3, 4, 7, 5, 6]   swap (4,5)
4   [1, 2, 3, 4, 5, 7, 6]   swap (5,6)
5   [1, 2, 3, 4, 5, 6, 7]
It took  swaps to sort the array.

# Function Description

Complete the function minimumSwaps in the editor below.

minimumSwaps has the following parameter(s):

int arr[n]: an unordered array of integers
Returns

int: the minimum number of swaps to sort the array
Input Format

The first line contains an integer, , the size of .
The second line contains  space-separated integers .

# Constraints

# Sample Input 0

4
4 3 1 2
Sample Output 0

3
Explanation 0

Given array
After swapping  we get
After swapping  we get
After swapping  we get
So, we need a minimum of  swaps to sort the array in ascending order.

# Sample Input 1

5
2 3 4 1 5
Sample Output 1

3
Explanation 1

Given array
After swapping  we get
After swapping  we get
After swapping  we get
So, we need a minimum of  swaps to sort the array in ascending order.

# Sample Input 2

7
1 3 5 2 4 6 7
Sample Output 2

3
Explanation 2

Given array
After swapping  we get
After swapping  we get
After swapping  we get
So, we need a minimum of  swaps to sort the array in ascending order.
*/
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the minimumSwaps function below.
func minimumSwaps(arr []int32) int32 {
	var swaps int32
	var reorder bool = true

	for reorder {
		reorder = false
		for i := 0; i < len(arr); {
			expectedIndex := arr[i] - 1
			if arr[i] != arr[expectedIndex] {
				arr[i], arr[expectedIndex] = arr[expectedIndex], arr[i]
				swaps++
				reorder = true
			} else {
				i++
			}
		}

	}

	return swaps
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	res := minimumSwaps(arr)

	fmt.Fprintf(writer, "%d\n", res)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
