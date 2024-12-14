package answer

import (
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	aPrice  = 3
	bPrice  = 1
	convErr = 10000000000000
)

type XY struct {
	x int
	y int
}

type Machine struct {
	a             XY
	b             XY
	prize         XY
	lowestCommonX int
	lowestCommonY int
}

type AB struct {
	a int
	b int
}

func lowestCommonProductAboveConvErr(a, b int, max int) int {
	start := convErr
	for prod := start; prod <= max+convErr; prod++ {
		if (prod%a == 0) && (prod%b == 0) {
			return prod
		}
	}

	return 0
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

		a := parseButtonSection(sections[0])
		b := parseButtonSection(sections[1])
		prize := parsePrizeSection(sections[2])
		lowestCommonX := lowestCommonProductAboveConvErr(a.x, b.x, prize.x)
		lowestCommonY := lowestCommonProductAboveConvErr(a.y, b.y, prize.y)

		machine := Machine{
			a:             a,
			b:             b,
			prize:         prize,
			lowestCommonX: lowestCommonX,
			lowestCommonY: lowestCommonY,
		}

		machines = append(machines, machine)
	}

	return machines
}

func potentialXCombos(machine Machine) []AB {
	potentials := []AB{}
	target := machine.prize.x
	for i := 0; i*machine.a.x <= target; i++ {
		if (target-(i*machine.a.x))%machine.b.x == 0 {
			potentialA := i
			potentialB := (target - i*machine.a.x) / machine.b.x
			potentials = append(potentials, AB{a: potentialA, b: potentialB})
		}
	}

	return potentials
}

func potentialYCombos(machine Machine) []AB {
	potentials := []AB{}
	target := machine.prize.y
	for i := 0; i*machine.a.y <= target; i++ {
		if (target-(i*machine.a.y))%machine.b.y == 0 {
			potentialA := i
			potentialB := (target - i*machine.a.y) / machine.b.y
			potentials = append(potentials, AB{a: potentialA, b: potentialB})
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

		fmt.Println(potentialX)
		fmt.Println(potentialY)
		fmt.Println()
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
