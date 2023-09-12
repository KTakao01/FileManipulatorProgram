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
	writtenContent []byte
	err            error
}

// mockFileWriterのWriteFileメソッドを実装
func (m *mockFileWriter) WriteFile(filename string, data []byte, perm os.FileMode) error {
	if m.err != nil {
		return m.err
	}
	m.writtenContent = data
	return nil
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

			if !tt.expectedErr {
				if err != nil {
					t.Errorf("expected no error but got %v", err)
				}
			}

			if tt.expectedOutput != string(writer.writtenContent) {
				t.Errorf("expected %s but got %s", tt.expectedOutput, string(writer.writtenContent))
			}

		})
	}

}

// copy関数のテストを実装
func TestCopy(t *testing.T) {
	tests := []struct {
		name           string
		inputContent   string
		expectedOutput string
		readerErr      error
		writerErr      error
		expectedErr    bool
	}{
		{
			name:           "Successful copy",
			inputContent:   "hello",
			expectedOutput: "hello",
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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := &mockFileReader{
				content: []byte(tt.inputContent),
				err:     tt.readerErr,
			}
			writer := &mockFileWriter{
				err: tt.writerErr,
			}

			err := copy("dummyInput", "dummyOutput", reader, writer)

			if tt.expectedErr {
				if err == nil {
					t.Errorf("expected an error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("expected no error but got %v", err)
				}
			}

			if tt.expectedOutput != string(writer.writtenContent) {
				t.Errorf("expected %s but got %s", tt.expectedOutput, string(writer.writtenContent))
			}
		})

	}

}

// duplicateContents関数のテストを実装
func TestDuplicateContents(t *testing.T) {
	tests := []struct {
		name           string
		inputContent   string
		expectedOutput string
		readerErr      error
		writerErr      error
		expectedErr    bool
		loopNumber     int
	}{
		{
			name:           "Successful duplication",
			inputContent:   "hello",
			expectedOutput: "hellohello",
			readerErr:      nil,
			writerErr:      nil,
			expectedErr:    false,
			loopNumber:     2,
		},
		{
			name:           "Reader error",
			inputContent:   "",
			expectedOutput: "",
			readerErr:      errors.New("read error"),
			writerErr:      nil,
			expectedErr:    true,
			loopNumber:     2,
		},
		{
			name:           "Writer error",
			inputContent:   "hello",
			expectedOutput: "",
			readerErr:      nil,
			writerErr:      errors.New("write error"),
			expectedErr:    true,
			loopNumber:     2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := &mockFileReader{
				content: []byte(tt.inputContent),
				err:     tt.readerErr,
			}
			writer := &mockFileWriter{
				err: tt.writerErr,
			}

			err := duplicateContents("dummyInput", tt.loopNumber, reader, writer)

			if tt.expectedErr {
				if err == nil {
					t.Errorf("expected an error but got none")
				}
			}

			if !tt.expectedErr {
				if err != nil {
					t.Errorf("expected no error but got %v", err)
				}

			}

			if tt.expectedOutput != string(writer.writtenContent) {
				t.Errorf("expected %s but got %s", tt.expectedOutput, string(writer.writtenContent))
			}

		})
	}

}

func TestReplaceString(t *testing.T) {
	tests := []struct {
		name           string
		inputContent   string
		searchKey      string
		newString      string
		expectedOutput string
		readerErr      error
		writerErr      error
		expectedErr    bool
	}{
		{
			name:           "Successful replacement",
			inputContent:   "hello",
			searchKey:      "l",
			newString:      "x",
			expectedOutput: "hexxo",
			readerErr:      nil,
			writerErr:      nil,
			expectedErr:    false,
		},
		{
			name:           "Reader error",
			inputContent:   "",
			searchKey:      "l",
			newString:      "x",
			expectedOutput: "",
			readerErr:      errors.New("read error"),
			writerErr:      nil,
			expectedErr:    true,
		},
		{
			name:           "Writer error",
			inputContent:   "hello",
			searchKey:      "l",
			newString:      "x",
			expectedOutput: "",
			readerErr:      nil,
			writerErr:      errors.New("write error"),
			expectedErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := &mockFileReader{
				content: []byte(tt.inputContent),
				err:     tt.readerErr,
			}

			writer := &mockFileWriter{
				err: tt.writerErr,
			}
			err := replaceString("dummyInput", "l", "x", reader, writer)

			if tt.expectedErr {
				if err == nil {
					t.Errorf("expected an error but got none")
				}
			}

			if !tt.expectedErr {
				if err != nil {
					t.Errorf("expected no error but got %v", err)
				}
			}

			if tt.expectedOutput != string(writer.writtenContent) {
				t.Errorf("expected %s ,but got %s", tt.expectedOutput, string(writer.writtenContent))
			}

		})
	}

}

//プロセス
//モック構造体を考える(フィールドに実装データのかわりとなるデータを持つ)
//テストケース考える（正常系、異常系）
//テストケースを実行する（テーブルドリブンテスト）
//テストケースの結果を確認する

//インターフェースを関数に渡すことでメソッドを内部で呼び出す処理ができる
//呼び出す際にモック構造体を作成、そのメソッドもモック構造体のフィールドを返すようにすると実際のデータを使用せずシミュレートできる

//実際の手順
//テストコードの作成方法
//テスト構造体かく
//スコープわけ
//モック構造体つくる
//メソッドつくる
