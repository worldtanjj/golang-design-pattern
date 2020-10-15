package iterator

type Iterator interface {
	haxNext() bool
	next()
	currentItem() interface{}
}

type Demo1Slice []interface{}

type Demo1Iterator struct {
	cursor  int
	arrlist Demo1Slice
}

func (this *Demo1Iterator) haxNext() bool {
	return this.cursor != len(this.arrlist) //注意这里，cursor在指向最后一个元素的时候，hasNext()仍旧返回true
}

func (this *Demo1Iterator) next() {
	this.cursor++
}
func (this *Demo1Iterator) currentItem() interface{} {
	if this.cursor >= len(this.arrlist) {
		panic("索引越界")
	}
	return this.arrlist[this.cursor]
}
