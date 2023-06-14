package game

import (
	"fmt"
	"math/rand"
	"time"
)

func generateCharactor() Charactor {

	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(3)

	return Charactors[n]
}

func Gacha() {
	fmt.Println("ガチャを回す。Enterを押してください。")
	fmt.Scanln()
	fmt.Println("ガチャン")
	time.Sleep(1500 * time.Millisecond)
	fmt.Println("ガラガラ...")
	time.Sleep(1500 * time.Millisecond)
	fmt.Println("ポン！")
	time.Sleep(1500 * time.Millisecond)

	charactor := generateCharactor()

	fmt.Println(charactor.Image)
	fmt.Println("なまえ : ", charactor.Name)
	fmt.Println("レアリティ : ", charactor.RaritytoStar())
}
