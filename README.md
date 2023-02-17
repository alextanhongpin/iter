# iter [![Go Reference](https://pkg.go.dev/badge/github.com/alextanhongpin/iter.svg)](https://pkg.go.dev/github.com/alextanhongpin/iter)


Implements generic `map`, `reduce` and `filter` function in golang.

## Examples
```go
type UserDB struct {
	FullName string
	Age      int64
}

type UserAPI struct {
	FirstName string
	LastName  string
	Age       int64
}

func main() {
	users := []UserDB{{"John AppleSeed", 24}, {"John Doe", 20}}

	res0 := Map(users, func(u UserDB) UserAPI {
		names := strings.SplitN(u.FullName, " ", 2)

		var lastName string
		if len(names) > 1 {
			lastName = names[1]
		}

		return UserAPI{
			FirstName: names[0],
			LastName:  lastName,
			Age:       u.Age,
		}
	})
	fmt.Printf("%#v\n", res0)

	res1 := Filter(users, func(u UserDB) bool {
		return u.Age > 20
	})
	fmt.Printf("%#v\n", res1)

	res2 := Reduce(users, 0, func(u UserDB, totalAge int64, _ int) int64 {
		return totalAge + u.Age
	})
	fmt.Printf("%#v\n", res2)

	res3 := Reduce(users, "", func(u UserDB, names string, i int) string {
		if names == "" {
			return u.FullName
		}
		return names + ", " + u.FullName
	})
	fmt.Printf("%#v\n", res3)

	res4 := Map(users, func(u UserDB) string {
		return u.FullName
	})
	fmt.Printf("%#v\n", strings.Join(res4, ", "))
}
```
