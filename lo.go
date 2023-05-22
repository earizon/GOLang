package main

import (
	"fmt"
	"strconv"

	"github.com/samber/lo"
)

func main() {
	INPUT := [6]int64{1, 2, 2, 3, 4, 4}
	R :=
		lo.Reduce( // ReduceRight iterates right-to-left
			lo.FlatMap(
				lo.Map( // <·· FilterMap can be used to in place
					// lop.Map for parallel processing
					lo.Filter( // .Reject  to
						//	lop.Filter for parallel processing
						//	lop.Reject to exclude (vs include) match
						lo.Uniq[int64](INPUT[:]), // [ 1 2 3 4 ]
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

	lo.ForEach(INPUT[:], // lop.ForEach: goroutine parallel proc
		func /*ForEach*/ (e int64, idx int) {
			print(fmt.Sprintf("%d->%d,", idx, e)) // 0->1,1->2,2->3,3->4
		})
}
