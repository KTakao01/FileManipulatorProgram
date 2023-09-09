package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//"Usage: go run file_manipulator.go reverse <inputfile> <outputfile>"

	//コマンドラインから引数を取得する。
	//反転させたいファイル（入力ファイル）を開く

		//入力ファイルから読み込みバイトスライスとして返す。

		//バイトスライスをstringに変換する
		//逆になった読み込みデータを出力ファイルを指定して書き込む。


//文字列を逆に変換する関数を作成する
//コードポイント単位(rune)で扱うとよい。文字列型、バイト型はUTF-8エンコードされた場においてUnicode文字はマルチバイトが多いので、不正な文字列が生成されやすい
