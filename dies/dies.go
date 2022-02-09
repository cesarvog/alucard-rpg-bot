				package main

				import (
					"math/rand"
					"math"
					"fmt"
					"time"
					"strings"
				)

				const (
					Miss = "❌"
					Critical_miss = "💀"
					Success = "➕"
					Critical_success = "'➕'"
				)

				func mapResult(r int, h bool) string {
					if r == 0 {
						return Critical_success
					}

					if r > 0 && r < 6 {
						if h && r == 1 {
							return Critical_miss
						}

						return Miss
					}

					return Success
				}

				func Hunger() string {
					rand.Seed(time.Now().UnixNano())

					if rand.Intn(10) <= 4 {
						return "Success"
					} else {
						return "Fail"
					}
				}

				func Roll(n int, h int) string { 
					rand.Seed(time.Now().UnixNano())

					npool := make([]string, 0, n) 
					hpool := make([]string, 0, h) 

					for i := 0; i < n; i++ {
						npool = append(npool, 
							mapResult(rand.Intn(10), false))
					}

					for i := 0; i < h; i++ {
						hpool = append(hpool, 
							mapResult(rand.Intn(10), true))
					}

					return addLabels(npool, hpool)
				}

				func addLabels(n []string, h []string) string {
					var	b strings.Builder //whattt
					criticals := 0
					successes := 0

					bestialFailure := false
					mess := false

					fmt.Printf("%d AAAAA", len(n))
					b.WriteString(strings.Join(n, "_"))
					b.WriteString(" [ ")
					b.WriteString(strings.Join(h, "_"))
					b.WriteString(" ] ")
	
	for i, v := range append(n, h...) {
		switch v {
		case Critical_success:
			criticals++
			successes++
			mess = (i > len(n))
		case Success:
			successes++
		case Critical_miss:
			bestialFailure = true
		}
	}

	//pairs
	criticals = int(math.Floor(float64(criticals)/float64(2)))
	successes += criticals * 2


	fmt.Fprintf(&b, "\n%d Successess", successes)
	if mess && criticals > 0 {
		b.WriteString("\nIt's a messy critical!")
	}

	if bestialFailure {
		b.WriteString("\nIt's may be a bestial failure!")
	}

	return b.String()
}


