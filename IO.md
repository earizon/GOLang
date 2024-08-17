# File I/O [[{io.101.files,io.101.persistence]]

<!-- Simple 'String' ←(read-from/write-to)→ 'File' _-->

## write in-memory 'String' to 'File'

  ```
  | import ( "os" )
  | func WriteFile() {
  |   F_NAME := "test.txt"                 
  |   const data []byte := []("lorem...") 
  |   if err := os.WriteFile(F_NAME,
  |      err != nil {
  |       panic(err)
  |   }
  |   defer file.Close()                  // Don't forget!!!
  |
  |   // no buffer, no looping, no async,...
  |   if writtenLen, err := file.WriteString("...") ;
  |   /* └────┬───┘*/err != nil {
  |   // effectively written string len.
  |       panic(err)
  |   }
  | }
  ```

## Read 'File' to in-memory 'String'

  ```
  | import ( "os" )
  | func ReadFile() {
  |     F_NAME := "test.txt"
  |     if data, err := os.ReadFile(F_NAME);  // go 1.16+
  |              err != nil {                 // ioutils.ReadFile in go 1.15-
  |        panic(err)
  |     }
  | }
  ```

## Read "pieces" of files

  ```
  | import ( "io" "os")
  |
  | func ReadPieceOfFile() {        //< WARN:  error control removed
  |   f, err := os.Open("/tmp/dat")
  |   ... error control ...
  |   defer  f.Close()              //< "defer": Don't forget to close
  |
  |   b1 := make([]byte, 5)
  |   n1, err := f.Read(b1)         //< read 5 bytes (@ start-of-file)
  |   ... error control ...         //   n1: effective bytes read
  |
  |   o2, err := f.Seek(6, 0)       //< Move in file for next read/write
  | //^                 ^  ^            Move relative to:
  | //·                 ·  └─·········· 0: file-origin
  | //·                 ·               1: current offset
  | //·                 ·               2: file-end
  | //·                 └·············· offset "jump" in bytes
  | //└································ new file offset
  |                                     For write-append behavior is undefined
  |   b2 := make([]byte, 2)
  |   n2, err := io.ReadAtLeast(f, b2, 2) ← TODO: Wait otherwise?
  |
  |   _, err = f.Seek(0, 0)
  |
  |   r4 := bufio.NewReader(f)
  |   b4, err := r4.Peek(5 /* bytes */)
  | }
  ```

## Read "pieces" of files using buffers: [[{io.101.buffers]]

  ```
  | import (
  |     "bufio"           //< bufio implements a buffered reader
  |     "io"              //  efficiency with many small reads
  |     "os"
  | )
  |
  | func ReadPieceOfFile() {           //< WARN: error control removed
  |     f, err := os.Open("/tmp/dat")
  |     ... error control ...
  |     defer  f.Close()               //< "defer": Don't forget to close
  |
  |     r4 := bufio.NewReader(f)         //< "wrap" File into buiffer
  |     b4, err := r4.Peek(5 /* bytes */)//< Operate on buffer
  | }                                             
  ```
[[io.101.buffers}]]

[[io.101.files,io.101.persistence}]]

## Sling HTTP client [[{IO.http.client]]
* HTTP client library.
* Sling stores HTTP Request properties to simplify sending requests and decoding responses.
[[IO.http}]]

## data structure <··> JSON [[{io.serialization.JSON]]

  ```
  | package main
  |
  | import (
  |     "encoding/json"                   <- Builtin library!!!
  |     "fmt"
  |     "os"
  | )
  |
  | type customStruct01 struct {
  |     Page   int
  |     Fruits []string
  | }
  |
  | type customStruct02 struct {
  |     Page   int      `json:"page_i"`   <- Only Fields starting with capital
  |     Fruits []string `json:"fruit_l"`     letters will be en|de-coded to/from
  | }                                        JSON
  |                      └···········┴─   <- tags allow to customized output
  |                                          JSON key names (otherwise same name
  |                                          is used).
  |
  | func main() {                            OBJECT to JSON string
  |                                          =====================
  |   bolB, _ := json.Marshal(true)       <- encoding basic data types to JSON
  |   fmt.Println(string(bolB))              (boolean, int, float, strings,..,
  |                                          as well as slices and maps.
  |
  |   res1D := &customStruct01{           <- No problem to encode custom structs.
  |       Page:   1,
  |       Fruits: []string{"...", ...}
  |   }
  |   res1B, _ := json.Marshal(res1D)
  |
  |                                          JSON string to OBJECT  (example 1)
  |                                          =====================
  |   byt := []byte(                      <- INPUT JSON
  |          `{"num":6.1,"strs":["a","b"]}`)
  |
  |   var dat map[string]interface{}      <- STEP 1) Variable to store decoded JSON
  |                                          map of strings to arbitrary data types.
  |
  |   if err := json.Unmarshal(byt, &dat);<- STEP 2) Unmarshal
  |      err != nil { panic(err) }        <- error control
  |   fmt.Println(dat)
  |
  |   num := dat["num"].(float64)         <- STEP 3) convert map values to suitable type.
  |
  |                                          nested data requires a series of conversions.
  |   strs := dat["strs"].([]interface{}) <- 1) Convert "strs" values to array of any-type
  |   str1 := strs[0].(string)            <- 2) Convert element to string
  |   fmt.Println(str1)
  |                                          JSON string to OBJECT  (example 2)
  |                                          =====================
  |   str :=                              <- INPUT JSON (to be decoded to custom data type)
  |       `{                                 - Adds additional type-safety
  |          "page": 1,                      - eliminates manual type assertions.
  |        "fruits": ["apple", "peach"]
  |       }`
  |
  |   res := customStruct02{}
  |   json.Unmarshal([]byte(str), &res)
  |   fmt.Println(res.Fruits[0])
  |
  |                                          USING STREAMS (os.Writers children)
  |                                          =============
  |
  |   enc := json.NewEncoder(os.Stdout)    <- Create JSON stream encoder (os.Stdout, HTTP response bodie,...)
  |   d := map[string]int{"...": 5, ...}
  |   enc.Encode(d)
  | }
  ```

## JSON API Summary  (https://pkg.go.dev/encoding/json)

* <https://go.dev/blog/json>
  ```
  | func Compact                type UnsupportedTypeError
  | func HTMLEscape             type UnsupportedValueError
  | func Indent                 type SyntaxError
  | func Marshal                type MarshalerError
  | func MarshalIndent          type InvalidUnmarshalError
  | func Unmarshal              type UnmarshalTypeError
  | func Valid /*isValid?*/
  |
  |
  | type Decoder             type Encoder     type   MarshalJSON
  | - NewDecoder             - NewEncoder     type UnmarshalJSON
  | - Buffered               - Encode         type Unmarshaler
  | - Decode v any           - SetEscapeHTML  type Marshaler
  | - DisallowUnknownFields  - SetIndent      - Error
  | - InputOffset                             - Unwrap
  | - More                                    - Error
  | - Token                                   - Unwrap
  | - UseNumber
  |
  |
  | type Number   type Delim  type RawMessage
  | - Float64     - String()  type Token
  | - Int64
  | - String
  ```

[[io.serialization.JSON}]]

## Mapstructure map<··>structure [[{IO.serialization,persistence,io.json,qa.error_control]]
<https://github.com/mitchellh/mapstructure>
* library for decoding generic map values to structures and vice versa,
  while providing helpful error handling.

   This library is most useful when decoding values from some data
  stream (JSON, Gob, etc.) where you don't quite know the structure of
  the underlying data until you read a part of it. You can therefore
  read a map[string]interface{} and use this library to decode it into
  the proper underlying native Go structure.
[[}]]

## Simple HTTP client/server [[{io.http]]
  • HTTP Simple client <https://gobyexample.com/http-clients>

    package main                              ← Exec like   $ go run http-clients.go

    import ( "bufio" "fmt" "net/http" )

    func main() {
      resp, err :=
         http.Get("http://....")              ←  convenient shortcut around
      if err != nil {                            creating http.Client instance
          panic(err)                             , then calling its Get method
      }                                           using http.DefaultClient with
                                                  sensible default settings

      defer resp.Body.Close()                 ←  Close I/O resource on defer.

      fmt.Println("Res.Status:", resp.Status)

      scanner := bufio.NewScanner(resp.Body)
      for i:= 0; scanner.Scan() && i<5; i++ { ← Print first 5-lines of the
        fmt.Println(scanner.Text())             res.body.
      }

      if err := scanner.Err(); err != nil {
          panic(err)
      }
    }

  • HTTP Simple server <https://gobyexample.com/http-servers>
    import ( "fmt" "net/http" )

    func hello( w    http.ResponseWriter,    ← handler 1
          req *http.Request) {
      fmt.Fprintf( w , "hello\n")
    }

    func headers(                             ← handler 2. echo Req.headers
        w    http.ResponseWriter,                          to Res.body
        req *http.Request) {
      for key, value_l
          := range req.Header {      ←  key: Header name, value: 1+ values
        for _, h := range value_l {
          fmt.Fprintf( w ,
            "%v: %v\n", key, h)
        }
      }
    }

    func main() {
        http.HandleFunc("/hello"  , hello  )
        http.HandleFunc("/headers", headers)
        http.ListenAndServe(":8090", nil) // nil: use def.router just defined
    }

  • See also: https://github.com/valyala/fasthttp
    - HTTP tuned for high performance.
    - Zero memory allocations in hot paths.
    - Up to 10x faster than net/http .

* Simple HTTP server using gin-gonic:
  https://github.com/gin-gonic/

  PRE-SETUP:
  go 1.16+
  $ go get -u github.com/gin-gonic/gin

  import "github.com/gin-gonic/gin"
  import "net/http" // Optional, needed if reusing constants like http.StatusOK

  func main() {
    router := gin.Default()
    router.Static("/assets", "./assets")
    router.StaticFS("/more_static", http.Dir("my_file_system"))
    router.StaticFile("/favicon.ico", "./resources/favicon.ico")
    router.StaticFileFS("/more_favicon.ico", "more_favicon.ico", http.Dir("my_file_system"))

    // Listen and serve on 0.0.0.0:8080
    router.Run(":8080")
  }
[[io.http}]]
