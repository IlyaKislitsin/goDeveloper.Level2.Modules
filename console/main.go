package main

import (
	"fmt"
	"github.com/IlyaKislitsin/goDeveloper.Level2.Modules/console/sort"
)

func main() {
	sortedNumbers := sort.ByInserts(45, -7654, 234, 9, -6543, 1, -591236, 93465, 123, 876)
	fmt.Println(sortedNumbers)
}
