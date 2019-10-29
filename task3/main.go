package main

import "fmt"

func arrayCommonPrefix(arr []string) string{
	prefix := arr[0]

	for i := 1; i < len(arr); i++ {
		prefix = commonPrefix(prefix, arr[i])
		if prefix == "" {
			return ""
		}
	}

	return prefix
}

func commonPrefix(word1, word2 string) string {
	var short,long string
	if len(word1) > len(word2){
		long = word1
		short = word2
	}else{
		long = word2
		short = word1
	}

	for i := len(short); i > 0; i -- {
		if long[0:i] == short[0:i]{
			return long[0:i]
		}
	}
	return ""
}

func main()  {
	test1:= []string{"qwe","qwesd", "qwsdfsdf", "tr"}
	test2:= []string{"qwe","qwesd", "qwsdfsdf"}
	test3:= []string{"qwe","qeasd", "qwsdfsdf"}
	fmt.Println(arrayCommonPrefix(test1))
	fmt.Println(arrayCommonPrefix(test2))
	fmt.Println(arrayCommonPrefix(test3))
}