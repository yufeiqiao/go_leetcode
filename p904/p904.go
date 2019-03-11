package main

const num_of_baskets = 2
const empty_flag = -1

/*
 * Sliding Window Algorithm
 * recording the sequence fruits was picked, 2 types currently in baskets, and max length to return
 */
type slidingBaskets struct {
	fruits     []int
	fruitTypes [num_of_baskets]int
	maxLength  int
}

func newBasket() slidingBaskets {

	types := [num_of_baskets]int{}
	for i := range types {
		types[i] = empty_flag
	}

	return slidingBaskets{
		fruits:     []int{},
		fruitTypes: types,
		maxLength:  0,
	}
}

/*
 * Check if fruitType is currently in any basket
 */
func (B *slidingBaskets) had(f int) bool {
	for _, t := range B.fruitTypes {
		if t == f {
			return true
		}
	}
	return false
}

/*
 * check if there is only one type of fruits
 */
func (B *slidingBaskets) hasEmptyBasket() bool {
	return B.had(empty_flag)
}

/*
 * When newBasket type is added, the type to empty is the one that is not fruits[size-1]
 * return nil if error
 */
func (B *slidingBaskets) typeToEmpty() *int {
	for i, t := range B.fruitTypes {
		if t != B.fruits[len(B.fruits)-1] {
			return &B.fruitTypes[i]
		}
	}
	return nil
}

/*
 * when a newBasket type of fruit is coming, empty one basket
 */
func (B *slidingBaskets) emptyOneBasket() {
	// figure out the last digit of the remaining sequence
	t := B.typeToEmpty()
	// faster, counting from the back
	for i := len(B.fruits) - 1; i >= 0; i-- {
		if *t == B.fruits[i] {
			// clear one type, now only one type in fruits
			*t = empty_flag

			B.fruits = B.fruits[i+1:]
			return
		}
	}
}

/*
 * adding a fruit to the sequence
 */
func (B *slidingBaskets) add(fruit int) {
	B.fruits = append(B.fruits, fruit)

	if len(B.fruits) > B.maxLength {
		B.maxLength = len(B.fruits)
	}
}

func (B *slidingBaskets) addTypeOf(fruit int) {
	for i, t := range B.fruitTypes {
		if t == empty_flag {
			B.fruitTypes[i] = fruit
			return
		}
	}
}

func TotalFruit(tree []int) int {
	B := newBasket()

	for _, fruit := range tree {
		if !B.had(fruit) {
			if !B.hasEmptyBasket() {
				B.emptyOneBasket()
			}
			B.addTypeOf(fruit)
		}
		B.add(fruit)
	}

	return B.maxLength
}

func main() {

}
