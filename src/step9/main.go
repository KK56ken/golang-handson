// Step9
// Interfaceを利用して，ポリモーフィズムを実装してみましょう．
// 以下のような構造体を定義してください．
// strongMan : 強い人間
//  - money : 所持金(int)
//  - power : 強さ(int)
//  - battle() : 戦闘を行うメソッドです．強い人間は戦闘を行うと所持金が増えます．(強いので)
//  -          : 戦闘を行ったとき，powerを含む文字列を出力してください．
//             : ※所持金の増加量は任意です．intの最大値を超えないようにしてください．
//  - getMoney() : 所持金を取得するメソッドです．
//  - getPower() : 強さを取得するメソッドです．
// otaku :  オタクくん
//  - money : 所持金(int)
//  - anime : オタクくんの好きなアニメ(string)
//  - battle() : 戦闘を行うメソッドです．オタクくんは戦闘を行うと所持金が減ります．(オタクなので)
//  -          : 戦闘を行ったとき，animeを含む文字列を出力してください．
//             : ※所持金の減少量は任意です．intの最小値を下回らないようにしてください．
//  - getMoney() : 所持金を取得するメソッドです．
//  - getAnime() : オタクくんの好きなアニメを取得するメソッドです．

// これらの構造体に対して適切な human インターフェースを実装し，
// humanインターフェースを受け取り，battleメソッドを実行するfightメソッドを実装してください．
// func fight(h human)

package main

import "fmt"

type human interface {
	battle()
}

type strongMan struct {
	money int
	power int
}

func (m *strongMan) getMoney() int {
	return m.money
}

func (m *strongMan) getPower() int {
	return m.power
}

func (m *strongMan) battle() {
	m.money += m.power
	fmt.Println("My power is", m.power)
}

type otaku struct {
	money int
	anime string
}

func (o *otaku) getMoney() int {
	return o.money
}

func (o *otaku) getAnime() string {
	return o.anime
}

func (o *otaku) battle() {
	o.money -= 100
	fmt.Println("My favorite anime is", o.anime)
}

func fight(h human) {
	h.battle()
}

func main() {
	s := strongMan{
		money: 100,
		power: 100,
	}
	o := otaku{
		money: 100,
		anime: "One Piece",
	}
	fight(&s)
	fight(&o)
	fmt.Println(s.getMoney())
	fmt.Println(o.getMoney())
	fmt.Println(s.getPower())
	fmt.Println(o.getAnime())
}
