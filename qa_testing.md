# GOlang Testing [[{qa.testing.101,qa.101]]
* <https://golang.org/cmd/go/#hdr-Test_packages>

  ```
  | $GOPATH/src/github.com/user/stringutil/reverse.go       <- File to test
  | $GOPATH/src/github.com/user/stringutil/reverse_test.go  <- Sibling of reverse.go
  |                                                            excluded in package builds
  | package stringutil
  | import "testing"                      <-  <https://golang.org/pkg/testing/>
  |
  | func TestReverse( t  *testing.T) {    <- TestXxx marks funct. as test routine
  |                                           var  t  injected by testing framework?
  |
  |   cases_l := []struct {               <- Define struct, init array of struct inline
  |     input          string,
  |     expectedOutput string
  |   }{
  |     {"Hello, world", "dlrow ,olleH"}, ← Single array element of type struct
  |     ...
  |   }
  |   for _, c := range cases_l {
  |     output := Reverse(c.input)        <- function being tested
  |     if output != c.expectedOutput {
  |     t .Errorf(                        <- alternatively t.Fail, t.Fatal or "panic"
  |       "Reverse(%q) == %q, Expected %q",
  |        c.input, output, c.expectedOutput)
  |     }
  |   }
  | }
  |
  | $ go test github.com/user/stringutil                   ← Running tests:
  | → ok github.com/user/stringutil 0.165s
  ```

## Benchmarks
  ```
  | func BenchmarkXxx(b *testing.B) {   <- Mark funct as Benchmarks
  |                                        run sequentially ($ go test -bench)
  |                                        (See official doc for running in parallel)
  |     ... some "slow" presetup ...
  |     b.ResetTimer()                  <- Reset benchmark after "slow" presetup.
  |     for i := 0; i < b.N; i++ {      <- @benchmark execution b.N b.N is adjusted
  |         rand.Int()                     until the benchmark function lasts long
  |     }                                  enough to be timed reliably.
  | }                                      ( Maybe tens of millions for short codes)
  ```
  
## FUZZ TESTING (https://go.dev/doc/fuzz)
  ```
  | func FuzzHex(f *testing.F) {        <- FuzzXxx marks func. as fuzz-test
  |                                        It will be called with random inputs in
  |                                        go test --fuzz
  |
  |   for _, seed := range [][]byte{    <- optional. seed corpus (inputs run by default)
  |      {}, {0}, {0x232}, {1, 2} }
  |   { f.Add(seed) }                   <- register seed input.
  |                                      alternative, use testdata/fuzz/"Name_of_fuzz_test"
  |
  |   f.Fuzz(
  |     func(t *testing.T               <- Fuzz test target function to be tested.
  |          in []byte ) {                 'in' random input, must match type of seed.
  |     enc := hex.EncodeToString(in)
  |     out, err := hex.DecodeString(enc)
  |     if err != nil {
  |       t.Fatalf("%v: decode: %v", in, err)
  |     }
  |     if !bytes.Equal(in, out) {
  |       t.Fatalf("%v: not equal after round trip: %v", in, out)
  |     }
  |   })
  | }
  ```

* Use t.Skip("...") to skip a test or a non-valid input in Fuzz tests.


## Subtests AND Sub-Benchmarks
  ```
  | func TestFoo(t *testing.T) {                 
  |     // <setup code>
  |     t.Run("A=1", func(t *testing.T) { ... }) <- Run allow to define subtest/sub-benchmarks
  |     t.Run("A=2", func(t *testing.T) { ... })    allowing to share common setup/tear-down
  |     t.Run("B=1", func(t *testing.T) { ... })
  |     // <tear-down code>
  | }
  | $ go test -run ''                   # <- Run all tests.
  | $ go test -run Foo                  # <- Run test  matching top-level "TestFoo*".
  | $ go test -run Foo/A=               # <- Run tests matching top-level "TestFoo*" && subtests matching "A="
  | $ go test -run /A=1                 # <- Run tests matching top-level "Test*"    && subtests matching "A=1"
  | $ go test -run=FuzzFoo/9ddb952d9814 # <- run Fuzz test matching top-level "FuzzFoo" with given input.
  ```

  ``` 
  func TestMain(m *testing.M) {    <- Optional, low-level primitive, only necessary for casual testing
                                      Control code run on main-thread, pre-setup, teardown
                                      - call flag.Parse() here if TestMain uses flags
    os.Exit(m.Run())                  If present it will be run (instead of Test*, Benchmarks*, Fuzz*)
  }
  ``` 
[[qa.testing.101}]]
