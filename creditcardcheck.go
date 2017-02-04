package creditcardcheck

import (
	"strconv"
	"strings"
)

type CreditCard struct {
	CardNumber string
}

type CreditCardStatus struct {
	CardType   string `json:"cardType"`
	CardNumber string `json:"cardNumber"`
	CardStatus string `json:"cardStatus"`
	Error      error  `json:"error"`
}

func sumOfTwoDigit(a int) (int, error) {
	t := strings.Split(strconv.Itoa(a), "")
	v1, err := strconv.Atoi(t[0])
	if err != nil {
		return 0, err
	}
	v2, err := strconv.Atoi(t[1])
	if err != nil {
		return 0, err
	}
	return v1 + v2, nil
}

func (c *CreditCard) retriveCreditCardType() string {
	Len := len(c.CardNumber)
	switch {
	case (c.CardNumber[0:2] == "34" || c.CardNumber[0:2] == "37") && Len == 15:
		return "AMEX"
	case c.CardNumber[0:4] == "6011" && Len == 16:
		return "Discover"
	case c.CardNumber[0:2] >= "51" && c.CardNumber[0:2] <= "55" && Len == 16:
		return "MasterCard"
	case c.CardNumber[0:1] == "4" && (Len == 13 || Len == 16):
		return "Visa"
	default:
		return "Unknown"
	}
}

func (c *CreditCard) creditCardNumberValidate() (bool, error) {
	sum := 0
	Len := len(c.CardNumber)

	for i, _ := range c.CardNumber {
		d, err := strconv.Atoi(string(c.CardNumber[Len-1-i]))
		if err != nil {
			return false, err
		}
		if i%2 != 0 {
			d = d * 2
			if d > 9 {
				var err error
				d, err = sumOfTwoDigit(d)
				if err != nil {
					return false, err
				}
			}
		}
		sum = sum + d
	}
	return sum%10 == 0, nil
}

func CheckCreditCard(cardNumber string) *CreditCardStatus {
	var status string
	card := &CreditCard{
		CardNumber: strings.Replace(cardNumber, " ", "", -1),
	}
	ct := card.retriveCreditCardType()
	n, err := card.creditCardNumberValidate()
	if err != nil {
		return &CreditCardStatus{
			CardNumber: card.CardNumber,
			Error:      err,
		}
	}
	if n == false {
		status = "invalid"
	} else {
		status = "valid"
	}
	return &CreditCardStatus{
		CardType:   ct,
		CardNumber: card.CardNumber,
		CardStatus: status,
	}
}
