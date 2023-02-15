package entity
import {
	"testing"
	"https://github.com"
	"github.com/asaskevich/govalidator"
	
}
func TestAll (t.testing*T){
	g :=NewgomegaWithT(t){
		customer := Customer{
			Name: "", //ผิด
			CustomerID: "M1234567", // รหัสลูกค้าขึ8นต้นด้วย L หรือ M หรือ H แล้วตามด้วยตัวเลขจํานวน 7 ตัว

		}
		ok,err := govalidator.ValidateStruct(customer)

		g.Expect(ok).ToNot(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("Name cannot be blank"))
	}
}