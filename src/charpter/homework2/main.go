package main

//现在是猜数字的游戏
import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	secret := rand.Intn(100) + 1
	fmt.Println("secret is ", secret)
	var guess int
	var i int
	var guesscorrect bool
	fmt.Println("猜数字游戏开始！ 可猜的数字的1-100之间")
outloop:
	for i = 0; i < 10; i++ {
		fmt.Println("现在是第", i+1, "次猜测")
		fmt.Println("请输入你的猜测：")
		fmt.Scanln(&guess)
		if guess != secret {
			continue
		} else {
			guesscorrect = true
			switch {
			case i == 0:
				fmt.Println("you are a genius!")
				break outloop
			case i < 3:
				fmt.Println("you are good , will catch me soon!")
				break outloop
			case i < 9:
				fmt.Println("you are just soso! ")
				break outloop
			case i == 9:
				fmt.Println("lucky boy ")
				break outloop
			}

		}

	}
	if i == 9 && !guesscorrect {
		fmt.Println("你已经猜了10次,游戏结束!")
	}

}
