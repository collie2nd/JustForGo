package main

import "testing"

func TestPrice_String(t *testing.T) {
	tests := []struct {
		name string
		i    Price
		want string
	}{
		{
			name: "Test case 1",
			i:    Price(100),
			want: "$100.00",
		},
		{
			name: "Test case 2",
			i:    Price(50),
			want: "$50.50",
		},
		{
			name: "Test case 3",
			i:    Price(0),
			want: "$0.00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
