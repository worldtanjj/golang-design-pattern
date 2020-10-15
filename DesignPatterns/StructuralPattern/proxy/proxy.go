package proxy

type Subject interface {
	Do() string
}

type RealSubject struct{}

func (this RealSubject) Do() string {
	return "real"
}

type Proxy struct {
	real RealSubject
}

func (this Proxy) Do() string {
	var res string

	// 在调用真实对象之前的工作，检查缓存，判断权限，实例化真实对象等。。
	res += "pre:"

	res += this.real.Do()

	//调用后
	res += "after:"

	return res
}
