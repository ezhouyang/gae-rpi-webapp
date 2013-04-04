package analytics

import (
	"log"
	"math/rand"
	"testing"
	"time"
)

func TestUV(t *testing.T) {
	uv := NewUV(30)
	count := 1
	visitors := []string{"a", "b", "c", "d", "e", "f"}
	go func() {
		for {
			log.Println(count, uv.Sum())
			count++
			time.Sleep(1 * time.Second)
		}
	}()
	go func() {
		for i := 0; i < 15; i++ {
			cookie := visitors[rand.Intn(len(visitors))]
			uv.AddOne(cookie, time.Now())
			time.Sleep(1 * time.Second)
		}
	}()
	time.Sleep(2 * time.Minute)
	t.Log(uv)
}

func TestPV(t *testing.T) {
	pv := NewPV(2, 10) //count pv in 1 mins
	count := 1
	go func() {
		for {
			log.Println(count, pv.Sum(), pv.slots)
			count++
			if pv.Sum() > 20 {
				t.Error("Impossible to exceed 20 PVs")
			}
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for i := 0; i < 15; i++ {
			pv.AddOne()
			time.Sleep(1 * time.Second)
		}
		time.Sleep(20 * time.Second)
		for i := 0; i < 15; i++ {
			pv.AddOne()
			time.Sleep(1 * time.Second)
		}
		time.Sleep(41 * time.Second)
		for i := 0; i < 15; i++ {
			pv.AddOne()
			time.Sleep(1 * time.Second)
		}
	}()

	time.Sleep(2 * time.Minute)
	t.Log(pv)
}