package main
import "fmt"
var i int     //integer of input number
var s int = 0  //sum of the imput numbers
var a int     //number of the test cases
var b int     //number of the input numbers

func square (k int) int{  //function of squaring the input numbers
if k>0
    {
    return k*k
    }
return 0
}

func inner(n int) int {   //recursive function of adding up the input numbers in a round
  if n == 0 
    {
    return 1
    }
fmt.Scan(&i)
i=square(i)
s=s+i
    return inner(n-1)
}

func outer(n, k int) int {  //recursive function for the number of the test cases
  if n == 0 
    {
    return 1
    }
    inner(k)
    fmt.Println(s)
    s=0
    if n>1 
        {
        fmt.Scan(&k)
        }
    return outer(n-1, k)
}

func main() {
fmt.Scan(&a)
fmt.Scan(&b)
    outer(a,b)
}
