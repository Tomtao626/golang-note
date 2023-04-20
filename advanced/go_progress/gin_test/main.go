package main

import (
	"github.com/DeanThompson/ginpprof"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	ginpprof.Wrap(r)
	r.Run()
}
