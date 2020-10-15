package ChainOfResponsibility

import (
	"fmt"
	"strings"
)

type Handler interface {
	handle(text string) string
	setnext(Handler)
}

type FilterBlackHandler struct {
	next Handler
}

func (this *FilterBlackHandler) handle(text string) string {
	var result = strings.TrimSpace(text)
	if this.next == nil {
		return result
	}
	return this.next.handle(result)
}

func (this *FilterBlackHandler) setnext(next Handler) {
	this.next = next
}

type AddLogoHandler struct {
	next Handler
}

func (this *AddLogoHandler) handle(text string) string {
	var result = fmt.Sprintf("%s_logo", text)
	if this.next == nil {
		return result
	}
	return this.next.handle(result)

}

func (this *AddLogoHandler) setnext(next Handler) {
	this.next = next
}

type HandlerChain struct {
	head Handler
	tail Handler
}

func (this *HandlerChain) AddHandler(h Handler) {
	if this.head == nil {
		this.head = h
		this.tail = h
		return
	}

	this.tail.setnext(h)
	this.tail = h
}

func (this *HandlerChain) handle(text string) string {
	if this.head == nil {
		return ""
	}

	return this.head.handle(text)

}
