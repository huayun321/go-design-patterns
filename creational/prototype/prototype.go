package prototype

import (
	"errors"
	"fmt"
	"reflect"
)

const (
	White = iota
	Black
	Blue
)

type ShirtCloner interface {
	GetClone(s int) (ItemInfoGetter, error)
}

func GetShirtsCloner() ShirtCloner {
	return &ShirtsCache{}
}

type ShirtsCache struct{}

func (sc *ShirtsCache) GetClone(m int) (ItemInfoGetter, error) {
	switch m {
	case White:
		//指针地址 指针类型 指针取值
		//a := &T
		//a 的类型为 *T
		//a 取值 *a
		//
		//& new 操作符 返回的都是指针类型变量

		newItem := *whitePrototype //新建一个shirt 类型变量 并拷贝 whitePrototype *shirt 指针类型中的shirt值
		fmt.Printf("newItem value %#v\n", newItem)
		fmt.Printf("newItem type %T\n", newItem)
		fmt.Printf("newItem addr %p\n", &newItem)
		fmt.Printf("newItem typeof %#v\n", reflect.ValueOf(newItem))
		fmt.Printf("newItem valueof %#v\n", reflect.TypeOf(newItem))

		fmt.Printf("whitePrototype value %#v\n", whitePrototype)
		fmt.Printf("whitePrototype type %T\n", whitePrototype)
		fmt.Printf("whitePrototype addr %p\n", whitePrototype)
		fmt.Printf("*whitePrototype value %#v\n", *whitePrototype)
		fmt.Printf("*whitePrototype type %T\n", *whitePrototype)
		fmt.Printf("*whitePrototype addr %p\n", &(*whitePrototype))

		return &newItem, nil
	case Black:
		newItem := *blackPrototype
		return &newItem, nil
	case Blue:
		newItem := *bluePrototype
		return &newItem, nil
	default:
		return nil, errors.New("Shirt model not recognized")
	}
}

type ItemInfoGetter interface {
	GetInfo() string
}

type ShirtColor byte

func (sc ShirtColor) String() string {
	switch sc {
	case White:
		return "white shirt"
	case Black:
		return "black shirt"
	case Blue:
		return "blue shirt"
	default:
		return "shirt color not recongnized!"
	}
}

type Shirt struct {
	Price float32
	SKU   string
	Color ShirtColor
}

func (s *Shirt) GetInfo() string {
	return fmt.Sprintf("Shirt with SKU '%s' and Color id %d that costs %f\n", s.SKU, s.Color, s.Price)
}

func (s *Shirt) GetPrice() float32 {
	return s.Price
}

var whitePrototype *Shirt = &Shirt{
	Price: 15.00,
	SKU:   "empty",
	Color: White,
}

var blackPrototype *Shirt = &Shirt{
	Price: 16.00,
	SKU:   "empty",
	Color: Black,
}

var bluePrototype *Shirt = &Shirt{
	Price: 17.00,
	SKU:   "empty",
	Color: Blue,
}
