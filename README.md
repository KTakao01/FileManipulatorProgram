# FileManipulatorProgram
Recursionのバックエンド・プロジェクト1の課題。</br>
課題に必要な基礎知識を学んだ後、0から課題を実装した。3つの課題がある。</br>

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
## 工夫した点
・Pythonで解説されていましたが、 Goに変換しながら課題を完了させたこと。</br>
・実装内容を最初に日本語で考えてコメントに残してから実装することで手戻りや不要な思考をなるべく排除するように努めたこと。</br>
・入力できる回数を制限したこと。最後に正答率を表示できるようにするなど仕様を考えたこと。</br>
・エラーを網羅できるように書いた（はず・・・）。</br>
## 苦労した点
以下の点は未知で多少手こずった。</br>
・乱数生成器の作り方がバージョンアップで変わっていたこと。</br>
・Sprintfのフォーマットに"%"を使う時は注意すること。</br>
・Ubuntu環境でコードを動作させたこと。</br>
## 改善点
・簡単な処理なので複数の関数に分割しなかった。今後は気をつけたい。</br>