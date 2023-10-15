# FileManipulatorProgram
Recursionのバックエンド・プロジェクト1の課題。</br>
課題に必要な基礎知識を学んだ後、0から課題を実装した。3つの課題がある。以下詳細。</br>

# 学習の目的
・Linuxでの開発</br>
・Goによるバックエンドブログラムの構築経験の獲得</br>
・低レイヤーの理解</br>

以下、RecursionのHPより引用します。カリキュラムは元Facebookエンジニアが作成しています。</br>
>このコースでは、アプリケーションプログラムがオペレーティングシステムとどのように相互作用するかを探ります。
パイプのようなメモリベースの手法とファイルのようなストレージベースの手法の両方を使用して、Linux ファイルシステムに格納されているデータにアクセスして操作するプログラムを構築する経験を積むことができます。
さらに、ソフトウェアエンジニアにとって重要なスキルである開発環境の構築と使用方法についても学びます。
このコースが終了する頃には、これらの基本的な概念をしっかりと身につけることができるでしょう。

## Guess the number game 
### 概要
・最大、最小の順で２数を入力するとその範囲で乱数を生成する。その乱数の値が何かを規定回数以内に当てるゲーム。
### 工夫した点
・Pythonで予備知識が解説されていたが、 Goに変換しながら課題を完了させたこと。</br>
・実装内容を最初に日本語で考えてコメントに残してから実装することで手戻りや不要な思考をなるべく排除するように努めたこと。</br>
・入力できる回数を制限したこと。最後に正答率を表示できるようにするなど仕様を考えたこと。</br>
・エラーを網羅できるように書いた（はず・・・）。</br>
### 苦労した点
初めて用いるパッケージの関数などがあり、仕様を理解するのに手間取った。詳細は以下の通り。</br>
・乱数生成器の作り方がバージョンアップで変わっていたこと。</br>
・Sprintfのフォーマットに"%"を使う時は注意すること。</br>
・Ubuntu環境でコードを動作させたこと。</br>
### 改善点
・簡単な処理なので複数の関数に分割しなかった。今後は気をつけたい。</br>

## File Manipulator Program
### 概要
・以下のコマンドを実装</br>

>reverse inputpath outputpath: inputpath にあるファイルを受け取り、outputpath に inputpath の内容を逆にした新しいファイルを作成します。</br>
copy inputpath outputpath: inputpath にあるファイルのコピーを作成し、outputpath として保存します。</br>
duplicate-contents inputpath n: inputpath にあるファイルの内容を読み込み、その内容を複製し、複製された内容を inputpath に n 回複製します。</br>
replace-string inputpath needle newstring: inputpath にあるファイルの内容から文字列 'needle' を検索し、'needle' の全てを 'newstring' に置き換えます。</br>
### 工夫した点
・Pythonで予備知識が解説されていたが、 Goに変換しながら課題を完了させたこと。</br>
・ユニットテストコードを書いたこと。(DI,モック構造体の作成,テストコードの実行,期待値と結果の比較)</br>
### 苦労した点
・ユニットテストコードを書くためにリファクタリングしたが、インターフェースや構造体、DIなどの概念を理解、実装するのが難しかった。</br>
### 改善点
・テストコードを作ったのでCIを構築したい。</br>

## Markdown to HTML Converter
### 概要
>タスクはマークダウンを HTML に変換するプログラムを作成し、シェルを通して python3 file-converter.py markdown inputfile outputfile というコマンドを実行させることです。</br>ここで、markdown は実行するコマンド、inputfile は .md ファイルへのパス、出力パスはプログラムを実行した後に作成される .html です。</br>

### 工夫した点
・Pythonで予備知識が解説されていたが、 Goに変換しながら課題を完了させたこと。</br>
### 苦労した点
・パッケージの選定を悩んだ。使用感は似ている２つのパッケージのどちらを使うか。定番で更新がとまっているblackfridayと比較的新しいものgomarkdownと。情報はいずれにしても公式ドキュメントがほとんどだったが、
目立った差異がみられなかったので、使いやすそうな前者を使用した。
### 改善点
・blackfridayのライブラリの中身をおいおい読みたい。</br>