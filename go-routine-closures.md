# Closures 

* Here the function activateGiftCard(), Having amount variables. 

* debitAmount is inner function, will subtract the amount he want to use from the amount on the giftcard. 


```
func activateGiftCard() func(int) int {
	amount := 100

	debitFunc := func(debitAmount int) int {
		amount -= debitAmount
		return amount
	}
	return debitFunc
}

```

* When we activate the Giftcard using the useGiftCard. 

* DebitFunc function will help to remove the amount of the giftcard. 

```
func main() {
	useGiftCard1 := activateGiftCard()
	useGiftCard2 := activateGiftCard()

	fmt.Println(useGiftCard1(10))
	fmt.Println(useGiftCard2(5))
}

```


* Here the amount is variable above the scope of the debitFunc function. 

* This happens due to closures. 

* Closures will help to access the variable outside the scope of the inner function. 

* This will reference on the memory. 


* Complete function looks like: 

```
func activateGiftCard() func(int) int {
	amount := 100

	debitFunc := func(debitAmount int) int {
		amount -= debitAmount
		return amount
	}
	return debitFunc
}

func main() {
	useGiftCard1 := activateGiftCard()
	useGiftCard2 := activateGiftCard()

	fmt.Println(useGiftCard1(10))
	fmt.Println(useGiftCard2(5))
}

```