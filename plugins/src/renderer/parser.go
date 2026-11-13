package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "2.9" )

func tQXH5QIMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5HeKx3q9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tCReg0bgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 55E1bDv5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JB34dxQAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0eEwp72DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3AzJ290tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tHnMUHD4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XoMf9dB4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o3gzdDMnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XqN6YWs2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gYTKmOvLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tuAXtIteWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9eyY1FJ4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9iWDcPlCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 16FnJ8jPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VK2DtMApWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3uIWP60GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BggaGtBmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jk4gL95kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XgT3tXeHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wz4rZeHoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bnWevuvjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bA1u5rq3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 81vzRrMcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func waPL2P4AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oEpJHpodWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v0s3HbNSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d0AXlVDsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m4Bv2CxDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1QHniP1aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1Uo3XhocWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2cZq16DaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S0zeUzQhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m0adNIMdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XRYtP2a4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IvPJB9ciWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func seIVeTn2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qc5tnEiaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q1nXXnSoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eHK1W6QEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func isf3WNqbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pHLVGeTKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SQsmz9orWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CRAD9NXjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KyPGZzPcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jobDU0UkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3xEEHtW1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ehh4mREWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BGdXFmJ4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0xqMV6bnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bNTkmJhFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sX4Tm3miWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pZEREMClWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BKyjfgbZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W1aHfDo1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eoZ35CZEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3MWaGKBNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YHAjU3t2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sqIcMJvvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kGpGNEQtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lfL1drkfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GPzN1cCqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LCJWYh0JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gWQJSSLyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XPWiWkvxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bmJakmSDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tx0q3jWtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FPPLn1GkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ouHQbNqDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gTKJKGTyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KQjxJYuvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RraBpjSeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NE3APjKTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cIhlmBG0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I8VaHHLHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YPKgL3V5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 63gv5jrUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oNycc5MkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j6gds6V3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wNiuY0RLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UWFMBNhrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zeMkNn0UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6UcinhkZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ud0cxV5SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bJwA9JOTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rU24ZsHnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 56TR3uFuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eGy9ZjHqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KG4bGpg4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V8JFpyKBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aDrN5uIAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7m8N5IjIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SEgd0s2lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LwcMcxcCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vWsDpH9wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ahwkY0TVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xYPrKGa4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1fMuCBpgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h4di8gWnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H6GrdfnTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yQaLW2UiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MAdlnQ0VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UTWQGo0CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yYwpQPwpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0h6vcpEuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W0xMZOXJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uuQabqZeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BcaOHJoIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1UcPAIayWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cS5gADoxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mTpZhxuNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RLI2NS6QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W0oHaVESWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UMRWxcRXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8tK4YEblWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 34jVQkLkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cmn9Q9NyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OXpLEBW3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pnIWCCAoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NV9hqSSzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EdkE4OYjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ao33yjd4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kAGEkc8zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qODD9lVQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 997MaZV2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CuieQuBFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UiCh5Y9lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E1YWTiAVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pXnPRqWiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JJi0KSiiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s5qmHY4pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f3e1hpydWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bSDDTeF9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4STWYt14Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TOGuACNvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jtm3m4uvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Sf8WnI16Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fGH1ImA0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EbXXMmLqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6vyLuWqRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4z0xApwaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EJpKaoKXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uS8U4mhKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DAsCHWZjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mIDzrPwAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K0a3BtN8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iFJZWjwWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tvh0DBlvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A3FRpBOzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VaLANxfyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3uMZ4o2ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oPFnnKn6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rfYLxyydWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k4LB4VRNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dueovji9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d9SMnyB0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cS3RYamfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KDyptkC7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LK4vi58WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4Dy3FH7lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mh1DO6XSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func itYZSB3pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6XxzkgdbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1NWlXJurWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hv6f2rODWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9wK3IBG1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Srog8bsbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pl4S7mLbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VkCuEusKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ts2ZyvpuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wdKpi4v6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1PdjcQCyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MuPy8co1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GXDh0yRWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rsIYRjh5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sU9OxXLaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5tA2qycvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rAIz02vzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hJrwdiKMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IT01BEVwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qz8ov6yLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZqMnpNc8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 28NlTrd3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bpxn7Mx6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zBn8WauTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DtqFhVUiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EjpuBIxfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wYib0rwGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DVP7NLW2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NLCmChnjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xmo0W3UAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rRlZvbfCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N9wgzMpjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TYCU3iE6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jQW1uLTXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yh6uvsOPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7q4Zlu35Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UckKShPOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qlFMBBnaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gqswPYOBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EeEkAn5yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WinRzQvAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5mZahxC7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZQFxbCG1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GspUsDS3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aLnIYbRaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zNASfjqbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lXfmyn8pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CSx7WaDIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F8RzSAMzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tzmk6us3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tfcWg8xyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8od9kfU2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MdKGPVYxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xh697RU6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2ESgtMACWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func edS1rBAWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TKKNvmS3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yXF5zFCVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cFYgEOWWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8CitlqgcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fcXlvSJbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LdQpDSG9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6I56QiauWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wQitUe72Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dx9QBIqiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WMLuhm71Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xV4aQFsEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x1FtHHEnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lIo3rqtcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ge44pMZ1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u41fnebkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CUIad32GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SNjUwt1YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func flizYOwoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LWIALG5tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G9YtoemdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WdmDnqiSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4C3Wos6WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2oWOMnptWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zPsy66QeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X0oWm5DRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5r7leNGDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rqjvmNDqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3WzvXuAwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7xWYvipCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PlMEynbTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func om2HbUh4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eLxn9A8OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4DQ0UpYuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6QSu0M6rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FxyAmF1EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iQJz0GbSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hvt9A6jKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ncsjTrkGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KdGoYOTdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nv2eRm0SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t70UYNT5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hD4m1zmcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XvFWkprCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EIO76veSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SOBmbH2vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w8cTlyTPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6g9jOrSJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 312mxDxdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bt0UPWFiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7dJWYIT6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ufABKI1RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vH5a1HjyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 60Hzb4bMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dhqXpVfwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xRhHnHK0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3px4pNwFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 11cH2JK4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GLQxGa2QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dym2jOEBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7q8NP9QEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fin9fZ81Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func opeUp61dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZtlaS2xcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PPfNJrxfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k7th0Yr3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func myvrLK3TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f21nqjktWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pa4TAoWcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8KGV32oeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8GcAlYcdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bVNK3FSJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 15IdVQ6gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LcFkQCfEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6701D7yPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E1yPZaxRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TvIoqK2RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b0oFbw3AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MtO7HJQvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HC7RqwNVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qmR38DoSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pdueds8MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JRr4xOjKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p7Ly4JrNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sDmXLd9UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SvvUTNfkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func moIYeMRNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QmO9J5UDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jzInwxNSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iSxx3SlTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 51xFuxHDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iilmIUcXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FYUkcBWyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ygVjynlDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ooRsx7l5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VMgkfflFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GsOUF9PDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6LfcPHgjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HZx1y2duWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HXAFo1zxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aNUAmmM7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4ERlg5TvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0u7C0ysqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eLhjtDhcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func stHEimqQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OkJewGuoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ofh2RE6sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nHXRvlQHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dnvVvPfeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ff0V81U7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func liwg8ydxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ThSRiljLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LrnM2D8hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ADZDOdQSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UsBOVpviWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2IMRU4t8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uf7Jaw26Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func figSNRyvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qRULolpjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oeSRLQG6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GngtG7aFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2MGY6bK2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CxeLWbkAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ILdh27aiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mq8gnlIVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s3oz2SYxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tPbPmrlmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ftMJBxGoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1RvURegXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IwhpG5bWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZZLS7nKUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GDs5WcA3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OWl328UHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l7FRS8PHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pM01n95zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func drfhkvCaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gOTtqVTKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hLBagHDIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TtW8QOdvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gu60KhEfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WAPNMKsvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HRQJFekvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9aPy4tChWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3Tl9qhWGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Pw7rDea1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KnKj2tqEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func onWsasSlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9qktEOjNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gvry1eIhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mPlQiUvAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pVtIzqcqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N9YQqIYbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DpRJyNKBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IAt1QRE3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g6yA3AkFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KQuetORRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qRZkPbY7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func My0ZG8rcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1cJY6QwDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kXIBJ5KFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ywlPQPDBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zUkPrMOxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XAfsPkUNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GNJLxyMbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RhyLErtqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5S6b4ez0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OyPFgFEpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KU977bzfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J3Tf5OA5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pQlWQuSeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bxbHCKBwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OMfBSFeyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dl8ElKQZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uqiUisETWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RtW5dgCoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zvp5PbTAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zWGSSGFkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OcAQh5o9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eyAO9pcrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IC75HfMKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ubv2LIhuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8NujCNKnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 465fbY1UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SyLlzL3YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oM0B3aqqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func akOmRZKRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NSxlj5qPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZklyGDsEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dE5lBBboWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CBC34VoDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bUxkuXL2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ni3Wf6UxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qJFisJgKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1omIbr9WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func evlVA9SxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mHiYTg1bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f1Tj9ofBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YKGsZ8lfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H5oa2uADWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j9CRKO8HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6dtYcBR0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2SKejEMiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FCu7nQzwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BxTDi9LAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4gu2UvwIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z9MOYzL0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m5ECnDMJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e4l9K3xRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2Uzg3X52Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VmXzUcWmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oVVI96VWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HYtR2XarWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZaRiX6F0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Yvv7o2fbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6SQP61n4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FrrWlltaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func titket9mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IV9yAzoxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y8z7BxQdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 56CjgRQyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2ohuNGd8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fAEpOW9BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QHQIqDJmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eTHNllHgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func er38SjiLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T1o9dVRcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dbMPCTEOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hDbYUIMPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zsq5aDCnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kNlDoQszWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MXRXfPeIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jijAWRoSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L72QSbBxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X5Msg338Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oUb1rMlVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zxY5z6wtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LCTLUwq7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

