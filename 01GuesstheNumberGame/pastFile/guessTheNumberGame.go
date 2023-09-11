package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	//1.標準出力に最大、最小を書き込むように指示
	instruction1 := []byte("Please enter the maximum and minimum values of the random number.\n" +
		"'m n',for example.Firstly input maximum,then input space,finally input minimum number.\n")
	os.Stdout.Write(instruction1)
	os.Stdout.Sync()

	//2.標準入力を作成して出力された最大、最小を読み込み
	//3.標準入力から読み込んだ値をintに変換
	var max, min int
	_, err := fmt.Scanf("%d %d", &max, &min)
	if err != nil {
		fmt.Println("The input number is invalid.")
		return
	}
	//4.標準入力から読み込んだ値が最大＞最小となっていることを確認
	//5-2.最大＜最小ならば、エラーを返す。エラーも出力する。
	if max < min {
		err := fmt.Errorf("The maximum value is less than the minimum value.")
		fmt.Println(err.Error())
		return
	}
	if max == min {
		fmt.Println("The maximum and minimum values cannot be the same. Please enter different values.")
		return
	}
	//5-1.最大＞最小ならば、整数の乱数を生成

	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	randomNumberToBeAnswered := r.Intn(max-min+1) + min

	NumberOfTimesYouCanTry := (max - min + 1) / 2
	var countPoint int = (max-min+1)/2 + 1
	var inputNumber int

	for i := 0; i < NumberOfTimesYouCanTry; i++ {
		//6.生成された乱数の値を保持して、標準出力に”乱数の値を推測して”などと書き込み
		instruction2 := []byte("Please guess the value of the random number.\n" +
			"Please input the number you think made of the random generator.\n")
		os.Stdout.Write(instruction2)
		os.Stdout.Sync()

		// 7.乱数の値を推測していく。
		fmt.Scanf("%d", &inputNumber)
		if inputNumber == randomNumberToBeAnswered {
			// 7-1.数字を１つ入力して正解ならば、”正解です”と標準出力に書き込み、終了する。
			fmt.Println("Correct!")
			countPoint--
			score := countPoint / NumberOfTimesYouCanTry * 100
			scoreMessage := fmt.Sprintf("You got percentage of correct answers: %d %%", score)
			fmt.Println(scoreMessage)
			break
		} else if inputNumber > randomNumberToBeAnswered {
			fmt.Println("Incorrect! The number you input is bigger than the answer.")
			countPoint--
		} else {
			// 7-2.数字を一つ入力して不正解ならば、"不正解です。その数値より大きいです。"or"不正解です。その数値より小さいです"と標準出力に書き込み、7-1に戻る。
			fmt.Println("Incorrect! The number you input is smaller than the answer.")
			countPoint--
		}

		//推測できるまで繰り返す。

	}
}
