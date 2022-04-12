package converter

import (
	"context"
	"testing"
)

var testsTableToTransformation = []struct {
	From   interface{}
	To     Di
	TestId int
}{
	{
		TestId: 1,
		From: Di{
			DB: mockStruct{},
		},
		To: Di{
			DB: mockStruct{},
		},
	},
}

func TestAdapterTransformationFunc(t *testing.T) {

	func() {
		for _, testTableT := range testsTableToTransformation {

			got, err := To[Di](testTableT.From)

			if err != nil {
				t.Error("Failed: The convert interface in struct fail.")
			}
			if got != testTableT.To {
				t.Error("Failed: The value return is different of the expected.")
			}
			if got.DB == nil {
				t.Error("Failed: The DB value loaded should to have value.")
			}
			result, err := got.DB.GetItem(context.TODO(), "")
			if err != nil {
				t.Error("Failed: The err mock return not sholud has value.")
			}
			if result != (Bar{}) {
				t.Error("Failed: The result mock return not sholud has value.")
			}
		}
	}()

	func() {
		_, err := To[Di]("foo")

		if err == nil {
			t.Error("Failed: Is expect an err when pass a arg different the T.")
		}
	}()

}

type IfooDb interface {
	GetItem(ctx context.Context, customerId string) (Bar, error)
}

type Di struct {
	DB IfooDb
}

type mockStruct struct{}

func (c mockStruct) GetItem(ctx context.Context, customerId string) (Bar, error) {
	return Bar{}, nil
}

type Bar struct {
}
