package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "4.8" )

func T7oyeBbrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tIq2d6kOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GzMI3gHHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2ESYRUNfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W6CgEhqdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KaUwh9fZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WOiyCvEkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sG3RxwTYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5VBMnX6iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GW4Q3OMzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FIfOPoe6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Iu5aTQyPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ICsFGfxkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QYkmN9DjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hC3RiZk4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VSCGYVLnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pbegS6UTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BMEWAWOuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jneDlSHaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vjDvAv3nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rVKxSmhjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zo1oh7QJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f565VUkmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cQSqetc3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JUuaTdqVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JRau3S56Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RCJH0yx0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0TiSJlqUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uLfoCqYEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WQsKzpzrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WNtGLEfdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nIC9y0erWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wr8qNQl7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zZXPddAhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ptZbRAUWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mu0ZXXH1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZbWGomo7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F1ROuanJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 71gpktfBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wrbmHuGKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v7jbpHyzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fukCxnQFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kddJLPuzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4NdXEEdlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7rvqQ8f9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HlY8gSNTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MRtFTDpUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DZ1IBN5hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GUj8tQj2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N98FjrotWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 18HpfDk4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jN8C13xlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U9sWy72FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func diIgXAkkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cfp0xQoKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sgvMi1tWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d1WjmfjCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a0jqAnZjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5YyUTZR5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LTdqooHZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P9Wfk16jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5VZsvc4oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HXHuOl04Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eVbW98bjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7mFrnRsbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EMAbin0OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0CVNGeaAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mh5XYiheWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mNYaLAulWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h2rgzTT2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wsPMSB08Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GZcwD0yIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JZOSNuXfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jZDWkjuzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func koMgb3kYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gz94mMS9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fHQLXwYLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mHcq0r5EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GigHrOWmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JPPBGfZuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qqxZsZ54Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZMNGrnUrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i45HaR7hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pVympQMCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kysVGcS0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zgr8R8WQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OWPWJpCjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EKSYc85aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pnvZlMzXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O4Ju2LBVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PPNwEe6HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B8kT9BsVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iu1dwHdMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3C1H1nsFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 63iINSMpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iPWQe2rlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 906DDpTiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lV1sNfCsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CRB3zWK0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func svr6y379Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q3W1yPPNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mlFR7gdPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2Vpk5yv9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BVamE8TNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P9GURnsGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CkYhGIGWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9nymHMFzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XiZnEZBBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3A3Qn2qTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IviQgdHmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dCsADfU7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2Cgrui2MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LeN1BYTDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BdNjXbrTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PrdE3zeuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2tgDbdGLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i8gpEk2jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pNO4zBwCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oMjplDv1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hN7EO3XhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V5xYSAkFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iu2JWQmAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BK5xT5pUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tdcEkcTvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7gFCjnZbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aNhteWlhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PNYcIgZYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cXslLFuoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GitdAFImWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rdaXZXq3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0OLrIz5eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Msb4mPS4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RNJYvjvZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3MtfkMhcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PUFwH2StWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YmDzR1I4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vmnfYTqUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2AjqPanuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func khpoMdLJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func moGfRAG0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4Pju4nmKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dtMEOzkzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qAMeQLVaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KMtgSnp1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SbAVmolBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EM6inP79Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8DxiBSQxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func POrsPETAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4JDXRFRmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k8E7EwQZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pVb3FKRfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9gxtFdEXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gs0ECAPPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4b908mJLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z64PS7reWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SkICJAbjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aH4WRQo4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cKvcSRi5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k2YLEYq5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B6M5NA0vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xedJQd2EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IyvSAtY2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dg3SF7PqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VDNZi5dhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l3gjFms6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cIuLY6fOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z9f3cVTHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6dVbgtc5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4Ko7eahfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GJJowXbCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x3A3KXc9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LY3sfDz8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func diuU9lvgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FAvNx0bJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lSVNpxOzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z9w0EwdaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ut6oohQyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func urM3VDUtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s5PoLbAlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kAinC74HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9md3QnItWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tCtD4mJuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MB4Qg8brWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WYwut8laWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O0ISyGJVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Yfdh8ltWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gDPcavrRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bS6GkmFcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0YSngzSWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zzV4X4dSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eRu0PIT8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f8osFOq3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t4ER83P1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ueMGaCveWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4Hbqpmj0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JrmJZQ2RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B3Qlowf5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j2wIUqpuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1BhzY1QeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ArxR6MiUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rzL5s6ulWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OLohifCuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ekSjHYa5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TEOdpEF3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lv4JEDZdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ljwHIYndWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yXRnZ3OlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7pMLDJlCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2k3pmVrTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0ClPPhdoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vmaSFwC7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3GQ7e71iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AZuPFngNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fSoBXcq5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hn1TkoR1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3WjQm6pCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4JFz7mTgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nNj40OmEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nDK9hCfXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E7sYcd0TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zWmzgFBUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OUZfxpAfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xpaxx8lvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cBT288aCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tUExsCtSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0RRw1Bs3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b7rvOhgXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wK89kkehWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RJFYyD77Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nZnxk8E8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lBRVC2UVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rhVDB6Y5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U4f1nVPRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0sjXp5cvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X8oA60tUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JodXx2HeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m7VwMn4aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lYOKibRuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L2fpUY5uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1Rg29V6VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func weynlckPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func maqePZynWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Nfw1cQ5VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u7ijJ0UGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LELLORzlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IAJMYgx6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9uhQTO6GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4qEfDovnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iYZESgXwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eu2lFEO2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SFrUL3KTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func evbfj8oEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nQc08a2IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4Rwk4HciWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q9lFTnvRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

