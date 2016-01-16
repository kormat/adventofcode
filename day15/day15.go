package main

import (
	"flag"
	"fmt"
	"github.com/kormat/adventofcode/util"
	"log"
	"os"
	"regexp"
	"strconv"
	"sync"
)

type Ingredient struct {
	capacity   int
	durability int
	flavor     int
	texture    int
	calories   int
}

type Ingredients map[string]Ingredient
type Recipe map[string]int

var best_score int
var best_recipe Recipe
var best_cal_score int
var best_cal_recipe Recipe
var wg sync.WaitGroup

const maxIngredients = 100

func main() {
	flag.Parse()
	lines, ok := util.ReadFileArg(flag.Args())
	if !ok {
		os.Exit(1)
	}
	ings := parseLines(lines)
	findBestRecipe(ings)
	fmt.Printf("Best at any calories: %v %v\n", best_recipe, best_score)
	fmt.Printf("Best at 500 calories: %v %v\n", best_cal_recipe, best_cal_score)
}

func parseLines(lines []string) *Ingredients {
	ings := make(Ingredients)
	rx := regexp.MustCompile(
		`^(\w+): capacity (-?\d+), durability (-?\d+), flavor (-?\d+), texture (-?\d+), calories (-?\d+)$`)
	for i, line := range lines {
		result := rx.FindStringSubmatch(line)
		if result == nil {
			log.Fatalf("Unable to parse line %d: %s", i, line)
		}
		capacity := parseInt(i, result[2], "capacity")
		durability := parseInt(i, result[3], "durability")
		flavor := parseInt(i, result[4], "flavor")
		texture := parseInt(i, result[5], "texture")
		calories := parseInt(i, result[6], "calories")
		ings[result[1]] = Ingredient{capacity, durability, flavor, texture, calories}
	}
	return &ings
}

func parseInt(line int, input, desc string) int {
	ret, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal("Unable to parse %s '%s' on line %d: %s", desc, input, line, err)
	}
	return ret
}

func findBestRecipe(ings *Ingredients) {
	c := make(chan Recipe)
	names := make([]string, 0, len(*ings))
	counts := make([]int, len(*ings))
	for name, _ := range *ings {
		names = append(names, name)
	}
	go gatherResults(c, ings)
	wg.Add(1)
	go findRecipe(c, names, counts, 0)
	wg.Wait()
	close(c) // Signals to gatherResults that all results have been sent.
}

func gatherResults(c chan Recipe, ings *Ingredients) {
	for recp := range c {
		total := 0
		for _, v := range recp {
			total += v
		}
		if total != maxIngredients {
			log.Fatalf("Invalid recipe, doesn't contain 100 teaspoons:  %v: %v teaspoons", recp, total)
		}
		score, calories := recp.score(ings)
		if score > best_score {
			best_score = score
			best_recipe = recp
		}
		if calories == 500 && score > best_cal_score {
			best_cal_score = score
			best_cal_recipe = recp
		}
	}
}

func findRecipe(c chan Recipe, names []string, counts []int, idx int) {
	defer wg.Done()
	total := countIngredients(counts)
	if idx >= len(counts) {
		if total != maxIngredients {
			return
		}
		r := Recipe{}
		for i := 0; i < len(names); i++ {
			r[names[i]] = counts[i]
		}
		c <- r
	}
	if total >= maxIngredients {
		return
	}
	for ; counts[idx] <= maxIngredients; counts[idx]++ {
		new_counts := append([]int(nil), counts...)
		wg.Add(1)
		findRecipe(c, names, new_counts, idx+1)
	}
}

func countIngredients(counts []int) int {
	total := 0
	for _, val := range counts {
		total += val
	}
	return total
}

func (recp Recipe) score(ings *Ingredients) (int, int) {
	var capacity, durability, flavor, texture, calories int
	for k, v := range recp {
		ing := (*ings)[k]
		capacity += ing.capacity * v
		durability += ing.durability * v
		flavor += ing.flavor * v
		texture += ing.texture * v
		calories += ing.calories * v
	}
	if capacity < 0 {
		capacity = 0
	}
	if durability < 0 {
		durability = 0
	}
	if flavor < 0 {
		flavor = 0
	}
	if texture < 0 {
		texture = 0
	}
	score := capacity * durability * flavor * texture
	return score, calories
}
