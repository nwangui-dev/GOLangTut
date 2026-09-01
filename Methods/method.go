package main

import "fmt"

//BankAccount is a struct (a type) that has a balance
type BankAccount struct{
	balance int
}

type Rect struct{
	width, height int
}

// a method is simply a function with a special receiver argument placed between the func keyword and the method name

// func (r ReceiverType) MethodName(parameters) returnTypes {
//     // method body
// }
func (r Rect) Area() int{
	return r.width * r.height
}

//deposit is a method that adds money to the account
//(a *BankAccount)simply means "a is a pointer to BankAccount" allowing the method to change the original account 
// Use a pointer receiver when a method needs to change the original value for flexibility unlike with the normal value receiver.
func (a *BankAccount) deposit(amount int) {
	a.balance += amount
}

func main() {
	myRect := Rect { width: 10, height: 5}
	fmt.Printf("Area: %v", myRect.Area())	
	fmt.Println(" ...hello chap chap on")

//create a bankaccount with variable account balance starting at 100
	account := BankAccount{balance: 100}
//then call the deposit method on account. this adds 50 to account's balance
	account.deposit(50)//tell the account to deposit 50
//'account.balance' means go inside account and get its balance, where the balance is now 150 
	fmt.Println(account.balance) //so this prints 150
}
