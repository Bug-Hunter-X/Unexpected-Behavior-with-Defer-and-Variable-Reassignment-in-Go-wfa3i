func main() {



  x := 10
  y := 5

  defer fmt.Println("x + y = ", x + y)

  x = 5

  fmt.Println("x + y = ", x + y)

}