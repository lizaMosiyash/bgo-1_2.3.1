package card

import (
	"reflect"
	"testing"
)

func TestSortBySum(t *testing.T) {
	type args struct {
		transactions []*Transaction
	}
	tests := []struct {
		name string
		args args
		want []*Transaction
	}{
		{
			name: "test",
			args: args{
				transactions: []*Transaction{
					{
						Id:  1,
						Sum: 99_00,
					},
					{
						Id:  2,
						Sum: 735_55,
					},
					{
						Id:  3,
						Sum: 1_000_00,
					},
				},
			},
			want: []*Transaction{
				{
					Id:  3,
					Sum: 1_000_00,
				},
				{
					Id:  2,
					Sum: 735_55,
				},
				{
					Id:  1,
					Sum: 99_00,
				},
			},
		},
	}
	for _, tt := range tests {
		if got := SortBySum(tt.args.transactions); !reflect.DeepEqual(tt.args.transactions, tt.want) {
			t.Errorf("Sum() = %v, want %v", got, tt.want)
		}
	}
}
