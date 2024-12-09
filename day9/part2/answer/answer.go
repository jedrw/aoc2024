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

func findFree(disk []int, size int, leftof int) []int {
	free := []int{}
	for i := 0; len(free) < size && i < leftof; i++ {
		if disk[i] >= 0 {
			free = []int{}
			continue
		}

		if disk[i] < 0 {
			free = append(free, i)
		}
	}

	if len(free) < size {
		return []int{}
	}

	return free
}

func Compute(disk []int) int {
	id := 0
	for i := len(disk) - 1; i > 0 && id >= 0; i-- {
		id := disk[i]
		fileBlocks := []int{i}
		for j := i - 1; disk[j] == id && j > 0; j-- {
			fileBlocks = append([]int{j}, fileBlocks...)
		}

		freeBlocks := findFree(disk, len(fileBlocks), fileBlocks[0])
		if len(freeBlocks) > 0 {
			for i := range freeBlocks {
				disk[freeBlocks[i]] = disk[fileBlocks[i]]
				disk[fileBlocks[i]] = -1
			}
		}

		i = i - len(fileBlocks) + 1
	}

	checksum := 0
	for i, block := range disk {
		if block >= 0 {
			checksum += block * i
		}
	}

	return checksum
}
