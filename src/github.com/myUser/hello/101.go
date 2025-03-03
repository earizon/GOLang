// [[{101.hello_world]]
/**
 * "Hello World" build and run:
 * ${GOPATH}/src/github.com/myUser/hello/hello.go 
 * $ go run  hello.go     # alt 1. Quick version
 * $ go build  && ./hello # alt 2.
 **/
// ┌─────┬───────────────────────── // 'package' statement must be first statement in file
// ·     · ┌──┬──────────────────── // 'main' tells that package should compile as an executable
// v     v v  v                     // (vs shared (source) library), having a "main" entry point.
   package main                     //  
                                    // pkg.names for libraries are short,clear,lowercase (without
                                    // conflictive /[_@^...]/ chars)
                                    // CONVENTION: function called 'New' returns the type *pkg.Pkg
                                    //                            '*' stands for Pointer ─┘
                                    // e.g:
                                    // list.New() @ ./list/list.go returns *list.List
                                    // Use New'OtherType' for OtherType != pkg.Pkg
                                    // 
import (                            // <·· importing several types and functions
  "fmt"                             //     <·· fmt.Println, fmt.Scanf,... [[strings.101]]
   pkg11 "golang-book/chapter11/math"// <·· ./golang-book/chapert11/math/math.go
   pkg12 "golang-book/chapter12/math"
// └───┴─·································· Avoid conflict with same file name in different packages
  "time"                            //  [[101.date_time]]
  "flag"                            //  [[qa.UX]]
  ... )
  //     ┌─ ignore 2nd returned value (error in this case)
  if m , _ == functionN() ; m == 0  //< declare/assign  m "inside" i , then pass m == n to if
   { return 0, errors.New("n==0") } //< return N values. Normally (result, errors)
                                    //  errors package is standard (no need to import) 
                                    //  NOTE: return values can be named. [[qa.error_control]]

  for i := 0; i < 10; i++ {         //< NO PARENTHESIS USED in 'for' or 'if'
    fmt.Println(n, ":", i)          //  break'|'continue' statements work as in C/Java/...
    amt :=                          //
      time.Duration(rand.Intn(250)) //  [[{101.date_time]]
    time.Sleep                      //  
      (time.Millisecond * amt)      //  [[101.date_time}]]
  }
  ... m.Average(...)
  rawString01 := `lorem ipsum...    //< Multiline string def. 
... lorem ipsum ...                 //                       
....`                               //
 x := 42.0
 //
 sumToX := func(b float) bool {     //< Declare+Define clouse, funct. within 
                                    //   current exec. ctx (sumToX will have stack
                                    //   == private Stack + "parent" stack at def. point
  return b == x                     //   x point to x defined above
 }                                  //   WARN: x value can be changed by some other code
 switch x {
   case x>10:                       //
     ...                            // 
     ... ; fallthrough              //< `fallthrough`: evaluate next case  (x>20)
                                    // WARN: This behaviour differs from (buggy)
                                    // C/JAVA/... (a break is needed to stop evaluating)
   case x>20: 
     ...
   case x>30:
   default:                         //< (optional) execute if no previous case matches.
     ... 
 }
 res, err := http.Get("...")        //< open I/O HTTP Resource [[{io.101,io.http]]
 defer res.Body.Close()             //< defer will close resources at fun.exit [[{qa.101}]]
 body, err :=                       //
   ioutil.ReadAll(resp.Body)        // [[}]]

 switch c := instance.(type) {      //< switch-case based on instance type
   case string:
       ...
 }

func main() {                       //< Entry point to App
  const (                           //< DEFINE CONSTANT VALUES [[{qa.101,qa.inmutable_code]]
     HELLO = "Hello"                //  (The more constants, the best)
     WORLD = "World"
  )                                 //  [[qa.101,qa.inmutable_code}]]

  type IPv4    [4]byte              //< Easy way to assign a type to a "bunch-of-bytes"
  type IPv6   [10]byte              // To convert back to []byte array: return myIPv4[:]



  var x string = HELLO+" "+WORLD    //< string concatenation [[strings.101]]
                                    //  x := HELLO+" "+WORLD will use compiler's type inference
                                    //
  parsingCommandLineValues()
  input101()
  builtInNumericTypes()
  builtInStrings()
  builtInFunctions()
  golangPointers(p1, p2 *int)
  Channels_And_Go_Routines_Interaction()
  goBuffers()
}

// [[{qa.UX]]
funct parsingCommandLineValues() {  
  //  (@package main only) 
  var maxp := flag.int(             // <·· maxp: pointer to flag structure. 
     "max", 6 /*def.value*/,        // <·· define CLI flag 'max'
     "max value" /*description*/)
  flag.parse()                      // <·· parse.
  fmt.println(...*maxp)             // <·· Referencing parsed value. 
}
// [[qa.UX}]]

funct input101() {
  var input01 float64               //  
  fmt .Scanf ("%f", &input01)       //< formated input from STDIN to 'input01' var.[[IO.101]]
  fmt .Printf("%f", input01 )       //< formated output to STDOUT                  [[strings.101]]
  fmt .Printf(x)                    //< Printf. alt 1
  fmt .Println("1+1=", 1.0+1.0)     //< Printf. alt 2

}

func builtInNumericTypes() {        
  // [[{data_structure.101]]
                                    // REF: <https://go.dev/ref/spec#Types>
  var a unit8 = 5                   // <·· (u)int8/16/32/64 
  var b rune  = 0x1111              // <·· rune type is an alias for int32
  var b byte  = byte('\n')          // <·· byte type is an alias for unit8
  var pi float32 = 22. / 7          // <·· float32/64
  var f1 = float64(len(int01))      // <·· type conversion (int to float64) 
  var c1 := 3 + 4i                  // <·· complex64/128
  var err := NaN                    // <·· NaN == positive/negative inifinity  
  // [[data_structure.101}]]
}

// [[{101]]
func builtInFunctions() {

  /* 
   * (<https://golang.org/ref/spec#Built-in_functions>)
   * <https://golang.org/pkg/builtin/>
   *
   * append elements to the end-of-slice. Returns new updated slice.
   * that must be stored, often in existing var. holding the original slice.
   * 
   * Signature: 
   * - append(slice []Type, elems ...Type) []Type
   */

  sliceRefInFunc  = append( sliceRefInFunc , elem1, elem2)  
  sliceRefInFunc  = append( sliceRefInFunc , anotherSlice...) // 
  slice           = append([]byte("hello "), "world"...) // appending string to byte slice is legal

  cap(v Type) int // <·· TODO:(0) return capacity of v, according to its type

  close(c chan<- Type) // <·· closes channel ¹
  /*
   * ¹ channel must be either bidirectional or send-only.
   * WARN: It should be executed only by the sender "controlling" the data
   *       entering the channel, never the receiver,
   *       effect: shuts down channel after last-sent-value is received.
   * WARN: After last-value has been received from a closed channel c,
   *      any receive from c will succeed without blocking, returning the
   *      zero value for the channel element.
   * WARN: "x, ok := <-c" will also set ok to false for a closed channel.
   */

  /*
   * copy elements from source-slice into destination-slice.
   * return number-of-elements copied, (it will be the minimum of len(src)|len(dst).
   * WARN: source and destination may overlap.
   */
  copy(dst, src []Type) int :                                               [[data_structure.slice]]

  delete(m map[Type]Type1, key Type) :                                      [[data_structure.map]]
  /*
   * delete element with specified key from map.
   * If m == nil or key not in map, it results in a no-op.
   * WARN: This could mean a programming error. Check first that key
   *       exists and throw like .
   */

   if !val, ok := map01["key"]; ok {
     // raise error
   }
   
   len(v Type) int // returns length-of-v, according to its type             [[data_structure]]

   /*
    * allocate+initialize object of type slice|map|chan
    *   Like "new", 1st argument is a type, not a value.
    * Unlike "new", return type is same as argument-type (vs pointer )
    */
   make(t Type, size ...IntegerType) Type :
    
   new(Type) *Type : // allocates memory and return pointer to newly allocated type's "zero-value"

  /*
   * panic: stop normal execution of current-goroutine
   * termination sequence can be controlled by the built-in func. recover() .
   */
   panic(v interface{}) :
   recover() interface{} :
  /*
   * allows program to manage behavior for panicking goroutine.
   * Executing a call to recover inside a deferred function                    [[PM.TODO]]
   * (but not any function called by it) stops the panic sequence
   * by restoring normal execution and retrieves the error
   * value passed to the call of panic.
   * If recover called outside deferred-function it will not stop panic sequence.
   */
}
// [[101.functions}]]

// [[{strings.101]]
func builtInStrings() {

  var s1 = "asdf" + `asdf`          // newlines allowed
  fmt.Printf(len("asdf"))           // 4
  var s2 = "asdf"[0]                // ← s2 == a

  // ------ String package --------------------------------
  strings.Contains("test", "es")    // true
  strings.Count("test", "t")        // 2
  strings.HasPrefix("test", "te")   // true
  strings.HasSuffix("test", "st")   // true
  strings.Index("test", "e")        // 1
  strings.Join([]string{"a","b"}, "-") // "a-b"
  strings.Repeat("a", 3)            // == "aaa"
  strings.Replace("aa","a","b",1)   // "ba"
  strings.Split("a-b-c", "-")       // []string {"a","b","c"}
  strings.ToLower("TEST")           // "test"
  strings.ToUpper("test")           // "TEST"
  arr := []byte("test")             // string to byte-array
  str := string([]byte{'t','e','s','t'}) // byte-array to string
  // See also: text_templates

  // ------ Check ¿string is empty? ----------------------
      strings.TrimSpace(" ")  != ""  // < Alt1: Prefered in newer versions of Go
  len(strings.TrimSpace(" ")) != 0   // < Alt2: micro-optimization.
                                     //         Harder to read, still used.
}
// [[strings.101}]]


// [[{101.pointers]]

func golangPointers(p1, p2 *int) {
/**
 * GOLANG IS GARBAGE COLLECTED. => returning pointers to function
 * local variables are "OK". Local variable will stay in memory while 
 * some external reference is in place:
 * NO POINTER ARITHMETIC EXISTS.
 * nil pointer and nil pointer errors can arise :(
 * pointers introduce some complexity, but allows for much faster 
 * and memory efficient code in complex data structure scenarios.
 * ex: <https://www.infoq.com/articles/go-pointers-references-graphs-tutorial/>
 */
 val1 := 1
 pVal1 *int = &val1    // <·· new pointer to val1
                       //     & reference operator returns a int pointer type, i.e,
                       //     a typed memory address, allowing direct access and 
                       //     manipulation of the data located at that address.
 *pVal1  == val1       // <·· always true.
                       //     * pointer operators takes typed pointers as inputs 
                       //     and return the (typed) value pointed to.

 fmt.Println("val1 val: ",  val1) // val1 val: 1
 fmt.Println("val1 add: ", pVal1) // val1 add: 0x140000b2008
 *pVal1 = 2             // <·· Similar to val1 = 2.

 ptr1 := new(int)      // <·· new Pointer unnasigned
 s := make([]int, 20)  // Allocate 20 ints as a single block of memory.
 s[3] = 7              // Assign one of them.
 return &s[3], &s;     // & fetch address of object.
}

// [[101.pointers}]]

// [[{101.channels,101.goroutines]]
func Channels_And_Go_Routines_Interaction() {

/**
 * channel: CONCURRENCY-SAFE COMMUNICATION OBJECT used for:
 * - Concurrent internal go routines running in parallel using
 *   the channel to synchronize work.
 * - Async/reactive implementation (Concurrent code sending / receiving
 *   data in another machine outside our control).
 */

 /**  
  * goRoutine1: Standard function to be run as go-routine
  * (vs called as part of a thread). Syntax is similar to
  * standard functions but we must pass 1+ channel/s 
  * object/s to allow the go routine to communicate with 
  * other threads
  */
  func goRoutine1(i int, c chan int /* ¹ */) {   
      // ...                             
      c <- i + 1 //  <- write to channel 
  }                                      

  /**
   * 
   * ¹ We can restrict whether function/go-routine 
   *   will be able to read/write from/to channel 
   *   with channel sub-types:
   * 
   *   c chan   int: code can read/write to/from channel
   *   c chan<- int: code can write      to      channel
   *   c <-chan int: code can read          from channel
   */

  funct funcSettingUpGoRoutint1() {
    c   := make(chan int)         // ← create new unbuffered    int channel
    cs  := make(chan string)      // ← create new unbuffered string channel
    ccs := make(chan chan string) // ← create new unbuffered string-channel channel
    //                                            └────┬───┘
    // ┌───────────────────────────────────────────────┘
    // - goroutines writting (trying to write) to Unbuffered channel will block
    //   when no 'peer' thread is waiting to read from such channel.
    // - Use make( chan T, N /*buffer capacity*/ ) for buffered channels.
    //   Writting goroutines will continue to writing until buffer is full.
  
    go goRoutine1(  0,  c) // ← start go rutine 1 ┐ Any go routine can be first
    go goRoutine1( 10,  c) // ← start go rutine 2 ├ writing to unique channel c
    go goRoutine1(─805, c) // ← start go rutine 3 ┘ and so output will be random
    fmt.Println(<-c, <-c, <-c) //                 ← in this line
  
    go func()   { c <- 84  }()  // inline go-routine writing value to c  channel.
    go func()   { cs<- "." }()  // inline go-routine writing value to cs channel.
    select {               // <· select waits until something is received from channel
      //   ┌───────┬─······      varX := <─ c translates to "wait trying to read from c"
      case i := <-c:       // <┐ 
        fmt.Println(i)     //  ├ One go-routine will win. All other "competing" goroutines
      case s := <-cs:      // <┘ will block. (No one will be reading what they write)
        fmt.Println(s)
      case sc := <-ccs:    // <· Empty channel, NOT ready for communication.
        fmt.Println(sc)
    }
  }
}
// [[101.channels,101.goroutines}]]


// [[{io.101.buffers,01_PM.TODO]]
func goBuffers() {
  /*
   * - <https://pkg.go.dev/bytes>
   * type Buffer        <··· Manage byte[] arrays with File-like interface: Read (with drain) & Write
   * -----------
   * func NewBuffer      (buf []byte) *Buffer
   * func NewBufferString(s string  ) *Buffer
   * func (b *Buffer) Bytes() []byte
   * func (b *Buffer) Cap()   int
   * func (b *Buffer) Grow(n int)
   * func (b *Buffer) Len() int
   * func (b *Buffer) Next(n int) []byte    <· returns slice containing next n-bytes from buffer,
   *                                           advancing the buffer as if bytes had been returned
   *                                           by Read
   * func (b *Buffer) Read(p []byte)        <· reads next-len(p) bytes from buffer 
   *                (n int, err error)         (or until buffer is drained).
   *                └────────────────┴─······· returns effective 'n' bytes read. 
   *                                           err == io.EOF if there are no more bytes to read,
   *                                           (unless len(p) is zero); otherwise it is nil.
   * func (b *Buffer) ReadByte() (byte, error)
   * func (b *Buffer) ReadBytes(delim byte) (line []byte, err error)
   * func (b *Buffer) ReadFrom(r io.Reader) (n int64, err error)
   * func (b *Buffer) ReadRune() (r rune, size int, err error)
   * func (b *Buffer) ReadString(delim byte) (line string, err error)
   * func (b *Buffer) Reset()
   * func (b *Buffer) String() string
   * func (b *Buffer) Truncate(n int)
   * func (b *Buffer) UnreadByte() error
   * func (b *Buffer) UnreadRune() error
   * func (b *Buffer) Write(p []byte) (n int, err error)
   * func (b *Buffer) WriteByte(c byte) error
   * func (b *Buffer) WriteRune(r rune) (n int, err error)
   * func (b *Buffer) WriteString(s string) (n int, err error)
   * func (b *Buffer) WriteTo(w io.Writer) (n int64, err error)
   *
   * var buf bytes.buffer // no need to init (make)
   * buf.write( []byte("test") )
   */ 
}
// [[io.101.buffers}]]



// [[101.hello_world}]]
