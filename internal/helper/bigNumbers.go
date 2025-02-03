package helper

import "strings"

const (

	Zero = `
███
█ █
█ █
█ █
███`

	One = `
 █ 
 █ 
 █ 
 █ 
 █ `

	Two = `
███
__█
███
█__
███`

	Three = `
███
  █
███
  █
███`

	Four = `
█ █
█ █
███
  █
  █`

	Five = `
███
█__
███
__█
███`

	Six = `
███
█__
███
█ █
███`

	Seven = `
███
  █
  █
  █
  █`

	Eight = `
███
█ █
███
█ █
███`

	Nine = `
███
█ █
███
__█
███`

	Colon = `
 █ 
 █ 
   
 █ 
 █ `

	// Small style for when terminal is too small
	SmallStyle = "%02d:%02d"

	// Minimum dimensions needed for big display
	MinWidth  = 40
	MinHeight = 15
)

func GetBigNumber(n int) []string {
	var NumStr string
	switch n {
	case 0:
		NumStr = Zero
	case 1:
		NumStr = One
	case 2:
		NumStr = Two
	case 3:
		NumStr = Three
	case 4:
		NumStr = Four
	case 5:
		NumStr = Five
	case 6:
		NumStr = Six
	case 7:
		NumStr = Seven
	case 8:
		NumStr = Eight
	case 9:
		NumStr = Nine
	default:
		return []string{}
	}

	return strings.Split(NumStr, "\n")[1:] // Skip the first empty line
}
