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
		if len(*parts) == 0 {
			return sum
		}
		if r.subject == "" && r.amount == 0 && r.op == "" {
			sum += evalRuleR(rules, parts, r.dest)
		}
		slices.SortFunc(*parts, func(i, j part) int {
			return cmp.Compare(i.vals[r.subject], j.vals[r.subject])
		})
		i := slices.IndexFunc(*parts, func(i part) bool {
			return i.vals[r.subject] > r.amount
		})

		var p []part

		if r.op == "<" {
			if i == -1 {
				continue
			}
			p = (*parts)[:i]
			(*parts) = (*parts)[i:]
		} else {
			if i == -1 {
				p = (*parts)
				(*parts) = (*parts)[:0]
			} else {
				if i == -1 {
					continue
				}
			}
		}

		sum += evalRuleR(rules, &p, r.dest)
	}
	// those go to the destination (e.g. do recursive call)
	// then we create a new version of the list, with only the remaining elements
	// and we do the same for the next rule
	// if we're at the last rule of the set don't do any conditional logic, just pass it to it's destination

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
