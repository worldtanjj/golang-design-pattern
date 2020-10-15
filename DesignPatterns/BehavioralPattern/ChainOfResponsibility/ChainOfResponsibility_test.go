package ChainOfResponsibility

import (
	"fmt"
	"testing"
)

func TestChainOfResponsibility(t *testing.T) {
	var chain = &HandlerChain{}
	chain.AddHandler(&FilterBlackHandler{})
	chain.AddHandler(&AddLogoHandler{})

	var rt = chain.handle("  aaaa  bb c ")
	fmt.Println(rt)
}
