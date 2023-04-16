package main

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"log"
	"strings"
)

func check(e *casbin.Enforcer, sub, obj, act string) {
	ok, _ := e.Enforce(sub, obj, act)
	if ok {
		fmt.Printf("%s CAN %s %s\n", sub, act, obj)
	} else {
		fmt.Printf("%s CANNOT %s %s\n", sub, act, obj)
	}
}

func KeyMatcher(key1, key2 string) bool {
	i := strings.Index(key2, "*")
	if i == -1 {
		return key1 == key2
	}
	if len(key1) > 1 {
		return key1[:i] == key2[:i]
	}
	return key1 == key2[:i]
}

func KeyMatcherFunc(args ...interface{}) (interface{}, error) {
	name1 := args[0].(string)
	name2 := args[1].(string)
	return bool(KeyMatcher(name1, name2)), nil
}

func main() {
	e, err := casbin.NewEnforcer("./model.conf", "./policy.csv")
	if err != nil {
		log.Fatalf("NewEnforecer failed:%v\n", err)
	}
	e.AddFunction("my_func", KeyMatcherFunc)

	check(e, "tomtao626", "data1", "read")
	check(e, "tomtao627", "data2", "write")
	check(e, "tomtao626", "data1", "write")
	check(e, "tomtao626", "data2", "read")
}
