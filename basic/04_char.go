// ascii2 码： a -- 97
//            A -- 65
package main

import "fmt"

func main() {
	var ch byte
	ch = 97
	fmt.Println("ch = ", ch)
	// %c 以字符方式打印， %d 以整形方式打印
	fmt.Printf("%c , %d\n", ch, ch)

	ch = 'a'
	fmt.Printf("%c, %d\n", ch, ch)

	//大写转小写，小写转大写,大小写相差32，小写大
	fmt.Printf("大写：%d, 小写：%d\n", 'A', 'a')

	fmt.Printf("大写转小写： %c\n", 'A'+32)
	fmt.Printf("小写转大写：%c\n", 'a'-32)

	str2 := "mike"
	fmt.Printf("str2 type is %T\n", str2)

}
