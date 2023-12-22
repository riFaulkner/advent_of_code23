package _19

import (
	"advent_of_code23/utils"
	"bytes"
	"cmp"
	"maps"
	"regexp"
	"slices"
	"strings"
)

type part struct {
	vals map[string]int
}

type rule struct {
	subject string
	op      string
	amount  int
	dest    string
}

func Problem1(inputFileName string) int {
	segs := bytes.Split(utils.GetFileContent(inputFileName), []byte("\n\n")) // split by empty line, seg0 is the workflows seg2 is the parts

	rules := make(map[string][]rule, len(segs[0]))
	serializeRules(string(segs[0]), &rules)

	parts := make([]part, 0)
	serializeParts(string(segs[1]), &parts)

	return evalRuleR(&rules, &parts, "in")
}

func Problem2(inputFileName string) int {
	segs := bytes.Split(utils.GetFileContent(inputFileName), []byte("\n\n")) // split by empty line, seg0 is the workflows seg2 is the parts

	rules := make(map[string][]rule, len(segs[0]))
	serializeRules(string(segs[0]), &rules)

	return findPermutationsR(&rules, "in", map[string]int{
		"xMin": 1,
		"xMax": 4000,
		"mMin": 1,
		"mMax": 4000,
		"aMin": 1,
		"aMax": 4000,
		"sMin": 1,
		"sMax": 4000,
	})
}

func findPermutationsR(rules *map[string][]rule, ruleKey string, parts map[string]int) int {
	if ruleKey == "R" {
		return 0
	}
	if ruleKey == "A" {
		xN := 1 + parts["xMax"] - parts["xMin"]
		mN := 1 + parts["mMax"] - parts["mMin"]
		aN := 1 + parts["aMax"] - parts["aMin"]
		sN := 1 + parts["sMax"] - parts["sMin"]
		return xN * mN * aN * sN
	}

	cRule := (*rules)[ruleKey]
	sum := 0
	for _, r := range cRule {
		if r.subject == "" && r.amount == 0 && r.op == "" { // this is the last rule, just go to the destination
			return sum + findPermutationsR(rules, r.dest, parts)
		}
		vMin := parts[r.subject+"Min"]
		vMax := parts[r.subject+"Max"]

		var p map[string]int
		if r.op == "<" {
			newMax := min(vMax, r.amount-1)
			if newMax < vMin {
				continue
			}
			p = maps.Clone(parts)
			p[r.subject+"Max"] = newMax
			// and remove those values from the original parts list
			parts[r.subject+"Min"] = newMax + 1 // double check this
		} else {
			newMin := max(vMin, r.amount+1)
			if newMin > vMax {
				continue
			}
			p = maps.Clone(parts)
			p[r.subject+"Min"] = newMin
			// and remove those values from the original parts list
			parts[r.subject+"Max"] = newMin - 1 // double check this
		}

		sum += findPermutationsR(rules, r.dest, p)
	}

	return sum
}

func evalRuleR(rules *map[string][]rule, parts *[]part, ruleKey string) int {
	if ruleKey == "R" {
		return 0
	}
	if ruleKey == "A" {
		sum := 0
		for p := range *parts {
			for _, v := range (*parts)[p].vals {
				sum += v
			}
		}
		return sum
	}

	cRule := (*rules)[ruleKey]
	sum := 0
	for _, r := range cRule {
		if r.subject == "" && r.amount == 0 && r.op == "" { // this is the last rule, just go to the destination
			return sum + evalRuleR(rules, parts, r.dest)
		}

		slices.SortFunc(*parts, func(i, j part) int {
			return cmp.Compare(i.vals[r.subject], j.vals[r.subject])
		})
		i := slices.IndexFunc(*parts, func(i part) bool {
			return i.vals[r.subject] > r.amount
		})

		var p []part

		if r.op == "<" {
			// if the op is < and we get -1 that means all elements qualify
			if i == -1 {
				i = len(*parts)
			}
			p = (*parts)[:i]
			(*parts) = (*parts)[i:]
		} else {
			// if the op is more than and we get -1 that means no elements qualify
			if i == -1 {
				continue
			}
			p = (*parts)[i:]
			(*parts) = (*parts)[:i]
		}
		if len(p) == 0 {
			continue
		}

		sum += evalRuleR(rules, &p, r.dest)
	}

	return sum
}

func serializeRules(s string, rules *map[string][]rule) {
	for _, rSeg := range strings.Split(s, "\n") {
		sg := regexp.MustCompile(`([^{}]+)\{([^{}]+)`).FindStringSubmatch(rSeg)
		ruleS := make([]rule, 0)
		for _, rS := range strings.Split(sg[2], ",") {
			r := regexp.MustCompile(`(\w)([<>])(\d+):(\w+)`).FindStringSubmatch(rS)
			if r != nil {
				ruleS = append(ruleS, rule{
					subject: string(r[1][0]),
					op:      string(r[2][0]),
					amount:  utils.ITS(r[3]),
					dest:    r[4],
				})
			} else {
				ruleS = append(
					ruleS,
					rule{
						dest: rS,
					})
			}

		}

		(*rules)[sg[1]] = ruleS
	}
}

func serializeParts(s string, parts *[]part) {
	for _, pS := range strings.Split(s, "\n") {
		sg := regexp.MustCompile(`(\d+)`).FindAllString(pS, -1)

		*parts = append(*parts, part{
			vals: map[string]int{
				"x": utils.ITS(sg[0]),
				"m": utils.ITS(sg[1]),
				"a": utils.ITS(sg[2]),
				"s": utils.ITS(sg[3]),
			},
		})
	}
	return
}
