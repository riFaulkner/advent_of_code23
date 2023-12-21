package _19

import (
	"advent_of_code23/utils"
	"bytes"
	"cmp"
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

func Problem1(inputFileName string, part2 bool) int {
	segs := bytes.Split(utils.GetFileContent(inputFileName), []byte("\n\n")) // split by empty line, seg0 is the workflows seg2 is the parts

	rules := make(map[string][]rule, len(segs[0]))
	serializeRules(string(segs[0]), &rules)

	parts := make([]part, 0)
	serializeParts(string(segs[1]), &parts)

	return evalRuleR(&rules, &parts, "in")
}

//func Problem2(inputFileName string) int {
//	// now we're calculating the total amount of permutations of acceptable parts
//
//	return 0
//}

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
}
