package aug312020

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

const fileHeader = `package aug312020

`

func generateSlice(n int) string {
	generated := make([]int, 0, n)
	for i := 0; i < n; i++ {
		generated = append(generated, rand.Intn(1200))
	}

	return fmt.Sprintf("%#v", generated)
}

func genFile() {
	file, err := os.OpenFile("benchSlices.go", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	sizes := []int{10, 100, 1000, 10000, 100000}

	var sb strings.Builder
	sb.Write([]byte(fileHeader))

	for _, v := range sizes {
		sb.WriteString(fmt.Sprintf("var slice%d = %s\n\n", v, generateSlice(v)))
	}

	written, err := file.Write([]byte(sb.String()))
	if err != nil {
		panic(err)
	}

	fmt.Printf("wrote %d bytes\n", written)
}
