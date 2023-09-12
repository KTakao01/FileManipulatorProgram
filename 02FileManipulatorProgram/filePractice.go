// インターフェースによる抽象化、DIになれるための練習
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
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
		fmt.Println("Usage: go run <file_name.go> <command_name> <inputfile> <outputfile>")
		return
	}

	command := os.Args[1]
	inputFile := os.Args[2]
	var outputFile, duplicateNumber, searchKey, newString string

	switch command {
	case "reverse", "copy":
		outputFile = os.Args[3]
	case "duplicate-contents":
		duplicateNumber = os.Args[3]
	case "replace-string":
		searchKey = os.Args[3]
		newString = os.Args[4]
	}

	//モックのために用意したインターフェースに便宜的な具体型の導入
	ops := &defaultFileOperations{}

	//コマンド別に呼び出す関数を切り分ける
	switch command {
	case "reverse":
		err := reverse(inputFile, outputFile, ops, ops)
		if err != nil {
			fmt.Println("Error processing reverse:", err)
		}
	case "copy":
		err := copy(inputFile, outputFile, ops, ops)
		if err != nil {
			fmt.Println("Error processing copy:", err)
		}
	case "duplicate-contents":
		loopNumber, err := strconv.Atoi(duplicateNumber)
		if err != nil {
			fmt.Println("error to change from string to int ", err)
		}
		err = duplicateContents(inputFile, loopNumber, ops, ops)
		if err != nil {
			fmt.Println("Error processing duplicate-contents:", err)
		}

	case "replace-string":
		err := replaceString(inputFile, searchKey, newString, ops, ops)
		if err != nil {
			fmt.Println("Error processing replace-string:", err)
		}
	}
}

// 個別コマンドの実装
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

func copy(inputFile, outputFile string, reader ReadFile, writer WriteFile) error {
	//ファイルを読み込む
	content, err := reader.ReadFile(inputFile)
	if err != nil {
		return fmt.Errorf("Error reading content: %w", err)
	} //読み込んだデータを出力用ファイルにかき出し
	err = writer.WriteFile(outputFile, content, 0644)
	if err != nil {
		return fmt.Errorf("Error writing to file: %w", err)
	}
	return nil
}

func duplicateContents(inputFile string, loopNumber int, reader ReadFile, writer WriteFile) error {
	originalContent, err := reader.ReadFile(inputFile)
	if err != nil {
		return fmt.Errorf("Error reading content: %w", err)
	}

	// 元の内容をloopNumber回複製する
	var duplicatedContent []byte
	for i := 0; i < loopNumber; i++ {
		duplicatedContent = append(duplicatedContent, originalContent...)
	}

	err = writer.WriteFile(inputFile, duplicatedContent, 0644)
	if err != nil {
		return fmt.Errorf("Error writing to file: %w", err)
	}

	return nil
}

func replaceString(inputFile, searchKey, newString string, reader ReadFile, writer WriteFile) error {
	content, err := reader.ReadFile(inputFile)
	if err != nil {
		return fmt.Errorf("Error reading content: %w", err)
	}

	replacedContent := strings.Replace(string(content), searchKey, newString, -1)

	err = writer.WriteFile(inputFile, []byte(replacedContent), 0644)
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

//実際の手順
//DIの作成方法
//関数の作成
//ライブラリに基づかないようにラップ関数の作成
//それをもつ構造体の作成
//その構造体の抽象化したインターフェース作成
