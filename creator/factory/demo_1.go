/*
 * @Author: your name
 * @Date: 2020-09-01 21:14:58
 * @LastEditTime: 2020-09-02 22:13:35
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \golang-GOF\creator\factory\demo_1.go
 */
package factory

// 实现一个抽象的物品

type Product interface {
	SetName(name string)
	GetName() string
}

// 产品实例一

type Product1 struct {
	name string
}

func (p1 *Product1) SetName(name string) {
	p1.name = name
}

func (p1 *Product1) GetName() string {
	return p1.name
}

// 产品实例二

type Product2 struct {
	name string
}

func (p2 *Product2) SetName(name string) {
	p2.name = name
}

func (p2 *Product2) GetName() string {
	return p2.name
}

type ProductType int

const (
	p1 ProductType = iota
	p2
)

// 实现工厂

type ProductFactory struct {
}

// Create 1
func (pf ProductFactory) Create(productType ProductType) Product {
	switch productType {
	case p1:
		return &Product1{}
	case p2:
		return &Product2{}
	default:
		return nil
	}
}
