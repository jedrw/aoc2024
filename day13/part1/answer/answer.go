package answer

import (
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	aPrice = 3
	bPrice = 1
)

type XY struct {
	x int
	y int
}

type Machine struct {
	a     XY
	b     XY
	prize XY
}

type AB struct {
	a int
	b int
}

func lowestCommonMultiple(a, b int) int {
	start := 10000000000000
	for mul := start; (mul%a != 0) && (mul%b != 0); mul++ {
		if (mul%a == 0) && (mul%b == 0) {
			return mul
		}
	}

	panic("fuck")
}

func parsePrizeSection(input string) XY {
	xy := strings.Split(input, " ")[1:]
	xString := strings.Split(strings.Trim(xy[0], ", "), "=")[1]
	yString := strings.Split(strings.Trim(xy[1], ", "), "=")[1]

	x, err := strconv.Atoi(xString)
	if err != nil {
		panic(err)
	}

	y, err := strconv.Atoi(yString)
	if err != nil {
		panic(err)
	}

	return XY{x: x, y: y}
}

func parseButtonSection(input string) XY {
	xy := strings.Split(input, " ")[2:]
	xString := strings.Split(strings.Trim(xy[0], ", "), "+")[1]
	yString := strings.Split(strings.Trim(xy[1], ", "), "+")[1]

	x, err := strconv.Atoi(xString)
	if err != nil {
		panic(err)
	}

	y, err := strconv.Atoi(yString)
	if err != nil {
		panic(err)
	}

	return XY{x: x, y: y}

}

func Parse(filePath string) []Machine {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	lines, err := io.ReadAll(file)
	machineStrings := strings.Split(strings.Trim(string(lines), "\n"), "\n\n")

	machines := []Machine{}
	for _, machineConfig := range machineStrings {
		sections := strings.Split(machineConfig, "\n")
		machine := Machine{
			a:     parseButtonSection(sections[0]),
			b:     parseButtonSection(sections[1]),
			prize: parsePrizeSection(sections[2]),
		}

		machines = append(machines, machine)
	}

	return machines
}

func potentialXCombos(machine Machine) []AB {
	potentials := []AB{}
	target := machine.prize.x
	for i := 0; i*machine.a.x <= target; i++ {
		if (target-i*machine.a.x)%machine.b.x == 0 {
			potentials = append(potentials, AB{a: i, b: (target - i*machine.a.x) / machine.b.x})
		}
	}

	return potentials
}

func potentialYCombos(machine Machine) []AB {
	potentials := []AB{}
	target := machine.prize.y
	for i := 0; i*machine.a.y <= target; i++ {
		if (target-i*machine.a.y)%machine.b.y == 0 {
			potentials = append(potentials, AB{a: i, b: (target - i*machine.a.y) / machine.b.y})
		}
	}

	return potentials
}

func Compute(machines []Machine) int {
	total := 0

	for _, machine := range machines {
		potentials := []AB{}
		potentialX := potentialXCombos(machine)
		potentialY := potentialYCombos(machine)

		for _, xCombo := range potentialX {
			for _, yCombo := range potentialY {
				if xCombo.a == yCombo.a && xCombo.b == yCombo.b {
					potentials = append(potentials, xCombo)
				}
			}
		}

		cheapest := math.MaxInt
		for _, combo := range potentials {
			total := 0
			total += combo.a * aPrice
			total += combo.b * bPrice
			if total < cheapest {
				cheapest = total
			}
		}

		if cheapest == math.MaxInt {
			cheapest = 0
		}

		total += cheapest
	}

	return total
}
