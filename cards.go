// package cards provides basic structures and functions for dealing with playing cards.
package cards

import ("math/rand"; "time")

// Card is a struct holding the suit and value associated with a card.
type Card struct {
  Suit  string
  Value string
}

// NewCard creates a new instance of the card struct.
func NewCard(suit, value string) Card {
  return Card{Suit: suit, Value: value}
}

// NewDeck creates a new deck of cards from a slice of suits and a slice of values.
func NewDeck(suits, values []string) []Card {
  var deck []Card
  for _, suit := range suits {
    for _, value := range values {
        card := NewCard(suit, value)
        deck = append(deck, card)
    }
  }
  return deck
}

// NewStandardDeck creates a new standard western deck of cards with French suits
// (Spades, Clubs, Hearts, and Diamonds).
func NewStandardDeck() []Card {
  suits  := []string{"♠", "♣", "♥", "♦"}
  values := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}
  return NewDeck(suits, values)
}

// NewFifthDimensionDeck creates a deck of 5°Dimension Playing Cards.
// 5°Dimension is unique in that it has five suits, a Princess, a '1' different from Ace,
// and a Joker for each of the five suits.
// The five suits are: Spades, Clubs, Hearts, Diamonds, and Stars.
// The 5°Dimension deck has a total of 80 cards, as opposed to the usual 52.
func NewFifthDimensionDeck() []Card {
  suits  := []string{"♠", "♣", "♥", "♦", "★"}
  values := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "P", "Q", "K", "A", "Joker"}
  return NewDeck(suits, values)
}

// Shuffle randomises the order of the cards in-place.
// It uses "math/rand" internally, and thus needs to be seeded before use.
// cards.Seed() and cards.SeedInt(int64) can be used for this
func Shuffle(deck []Card) {
  for i := range deck {
    j := rand.Intn(i + 1)
    if i != j {
      deck[i], deck[j] = deck[j], deck[i]
    }
  }
}

// Shuffled is like Shuffle, except it's not destructive
func Shuffled(deck []Card) []Card {
  cp := make([]Card, len(deck))
  copy(cp, deck)
  Shuffle(cp)

  return cp
}

// SeedInt seeds the card randomiser using an int64 value as its seed.
func SeedInt(value int64) {
  rand.Seed(value)
}

// Seed seeds the card randomiser using the current time in nanoseconds
func Seed() {
  SeedInt(time.Now().UTC().UnixNano())
}
