package amount

var belowTwenty = [19]string{"One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}

var tens = [9]string{"Ten", "Twenty", "Thirty", "Fourty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}
var units = [4]string{"Hundred", "Thousand", "Lakh", "Crore"}

const twenty = 20
const hundred = 100
const thousand = 1000
const lakh = 100000
const crore = 10000000

// Converts an uint to string (corresponding amount in words). Suitable for indian style i.e uses Lakhs and Crores instead of Millions.
func ToWords(amount uint) string {
	if amount == 0 {
		return "Zero"
	}

	return checkAmount(amount)
}

func checkAmount(amount uint) string {
	space := " "
	switch {
	case amount == 0:
		return ""

	case amount < twenty:
		return belowTwenty[amount-1]

	case amount < hundred:
		if amount%10 == 0 {
			return tens[(amount/10)-1]
		} else {
			return tens[(amount/10)-1] + space +belowTwenty[amount%10 - 1]
		}

	case amount < thousand:
		return checkAmount(amount/hundred) + space + units[0] + space + checkAmount(amount%hundred)

	case amount < lakh:
		return checkAmount(amount/thousand) + space + units[1] + space + checkAmount(amount%thousand)

	case amount < crore:
		return checkAmount(amount/lakh) + space + units[2] + space + checkAmount(amount%lakh)

	default:
		return checkAmount(amount/crore) + space + units[3] + space + checkAmount(amount%crore)
	}
}
