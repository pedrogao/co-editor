package ot

import "testing"

func TestApply(t *testing.T) {
	inital := "gos"

	opA := []Applier{Retain(2), Insert("a"), Delete(-1)}
	opT := []Applier{Retain(2), Insert("t"), Delete(-1)}

	strA, err := Apply(inital, opA...)
	if err != nil {
		t.Error(err)
	}
	strT, err := Apply(inital, opT...)
	if err != nil {
		t.Error(err)
	}

	if strA != "goa" || strT != "got" {
		t.Fail()
	}

	t.Run("when ops do not span string", func(t *testing.T) {
		_, err := Apply("hello")
		if err == nil {
			t.Error("it should error")
		}
	})
}

func TestInvert(t *testing.T) {
	inital := "hello"

	// ops := []Applier{Delete(-1), Insert("h"), Retain(4), Delete(-1), Retain(6), Delete(-1)}
	ops := []Applier{Insert("H"), Delete(-1), Retain(3), Delete(-1)}

	result, err := Apply(inital, ops...)
	if err != nil {
		t.Error(err)
	}
	inverted, err := Apply(result, Invert(inital, ops)...)
	if err != nil {
		t.Error(err)
	}

	if inital != inverted {
		t.Fail()
		t.Logf("%q %q", string(inital), string(inverted))
	}
}

func TestCompose(t *testing.T) {
	inital := "hello"
	ops1 := []Applier{Insert("H"), Retain(4), Delete(-1)}
	ops2 := []Applier{Delete(-1), Retain(4)}
	composed, err := Compose(ops1, ops2)
	if err != nil {
		t.Error("it should not err")
		t.Log(err)
		t.Log(ops1, ops2, composed)
	}
	decomposedResult, err := Apply(inital, ops1...)
	if err != nil {
		t.Error("it should not err")
		t.Log(err)
		t.Log(ops1, ops2, composed, decomposedResult)
	}
	decomposedResult, err = Apply(decomposedResult, ops2...)
	if err != nil {
		t.Error("it should not err")
		t.Log(err)
		t.Log(ops1, ops2, composed, decomposedResult)
	}
	composedResult, err := Apply(inital, composed...)
	if err != nil {
		t.Error("it should not err")
		t.Log(err)
		t.Log(ops1, ops2, composed, composedResult, decomposedResult)
	}
	if composedResult != decomposedResult {
		t.Fail()
		t.Log(ops1, ops2, composed, composedResult, decomposedResult)
	}

	t.Run("bad input", func(t *testing.T) {
		_, err := Compose([]Applier{Insert("9000")}, []Applier{Delete(-1)})
		if err == nil {
			t.Fail()
		}
	})
}

func TestTransform(t *testing.T) {
	inital := "gos"
	opA := []Applier{Retain(2), Insert("a"), Delete(-1)}
	opB := []Applier{Retain(2), Insert("t"), Delete(-1)}
	opAPrime, opBPrime, err := Transform(opA, opB)
	if err != nil {
		t.Error("it should not err")
		t.Log(err)
	}
	composedA, err := Compose(opA, opBPrime)
	if err != nil {
		t.Error("it should not err")
		t.Log(err)
		t.Logf("A: %v", opA)
		t.Logf("B': %v", opBPrime)
		t.Logf("Compose(A, B'): %v", composedA)
	}
	composedB, err := Compose(opB, opAPrime)
	if err != nil {
		t.Error("it should not err")
		t.Log(err)
		t.Logf("B: %v", opB)
		t.Logf("A': %v", opAPrime)
		t.Logf("Compose(A', B): %v", composedA)
	}
	resA, err := Apply(inital, composedA...)
	if err != nil {
		t.Error("it should not err")
		t.Log(err)
	}
	resB, err := Apply(inital, composedB...)
	if err != nil {
		t.Error("it should not err")
		t.Log(err)
	}
	if resA != resB {
		t.Error("it should return A' and B' such that Apply(Apply(S, A), B') = Apply(Apply(S, B), A')")
		t.Logf("A: %v", opA)
		t.Logf("A': %v", opAPrime)
		t.Logf("B: %v", opB)
		t.Logf("B': %v", opBPrime)
		t.Logf("results: %q, %q", resA, resB)
	}
}

// func TestTransform1(t *testing.T) {
// 	inital := ""
// 	opA := []Applier{Insert("111")}
// 	opB := []Applier{Insert("222")}
// 	opAPrime, opBPrime, err := Transform(opA, opB)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	composedA, err := Compose(opA, opBPrime)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	composedB, err := Compose(opB, opAPrime)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	resA, err := Apply(inital, composedA...)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	resB, err := Apply(inital, composedB...)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	if resA != "钢铁侠雷神" || resB != "钢铁侠雷神" {
// 		t.Fatal("it should return A' and B' such that Apply(Apply(S, A), B') = Apply(Apply(S, B), A')")
// 	}
// }
