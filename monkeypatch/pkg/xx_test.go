package pkg

import (
	"testing"

	"reflect"

	"fmt"

	"github.com/bouk/monkey"
	. "github.com/smartystreets/goconvey/convey"
)

func TestOnet(t *testing.T) {

	Convey("a", t, func() {
		gurad := monkey.Patch(UsedForTestMoneyPatch, PatchFunc)

		defer gurad.Unpatch()

		UsedForTestMoneyPatch()

		a := &AAAA{}
		gurad1 := monkey.PatchInstanceMethod(reflect.TypeOf(a), "UseForPatch", func(*AAAA) {
			fmt.Println("instance patch")
		})

		defer gurad1.Unpatch()
		a.UseForPatch()
	})
}
