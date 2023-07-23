/*
A person wants to determine the most expensive computer keyboard and USB drive that can be purchased with a give budget. Given price lists for keyboards and USB drives and a budget, find the cost to buy them. If it is not possible to buy both items, return .

# Example

The person can buy a , or a . Choose the latter as the more expensive option and return .

# Function Description

Complete the getMoneySpent function in the editor below.

getMoneySpent has the following parameter(s):

int keyboards[n]: the keyboard prices
int drives[m]: the drive prices
int b: the budget
Returns

int: the maximum that can be spent, or  if it is not possible to buy both items
Input Format

The first line contains three space-separated integers , , and , the budget, the number of keyboard models and the number of USB drive models.
The second line contains  space-separated integers , the prices of each keyboard model.
The third line contains  space-separated integers , the prices of the USB drives.

# Constraints

The price of each item is in the inclusive range .
Sample Input 0

10 2 3
3 1
5 2 8
Sample Output 0

9
Explanation 0

Buy the  keyboard and the  USB drive for a total cost of .

# Sample Input 1

5 1 1
4
5
Sample Output 1

-1
Explanation 1

There is no way to buy one keyboard and one USB drive because , so return .
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

/*
 * Complete the getMoneySpent function below.
 */
func getMoneySpent(keyboards []int32, drives []int32, b int32) int32 {
	var max int32 = -1
	for _, keyboard := range keyboards {
		for _, drive := range drives {
			if keyboard+drive > max && keyboard+drive <= b {
				max = keyboard + drive
			}
		}
	}

	return max
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	bnm := strings.Split(readLine(reader), " ")

	bTemp, err := strconv.ParseInt(bnm[0], 10, 64)
	checkError(err)
	b := int32(bTemp)

	nTemp, err := strconv.ParseInt(bnm[1], 10, 64)
	checkError(err)
	n := int32(nTemp)

	mTemp, err := strconv.ParseInt(bnm[2], 10, 64)
	checkError(err)
	m := int32(mTemp)

	keyboardsTemp := strings.Split(readLine(reader), " ")

	var keyboards []int32

	for keyboardsItr := 0; keyboardsItr < int(n); keyboardsItr++ {
		keyboardsItemTemp, err := strconv.ParseInt(keyboardsTemp[keyboardsItr], 10, 64)
		checkError(err)
		keyboardsItem := int32(keyboardsItemTemp)
		keyboards = append(keyboards, keyboardsItem)
	}

	drivesTemp := strings.Split(readLine(reader), " ")

	var drives []int32

	for drivesItr := 0; drivesItr < int(m); drivesItr++ {
		drivesItemTemp, err := strconv.ParseInt(drivesTemp[drivesItr], 10, 64)
		checkError(err)
		drivesItem := int32(drivesItemTemp)
		drives = append(drives, drivesItem)
	}

	/*
	 * The maximum amount of money she can spend on a keyboard and USB drive, or -1 if she can't purchase both items
	 */

	moneySpent := getMoneySpent(keyboards, drives, b)

	fmt.Fprintf(writer, "%d\n", moneySpent)

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
