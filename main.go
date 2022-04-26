package main

//import "fmt"

//import "fmt"

func main() {
	//var card string = "Ace of Spades"
	//	card := newCard() //prothombar : dite hobe
	//	card = "five diamonds"
	// //	cards := deck{newCard(), "hi"}
	// //	cards = append(cards, "six of spades")
	// cards := newDeck()
	// hand, remainingCards := deal(cards, 5)
	// //fmt.Println(cards)

	// // for i, card := range cards {
	// // 	fmt.Println(i, card)
	// // }
	// //cards.print()
	// hand.print()
	// remainingCards.print()
	// greeting := "Hi"
	// fmt.Println([]byte(greeting))

	//cards:=newDeck()
	//fmt.Println(cards.toString())
	//cards.saveToFile("my_cards")

	//cards := newDeckFromFile("my_cards")
	cards := newDeck() 
	cards.shuffle()
	cards.print()
}

// func newCard() string {
// 	return "Five of diamonds"
// }

