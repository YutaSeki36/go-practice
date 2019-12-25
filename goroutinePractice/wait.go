package goroutinePractice

import (
	"fmt"
	"time"
)

type WaitGo struct {
}

func (w WaitGo) Execu() {
	const totalExecuteNum = 6   // 合計実行数
	const maxConcurrencyNum = 3 // 同時実行数
	// バッファの最大長が3のチャンネル作成(3つまでデータ格納可能)
	sig := make(chan string, maxConcurrencyNum)
	// バッファの最大長が6のチャンネル作成(6つまでデータ格納可能)
	res := make(chan string, totalExecuteNum)
	defer close(sig)
	defer close(res)

	fmt.Println("start concurrency execute  %s", time.Now())
	for i := 0; i < totalExecuteNum; i++ {
		go wait6sec(sig, res, fmt.Sprintf("no%d", i))
	}
	// 無限ループで処理終了を待っている
	for {
		// 全部が終わるまで待つ
		// len(channel)で格納されているデータ数を取得できる
		// データが6つ入ったら次の処理へ
		if len(res) >= totalExecuteNum {
			break
		}
	}

	fmt.Println("end concurrency execute %s", time.Now())
}

func wait6sec(sig chan string, res chan string, name string) {
	// 最大長が3なので、3つデータが入った段階で次の処理をブロックする
	sig <- fmt.Sprintf("sig %s", name)
	time.Sleep(6 * time.Second)
	fmt.Println("%s:end wait 6sec", name)
	res <- fmt.Sprintf("sig %s", name)
	// データを送信することでバッファに空きができる(次の処理が実行できる)
	<-sig
}
