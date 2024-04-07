package funcs

import (
	"context"

	"github.com/nevalang/neva/internal/runtime"
	"golang.org/x/exp/slices"
)

type listSortFloat struct{}

func (p listSortFloat) Create(io runtime.FuncIO, _ runtime.Msg) (func(ctx context.Context), error) {
	dataIn, err := io.In.Port("data")
	if err != nil {
		return nil, err
	}

	resOut, err := io.Out.Port("res")
	if err != nil {
		return nil, err
	}

	return func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				return
			case data := <-dataIn:
				select {
				case <-ctx.Done():
					return
				default:
					lst := data.List()
					arr := []float64{}

					for i := 0; i < len(lst); i++ {
						arr = append(arr, lst[i].Float())
					}
					slices.Sort(arr)

					finalArr := []runtime.Msg{}
					for i := 0; i < len(arr); i++ {
						finalArr = append(finalArr, runtime.NewFloatMsg(arr[i]))
					}
					resOut <- runtime.NewListMsg(finalArr...)
				}
			}
		}
	}, nil
}
