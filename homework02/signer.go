package main

import (
	"fmt"
	"strconv"
)

func ExecutePipeline() {

}
func SingleHash(data string) string {
	return DataSignerCrc32(data) + "~" + DataSignerCrc32(DataSignerMd5(data))
}

func MultiHash(data string) (result string) {
	for th := 0; th < 6; th++ {
		result += DataSignerCrc32(strconv.Itoa(th) + data)
	}
	return
}

func CombineResults(hashes []string) (result string) {
	for n, hash := range hashes {
		result += hash
		if n < len(hashes)-1 {
			result += "_"
		}
	}
	return
}

func main() {
	hashes := []string{MultiHash(SingleHash("0")), MultiHash(SingleHash("1"))}
	fmt.Println(CombineResults(hashes))
}
