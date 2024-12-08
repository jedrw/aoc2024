package answer

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Parse(filePath string) [][]string {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	grid := [][]string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, strings.Split(strings.Trim(scanner.Text(), " "), ""))
	}

	return grid
}

type XY struct {
	x int
	y int
}

func Compute(input [][]string) int {
	antennas := map[string][]XY{}

	for y, row := range input {
		for x, cell := range row {
			if cell != "." {
				antennas[cell] = append(antennas[cell], XY{x, y})
			}
		}
	}

	inBounds := func(pos XY) bool {
		return pos.y >= 0 &&
			pos.y < len(input) &&
			pos.x >= 0 &&
			pos.x < len(input[0])

	}

	antinodes := map[string]struct{}{}
	for _, antennas := range antennas {
		for i, antenna := range antennas {
			for j, dup := range antennas {
				if i == j {
					continue
				}

				offset := XY{x: dup.x - antenna.x, y: dup.y - antenna.y}
				antinode := XY{x: dup.x + offset.x, y: dup.y + offset.y}
				if inBounds(antinode) {
					antinodes[fmt.Sprintf("%d,%d", antinode.x, antinode.y)] = struct{}{}
				}
			}
		}
	}

	return len(antinodes)
}
