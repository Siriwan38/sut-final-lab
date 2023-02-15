package entity
import {
	"gorm.io/gorm"
	"github.com/asaskevich/govalidator"
}
type Customer struct {
	gorm.Model
	Name string 		`valid:"required~Name cannot be blank"`// ต้องไม่เป็นค่าว่าง 
	Email string
	CustomerID string 		`valid:"required~CustomerID cannot be blank"`// รหัสลูกค้าขึ8นต้นด้วย L หรือ M หรือ H แล้วตามด้วยตัวเลขจํานวน 7 ตัว
}
// func init(){
// 	govalidator.CustomerTypeTagMap ("Customer", (i interface{} context interface{})) bool{
// 		return float64.i>=0
// 	}
// }