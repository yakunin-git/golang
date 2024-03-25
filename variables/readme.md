#### variables


```go
package main
import "fmt"

func main() {
	var a, b, c int     // Объявление переменных
	
	a = 2               // Инициализация переменных
	b = 2
	c = a + b

	fmt.Println(c)      // Вывод 
}
```

```go
package main
import "fmt"

func main() {
	a, b := 2, 2        // Объявление и инициализация переменных
	c := a + b          // Сумма значений
	fmt.Println(c)      // Вывод
}
```

```go
package main
import "fmt"

func main() {

	var (
		age    int
		weight float32
		name   string
	)

	age = 42
	weight = 90.4
	name = "Yakunin Vasily"

	fmt.Printf("Name: %s\nAge: %d\nWeight: %f", name, age, weight)

}

```
