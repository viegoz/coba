package main

import "fmt"

type set [10]int

func main() {
 var arr set
 var n, num, i, idx, temp int
 var pass int

 fmt.Scan(&num)

 for num != -5313541 {
  pass = 1
  if num != 0 {
   arr[i] = num
   i++
  } else {
   for pass <= i-1 {
    idx = pass - 1
    n = pass
    for n < i {
     if arr[idx] > arr[n] {
      idx = n
     }
     n++
    }
    temp = arr[pass-1]
    arr[pass-1] = arr[idx]
    arr[idx] = temp
    pass++
   }
   if i%2 != 0 {
    fmt.Println(arr[(i / 2)])
   } else if i%2 == 0 {
    fmt.Println(float64(arr[i/2]+arr[(i/2)-1]) / 2)
   }
  }
  fmt.Scan(&num)
 }
}