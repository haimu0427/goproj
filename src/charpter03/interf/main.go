package main

import (
	"fmt"
	"sort"
)

type Heros struct {
	Name   string
	Number int
}
type HeroSlice []Heros

func (h Heros) String() string {
	return fmt.Sprintf("Hero Name: %s, Number: %d", h.Name, h.Number)
}
func (liangMen HeroSlice) Len() int {
	return len(liangMen)
}
func (liangMen HeroSlice) Less(i, j int) bool {
	return liangMen[i].Number < liangMen[j].Number
}

func (liangMen HeroSlice) Swap(i, j int) {
	liangMen[i], liangMen[j] = liangMen[j], liangMen[i]
}

func main() {
	liangMen := HeroSlice{}
	hero1 := Heros{Name: "Superman", Number: 1}
	hero2 := Heros{Name: "Batman", Number: 9}
	hero3 := Heros{Name: "Wonder Woman", Number: 3}
	fmt.Println("Adding heroes to the list...")
	liangMen = append(liangMen, hero1, hero2, hero3)
	fmt.Println("Heroes added:", liangMen)
	sort.Sort(liangMen)
	fmt.Println("Heroes sorted:", liangMen)

}
