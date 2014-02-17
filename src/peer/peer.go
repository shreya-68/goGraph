
package peer

import (

    "fmt"
    "time"
    "os"
    "net"
    "strconv"
    )

//Client Node with name, nbr are the first hop neighbours and status is current running status
type Node struct {
    name   string
    nbr    []*net.TCPAddr
    status string
    addr   *net.TCPAddr
    //connections []*net.TCPConn

}

func checkErr(err error) {
    if err != nil {
        fmt.Printf("Fatal error: %s \n", err)
        os.Exit(1)
    }
}

func handleClient(conn net.Conn) {
    var buf [256]byte
    n, err := conn.Read(buf[0:])
    checkErr(err)
    fmt.Println("read")
    fmt.Println(string(buf[0:n]))
    conn.Close() 
}

func (node *Node) accept(listener *net.TCPListener) {
    for {
            conn, err := listener.Accept()
            if err != nil {
                continue
            }
            go handleClient(conn)
            //node.connections[0] = &net.TCPConn(conn)
        }
}

func (node *Node) listen() {
    listener, err := net.ListenTCP("tcp", node.addr)
    checkErr(err)
    go node.accept(listener)
}

func (node *Node) openTCPconn(rcvr *net.TCPAddr) {
    conn, err := net.DialTCP("tcp", nil, rcvr)
    checkErr(err)
    write := "Hi" + strconv.Itoa(rcvr.Port)
    fmt.Printf("Writing %s\n", write)
    _, err = conn.Write([]byte(write))
    checkErr(err)
}


func (node *Node) broadcast() {
    for _, nbr := range node.nbr {
        node.openTCPconn(nbr)
    }
}

func Client(port string, nbrs []string) {
    node := Node{name: port, status: "Init"}
    var err error
    node.addr, err = net.ResolveTCPAddr("tcp", port)
    checkErr(err)
    tcpAddrNbr := make([]*net.TCPAddr, len(nbrs))
    for i, val := range nbrs {
        addr, err := net.ResolveTCPAddr("tcp", val)
        checkErr(err)
        tcpAddrNbr[i] = addr
    }
    node.nbr = tcpAddrNbr 
    fmt.Printf("Hi my name is %s\n", node.name)
    //for _, val := range node.nbr {
    //    fmt.Printf("I am %s. And my neighbour is %d", node.name, val.Port)
    //}
    node.listen()
    time.Sleep(200*time.Millisecond) 
    node.broadcast()
}
