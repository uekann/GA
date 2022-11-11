package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

const (
	GENOM_LENGTH        = 100
	INDIVIDUAL_NUMBER   = 50
	CHILD_NUMBER        = 50
	INDIVIDUAL_MUTATION = 0.05
	GENOM_MUTATION      = 0.01
)

type Genom []int

func (genom *Genom) Evaluate() int {
	// 遺伝子の評価関数(適応度)
	// 今回は遺伝子の合計値
	sum := 0
	for _, x := range *genom {
		sum += x
	}
	return sum
}

func InitGenom() *Genom {
	// 疑似乱数による遺伝子の初期化
	genom := make(Genom, GENOM_LENGTH)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < GENOM_LENGTH; i++ {
		genom[i] = rand.Intn(2)
	}
	return &genom
}

type GA []*Genom

func InitGA() *GA {
	// 世代の初期化
	ga := make(GA, INDIVIDUAL_NUMBER)
	for i := 0; i < INDIVIDUAL_NUMBER; i++ {
		ga[i] = InitGenom()
	}
	return &ga
}

func crossOver(parent []*Genom) (child []*Genom) {
	// 二点交叉
	// 親と同じ数の子を返す
	rand.Seed(time.Now().UnixMicro())
	for i := 0; i < len(parent)/2; i++ {
		g1 := *parent[2*i]
		g2 := *parent[2*i+1]
		p1 := rand.Intn(len(g1)-3) + 1
		p2 := rand.Intn(len(g1)-p1-2) + p1 + 1
		c1 := make(Genom, len(g1))
		c2 := make(Genom, len(g2))
		copy(c1, g1[:p1])
		copy(c2, g2[:p1])
		copy(c1[p1:], g2[p1:p2])
		copy(c2[p1:], g1[p1:p2])
		copy(c1[p2:], g1[p2:])
		copy(c2[p2:], g2[p2:])
		child = append(child, &c1, &c2)
	}
	if len(parent)/2 != 0 {
		child = append(child, parent[len(parent)-1])
	}
	return
}

func (ga *GA) mutation() {
	// 突然変異
	rand.Seed(time.Now().UnixMicro())
	for _, gene := range *ga {
		if rand.Intn(1000) <= INDIVIDUAL_MUTATION*10000 {
			for i := 0; i < len(*gene); i++ {
				if rand.Intn(1000) <= GENOM_MUTATION*1000 {
					(*gene)[i] = rand.Intn(2)
				}
			}
		}
	}
}

func (ga *GA) AlternateGene() {
	sort.Slice(*ga, func(i, j int) bool { return (*ga)[i].Evaluate() > (*ga)[j].Evaluate() })
	child := crossOver((*ga)[:CHILD_NUMBER])
	copy((*ga)[:CHILD_NUMBER], child)
	ga.mutation()
}

func main() {
	ga := InitGA()

	for i := 0; i < 50000; i++ {
		ga.AlternateGene()
	}

	fmt.Println(*((*ga)[0]))
}
