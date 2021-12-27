package optional

import (
	"testing"
)

func TestOptional_Value(t *testing.T) {
	intData := 1
	op := New(intData)
	intResData, ok := op.Value().(int)
	if !ok {
		t.Error("Not a int")
	}
	if intResData != 1 {
		t.Error("want a int 1, got ", intResData)
	}

	stringData := "TestString"
	op = New(stringData)
	stringResData, ok := op.Value().(string)
	if !ok {
		t.Error("Not a string")
	}
	if stringResData != "TestString" {
		t.Error("want a string \"TestString\", got ", stringResData)
	}
}

func TestOptional_ValueOr(t *testing.T) {
	var op Optional
	intResData, ok := op.ValueOr(1).(int)
	if !ok {
		t.Error("Not a int")
	}
	if intResData != 1 {
		t.Error("want a int 1, got ", intResData)
	}

	stringData := "TestString"
	op = Nil()
	stringResData, ok := op.ValueOr(stringData).(string)
	if !ok {
		t.Error("Not a string")
	}
	if stringResData != stringData {
		t.Error("want a string \"TestString\", got ", stringResData)
	}
}

func TestOptional_Has(t *testing.T) {
	intData := 1
	op := New(intData)
	if op.HasValue() != true {
		t.Error("op should be present")
	}

	op = Nil()
	if op.HasValue() != false {
		t.Error("op should be not present")
	}
}

func TestOptional_Equal(t *testing.T) {
	caseDataX := make(map[int]string)
	caseDataX[0] = "Com"
	caseDataX[1] = "Jerry"

	caseDataY := make(map[int]string)
	caseDataY[0] = "Anita"
	caseDataY[1] = "Robin"

	type Arg struct {
		X interface{}
		Y interface{}
	}

	tests := []struct {
		name string
		args Arg
		want bool
	}{
		{"TestOptional_Equal_case1", Arg{10, 10}, true},
		{"TestOptional_Equal_case2", Arg{10, 20}, false},
		{"TestOptional_Equal_case3", Arg{caseDataX, caseDataX}, true},
		{"TestOptional_Equal_case4", Arg{caseDataX, caseDataY}, false},
		{"TestOptional_Equal_case5", Arg{[]string{"a", "b", "c"}, []string{"a", "b", "c"}}, true},
		{"TestOptional_Equal_case6", Arg{[]string{"a", "b", "c"}, []string{"a", "m", "d"}}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			op := New(tt.args.X)
			if got := op.Equal(tt.args.Y); got != tt.want {
				t.Errorf("Optional.Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOptional_Assign(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want args
	}{
		{"TestOptional_Assign_case1", args{100}, args{100}},
		{"TestOptional_Assign_case2", args{"textString"}, args{"textString"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var op Optional
			got := op.Assign(tt.args.value)
			if !got.Equal(tt.want.value) {
				t.Errorf("Optional.Equal() = %v, want %v", got, tt.want.value)
			}
		})
	}
}
