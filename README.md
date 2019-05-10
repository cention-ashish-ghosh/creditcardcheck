# creditcardcheck
#To install creditcardcheck package run this command in terminal
#go get github.com/ashish041/creditcardcheck


```
package main

import (
	"fmt"

	"github.com/ashish041/creditcardcheck"
)

func main() {
	cardList := []string{
		"4111111111111111",
		"4111111111111",
		"4012888888881881",
		"378282246310005",
		"6011111111111117",
		"5105105105105100",
		"5105 1051 0510 5106",
		"9111111111111111",
	}
	for _, ci := range cardList {
		r := creditcardcheck.CheckCreditCard(ci)
		if r.Error != nil {
			fmt.Printf("! Error when checking credit card number %s : %v\n", r.CardNumber, r.Error)
		} else {
			fmt.Printf("%s: %s    (%s)\n", r.CardType, r.CardNumber, r.CardStatus)
		}
	}
}
```
