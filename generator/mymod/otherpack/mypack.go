package otherpack

//go:generate go run gen.go arg1 arg2

func PackFunc() string {
    return "other pack"
}
