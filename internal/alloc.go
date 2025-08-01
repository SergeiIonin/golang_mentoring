
package internal

import (
    "fmt"
    "runtime"
)

func PrintAlloc() {  
    var m runtime.MemStats  
    runtime.ReadMemStats(&m)  
    fmt.Printf("%d KB\n", m.Alloc/1024)  
} 
