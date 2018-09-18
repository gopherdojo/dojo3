package main

import (
	"net/http/httptest"
	"testing"
	"time"
)

func TestHandler(t *testing.T) {
	t.Helper()
	dateSet := []time.Time{
		time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local),
		time.Date(2019, 1, 3, 23, 59, 59, 0, time.Local),
		time.Now(),
	}
	omikujis := make([]*Omikuji, 3)
	for i, v := range dateSet {
		omikujis[i] = NewOmikuji(v)
	}
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	N := 100
	daikichiCount := 0
	t.Run("1/1 test", func(t *testing.T) {
		for i := 0; i < N; i++ {
			omikujis[0].handle(w, r)
			if omikujis[0].Result == "大吉" {
				daikichiCount++
			}
		}
		if daikichiCount != N {
			t.Errorf("1/1 daikichiCount: %d\n", daikichiCount)
		}
		daikichiCount = 0
	})

	t.Run("1/3 test", func(t *testing.T) {
		for i := 0; i < N; i++ {
			omikujis[1].handle(w, r)
			if omikujis[1].Result == "大吉" {
				daikichiCount++
			}
		}
		if daikichiCount != N {
			t.Errorf("1/3 daikichiCount: %d\n", daikichiCount)
		}
		daikichiCount = 0
	})

	t.Run("normal day test", func(t *testing.T) {
		for i := 0; i < N; i++ {
			omikujis[2].handle(w, r)
			if omikujis[2].Result == "大吉" {
				daikichiCount++
			}
		}
		if daikichiCount == N {
			t.Errorf("daikichiCount: %d\n", daikichiCount)
		}
	})

}
