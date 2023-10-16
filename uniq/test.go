package uniq

import (
	"fmt"
	"testing"
	"uniq"
)

func TestUniqDefault(t *testing.T) {


	{
		c: false,
		d: false,
		u: false,
		f: 0,
		c: 0,
		i: false,
	}
	input := []string{
		"111\n",
		"222\n",
		"222\n",
		"111\n",
	}

	expectedRes := []string{
		"111\n",
		"222\n",
		"111\n",
	}

	receivedRes := uniq.CollapseLines(input, opts)
	fmt.Println(expectedRes)
	fmt.Println(receivedRes)

	//fmt.Println(reflect.DeepEqual(expectedRes, receivedRes))

}
