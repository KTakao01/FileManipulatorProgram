// インターフェースによる抽象化、DIになれるための練習
package main

import (
	"fmt"
	"os"
)

//main,reverseへの切り分け

// 具体型の定義
// os.ReadFile,os.WriteFileのラッパー関数ReadFile,WriteFile
// osパッケージのfunc ReadFile(filename string) ([]byte, error)
// osパッケージのfunc WriteFile(filename string, data []byte, perm FileMode) error
// ReadFileとWriteFileをメソッドとしてもつ型の定義defaultFileOperations
type defaultFileOperations struct{}

func (d *defaultFileOperations) ReadFile(filename string) ([]byte, error) {
	return os.ReadFile(filename)
}

func (d *defaultFileOperations) WriteFile(filename string, data []byte, perm os.FileMode) error {
	return os.WriteFile(filename, data, perm)
}

//ReadFile,WriteFileの抽象化

type ReadFile interface {
	ReadFile(filename string) ([]byte, error)
}

type WriteFile interface {
	WriteFile(filename string, data []byte, perm os.FileMode) error
}

//インターフェースの導入//メソッドを持つ型の集まりを表す型

// コマンド名による動作の切り分け(共通)
func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: go run <file_name.go> reverse <inputfile> <outputfile>")
		return
	}

	command := os.Args[1]
	inputFile := os.Args[2]
	outputFile := os.Args[3]
	//モックのための便宜的な具体型の導入
	ops := &defaultFileOperations{}

	//コマンド別に呼び出す関数を切り分ける
	switch command {
	case "reverse":
		err := reverse(inputFile, outputFile, ops, ops)
		if err != nil {
			fmt.Println("Error processing reverse:", err)
		}
	}

}

// 個別コマンド（ここではreverse)の実装
// 入力、出力ファイルの指定が必要
func reverse(inputFile, outputFile string, reader ReadFile, writer WriteFile) error {
	// ファイルを読み込む
	content, err := reader.ReadFile(inputFile)
	if err != nil {
		return fmt.Errorf("Error reading content: %w", err)
	}

	// 読み込んだデータを逆順にする
	reverseData := reverseString(string(content))
	// 逆順にしたデータを出力用ファイルにかき出し
	err = writer.WriteFile(outputFile, []byte(reverseData), 0644)
	if err != nil {
		return fmt.Errorf("Error writing to file: %w", err)
	}
	return nil
}

// データを逆順にする関数
func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

//設計プロセス

//没案1 reverse.go
//mainの中に全ての関数をベタガキする。

//ボツ案2 reverse2.go
//DIせずテストコードも書かない

//ボツ案3 reverse3.go
//パッケージ違いでややこしい。遅い。メソッドが増える（openしないといけない）と面倒。
//osパッケージのfunc Open(name string) (*File, error)
//osパッケージのfunc WriteFile(filename string, data []byte, perm FileMode) error
//func (f *File) Read(b []byte) (n int, err error):File型はReaderインターフェースを実装している
