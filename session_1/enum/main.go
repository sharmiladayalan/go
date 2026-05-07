/*
▶ WHAT IS ENUM?
- Enum = a group of fixed named values
- Used when a variable should have limited options
- Example: Days (Sunday, Monday, Tuesday...)

----------------------------------------
▶ HOW TO USE ENUM IN GO?
Go does NOT have a built-in enum keyword.

We create enum using:
1. type  → define custom type
2. const → define fixed values
3. iota  → auto-increment numbers

Example:
type Day int

const (
    Sunday Day = iota + 1
    Monday
    Tuesday
)

----------------------------------------

▶ WHY TO USE ENUM?
- Makes code easy to read
- Avoids using random numbers
- Improves maintainability
- Represents fixed set of values clearly

Bad:
var day int = 2

Good:
var day Day = Monday

----------------------------------------

▶ ABOUT IOTA
- iota is a special keyword in Go
- Works like a counter
- Starts from 0
- Increases by 1 automatically

Example:
const (
    A = iota  // 0
    B         // 1
    C         // 2
)

With offset:
const (
    Sunday = iota + 1  // 1
    Monday             // 2
)

----------------------------------------

▶ IMPORTANT NOTE
- Enums are just integers internally
- Go does NOT strictly restrict values

Example:
var d Day = 100  // Allowed (but not recommended)

*/

package main

import "fmt"

type Day int

const (
	Sunday Day = iota + 1
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {
	fmt.Println(Sunday)
	fmt.Println(Monday)
	fmt.Println(Tuesday)
	fmt.Println(Wednesday)
	fmt.Println(Thursday)
	fmt.Println(Friday)
	fmt.Println(Saturday)
}