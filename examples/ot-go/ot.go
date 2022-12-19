package ot

import "errors"

// Example:
// ln := len(str)
// opLn := 0
// for _, op := range ops {
//   switch o := op.(type) {
//   case Retain:
//     opLn += int(o)
//   case Delete:
//     opLn += len(o)
//   case Insert:
//     opLn += 0
//   }
// }
// if ln != opLn {
//   return str, errors.New("operations could not be applied to input")
// }

// Basic Operation
type (
	Retain int
	Delete int
	Insert string
)

type Applier interface {
	Apply(cursor int, str []rune) (int, []rune)
	baseLength() int   // 操作之前字符串的长度
	targetLength() int // 操作完成字符串的长度
}

func (op Retain) Apply(cursor int, str []rune) (int, []rune) {
	return cursor + int(op), str
}

func (op Retain) baseLength() int {
	return int(op)
}

func (op Retain) targetLength() int {
	return int(op)
}

func (op Delete) Apply(cursor int, str []rune) (int, []rune) {
	var res []rune
	res = append(res, str[:cursor]...)
	res = append(res, str[cursor-int(op):]...)
	return cursor, res
}

func (op Delete) baseLength() int {
	return -int(op)
}

func (op Delete) targetLength() int {
	return 0
}

func (op Insert) Apply(cursor int, str []rune) (int, []rune) {
	var res []rune
	res = append(res, str[:cursor]...)
	res = append(res, []rune(op)...)
	res = append(res, str[cursor:]...)
	return cursor + len(op), res
}

func (op Insert) baseLength() int {
	return 0
}

func (op Insert) targetLength() int {
	return len(op)
}

// Actions
func Apply(str string, ops ...Applier) (string, error) {
	var (
		cursor int
		result = []rune(str)
	)

	if len(result) != baseLength(ops) {
		return "", errors.New("can't apply operations to string")
	}

	for _, op := range ops {
		cursor, result = op.Apply(cursor, result)
	}
	return string(result), nil
}

// Invert operation
// eg: Insert => Delete
func Invert(str string, ops []Applier) []Applier {
	buf := []rune(str)
	cursor := 0

	var inverse []Applier
	for _, op := range ops {
		switch o := op.(type) {
		case Retain:
			inverse = append(inverse, o)
			cursor += int(o)
		case Delete:
			inverse = append(inverse, Insert(string(buf[cursor:cursor+(-int(o))])))
			cursor -= int(o)
		case Insert:
			inverse = append(inverse, Delete(-len(o)))
		}
	}
	return inverse
}

// Compose merges two consecutive operation lists into one
// such that: Apply(Apply(S, listA...), listB...)) == Apply(S, Compose(listA, listB)...)
func Compose(opListA, opListB []Applier) ([]Applier, error) {
	if targetLength(opListA) != baseLength(opListB) {
		return nil, errors.New("the base length of the second operation must be the target length of the first operation")
	}

	var (
		composed []Applier // 合并后的operation

		opA            = opListA[0]
		opB            = opListB[0]
		aIndex, bIndex = 1, 1
	)

	opListA = append(opListA, nil)
	opListB = append(opListB, nil)

	for {
		if opA == nil && opB == nil {
			break
		}
		if op, ok := opA.(Delete); ok {
			composed = append(composed, op)
			opA = opListA[aIndex]
			aIndex++
			continue
		}
		if op, ok := opB.(Insert); ok {
			composed = append(composed, op)
			opB = opListB[bIndex]
			bIndex++
			continue
		}
		if len(opListA) == 1 {
			return nil, errors.New("cannot compose: the first list is too short")
		}
		if len(opListB) == 1 {
			return nil, errors.New("cannot compose: the second list is too long")
		}
		switch a := opA.(type) {
		case Retain:
			switch b := opB.(type) {
			case Retain:
				if a > b {
					composed = append(composed, Retain(b))
					opA = Retain(a - b)
					opB = opListB[bIndex]
					bIndex++
				} else if a < b {
					composed = append(composed, Retain(a))
					opB = Retain(b - a)
					opA = opListA[aIndex]
					aIndex++
				} else { // a == b
					composed = append(composed, Retain(a))
					opA = opListA[aIndex]
					aIndex++
					opB = opListB[bIndex]
					bIndex++
				}
			case Delete:
				if int(a) > -int(b) {
					composed = append(composed, Delete(b))
					opA = Retain(int(a) + int(b))
					opB = opListB[bIndex]
					bIndex++
				} else if int(a) < -int(b) {
					composed = append(composed, Delete(a))
					opB = Retain(int(b) + int(a))
					opA = opListA[aIndex]
					aIndex++
				} else {
					composed = append(composed, Delete(b))
					opA = opListA[aIndex]
					aIndex++
					opB = opListB[bIndex]
					bIndex++
				}
			}
		case Insert:
			switch b := opB.(type) {
			case Retain:
				if len(a) > int(b) {
					composed = append(composed, Insert(a[:b]))
					opA = Insert(a[:b])
					opB = opListB[bIndex]
					bIndex++
				} else if len(a) < int(b) {
					composed = append(composed, Insert(a))
					opB = Retain(int(b) - len(a))
					opA = opListA[aIndex]
					aIndex++
				} else { // a == b
					composed = append(composed, Insert(a))
					opA = opListA[aIndex]
					aIndex++
					opB = opListB[bIndex]
					bIndex++
				}
			case Delete:
				if len(a) > -int(b) {
					opA = Insert(a[:-b])
					opB = opListB[bIndex]
					bIndex++
				} else if len(a) < -int(b) {
					opB = Delete(int(b) + len(a))
					opA = opListA[aIndex]
					aIndex++
				} else {
					opA = opListA[aIndex]
					aIndex++
					opB = opListB[bIndex]
					bIndex++
				}
			}
		}
	}
	return composed, nil
}

// Transform A => A'
func Transform(opListA, opListB []Applier) ([]Applier, []Applier, error) {
	if baseLength(opListA) != baseLength(opListB) {
		return nil, nil, errors.New("the base lengths of the operations must be equal")
	}
	opListA = append(opListA, nil)
	opListB = append(opListB, nil)
	var (
		opAPrime       []Applier
		opBPrime       []Applier
		opA            = opListA[0]
		opB            = opListB[0]
		aIndex, bIndex = 1, 1
	)
	for {
		if opA == nil && opB == nil {
			break
		}
		if op, ok := opA.(Insert); ok {
			opAPrime = append(opAPrime, op)                      // A insert
			opBPrime = append(opBPrime, Retain(len([]rune(op)))) // 那么 B retain
			opA = opListA[aIndex]
			aIndex++
			continue
		}
		if op, ok := opB.(Insert); ok {
			opAPrime = append(opAPrime, Retain(len([]rune(op)))) // B insert
			opBPrime = append(opBPrime, opB)                     // 那么 A retain
			opB = opListB[bIndex]
			bIndex++
			continue
		}
		if len(opListA) == 1 {
			return nil, nil, errors.New("cannot compose: the first list is too short")
		}
		if len(opListB) == 1 {
			return nil, nil, errors.New("cannot compose: the second list is too long")
		}
		switch a := opA.(type) {
		case Retain:
			switch b := opB.(type) {
			case Retain:
				var minlen int
				if a > b {
					minlen = int(b)
					opA = Retain(a - b)
					opB = opListB[bIndex]
					bIndex++
				} else if a < b {
					minlen = int(a)
					opB = Retain(b - a)
					opA = opListA[aIndex]
					aIndex++
				} else { // a == b
					minlen = int(b)
					opA = opListA[aIndex]
					aIndex++
					opB = opListB[bIndex]
					bIndex++
				}
				opAPrime = append(opAPrime, Retain(minlen))
				opBPrime = append(opBPrime, Retain(minlen))
			case Delete:
				var minlen int
				if int(a) > -int(b) {
					minlen = -int(b)
					opA = Retain(int(a) + int(b))
					opB = opListB[bIndex]
					bIndex++
				} else if int(a) < -int(b) {
					minlen = int(a)
					opB = Delete(int(b) + int(a))
					opA = opListA[aIndex]
					aIndex++
				} else { // a == b
					minlen = int(a)
					opA = opListA[aIndex]
					aIndex++
					opB = opListB[bIndex]
					bIndex++
				}
				opBPrime = append(opBPrime, Retain(minlen))
			}
		case Delete:
			switch b := opB.(type) {
			case Delete:
				if -a > -b {
					opA = Delete(a - b)
					opB = opListB[bIndex]
					bIndex++
				} else if -a < -b {
					opB = Delete(b - a)
					opA = opListA[aIndex]
					aIndex++
				} else { // a == b
					opA = opListA[aIndex]
					aIndex++
					opB = opListB[bIndex]
					aIndex++
				}
			case Retain:
				var minlen int
				if int(a) > -int(b) {
					minlen = -int(b)
					opA = Retain(int(a) + int(b))
					opB = opListB[bIndex]
					bIndex++
				} else if int(a) < -int(b) {
					minlen = int(a)
					opB = Delete(int(b) + int(a))
					opA = opListA[aIndex]
					aIndex++
				} else { // a == b
					minlen = int(a)
					opA = opListA[aIndex]
					aIndex++
					opB = opListB[bIndex]
					bIndex++
				}
				opBPrime = append(opBPrime, Retain(minlen))
			}
		}
	}
	return opAPrime, opBPrime, nil
}

// priavate helper functions
func baseLength(ops []Applier) int {
	l := 0
	for _, op := range ops {
		l += op.baseLength()
	}
	return l
}

func targetLength(ops []Applier) int {
	l := 0
	for _, op := range ops {
		l += op.targetLength()
	}
	return l
}
