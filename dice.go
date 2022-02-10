package main

import (
	"math/rand"
	"math"
	"fmt"
	"time"
	"strings"
)

const (
	miss = "❌"
	criticalMiss = "💀"
	success = "➕"
	criticalSuccess = "'➕'"
	dieType = 10
)

type Dice struct {}

func transformDie(r int, h bool) string {
	if r == 0 {
		return criticalSuccess
	}

	if r > 0 && r < 6 {
		if h && r == 1 {
			return criticalMiss
		}

		return miss
	}

	return success
}

func (d *Dice) Hunger() string {
	rand.Seed(time.Now().UnixNano())

	if rand.Intn(dieType) <= 4 {
		return "success"
	} else {
		return "Fail"
	}
}

func (d *Dice) Roll(n int, h int) string { 
	rand.Seed(time.Now().UnixNano())

	npool := make([]string, 0, n) 
	hpool := make([]string, 0, h) 

	for i := 0; i < n; i++ {
		npool = append(npool, 
			transformDie(rand.Intn(dieType), false))
	}

	for i := 0; i < h; i++ {
		hpool = append(hpool, 
			transformDie(rand.Intn(dieType), true))
	}

	return addLabels(npool, hpool)
}

func addLabels(n []string, h []string) string {
	var	b strings.Builder //whattt
	criticals := 0
	successes := 0

	bestialFailure := false
	mess := false

	b.WriteString(strings.Join(n, "_"))
	b.WriteString(" [ ")
	b.WriteString(strings.Join(h, "_"))
	b.WriteString(" ] ")
	
	for i, v := range append(n, h...) {
		switch v {
		case criticalSuccess:
			criticals++
			successes++
			mess = (i > len(n))
		case success:
			successes++
		case criticalMiss:
			bestialFailure = true
		}
	}

	//pairs
	criticals = int(math.Floor(float64(criticals)/float64(2)))
	successes += criticals * 2


	fmt.Fprintf(&b, "\n%d successes", successes)
	if mess && criticals > 0 {
		b.WriteString("\nIt's a messy critical!")
	}

	if bestialFailure {
		b.WriteString("\nIt's may be a bestial failure!")
	}

	return b.String()
}


