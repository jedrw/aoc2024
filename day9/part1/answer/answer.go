package answer

import (
	"bufio"
	"os"
)

func Parse(filePath string) []int {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	disk := []int{}
	scanner := bufio.NewScanner(file)
	id := 0
	for scanner.Scan() {
		diskMap := scanner.Text()
		for i := 0; i <= len(diskMap)-1; i++ {
			numBlocks := int(diskMap[i] - '0')
			if i%2 == 0 {
				for range numBlocks {
					disk = append(disk, id)
				}

				id++
			} else {
				for range numBlocks {
					disk = append(disk, -1)
				}
			}
		}
	}

	return disk
}

func countFree(disk []int) int {
	free := 0
	for _, block := range disk {
		if block < 0 {
			free++
		}
	}

	return free
}

func fillBlock(index int, start int, disk []int) int {
	for j := start; j >= (len(disk) - countFree(disk)); j-- {
		if disk[j] >= 0 {
			disk[index] = disk[j]
			disk[j] = -1
			return j
		}
	}

	panic("fuck")
}

func Compute(disk []int) int {
	rStart := len(disk) - 1
	for i := 0; i < (len(disk) - countFree(disk)); i++ {
		if disk[i] < 0 {
			rStart = fillBlock(i, rStart, disk)
		}
	}

	checksum := 0
	for i, block := range disk {
		if block >= 0 {
			checksum += block * i
		}
	}

	return checksum
}
