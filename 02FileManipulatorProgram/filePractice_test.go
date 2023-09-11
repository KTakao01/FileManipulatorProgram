// 必要なパッケージをインポート
package main

// ファイル読み込みのモックを作成
import (
	"errors"
	"os"
	"testing"
)

// mockFileReaderのReadFileメソッドを実装
type mockFileReader struct {
	content []byte
	err     error
}

func (m *mockFileReader) ReadFile(filename string) ([]byte, error) {
	return m.content, m.err
}

// ファイル書き込みのモックを作成
type mockFileWriter struct {
	err error
}

// mockFileWriterのWriteFileメソッドを実装
func (m *mockFileWriter) WriteFile(filename string, data []byte, perm os.FileMode) error {
	return m.err
}

// reverse関数のテストを実装
func TestReverse(t *testing.T) {
	//正常系と異常系のテストケースを定義
	tests := []struct {
		name           string
		inputContent   string
		expectedOutput string
		readerErr      error
		writerErr      error
		expectedErr    bool
	}{
		{
			name:           "Successful reversal",
			inputContent:   "hello",
			expectedOutput: "olleh",
			readerErr:      nil,
			writerErr:      nil,
			expectedErr:    false,
		},
		{
			name:           "Reader error",
			inputContent:   "",
			expectedOutput: "",
			readerErr:      errors.New("read error"),
			writerErr:      nil,
			expectedErr:    true,
		},
		{
			name:           "Writer error",
			inputContent:   "hello",
			expectedOutput: "",
			readerErr:      nil,
			writerErr:      errors.New("write error"),
			expectedErr:    true,
		},
	}

	// 各テストケースを実行
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//モックオブジェクトのインスタンスを生成
			reader := &mockFileReader{
				content: []byte(tt.inputContent),
				err:     tt.readerErr,
			}

			writer := &mockFileWriter{
				err: tt.writerErr,
			}
			// reverse関数を呼び出し
			err := reverse("dummyInput", "dummyOutput", reader, writer)
			// 出力結果と期待値を確認
			// 異常系の場合の確認
			// エラーの期待値を確認
			if tt.expectedErr {
				if err == nil {
					t.Errorf("expected an error but got none")
				}
			}

			// 正常系の場合の確認
			// 予期しないエラーが発生した場合の確認
			if !tt.expectedErr {
				if err != nil {
					t.Errorf("expected no error but got %v", err)
				}
			}

		})
	}

}

//プロセス
//モック構造体を考える
//テストケース考える
//テストケースを実行する
//テストケースの結果を確認する

//インターフェースを関数に渡すことでメソッドを内部で呼び出す処理ができる
//呼び出す際にモック構造体を作成、そのメソッドもモック構造体のフィールドを返すようにすると実際のデータを使用せずシミュレートできる
