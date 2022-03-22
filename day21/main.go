package main

import (
	"adventofcode2020/util"
	"fmt"
	"sort"
	"strings"
)

type Food struct {
	Ingredients []string
	Allergens   []string
}

func parseFoods(input []string) (foods []Food) {
	foods = make([]Food, 0)
	for _, line := range input {
		parts := strings.Split(line, "(contains ")
		ingredients := strings.Fields(parts[0])
		allergens := strings.Split(parts[1][:len(parts[1])-1], ", ")
		foods = append(foods, Food{
			Ingredients: ingredients,
			Allergens:   allergens,
		})
	}

	return
}

type Ingredient struct {
	Name     string
	Allergen string
}

type ByAllergen []Ingredient

func (a ByAllergen) Len() int           { return len(a) }
func (a ByAllergen) Less(i, j int) bool { return a[i].Allergen < a[j].Allergen }
func (a ByAllergen) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func filterIngredients(foods []Food) (allergenFree []string, withAllergens []Ingredient) {
	ingredients := util.NewSet[string]()
	for _, food := range foods {
		for _, ingredient := range food.Ingredients {
			ingredients.Add(ingredient)
		}
	}

	allergens := map[string]util.Set[string]{}

	for _, food := range foods {
		for _, allergen := range food.Allergens {
			if i, ok := allergens[allergen]; ok {
				allergens[allergen] = i.Intersect(util.NewSetFromSlice(food.Ingredients))
			} else {
				allergens[allergen] = ingredients.Intersect(util.NewSetFromSlice(food.Ingredients))
			}
		}
	}

	singleFields := util.NewQueue[string]()
	for _, allergenIngredients := range allergens {
		if allergenIngredients.Len() == 1 {
			singleFields.PushBack(allergenIngredients.First())
		}
	}

	for i := singleFields.Pop(); i != nil; i = singleFields.Pop() {
		for _, val := range allergens {
			if val.Len() > 1 {
				val.Remove(*i)
				if val.Len() == 1 {
					singleFields.PushBack(val.First())
				}
			}
		}
	}

	allergensIngredients := make([]Ingredient, 0, len(allergens))
	for k, v := range allergens {
		allergensIngredients = append(allergensIngredients, Ingredient{
			Name:     v.First(),
			Allergen: k,
		})
	}

	for _, allergenIngredients := range allergens {
		for _, i := range allergenIngredients.Values() {
			ingredients.Remove(i)
		}
	}

	return ingredients.Values(), allergensIngredients
}

func part1(foods []Food) int {
	ingredients, _ := filterIngredients(foods)

	total := 0
	for _, food := range foods {
		foodSet := util.NewSetFromSlice(food.Ingredients)
		for _, ingredient := range ingredients {
			if foodSet.Contains(ingredient) {
				total += 1
			}
		}
	}

	return total
}

func part2(foods []Food) string {
	_, ingredients := filterIngredients(foods)

	sort.Sort(ByAllergen(ingredients))

	names := make([]string, 0, len(ingredients))
	for _, i := range ingredients {
		names = append(names, i.Name)
	}
	return strings.Join(names, ",")
}

func main() {
	defer util.Stopwatch("Run")()
	input, _ := util.ReadInputLines("./input/day21.txt")
	foods := parseFoods(input)

	fmt.Println("(p1)", part1(foods)) //1685
	fmt.Println("(p2)", part2(foods)) //1685
}
