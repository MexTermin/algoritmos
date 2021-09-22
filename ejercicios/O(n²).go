package ejercicios

import (
    "sort"
)

func MakeArrayConsecutive2(statues []int)int {
    sort.Ints(statues) // O(n²)
    var result int // 1
    for i:= range statues{ // O(n)
        if i+1  >= len(statues){ break } // 1
        result += (statues[i+1] - statues[i])-1 // 1
    }
    return result
}
