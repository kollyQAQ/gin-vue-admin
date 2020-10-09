package spider

import (
	"fmt"
	"testing"
)

func TestGetSku(t *testing.T) {
	//skuList, err := GetJdSku("111")
	skuList, err := GetJdSku("111,67716911235,52044853364")

	if err != nil {
		panic(err)
	}

	for _, sku := range skuList {
		fmt.Println(sku)
	}
}
