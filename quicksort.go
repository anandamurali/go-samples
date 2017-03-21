package main
import(
  "fmt"
)

func main1(){
  fmt.Println("QuickSort program")
  a := []int{3,1,2,3,8,1,4,5}
  b:=quickSort(a,0,len(a)-1)
  fmt.Println(b)
}


func quickSort(a []int,start,end int) []int {
    if start >= end {
        return a
      }

    index,a := partition(a,start,end)
    fmt.Println("Index:",index)
    a = quickSort(a,start,index-1)
    a = quickSort(a,index+1,end)
    return a
}

func partition(a []int,start,end int) (int, []int) {
    pivot := a[end]
    fmt.Println("Pivot",pivot)
    i := start

    for j:= start;j < end-1;j++{
        if a[j]<pivot {
          i = i +1
          temp:= a[i]
          a[i] = a[j]
          a[j] = temp
        }
    }
    temp:= a[i+1]
    a[i+1] = a[end]
    a[end] = temp
    return i+1,a
}
