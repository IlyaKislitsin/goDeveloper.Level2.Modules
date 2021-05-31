package main

import (
	"fmt"

	"github.com/IlyaKislitsin/goDeveloper.Level2.Modules/v2/console/sort"
)

func main() {
	ar := []int{45, -7654, 234, 9, -6543, 1, -591236, 93465, 123, 876}
	sort.Quicksort(ar)
	fmt.Println(ar)
}
