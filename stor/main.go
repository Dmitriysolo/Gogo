package main

import (
	"fmt"
	"github.com\zhashkevych\go-basics\06@v0.0.0-20210613092302-c14ff6043162\storage\storage"
)

func main() {
	ms := newMemoryStorage()
	ds := newDumbStorage()

	spawnEmployees(ms)
	fmt.Println(ms.get(3))

	spawnEmployees(ds)
}

func spawnEmployees(s Storage) {
	for i := 1; i <= 10; i++ {
		s.Insert(Employee{id: i})
	}
}
