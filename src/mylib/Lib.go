package mylib

import (
	"math/rand"
	"time"
)

func StudentScore(slice []string) (studentMap map[string]int) {
	myRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	studentMap = make(map[string]int)
	for _, value := range slice {
		studentMap[value] = myRand.Intn(100)
	}
	return studentMap
}

func SoreStudentScore(scoreFrom map[string]int) (sortKey []string) {
	for name := range scoreFrom {
		sortKey = append(sortKey, name)
	}
	for i := 0; i < len(sortKey)-1; i++ {
		for j := 0; j < len(sortKey)-1-i; j++ {
			if scoreFrom[sortKey[j]] < scoreFrom[sortKey[j+1]] {
				sortKey[j], sortKey[j+1] = sortKey[j+1], sortKey[j]
			}
		}
	}
	return sortKey
}
