package models

import (
	"math/rand"
	"moodly/utils"
	"time"
)

type Emotion struct {
	Feeling string
}

func GetRandomFeeling(category string) Emotion {
	var em Emotion
	rand.Seed(time.Now().UnixNano())

	emotions := GetEmotionsByCategory(category)

	r := rand.Intn(len(emotions))

	em.Feeling = "feeling " + emotions[r]
	return em

}

func GetEmotionsByCategory(category string) []string {
	if category == "anger" {
		return utils.TransformFile("./wordlists/anger.txt")
	} else if category == "fear" {
		return utils.TransformFile("./wordlists/fear.txt")
	} else if category == "happiness" {
		return utils.TransformFile("./wordlists/happiness.txt")
	} else if category == "sadness" {
		return utils.TransformFile("./wordlists/sadness.txt")
	} else {
		return nil
	}
}
