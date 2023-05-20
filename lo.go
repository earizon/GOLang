package main

import (
	"fmt"
	"strconv"

	// replace samber/lo ··> samber/lop for goroutine parallel processing.
	"github.com/samber/lo"
)

func main() {
	INPUT := [4]int64{1, 2, 3, 4}
	R :=
		lo.Reduce( // ReduceRight iterates right-to-left
			lo.FlatMap(
				lo.Map( // <·· FilterMap can be used to in place
					lo.Filter(INPUT[:],
						func /*Filter*/ (x int64, _ int) bool {
							return x%2 == 0
						}), //  [ 2 4 ]
					func /*Map*/ (x int64, _ int) string {
						return fmt.Sprintf(">%s<",
							strconv.FormatInt(x, 10))
					}), // [">2<" ">4<" ]
				func /*FlatMap*/ (s string, _ int) []string {
					return []string{s, s}
				}), // [">2<" ">2<" ">4<" ">4<" ]

			func(aggregate string, item string, _ int) string {
				return aggregate + " - " + item
			}, "") // - >2< - >2< - >4< - >4<
	fmt.Println(R)

	lo.ForEach(INPUT[:],
		func /*ForEach*/ (e int64, idx int) {
			print(fmt.Sprintf("%d->%d,", idx, e)) // 0->1,1->2,2->3,3->4
		})
}
