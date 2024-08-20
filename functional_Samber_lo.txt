# Samber Lo Functional Programming [[{lo]]

## SPEC [https://godoc.org/github.com/samber/lo](https://godoc.org/github.com/samber/lo)

* <https://github.com/samber/lo>

## **SLICES!!!**

  ```
  | Times Uniq UniqBy GroupBy Chunk PartitionBy Flatten 
  | Interleave Shuffle Reverse Fill Repeat RepeatBy KeyBy 
  | Associate / SliceToMap Drop DropRight DropWhile 
  | DropRightWhile Reject Count CountBy CountValues 
  | CountValuesBy Subset Slice Replace ReplaceAll Compact 
  | IsSorted IsSortedByKey
  |
  | lo.Times( // Iterate N times, return array of intermediate results
  |           // use lop.Times parallel goroutine processing
  |   3 /*Number of iterations*/,
  |   func(index int) string {
  |     return strconv.FormatInt(int64(i), 10)
  | }) // result []string{"0", "1", "2"}
  | 
  | list0_5 :=  []int{0, 1, 2, 3, 4, 5}
  | uniqValues := lo.UniqBy( list0_5,
  |    func /*criteria of-uniqueness*/ (i int) int { return i%3 }
  | ) // []int{0, 1, 2}
  | 
  | groups := lo.GroupBy( // lop.GroupBy for parallel processing.
  |   list0_5,
  |   func /*GroupBy critery*/(i int) int { return i%3 })
  | //     ┌──────────────────────────┘
  | // map[int][]int{0: []int{0, 3}, 1: []int{1, 4}, 2: []int{2, 5}}
  | 
  | Parallel processing: like `lo.GroupBy()`, but callback is called in goroutine.
  | 
  | lo.Chunk(list0_5, 2 /*size*/) split into groups of length `size`.
  | // [][]int{{0, 1}, {2, 3}, {4, 5}} // Final chunk with up to size elements
  | 
  | partitions := lo.PartitionBy( // Split by criteria
  | //            lop.PartitionBy(
  |   []int{-2, -1, 0, 1, 2, 3, 4, 5},
  |   func /*partition criteria*/(x int) string {
  |     if x < 0 {
  |         return "negative"
  |     } else if x%2 == 0 {
  |         return "even"
  |     }
  |     return "odd"
  | })
  | // [][]int{{-2, -1}, {0, 2, 4}, {1, 3, 5}}
  | 
  | 
  | flat := lo.Flatten([][]int{{0, 1}, {2, 3}}) // []int{0, 1, 2, 3}
  | 
  | interleaved := lo.Interleave( // Round-robin input, seq append
  |   []int{1, 4, 6, 8},
  |   []int{2, 5, 7},
  |   []int{3, } )
  | // []int{1, 2, 3, 4, 5, 6, 7, 8}
  | 
  | randomOrder := lo.Shuffle(list0_5) // Fisher-Yates shuffle algorithm
  | // []int{1, 4, 0, 3, 5, 2}
  | 
  | reverseOrder := lo.Reverse(list0_5) // []int{5, 4, 3, 2, 1, 0}
  | helper is **mutable**. behavior might change in `v2.0.0`
  | 
  | initializedSlice := lo.Fill([5]int, 3 /*init value*/) // [3 3 3 3 3]
  | 
  | 
  | Builds a slice with N copies of initial value.
  | slice := lo.Repeat(2, foo{"a"}) // []foo{foo{"a"}, foo{"a"}}
  | slice := lo.RepeatBy(5, // [0 1 4 9 16]
  |    func(i int) int { return math.Pow(i, 2) }
  | )
  | 
  | 
  | m := lo.KeyBy( // slice||array ··> map, based on pivot func
  |   []string{"a", "aa", "aaa"},
  |   func /*pivot */(str string) int { return len(str) }
  | ) // map[int]string{1: "a", 2: "aa", 3: "aaa"}
  | 
  | (if two pairs have the same key last one gets added)
  | in := []*foo{
  |     {baz: "apple" , bar: 1, biz: 23.0},
  |     {baz: "banana", bar: 2, biz: 23.0},
  | }
  | aMap := lo.Associate( // slice ··> |transform func| ··> key-value map
  |     in,
  |     func /*transform*/ (f *foo) (string, int) {
  |     return f.baz, f.bar // <·· transform by "projecting") struct in map
  | }) // map[string][int]{ "apple":1, "banana":2 }
  | 
  | l := lo.Drop     (list0_5, 2/* drop N from start*/) // []int{2, 3, 4, 5}
  | l := lo.DropRight(list0_5, 2/* drop N from end  */) // []int{0, 1, 2, 3}
  | l := lo.DropWhile(list0_5, func(i int) bool { i < 3 }) // []int{3, 4, 5}
  | //     .DropRightWhile ...                     i > 3    // []int{0, 1, 2}
  | 
  | count := lo.Count(list0_5, 1 /* ref */ ) // 1 (#Elements matching ref value)
  | count := lo.CountBy(list0_5,
  |     func /* match criteria */(i int) bool { // 5: #Elem. matching criteria
  |       return i < 4
  |     })
  | 
  | lo.CountValues([]string{"foo", "bar", "bar"}) // map[string]int{"foo": 1, "bar": 2}
  | 
  | isEven := func(v int) bool {
  |     return v%2==0
  | }
  | 
  | lo.CountValuesBy(list0_5, isEvenFunc) // map[bool]int{false: 3, true: 3}
  |    └───────────┴─ equal to chaining lo.Map ··> lo.CountValues
  | 
  | sub := lo.Subset(list0_5, 2 /*offset*/, 3/*length*/) // Like `slice[start:start+length]`,
  | // returns new fresh "copy" []int{2, 3, 4}           // but does not panic on overflow.
  | // negative offset start from right.
  | 
  | slice := lo.Slice(list0_5, 0/*start*/, 5/*end, no-incl.*/) // []int{0, 1, 2, 3, 4}
  | 
  | slice := lo.Replace(in, 0, 42, 1) []int{42, 1, 2, 3, 4, 5}
  |                         ·   ·  └ replace first n ocurrences
  |                         ·   └··· new value
  |                         └······· matching element to replace
  |            .ReplaceAll (Replaces all occurrences)
  | 
  | slice := lo.Compact[string]( // return slice of all non-"zero" elements
  |         []string{"", "foo", "", "bar", ""})
  | 
  | slice := lo.IsSorted(list0_5)          // true
  | slice := lo.IsSortedByKey(list0_5,
  |            func(i int) { return -i } ) // false
  | 
  | map01= map[string]int{"foo": 1, "bar": 2, "baz": 3}
  ```
  
## **MAPS!!!**

  ```
  | Keys|Values ValueOr PickBy(Keys/Values) 
  | OmitBy(Keys/Values) (From|To)Entries/(From|To)Pairs 
  | Invert Assign Map(Keys|Values|Entries) MapToSlice
  |
  | keys := lo.Keys  [string, int](map01) // []string{"foo", "bar", "baz"}
  | vals := lo.Values[string, int](map01) // []int{1, 2, 3}
  | 
  | val  := lo.ValueOr[string, int](map01, "foo", 42) // 1
  | val  := lo.ValueOr[string, int](map01, "buf", 42) // 42
  | 
  | Returns same map type filtered by given predicate.
  | 
  | m := lo.PickBy(map01,  // map[string]int{"foo": 1, "baz": 3}
  |        .OmitBy to omit (vs pick)
  |    func /*filter predicate*/(key string, value int) bool {
  |     return value%2 == 1
  | })
  | 
  | m := lo.PickByKeys  (map01, []string{"foo"}) // ...{"foo": 1}
  |        .OmitByKeys
  | m := lo.PickByValues(map01, []int{1, 3}) // ..{"foo": 1, "baz": 3}
  |        .OmitByValues
  |  
  | entries := lo.Entries(map01) // transform map to list of key/value pairs
  | // []lo.Entry[string, int]{  // (alias: ToPairs)
  | //     { Key: "foo", Value: 1 },
  | //     { Key: "bar", Value: 2 },
  | // }
  | 
  | m := lo.FromEntries(...) // transform key/value list to map
  | 
  | m1 := lo.Invert(map01) // map[int]string{1: "foo", 2: "bar", 3: "baz"}
  |          └ "switch" key <··> value. For colling values, last one is
  |                                     taken as key
  | 
  | mergedMaps := lo.Assign[string, int]( // "merge" arrays
  |     map[string]int{"a": 1, "b": 2},
  |     map[string]int{"b": 3, "c": 4},
  | )
  |  // map[string]int{"a": 1, "b": 3, "c": 4}
  | 
  | m2 := lo.MapKeys( // transform keys
  |    // lo.MapValues to transform values
  |   map01, func /*key transformer*/(idx int, v string) string {
  |     return fmt.Sprintf("%d %s, idx, strings.ToUpper(v))
  | }) // map[string]int{"0 FOO": 1, "1 BAR": 2, "2 BAZ": 3}
  | 
  | out := lo.MapEntries( map01,
  |    func(k string, v int) (int, string) {
  |     return v,k
  | }) // map01= map[int]string{1: "foo", 2: "bar", 3: "baz"}
  | 
  | s := lo.MapToSlice(map01,
  |     func(k int, v string) string {
  |     return fmt.Sprintf("%d_%s", k, v)
  | }) // []string{"foo_1", "bar_2", "baz_3"}
  ```
  
  
## **MATH!!!**

  ```
  | Range(From|WithSteps), RangeWithSteps, Clamp, Sum(By)
  | 
  | result := lo.Range(4) // [0, 1, 2, 3]
  | result := lo.RangeFrom(1, 5) // [1, 2, 3, 4, 5]
  | result := lo.RangeFrom[float64](1.0, 4) // [1.0, 2.0, 3.0, 4.0 ]
  | result := lo.RangeWithSteps(0, 20, 5) // [0, 5, 10, 15]
  | result := lo.RangeWithSteps[float32](-1.0, -4.0, -1.0) // [-1.0, -2.0, -3.0]
  | 
  | r1 := lo.Clamp(0, -10, 10)   // 0
  | r2 := lo.Clamp(-42, -10, 10) // -10
  | r3 := lo.Clamp(42, -10, 10)  // 10
  | 
  | sum := lo.Sum(list0_5) // 15
  | sum := lo.SumBy(list0_5, func(e int) int { return e*2 }) // 30
  ```
  
## **STRINGS!!!**

  ```
  | Random|Sub|Chunk|String, RuneLength [[{strings.lo]]
  | str := lo.RandomString(5, lo.LettersCharset) // example: "eIGbt"
  | sub := lo.Substring("hello", 2, 3) // "llo"
  | sub := lo.Substring("hello", -2, math.MaxUint) // "lo"
  | lo.ChunkString("123456", 2) // []string{"12", "34", "56"}
  | [[}]]
  | 
  | sub := lo.RuneLength("hellô") // 5
  | sub := len("hellô") // 6
  ```
  
## **TUPLES!!!**: 

  ```
  | tuple2 := lo.T2("x", 1) // Tuple2[string, int]{A: "x", B: 1}
  |               └ 2...9
  | r1, r2 := lo.Unpack2(tuple2) // "x", 1
  |                    └ 2...9
  |                 ┌····┘
  | tuples := lo.Zip2(       // []Tuple2[string, int]{
  |         // UnZip2
  |   []string{"a", "b"},    //   {A: "a", B: 1},
  |   []int   {1  , 2  } )   //   {A: "b", B: 2}     }
  ```

## **CHANNELS!!!**: 

  ```
  | ChannelDispatcher SliceToChannel Generator Buffer(WithTimeout) Fan(In|Out)
  | input_ch := make(chan int, 42);
  | consumerFunc := func(c <-chan int) { ... }
  | children := lo.ChannelDispatcher(  // Distributes messages into N child channels.
  |   input_ch,
  |   5 /*child number*/,
  |   10 /*buffer capacity, 0=> unbuffered*/,
  |   lo.DispatchingStrategyRoundRobin[int])
  | //   └──────────────┬────────────┘
  | //   DispatchingStrategyRoundRobin
  | //   DispatchingStrategyRandom
  | //   DispatchingStrategyWeightedRandom
  | //   DispatchingStrategyFirst /* to first non full channel */
  | //   DispatchingStrategyLeast /* Distribute to emptiest channel */
  | //   DispatchingStrategyMost /* Distribute to the fullest channel. */
  | 
  | for i := range children { go consumerFunc(children[i]) }
  | for i := 0; i <= 10; i++ { input_ch <- i }
  | 
  | NOTE: Some strategies bring fallback, in order to favor non-blocking behaviors.  See implementations.
  |       Custom strategies can be defined with `lo.DispatchingStrategy` prototype.
  | 
  | 
  | for v := range lo.SliceToChannel(2, list0_5) { println(v) } // prints 0, then 1, ...
  |                   ^^^^^
  | Returns a read-only channels of collection elements.
  | Channel is closed after last element. Channel capacity can be customized.
  | 
  | ch_list := lo.SliceToChannel(2, list0_5)
  | items :=   lo.ChannelToSlice(ch_list) // []int{0, 1, 2, 3, 4, 5}
  | 
  | generator := func(yield func(int)) {
  |     yield(1)
  |     yield(2)
  |     yield(3)
  | }
  | for v := range lo.Generator(2, generator) { println(v) } // prints 1, then 2, then 3
  | 
  | 
  | ch_list := lo.SliceToChannel(2, list0_5)
  | items1, length1, duration1, ok1 := lo.Buffer(ch_list, 3) // []int{0, 1, 2}, 3, 2s, true
  | items2, length2, duration2, ok2 := lo.Buffer(ch_list, 3) // []int{3, 4, 5}, 3, 0s, false
  | Creates slice of n el. from channel  ─┴────┘
  | Returns the slice, slice length,
  | read time and channel status (opened/closed).
  | A lo.BufferWithTimeout(ch_list, N, 100*time.Milliseconds) version exists.
  | 
  | Example: RabbitMQ consumer !!!!
  | ch := readFromQueue()
  | for {
  |     items, length, _, ok := lo.Buffer(ch, 1000 /* read 1k items */)       // Alt 1.
  |     //                   := lo.BufferWithTimeout(ch, 1000, 1*time.Second) // Alt 2. With timeouts
  |     // ... batch stuff
  |     if !ok { break }
  | }
  | 
  | Example: Multithreaded RabbitMQ consumer
  | INPUT_CH := readFromQueue()
  | 
  | children := lo.ChannelDispatcher(
  |      INPUT_CH, 5, 1000 /* 1k per worker */, lo.DispatchingStrategyFirst[int])
  | 
  | consumer := func(c <-chan int) {
  |     for {
  |         items, length, _, ok := lo.BufferWithTimeout(ch, 1000, 1*time.Second) // read 1k items
  |         // ... batch stuff
  |         if !ok { break }
  |     }
  | }
  | for i := range children { go consumer(children[i]) }
  | 
  | 
  | all := lo.FanIn (100, ch1, ch2, ch3) Merge N input channels into 1 buffered channel
  | all := lo.FanOut(5, 100, input_channel) // [5]<-chan int
  ```
  
## **INTERSECT!!!**

  ```
  | Contains(By) Every(By) Some(By) None(By) Intersect, Diff. Union Without(Empty)
  | 
  | present := lo.Contains(list0_5, 5) // true
  | present := lo.ContainsBy(list0_5, func(x int) bool { return x == 3 })
  | 
  | ok := lo.Every  (list0_5, list0_2) // true (every element of subset are contained in collection
  | ok := lo.Every  (list0_5, list0_7) // false
  |         .EveryBy(collection, predicateFun)
  | ok := lo.Some   (list0_5, list0_7) // true (at least 1 element is contained in collection)
  |         .SomeBy (collection, predicateFun)
  | ok := lo.None   (list0_5, list0_2) // false (true if none element is contained)
  |         .NoneBy (collection, predicateFun)
  | 
  | lo.Intersect (list0_5, list0_2) // []int{0, 2}
  | lo.Difference(list0_5, list0_2) // []int{3, 4, 5}
  | lo.Union     (list0_5, list0_7) // []int{0, 1, 2, 3, 4, 5, 6, 7}
  | lo.Without   (list0_5, 0, 1, 2) // []int{3, 4, 5}
  |   .WithoutEmpty(list0_5)        // []int{1, 2, 3, 4, 5}
  ```
  
  
## **SEARCH!!!**: 

  ```
  | (Last)IndexOf , Find(|(Last)IndexOfOrElse
  | Key|KeyBy     , Uniques|UniquesBy, Duplicates(By)
  | 
  | Min MinBy Max MaxBy Last Nth Sample Samples
  | lo.IndexOf  (list0_5, 2)        // 2 (-1 if not found)
  |   .LastIndexOf ...
  | lo.Find(list0_5,     // 0, true (return first element found)
  |         func(i int) bool { return i%2 == 0 })
  | lo.FindIndexOf(...) // return element, index and true|fals
  | lo.FindLastIndexOf
  | lo.FindOrElse       // Like Find, if not found, returns default value
  | lo.FindKey
  | lo.FindKeyBy
  | lo.FindUniques   ([]int{1, 2, 2, 1, 2, 3}) // []int{3}
  |   .FindUniquesBy
  | lo.FindDuplicates([]int{1, 2, 2, 1, 2, 3}) // []int{1, 2}
  |   .FindDuplicatesBy
  | 
  | min := lo.Min([]int{1, 2, 3}) // 1
  |          .MinBy(collection, comparison function)
  |          .Max
  |          .MaxBy(collection, comparison function)
  | last, err := lo.Last(list0_5) // 5, false (error if empty)
  | nth, err := lo.Nth(list0_5, 2) // 1, false return at index N
  | lo.Sample(list0_5)    // 3,    fetch random element
  | lo.Samples(list0_5,3) // 1,5,0 fetch random elements
  ```
  
## **CONDITIONAL!!!**

  ```
  | Ternary(F) If/ElseIf/Else Switch/Case/Default
  |
  | result := lo.Ternary(true, "a", "b") // "a"
  | result := lo.TernaryF(true,          // "a"
  |                       func() string { return "a" },
  |                       func() string { return "b" })
  |  Very useful to avoid nil-pointer dereferencing in intializations, [[{QA]]
  |   or avoid running unnecessary code                                [[}]]
  | 
  | 
  | lo.If(true, 1).ElseIf(false, 2).Else(3) // 1
  | lo.IfF(true, func1).ElseIfF(false, func2).ElseF(func3)
  | lo.IfF(true, func1).ElseIfF(false, func2).Else(42)  // If/IfF can mixed
  | 
  | result := lo.Switch(1).Case (1, "1"). Case (2, "2"). Default ("3")
  | result := lo.Switch(1).CaseF(1,  f1). CaseF(2, f2 ). DefaultF(f3)
  ```
  
## **TYPE MNG!!!**

  ```
  | (Emptyable)ToPtr FromPtr(Or) To|From(SlicePtr|AnySlice) (Is)(Not)Empty Coalesce
  | ptr := lo.ToPtr("hello world") // *string{"hello world"}
  |          .EmptyableToPtr       // nil pointer if element is zero
  | value := lo.FromPtr(&strPtr) // "hello world"
  |            .FromPtrOr        // value or default value if prt is nill.
  | Returns a slice of pointer copy of value.
  | ptr := lo.ToSlicePtr([]string{...}) // []*string{"hello", "world"}
  | elements := lo.ToAnySlice([]int{1, 5, 1}) // []any{1, 5, 1}. All elements mapped to `any` type
  | elements, ok := lo.FromAnySlice([]any{"foobar", 42})
  | elements, ok := lo.FromAnySlice([]any{"foobar", "42"}) // []string{"foobar", "42"}, true
  |                    └ maps all `any` elements to a type or fails
  | lo.Empty[int](), lo.Empty[string](), lo.Empty[bool]() // 0, "", false
  | lo.IsEmpty(0), lo.IsEmpty(""), lo.IsEmpty(test{foobar: ""}) // true, true, true
  |   .IsNotEmpty
  | 
  | result, ok := lo.Coalesce(0, 1, 2, 3) // 1 true. Returns the 1st non-empty argument
  ```
  
## **FUNCTION!!!**: Partial
  ```
  | add := func(x, y int) int { return x + y }
  | f := lo.Partial(add, 5)     // Returns new func. that,
  |                             // when called, has its 1st arg 
  |                             // set to provided value.
  ```
  
## **CONCURRENCY!!!**

  ```
  | Attempt(While)(WithDelay) Debounce(By) Synchronize Async Transaction
  | iter, err := lo.Attempt( // Invokes func. N times until it returns valid output.
  |             // .AttemptWithDelay
  |             // .AttemptWhile
  |             // .AttemptWhileWithDelay
  |     42, // number of trials, 0 or negative to attempt forever
  |     func(i int) error {
  |        if i == 5 { return nil }
  |        return fmt.Errorf("failed") // Fail after N attemps
  |     }) // 6 nil
  | 
  | NOTE: For more advanced retry strategies (delay, exponential backoff...),
  |       please take a look on [cenkalti/backoff](https://github.com/cenkalti/backoff).
  | 
  | debounceWrap, cancel := lo.NewDebounce(    // creates wrapper around func. that delays invoking
  |     100 * time.Millisecond, f1)        // the func. until after wait milliseconds,
  | for j := 0; j < 10; j++ { debounce() } // or until `cancel` is called.
  | time.Sleep(1 * time.Second)
  | cancel()
  | 
  | debounceWrap, cancel := lo.NewDebounceBy(  // create new debounce by key
  |      100 * time.Millisecond, f1)
  | for j := 0; j < 10; j++ {
  |     debounceWrap("first key")
  |     debounceWrap("second key")
  | }
  | time.Sleep(1 * time.Second)
  | cancel("first key")
  | cancel("second key")
  | 
  | 
  | s := lo.Synchronize() // Wraps underlying callback in a mutex. (It receives an optional mutex)
  | for i := 0; i < 10; i++ {
  |     go s.Do(func () { println("will be called sequentially") })
  | }
  | Equivalent to:
  | mu := sync.Mutex{}
  | func foobar() {
  |     mu.Lock()
  |     defer mu.Unlock()
  |     // ...
  | }
  | 
  | ch := lo.Async2(func() (int, string) { // chan lo.Tuple2[int, string] ({42, "Hello"})
  |   time.Sleep(10 * time.Second);
  |   return 42, "Hello"
  | })
  | 
  | Implements a Saga pattern
  | transaction := NewTransaction[int]().
  |       Then( func(state int) (int, error) { fmt.Println("step 1"); return state + 10, nil },
  |             func(state int)          int { fmt.Println("rollback 1"); return state - 10 },
  |     ).Then( func(state int) (int, error) { fmt.Println("step 2"); return state + 15, nil },
  |             func(state int)          int { fmt.Println("rollback 2") ; return state - 15 },
  |     ).Then( func(state int) (int, error) { return state, fmt.Errorf("error") } // return error
  |             func(state int)          int { fmt.Println("rollback 3") ; return state - 42 } )
  | _, _ = transaction.Process(-5) // step 1 > step 2 > rollback 2 > rollback 1
  ```

  
## **ERROR MNG!!!**

  ```
  | Try(Or){0..6} Try(WithErrorValue)(Catch(WithErrorValue)))))))))) ErrorsAs
  | val := lo.Validate( // Error helper func.
  |       len(someSlice) == 0,
  |        "Slice should be empty but contains %v", slice)
  | // error("Slice should be empty but contains [a]")
  | 
  | lo.Must // Must{0-6}
  | if 2nd argument is `error`||`false` panic, Otherwise returns value
  | 
  | lo.Must( time.Parse("2006-01-02", "2022-01-15")) // == 2022-01-15
  | lo.Must( time.Parse("2006-01-02", "bad-value" )) // panics
  | lo.Must0(...,"panic hint %s", someStringParam)
  | lo.Must0(ok, "'%s' must always contain '%s'", myString, requiredChar)
  | val1 := lo.Must1(example1())    // alias to Must
  | val1, val2 := lo.Must2(example2())
  | val1, val2, val3 := lo.Must3(example3())
  | val1, val2, val3, val4 := lo.Must4(example4())
  | val1, val2, val3, val4, val5 := lo.Must5(example5())
  | val1, val2, val3, val4, val5, val6 := lo.Must6(example6())
  | 
  | lo.Try(func() error { return nil })                 // true
  | lo.Try(func() error { return fmt.Errorf("error") }) // false
  | 
  | Like `Try`, but callback returns 0,1,2,...6 variables.
  | 
  | Like Try but returns default value in case of error||panic.
  | str, ok := lo.TryOr(
  |    func() (string, error) { panic("error") ; return nil },
  |    "defVal") // world, false
  | str, ok := lo.TryOr(
  |    func() error { return "hello", nil },
  |    "world") hello true
  | 
  | Like `Try`, but also returns value passed to panic
  | err, ok := lo.TryWithErrorValue(
  |   func() error { panic("error") ; return nil }) // "error", false
  | 
  | Like `Try`, but calls catch function in case of error.
  | caught := false
  | ok := lo.TryCatch(
  |    func() error { panic("error") ; return nil },
  |    func() { caught = true }) // false
  | 
  | The same behavior than `TryWithErrorValue`, but calls the catch function in case of error.
  | 
  | caught := false
  | ok := lo.TryCatchWithErrorValue(func() error {
  |     panic("error")
  |     return nil
  | }, func(val any) { caught = val == "error" }) // false; (caught == true)
  | 
  | A shortcut for:
  | err := doSomething()
  | 
  | var rateLimitErr *RateLimitError
  | if ok := errors.As(err, &rateLimitErr); ok {
  |     // retry later
  | }
  | 
  | err := doSomething()
  | if rateLimitErr, ok := lo.ErrorsAs[*RateLimitError](err); ok {
  |     // retry later
  | }
  ```
[[lo}]]
