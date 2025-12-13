package fixtures

import "context"

type Mock struct{}

func (m Mock) EXPECT() Mock                                     { return m }
func (m Mock) PublishWeightedChargingForecast(a, b, c any) Mock { return m }
func (m Mock) DoAndReturn(fn any) Mock                          { return m }

type WeightedChargingForecast struct{}

var gomock = struct {
	Any func() any
	Eq  func(any) any
}{}

var charging = struct {
	WeightedChargingForecast *WeightedChargingForecast
}{}

func testFuncLitParams() {
	parallax := Mock{}

	// Chain with function literal that has long parameter list
	// The body is already expanded, so params should be split
	parallax.EXPECT().
		PublishWeightedChargingForecast(gomock.Any(), gomock.Eq("123"), gomock.Any()).
		DoAndReturn(func(ctx context.Context, target string, forecast *charging.WeightedChargingForecast) error {
			return nil
		})
}
