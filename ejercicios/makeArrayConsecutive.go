// Package morestrings implements additional functions to manipulate UTF-8
// encoded strings, beyond what is provided in the standard "strings" package.
package ejercicios

import (
    "sort"
)

func MakeArrayConsecutive2(statues []int)int {
    sort.Ints(statues)
    var result int
    for i:= range statues{
        if i+1  >= len(statues){ break }
        result += (statues[i+1] - statues[i])-1
    }
    return result
}
