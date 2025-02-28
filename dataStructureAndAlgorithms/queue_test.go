package dataStructureAndAlgorithms

import (
	"reflect"
	"sync"
	"testing"
)

func TestNewQueue(t *testing.T) {
	type args struct {
		cap int
	}
	tests := []struct {
		name string
		args args
		want QueueInterface
	}{
		{
			name: "newQueue",
			args: args{
				cap: 10,
			},
			want: &Queue{
				Mutex: sync.Mutex{},
				len:   0,
				cap:   10,
				arr:   nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewQueue(tt.args.cap); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewQueue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueue_Poll(t *testing.T) {
	type fields struct {
		Mutex sync.Mutex
		len   int
		cap   int
		arr   []int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "getPoll",
			fields: fields{
				Mutex: sync.Mutex{},
				len:   0,
				cap:   10,
				arr:   nil,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Queue{
				Mutex: tt.fields.Mutex,
				len:   tt.fields.len,
				cap:   tt.fields.cap,
				arr:   tt.fields.arr,
			}
			if got := q.Poll(); got != tt.want {
				t.Errorf("Poll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueue_Push(t *testing.T) {
	type fields struct {
		Mutex sync.Mutex
		len   int
		cap   int
		arr   []int
	}
	type args struct {
		i int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "PushElem",
			fields: fields{
				Mutex: sync.Mutex{},
				len:   0,
				cap:   10,
				arr:   make([]int, 0, 10),
			},
			args: args{
				i: 10,
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := NewQueue(10)
			if got := q.Push(tt.args.i); got != tt.want {
				t.Errorf("Push() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueue_getCap(t *testing.T) {
	type fields struct {
		Mutex sync.Mutex
		len   int
		cap   int
		arr   []int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "getCap",
			fields: fields{
				Mutex: sync.Mutex{},
				len:   0,
				cap:   10,
				arr:   nil,
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Queue{
				Mutex: tt.fields.Mutex,
				len:   tt.fields.len,
				cap:   tt.fields.cap,
				arr:   tt.fields.arr,
			}
			if got := q.getCap(); got != tt.want {
				t.Errorf("getCap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueue_getLen(t *testing.T) {
	type fields struct {
		Mutex sync.Mutex
		len   int
		cap   int
		arr   []int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "getLen",
			fields: fields{
				Mutex: sync.Mutex{},
				len:   5,
				cap:   10,
				arr:   nil,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Queue{
				Mutex: tt.fields.Mutex,
				len:   tt.fields.len,
				cap:   tt.fields.cap,
				arr:   tt.fields.arr,
			}
			if got := q.getLen(); got != tt.want {
				t.Errorf("getLen() = %v, want %v", got, tt.want)
			}
		})
	}
}
