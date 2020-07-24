package card

import (
	"reflect"
	"testing"
	"time"
)

func TestSortBySum(t *testing.T) {
	testCardNew := &Card{
		Id:           0,
		Balance:      0,
		Currency:     "RUB",
		Number:       "000000",
		Transactions: []*Transaction{
			&Transaction{
				Id:      1,
				Sum:     99_00,
				Date:    time.Date(2020, time.June, 5, 11, 06, 00, 0, time.UTC).Unix(),
				MccCode: "5411",
				Status:  "Операция в обработке",
			},
			&Transaction{
				Id:      1,
				Sum:     735_55,
				Date:    time.Date(2020, time.June, 9, 22, 06, 23, 0, time.UTC).Unix(),
				MccCode: "5411",
				Status:  "Операция в обработке",
			},
			&Transaction{
				Id:      1,
				Sum:     1_000_00,
				Date:    time.Date(2020, time.June, 19, 21, 46, 11, 0, time.UTC).Unix(),
				MccCode: "5411",
				Status:  "Операция в обработке",
			},
		},
	}
	testCard := &Card{
		Id:           0,
		Balance:      0,
		Currency:     "RUB",
		Number:       "000000",
		Transactions: []*Transaction{
			&Transaction{
				Id:      1,
				Sum:     99_00,
				Date:    time.Date(2020, time.June, 5, 11, 06, 00, 0, time.UTC).Unix(),
				MccCode: "5411",
				Status:  "Операция в обработке",
			},
			&Transaction{
				Id:      1,
				Sum:     735_55,
				Date:    time.Date(2020, time.June, 9, 22, 06, 23, 0, time.UTC).Unix(),
				MccCode: "5411",
				Status:  "Операция в обработке",
			},
			&Transaction{
				Id:      1,
				Sum:     1_000_00,
				Date:    time.Date(2020, time.June, 19, 21, 46, 11, 0, time.UTC).Unix(),
				MccCode: "5411",
				Status:  "Операция в обработке",
			},
		},
	}

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
				testCard.Transactions,
			},
			want: testCardNew.Transactions,
		},
	}
	for _, tt := range tests {
		if got := SortBySum(tt.args.transactions); !reflect.DeepEqual(tt.args.transactions, tt.want) {
			t.Errorf("Sum() = %v, want %v", got, tt.want)
		}
	}
}