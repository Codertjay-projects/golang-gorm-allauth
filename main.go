package main

import "log"

type Test1 struct {
	id   string `json:"first_name"`
	name string `json:"first_names"`
}

func (t *Test1) Name() string {
	return t.name
}

type Test2 struct {
	*Test1
	id   string `json:"first_name"`
	//id string
}

func main() {
	/* Just using this to test configuration*/

	aa := Test2{
		id: "aa",
	}
	log.Panicln(aa)


}
