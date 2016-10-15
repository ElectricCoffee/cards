# cards
--
    import "github.com/ElectricCoffee/cards"

package cards provides basic structures and functions for dealing with playing
cards.

## Usage

#### func  NewDeck

```go
func NewDeck(suits, values []string) []Card
```
NewDeck creates a new deck of cards from a slice of suits and a slice of values.

#### func  NewFifthDimensionDeck

```go
func NewFifthDimensionDeck() []Card
```
NewFifthDimensionDeck creates a deck of 5°Dimension Playing Cards. 5°Dimension
is unique in that it has five suits, a Princess, a '1' different from Ace, and a
Joker for each of the five suits. The five suits are: Spades, Clubs, Hearts,
Diamonds, and Stars. The 5°Dimension deck has a total of 80 cards, as opposed to
the usual 52.

#### func  NewStandardDeck

```go
func NewStandardDeck() []Card
```
NewStandardDeck creates a new standard western deck of cards with French suits
(Spades, Clubs, Hearts, and Diamonds).

#### func  Randomize

```go
func Randomize(deck *[]Card)
```
Randomize is like Shuffle, except it's destructive

#### func  Seed

```go
func Seed()
```
Seed seeds the card randomiser using the current time in nanoseconds

#### func  SeedInt

```go
func SeedInt(value int64)
```
SeedInt seeds the card randomiser using an int64 value as its seed.

#### func  Shuffle

```go
func Shuffle(deck []Card) []Card
```
Shuffle randomises the order of the cards in-place. It uses "math/rand"
internally, and thus needs to be seeded before use. cards.Seed() and
cards.SeedInt(int64) can be used for this

#### type Card

```go
type Card struct {
	Suit  string
	Value string
}
```

Card is a struct holding the suit and value associated with a card.
