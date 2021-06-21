package main

import (
	"fmt"
	"sort"
)

type People struct {
	Name string
	Age int
}

type Peoples []People

func (p Peoples) Len() int {
	return len(p)
}

func (p Peoples) Swap(i, j int)  {
	p[i], p[j] = p[j], p[i]
}

func (p Peoples) Less(i, j int) bool {
	return p[i].Age < p[j].Age
}

func main()  {

	example()

	exampleStable()

	// 对任意 函数实现版 切片类型 反向排序
	exampleReverse()
	// []string排序
	exampleString()
	// []int排序
	exampleInt()

	exampleFloat()

	exampleSlice()

	exampleSliceStable()
}

func example()  {

	peoples := Peoples{
		{"Gopher", 7},
		{"Alice", 55},
		{"Vera", 23},
		{"Bob", 75},
	}

	sort.Sort(peoples)
	fmt.Println(peoples)

	fmt.Println(sort.IsSorted(peoples))

}

func exampleStable()  {

	peoples := Peoples{
		{"Gopher", 7},
		{"Alice", 55},
		{"Vera", 24},
		{"Bob", 75},
	}

	sort.Stable(peoples)
	fmt.Println(peoples)

	fmt.Println(sort.IsSorted(peoples))

}

func exampleReverse()  {

	peoples := Peoples{
		{"Gopher", 7},
		{"Alice", 55},
		{"Vera", 24},
		{"Bob", 75},
	}

	p := sort.Reverse(peoples)

	sort.Sort(p)
	fmt.Println(peoples)

}

func exampleString()  {

	s := []string{"Go", "Bravo", "Gopher", "Alpha", "Grin", "Delta"}

	sort.Strings(s)
	fmt.Println(s)

	s = []string{"Go", "Bravo", "Gopher", "Alpha", "Grin", "Delta"}

	sort.Sort(sort.StringSlice(s))
	fmt.Println(s)

	fmt.Println(sort.StringsAreSorted(s))

	index := sort.SearchStrings(s, "Hello")
	fmt.Println(index)
	fmt.Println(s)
}

func exampleInt()  {

	s := []int{5,2,6,3,1,4}
	sort.Ints(s)
	fmt.Println(s)

	s = []int{5,2,6,3,1,4}
	sort.Sort(sort.IntSlice(s))
	fmt.Println(s)

	fmt.Println(sort.IntsAreSorted(s))

	index := sort.SearchInts(s, 5)
	fmt.Println(index)

}

func exampleFloat()  {

	s := []float64{5.2, -1.3, 0.7, -3.0, 2.6}
	sort.Float64s(s)
	fmt.Println(s)

	s = []float64{5.2, -1.3, 6.7, -3.0, 2.6}
	sort.Sort(sort.Float64Slice(s))
	fmt.Println(s)

	fmt.Println(sort.Float64sAreSorted(s))

	index := sort.SearchFloat64s(s, 2)
	fmt.Println(index)

}

func exampleSlice()  {

	people := []struct{
		Name string
		Age int
	}{
		{"Gopher", 7},
		{"Alice", 55},
		{"Vera", 24},
		{"Bob", 75},
	}

	sort.Slice(people, func(i, j int) bool {
		return people[i].Name < people[j].Name

	})
	fmt.Println("By name:", people)

	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age

	})
	fmt.Println("By age", people)

}

func exampleSliceStable()  {

	people := []struct{
		Name string
		Age int
	}{
		{"Alice", 25},
		{"Elizabeth", 75},
		{"Alice", 75},
		{"Bob", 75},
		{"Alice", 75},
		{"Bob", 25},
		{"Colin", 25},
		{"Elizabeth", 25},
	}

	sort.SliceStable(people, func(i, j int) bool {
		return people[i].Name < people[j].Name
	})
	fmt.Println("By name", people)

	sort.SliceStable(people, func(i, j int) bool {
		return people[i].Age < people[j].Age

	})
	fmt.Println("By age.name", people)
}