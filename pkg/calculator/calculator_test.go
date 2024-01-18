package calculator

import (
	"testing"

	"repartners/pkg/calculator/mocks"
	"repartners/pkg/tests"

	"go.uber.org/mock/gomock"
)

func TestPacksCalculator_Calculate(t *testing.T) {
	t.Parallel()

	type fields struct {
		packs []int
	}
	type args struct {
		order int
	}
	tc := []struct {
		name   string
		fields fields
		args   args
		want   []int
	}{
		{
			name: "No values",
			args: args{},
			fields: fields{
				packs: []int{},
			},
			want: []int{},
		},
		{
			name: "Optimistic case - exact value - highest",
			args: args{
				order: 100,
			},
			fields: fields{
				packs: []int{20, 100},
			},
			want: []int{100},
		},
		{
			name: "Optimistic case - exact value - lowest",
			args: args{
				order: 100,
			},
			fields: fields{
				packs: []int{20, 200},
			},
			want: []int{20, 20, 20, 20, 20},
		},
		{
			name: "Optimistic case - exact value - mixed",
			args: args{
				order: 123,
			},
			fields: fields{
				packs: []int{20, 3, 100},
			},
			want: []int{20, 100, 3},
		},
		{
			name: "No exact value - v1",
			args: args{
				order: 10,
			},
			fields: fields{
				packs: []int{3, 8},
			},
			want: []int{8, 3},
		},
		{
			name: "No exact value - v2",
			args: args{
				order: 11,
			},
			fields: fields{
				packs: []int{2, 4, 5},
			},
			want: []int{5, 4, 2},
		},
		{
			name: "No exact value - v3",
			args: args{
				order: 13,
			},
			fields: fields{
				packs: []int{3, 8},
			},
			want: []int{8, 3, 3},
		},
		{
			name: "No exact value - v4",
			args: args{
				order: 21,
			},
			fields: fields{
				packs: []int{6, 5},
			},
			want: []int{6, 5, 5, 5},
		},
	}
	for _, tt := range tc {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			provider := mocks.NewMockPacksGetAller(ctrl)
			provider.EXPECT().GetAll().Times(1).Return(tt.fields.packs)

			packsCalc := New(provider)
			got := packsCalc.Calculate(tt.args.order)

			if !tests.CompareUnorderedSlice(tt.want, got) {
				t.Errorf("PacksCalculator.Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
