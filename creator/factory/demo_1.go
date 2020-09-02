/*
 * @Author: your name
 * @Date: 2020-09-01 21:14:58
 * @LastEditTime: 2020-09-01 21:59:52
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

type product1 struct {
	name string
}

func (p1 *product1) SetName(name string) {
	p1.name = name
}

func (p1 *product1) GetName() string {
	return p1.name
}

// 产品实例二

type product2 struct {
	name string
}

func (p2 *product2) SetName(name string) {
	p2.name = name
}

func (p2 *product2) GetName() string {
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

func (pf ProductFactory) Create(productType ProductType) Product {
	switch productType {
	case p1:
		return &product1{}
	case p2:
		return &product2{}
	default:
		return nil
	}
}
