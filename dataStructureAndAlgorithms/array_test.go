package dataStructureAndAlgorithms

import (
	"fmt"
	"reflect"
	"testing"
)

func TestArray_Insert(t *testing.T) {
	type fields struct {
		len int
		cap int
		arr []int
	}
	type args struct {
		index []int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "",
			args: args{
				index: []int{1, 10},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := NewArray(10)
			if got := a.Insert(tt.args.index...); got != tt.want {
				t.Errorf("Insert() = %v, want %v", got, tt.want)
			}
			fmt.Println("a", a)
		})
	}

}

func TestArray_Lookup(t *testing.T) {
	type fields struct {
		len int
		cap int
		arr []int
	}
	type args struct {
		index int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "",
			fields: fields{
				len: 5,
				arr: nil,
			},
			args: args{
				index: 0,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Array{
				len: tt.fields.len,
				arr: tt.fields.arr,
			}
			if got := a.Lookup(tt.args.index); got != tt.want {
				t.Errorf("Lookup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArray_delete(t *testing.T) {
	type fields struct {
		len int
		cap int
		arr []int
	}
	type args struct {
		index int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "",
			fields: fields{
				len: 5,
				cap: 0,
				arr: nil,
			},
			args: args{
				index: 0,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Array{
				len: tt.fields.len,
				arr: tt.fields.arr,
			}
			if got := a.delete(tt.args.index); got != tt.want {
				t.Errorf("delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArray_update(t *testing.T) {
	type fields struct {
		len int
		cap int
		arr []int
	}
	type args struct {
		index int
		value int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Array{
				len: tt.fields.len,
				arr: tt.fields.arr,
			}
			if got := a.update(tt.args.index, tt.args.value); got != tt.want {
				t.Errorf("update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewArray(t *testing.T) {
	type args struct {
		cap int
	}
	tests := []struct {
		name string
		args args
		want ArrayInterface
	}{
		{
			name: "success",
			args: args{
				cap: 10,
			},
			want: &Array{
				len: 0,
				arr: make([]int, 10),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewArray(tt.args.cap); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
