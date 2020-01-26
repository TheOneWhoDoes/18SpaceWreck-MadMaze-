package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var target int
var arrvisit [1000][1000]int
var flag int

type Node struct {
	num   int
	color string
	head  []*Node
	child map[*Node]string
}

func (list *Node) addhead(head *Node) {
	list.head = append(list.head, head)
}

func (list *Node) addtail(tail *Node, color string) {
	list.child[tail] = color
}

func recursive(rocket *Node, lucky *Node) int {
	arrvisit[lucky.num][rocket.num] = 1
	if rocket.num == target || lucky.num == target {
		fmt.Println("seen")
		flag = 1
		return -2
	}

	for index, val := range rocket.child {

		if val == lucky.color && arrvisit[lucky.num][index.num] == 0 {
			a := recursive(index, lucky)
			if a == -2 {
				fmt.Println("Rocket ", rocket.num+1, "to ", index.num+1, " edge color: ", val, "//   current Lucky ", lucky.num+1, " color: ", lucky.color)
				return -2
			}
		}
	}

	for index, val := range lucky.child {
		if val == rocket.color && arrvisit[index.num][rocket.num] == 0 {
			a := recursive(rocket, index)
			if a == -2 {
				fmt.Println("Lucky ", lucky.num+1, "to ", index.num+1, " edge color: ", val, "//   current Rocky ", rocket.num+1, " color: ", rocket.color)
				return -2
			}
		}
	}

	return -1
}

func main() {

	file, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan() // internally, it advances token based on sperator
	nm := scanner.Text()
	nms := strings.Split(nm, " ")
	n, _ := strconv.Atoi(nms[0])
	m, _ := strconv.Atoi(nms[1])
	target = n - 1

	scanner.Scan()
	colors := scanner.Text()

	s := strings.Split(colors, " ")

	scanner.Scan() // internally, it advances token based on sperator
	ss := scanner.Text()
	sss := strings.Split(ss, " ")
	s1, _ := strconv.Atoi(sss[0])
	s2, _ := strconv.Atoi(sss[1])
	s1--
	s2--
	nodearray := make([]*Node, 0, 64)
	var ndStartR *Node
	var ndStartL *Node
	for i := 0; i < (n - 1); i++ {
		nd := new(Node)
		nd.color = s[i]
		nd.num = i
		nd.head = make([]*Node, 0, 0)
		nd.child = make(map[*Node]string)

		nodearray = append(nodearray, nd)
		if s1 == i {
			ndStartR = nd
		}
		if s2 == i {
			ndStartL = nd
		}
	}
	nd := new(Node)
	nd.color = " "
	nd.num = n - 1
	nd.head = make([]*Node, 0, 0)
	nd.child = make(map[*Node]string)
	nodearray = append(nodearray, nd)

	var edges []string
	for scanner.Scan() {
		edges = append(edges, scanner.Text())
	}

	for i := 0; i < m; i++ {
		p := strings.Split(edges[i], " ")
		i1, _ := strconv.Atoi(p[0])
		i1--
		i2, _ := strconv.Atoi(p[1])
		i2--
		ndhead := nodearray[i1]
		ndtail := nodearray[i2]
		ndhead.addtail(ndtail, p[2])
		ndtail.addhead(ndhead)
	}
	numf := recursive(ndStartR, ndStartL)
	if numf == -1 {
		fmt.Println("it can not be visited")
	}
}
