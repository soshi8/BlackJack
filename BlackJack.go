package main

// ランダム
// https://takeshiyako.blogspot.com/2015/10/go-lang-rand.html
//
// スライスソート
// https://qiita.com/Sekky0905/items/2d5ccd6d076106e9d21c

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// カードをシャッフルさせるために使用
// 独自の構造体
type Card struct {
	idx       uint8
	intRandum int32
}
type allCard []Card

func main() {

	myCard := make([]uint8, 0)
	yourCard := make([]uint8, 0)

	//カードをシャッフル
	YamaCard := shufleCard()

	//自分、相手ともに２枚づつ取得
	myCard, YamaCard = getCard(myCard, YamaCard)
	yourCard, YamaCard = getCard(yourCard, YamaCard)
	myCard, YamaCard = getCard(myCard, YamaCard)
	yourCard, YamaCard = getCard(yourCard, YamaCard)

	fmt.Printf("自分のカード\n")
	fmt.Printf("1枚目は %s の %d\n", getMarkName(getMarkNo(myCard[0])), getNumber(myCard[0]))
	fmt.Printf("2枚目は %s の %d\n\n", getMarkName(getMarkNo(myCard[1])), getNumber(myCard[1]))
	fmt.Printf("相手のカード\n")
	fmt.Printf("1枚目は %s の %d\n", getMarkName(getMarkNo(yourCard[0])), getNumber(yourCard[0]))
	fmt.Printf("2枚目は %s の %d\n\n", getMarkName(getMarkNo(yourCard[1])), getNumber(yourCard[1]))

	var boolIget bool = false
	var boolYouget bool = false
	var ansMessage string = ""

	for !boolIget || !boolYouget {
		for !boolIget {
			ansMessage = ""
			fmt.Printf("1枚めくりますか？(Y/N)\n")
			fmt.Scanln(&ansMessage)
			switch ansMessage {
			case "y", "Y":
				myCard, YamaCard = getCard(myCard, YamaCard)
				fmt.Printf("自分のカード\n")
				for i := 0; i < len(myCard); i++ {
					fmt.Printf("%d 枚目は %s の %d\n", i+1, getMarkName(getMarkNo(myCard[i])), getNumber(myCard[i]))
				}
				fmt.Printf("自分のポイント = %d\n", getPoint(myCard))

			case "n", "N":
				fmt.Printf("自分のターンを終了します\n")
				boolIget = true

			default:
				ansMessage = ""
			}

			// 得点をカウント、バーストしたら終了
			if getPoint(myCard) > 21 {
				fmt.Println("Burst!!")
				boolIget = true
			}
		}

		for !boolYouget {
			if getPoint(yourCard) < 17 {
				yourCard, YamaCard = getCard(yourCard, YamaCard)
			} else {
				boolYouget = true
			}
		}
	}
	fmt.Println("")
	fmt.Println("＜最終結果＞")

	fmt.Println("自分のカード")
	for i := 0; i < len(myCard); i++ {
		fmt.Printf("%d 枚目は %s の %d\n", i+1, getMarkName(getMarkNo(myCard[i])), getNumber(myCard[i]))
	}
	fmt.Printf("自分のポイント = %d\n\n", getPoint(myCard))

	fmt.Println("相手のカード")
	for i := 0; i < len(yourCard); i++ {
		fmt.Printf("%d 枚目は %s の %d\n", i+1, getMarkName(getMarkNo(yourCard[i])), getNumber(yourCard[i]))
	}
	fmt.Printf("相手のポイント = %d\n\n", getPoint(yourCard))

	fmt.Printf("勝負は %s\n", resultGame(getPoint(myCard), getPoint(yourCard)))
}

/* カードの数字を取得 */
func getNumber(cardNo uint8) uint8 {
	ret := cardNo%13 + 1
	return ret
}

/* カードのマークを取得 */
// 0：ハート
// 1：ダイヤ
// 2:スペード
// 3:グラブ
func getMarkNo(cardNo uint8) uint8 {
	ret := cardNo / 13
	return ret
}

/* カードのマーク（名称）を取得 */
func getMarkName(markNo uint8) string {

	var ret string = ""
	switch markNo {
	case 0:
		ret = "ハート"
	case 1:
		ret = "ダイヤ"
	case 2:
		ret = "スペード"
	case 3:
		ret = "クラブ"
	}
	return ret
}

/* カードをシャッフルした山積みを取得 */
func shufleCard() []uint8 {

	wkyama := make([]Card, 52)
	var i uint8
	retyama := make([]uint8, 52)

	// idxには連番を（０～５１）を設定する
	// これをカードと見立てる。
	// intRandumにはランダムの数字を設定する。
	rand.Seed(time.Now().UnixNano())
	for i = 0; i < 52; i++ {
		wkyama[i].idx = i
		wkyama[i].intRandum = rand.Int31n(10000)
	}

	//ソート
	sort.Slice(wkyama, func(i, j int) bool { return wkyama[i].intRandum < wkyama[j].intRandum })
	for j := 0; j < 52; j++ {
		//	fmt.Printf("idx-%d  val-%d\n", j, wkyama[j].idx)
		retyama[j] = wkyama[j].idx
	}
	return retyama
}

/* 山積みからカードを１枚取得し手札に加える */
func getCard(tefuda []uint8, yama []uint8) ([]uint8, []uint8) {
	// 引いたあとは99としておく。
	for i := range yama {
		if yama[i] == 99 {
			continue
		} else {
			tefuda = append(tefuda, yama[i])
			yama[i] = 99
			break
		}
	}
	return tefuda, yama
}

/* 手札から得点を取得 */
func getPoint(tefuda []uint8) uint8 {
	var cntAce uint8 = 0
	var sumP uint8 = 0
	for i := 0; i < len(tefuda); i++ {

		switch {
		case getNumber(tefuda[i]) == 1:
			sumP += 11
			cntAce++
		case getNumber(tefuda[i]) >= 2 && getNumber(tefuda[i]) <= 10:
			sumP += getNumber(tefuda[i])
		case getNumber(tefuda[i]) >= 11:
			sumP += 10
		}
	}

	// エース独自処理。いったん11とカウントしておき、
	// 最大でエースの枚数分10を引けるようにする
	for i := 0; i < int(cntAce); i++ {
		if sumP > 21 {
			sumP = sumP - 10
		}
	}
	return sumP
}

/* 勝敗のリテラルを取得 */
func resultGame(myPoint uint8, YourPoint uint8) string {
	var ret string = ""
	switch {
	case myPoint > 21 && YourPoint > 21:
		ret = "両者バースト"
	case myPoint == YourPoint:
		ret = "引き分け"
	case YourPoint > 21:
		ret = "自分の勝ち"
	case myPoint > 21:
		ret = "自分の負け"
	case myPoint > YourPoint:
		ret = "自分の勝ち"
	default:
		ret = "自分の負け"

	}
	return ret
}
