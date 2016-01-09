package tour2

import (
	"fmt"
	"runtime"
	"time"
)

func Demo_basic_loop() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func Demo_while() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

func Demo_if(age int) {
	if age < 18 {
		fmt.Println("你未成年")
	} else {
		fmt.Println("你已成年")
	}
}

func Demo_if_short(fake_age, required_age int) {
	if real_age := fake_age - 1; real_age >= required_age {
		fmt.Println("你符合年龄要求")
	} else if real_age-1 >= required_age {
		fmt.Println("你还差一年就能符合年龄要求")
	} else {
		fmt.Println("您不符合年龄要求")
	}
}

func Demo_switch() {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}
}

func Demo_switch_evaluation() {
	fmt.Println("哪天是周六")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("今天")
	case today + 1:
		fmt.Println("明天")
	case today + 2:
		fmt.Println("2天之后")
	default:
		fmt.Println("好遥远")
	}
}

func Demo_switch_if_like() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("早上好")
	case t.Hour() < 17:
		fmt.Println("下午好")
	default:
		fmt.Println("晚上好")
	}
}

func Demo_defer_function() {
	fmt.Println("i`m up")
	defer fmt.Println("world")
	defer_test()
	fmt.Println("hello")
}

func defer_test() {
	defer fmt.Println("world~~~~~~~")
	fmt.Println("helloooooooo")
}

func Demo_stack_defer() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
