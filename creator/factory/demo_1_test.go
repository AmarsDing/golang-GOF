/*
 * @Author: your name
 * @Date: 2020-09-01 22:00:19
 * @LastEditTime: 2020-09-02 22:22:05
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \golang-GOF\creator\factory\demo_1_test.go
 */
package factory

import (
	"fmt"
	"testing"
)

// 普通创建
func TestProduct_Create(t *testing.T) {
	product1 := &Product1{}
	product1.SetName("p1")
	fmt.Println(product1.GetName())

	product2 := &Product2{}
	product2.SetName("p2")
	fmt.Println(product2.GetName())
}

// 工厂创建

func TestFactory_Create(t *testing.T) {
	factory := ProductFactory{}
	product1 := factory.Create(p1)
	product1.SetName("p1-factory")
	fmt.Println(product1.GetName())

	product2 := factory.Create(p2)
	product2.SetName("p2-factory")
	fmt.Println(product1.GetName())
}
