package _20

import (
	"advent_of_code23/utils"
	"slices"
	"strings"
)

type pulse struct {
	dest     string // name of the destination module
	source   string // name of the source that sent the pulse
	strength int    // 0 or 1 for low and high pulses
}

type module struct {
	name  string
	t     string
	state int
	dests []string
	hist  map[string]int
}

func (m *module) processPulse(p pulse, queue chan pulse) int {
	added := 0
	switch m.t {
	case "%":
		if p.strength == 0 { // ignore high pulses, just focus on low pulses
			// flip between on and off
			m.state = (m.state + 1) % 2
			for _, d := range m.dests {
				queue <- pulse{dest: d, source: m.name, strength: m.state}
				added++
			}
		}
	case "&":
		m.hist[p.source] = p.strength
		pulseToSend := 0
		for _, v := range m.hist {
			if v == 0 {
				pulseToSend = 1
				break
			}
		}
		for _, d := range m.dests {
			queue <- pulse{dest: d, source: m.name, strength: pulseToSend}
			added++
		}

	default:
		for _, d := range m.dests {
			queue <- pulse{dest: d, source: m.name, strength: p.strength}
			added++
		}
	}
	return added
}

func Problem1(inputFileName string) int {
	lines := utils.GetFileAsString(inputFileName)
	modules := make(map[string]module)
	serializeModules(lines, &modules)

	// We're building a state machine! Yay
	counter := []int{0, 0}
	for i := 0; i < 1000; i++ {
		pushButton(&counter, &modules, false)
	}

	return counter[0] * counter[1]
}

func Problem2(inputFileName string) int {
	lines := utils.GetFileAsString(inputFileName)
	modules := make(map[string]module)
	serializeModules(lines, &modules)

	counter := []int{0, 0}

	for i := 0; ; i++ {
		if pushButton(&counter, &modules, true) {
			return i
		}
	}
}

func pushButton(counter *[]int, modules *map[string]module, part2 bool) bool {
	c := make(chan pulse, 1000) // try this and if we run into issues implement a queue instead
	rxPulse := 0

	added, processed := 1, 0
	// the first thing we do is send a pulse from the "button" to the broadcaster
	// that counts toward the totals passed in
	c <- pulse{dest: "broadcaster", source: "button", strength: 0}

	// from there we loop until the channel is empty.
	for ; processed < added; processed++ {
		p := <-c
		(*counter)[p.strength] += 1

		m := (*modules)[p.dest]
		added += m.processPulse(p, c)
		(*modules)[p.dest] = m
		if part2 && p.dest == "rx" && p.strength == 0 {
			rxPulse++
		}
	}
	return rxPulse == 1
}

func serializeModules(s string, modules *map[string]module) {
	// I MAY need to go back through and find the connections between modules that are &
	// save a list of modules that are &s and need to be mapped to.
	i := make([]string, 0)
	for _, l := range strings.Split(s, "\n") {
		parts := strings.Split(strings.Replace(l, " ", "", -1), "->")
		name := parts[0] // name, may or may not have a & or % in front of it
		t := ""
		if name[0] == '&' || name[0] == '%' {
			t = string(name[0])
			name = name[1:]
			if t == "&" {
				i = append(i, name)
			}
		}
		dests := strings.Split(parts[1], ",")
		(*modules)[name] = module{name: name, t: t, state: 0, dests: dests, hist: make(map[string]int)}

		// now iterate though i and make all the mappings
		for k, v := range *modules {
			for _, d := range v.dests {
				if slices.Contains(i, d) {
					(*modules)[d].hist[k] = 0
				}
			}
		}
	}

}
