package sublist

import "fmt"

// Relation is the type of relationship between two lists
type Relation string

var comparison Relation

const (
	equal    Relation = "equal"
	subset   Relation = "sublist"   //list1 is a subset of list2
	superset Relation = "superlist" //list1 is a superset of list2
	unequal  Relation = "unequal"
)

//Sublist returns the relationship between list1 and list2
func Sublist(list1 []int, list2 []int) Relation {
	comparison = unequal //list1 and list2 are unequal unless othewise is proven
	if len(list1) == len(list2) && isSubset(list1, list2) {
		comparison = equal //list1 and list2 are equal
	} else if len(list1) < len(list2) && isSubset(list1, list2) {
		comparison = subset //list1 is a subset of list2
	} else if len(list1) > len(list2) && isSubset(list2, list1) {
		comparison = superset //list1 is a superset of list2
	}
	var text = comparisonText(comparison)
	fmt.Println(text)
	return comparison
}

func isSubset(list1 []int, list2 []int) (list1IsSubset bool) {
	if len(list1) > len(list2) {
		list1IsSubset = false
		return
	}
	for i := 0; i <= len(list2)-len(list1); i++ {
		list1IsSubset = true
		list2slice := list2[i : len(list1)+i]
		for index, value := range list2slice {
			if value != list1[index] {
				list1IsSubset = false
				break
			}
		}
		if list1IsSubset {
			return
		}
	}
	return list1IsSubset
}

// == code written for only debugging purposes ==

//Compare is for calling the Sublist function from the main function
func Compare() {
	var list1 = []int{1, 2, 5}
	var list2 = []int{0, 1, 2, 3, 1, 2, 5, 6}
	Sublist(list1, list2)
}

//comparisonText is for printing to console for debugging
func comparisonText(comparison Relation) string {
	if comparison == equal {
		return "list1 is equal to list2"
	} else if comparison == subset {
		return "list1 is a subset of list2"
	} else if comparison == superset {
		return "list1 is a superset of list2"
	} else if comparison == unequal {
		return "list1 and list2 have no relationship"
	} else {
		return "case not considered"
	}
}
