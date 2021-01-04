package composition

//直接组合 接口组合 嵌入组合
//has a 代替 is a
//做到可替换方法和属性

type Athlete struct{}

func (a *Athlete) Train() {
	println("Training")
}

func Swim() {
	println("Swimming!")
}

type CompositeSwimmerA struct {
	MyAthlete Athlete
	MySwim    *func() //不使用指针类型也是可以的
}

//--------------------------------------------------------------------------

type Trainer interface {
	Train()
}
type Swimmer interface {
	Swim()
}

type SwimmerImplementor struct{}

func (s *SwimmerImplementor) Swim() {
	println("Swimming!")
}

type CompositeSwimmerB struct {
	Trainer
	Swimmer
}

//--------------------------------------------------------------------------

type Animal struct{}

func (r *Animal) Eat() {
	println("Eating")
}

type Shark struct {
	Animal
	Swim func()
}

//--------------------------------------------------------------------------

type Tree struct {
	LeafValue int
	Right     *Tree
	Left      *Tree
}

//--------------------------------------------------------------------------
//子类型不能当作父类型使用
type Parent struct {
	SomeField int
}

type Son struct {
	P Parent
}

func GetParentField(p Parent) int {
	return p.SomeField
}
