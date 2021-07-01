# amount

amount is a simple go package that converts financial amount to words. Catering to indian style i.e lakh and crore instead of millions or billions.

## Installation
```
go get github.com/jethi/amount
```

## Use
Simply use the ToWords function after installing and importing.

### Import:
```
import "github.com/jethi/amount"
```

### ToWords
ToWords accept a uint and returns corresponding amount in words.
```
amount.ToWords(uint) string
```

## Example
```
package main

import (
	"fmt"
	"github.com/jethi/amount"
)

func main() {
	var myNum uint
	var ans string
	
	myNum = 561085
	ans = amount.ToWords(myNum)
	fmt.Println(ans)
}
```
### Output:
```
Five Lakh Sixty One Thousand Eighty Five
```

