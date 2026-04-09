package main
import "fmt"
func main() {
    var M, A, B int
    fmt.Scan(&M)
    fmt.Scan(&A)
    fmt.Scan(&B)
    filho := M - (A + B)
    if filho > A && filho > B {
    fmt.Println(filho)
}
if A > filho && A > B {
fmt.Println(A)
}
if B > filho && B > A {
    fmt.Println(B)
}
}