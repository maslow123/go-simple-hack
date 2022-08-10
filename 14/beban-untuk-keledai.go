package beban_untuk_keledai

import (
	"fmt"
	"sort"
	"strings"

	"github.com/maslow123/go-simple-hack/utils"
)

func BebanUntukKeledai(item string) int {
	items := strings.Split(item, " ")

	if len(items) < 2 || len(items) > 100 {
		fmt.Println("Invalid input the data")
		return -1
	}

	listItem := utils.SliceStringToSliceInt(items)
	sort.Slice(listItem, func(i, j int) bool {
		return listItem[j] < listItem[i]
	})
	leftBucketTotal, rightBucketTotal := getBucketTotal(listItem)

	diff := leftBucketTotal - rightBucketTotal
	if leftBucketTotal < rightBucketTotal {
		diff = rightBucketTotal - leftBucketTotal
	}

	fmt.Printf("Diff: %d\n", diff)
	return diff
}

func sumBucketValue(bucket []int) int {
	total := 0
	for _, val := range bucket {
		total += val
	}

	return total
}

func getBucketTotal(listItem []int) (int, int) {
	var leftBucket, rightBucket []int
	var leftBucketTotal, rightBucketTotal int

	for i, val := range listItem {
		if i%2 == 0 && (leftBucketTotal == 0) || (leftBucketTotal <= rightBucketTotal) {
			leftBucket = append(leftBucket, val)
			leftBucketTotal = sumBucketValue(leftBucket)
			continue
		}
		rightBucket = append(rightBucket, val)
		rightBucketTotal = sumBucketValue(rightBucket)
	}

	fmt.Printf("Left bucket: %v\n", leftBucket)
	fmt.Printf("Left bucket total: %d\n", leftBucketTotal)

	fmt.Printf("Right bucket: %v\n", rightBucket)
	fmt.Printf("Right bucket total: %d\n", rightBucketTotal)

	return leftBucketTotal, rightBucketTotal
}
