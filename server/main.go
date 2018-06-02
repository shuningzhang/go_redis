
package main

import ("fmt"
    "net"
    "log"
    "os"
    "container/list"
    "github.com/go_redis/libs"
)


type RedisServer struct {
    Port int
}

type RedisCommand struct {
    Name string
    Arity int
    Type int
}

type RedisClient struct {
    fd int

}

var server RedisServer


func initServer(){

}


func main(){

    fmt.Print("Hello World\n")
    CmdList()
    libs.StorList()
}
