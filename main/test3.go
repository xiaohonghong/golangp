package main

import (
	"fmt"
	"reflect"
)

func main(){
	user1 := user{"sss",1}
	toSpeak(user1)
}
func toSpeak(u interface{}){
	//u.(*user);
	peoper1 := reflect.ValueOf(u)

	//peoper1 := u.(*user)
	//peoper1.speak();
	fmt.Println("get Type is :", peoper1.FieldByName("name"))
	//ss := peoper1.MethodByName("speak")

	for i := 0; i < peoper1.NumField(); i++ {
		field := peoper1.Field(i)
		value := peoper1.Field(i).Interface()
		fmt.Printf("%s: %v = %v\n", field, field.Type, value)
	}

}
type peoper interface {
	speak()
	drink(a string)(b string)
}

type user struct {
	name string
	sex int
}
func (u *user)speak(){
	fmt.Println("%d说%d想喝",u.drink("ddd"))
}
func (u *user)drink(a string) (b string){
	return a;
}
