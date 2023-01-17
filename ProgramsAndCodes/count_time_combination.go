package main
import (
    "fmt"
    "strconv"
    "strings"
    "time"
)
func main() {
    ProcessTime()
}
func ProcessTime() {
    //1. Given String
    var input = "3 1 2 8"
    //2. convert it into slice using strings.Fields() method
    inputSlice := strings.Fields(input)
    //3. convert string array to int array
    var sliceInputINT []int
    for _, v := range inputSlice {
        j, er := strconv.Atoi(v)
        if er != nil {
            panic(er)
        }
        //4. integer array is stored in sliceInputINT array variable
        sliceInputINT = append(sliceInputINT, j)
    }
    //5. Find all the possible combinations of input integer array which is array
    var pArray [][]int
    pArray = permutations(sliceInputINT, 4, pArray)
    //6. variable to check valid time counts in input string
    validTimeCounts := 0
    //7. maps to store unique time in input string
    map1 := map[string]int{}
    //8. loop to check all valid time in permutation and combination array
    for _, numbers := range pArray {
        //8.1. Convert them into this format "HH:MM"
        localTime := ""
        localTime += fmt.Sprintf("%v%v:%v%v", numbers[0], numbers[1], numbers[2], numbers[3])
        //8.2 Parse time in 24 hour or military time format
        // layout of valid time should be for 24 hours format
        layoutTime24 := "15:04"
        _, err := time.Parse(layoutTime24, localTime)
        if err == nil {
            _, ok := map1[localTime]
            if !ok {
                // to print valid times
                //fmt.Println(t.Format(layoutTime24))
                map1[localTime]++
                validTimeCounts++
            }
        }
    }
    fmt.Println(validTimeCounts)
}
func permutations(arr []int, l int, p [][]int) [][]int {
    if l == 1 {
        p = append(p, append([]int{}, arr...))
    }
    for i := 0; i < l; i++ {
        p = permutations(arr, l-1, p)
        if l%2 == 1 {
            arr[0], arr[l-1] = arr[l-1], arr[0]
        } else {
            arr[i], arr[l-1] = arr[l-1], arr[i]
        }
    }
    return p
}