package error

import (
	"fmt"
)

const (
	//ExitCodeOK 正常終了
	ExitCodeOK = 0
	//ExitCodeWarn 警告終了
	ExitCodeWarn = 1
	// ExitCodeError 異常終了
	ExitCodeError = 2
)

// ExitError エラーインターフェースを実装した構造体
type ExitError struct {
	exitCode  int    // 終了コードを定義 const化したほうがいい
	messageID string // メッセージIDの定義
	err       error  //エラーをそのまま入れる
}

func (ee *ExitError) Error() string {
	if ee.err == nil {
		return ""
	}
	return fmt.Sprintf("%v", ee.err)
}

// NewExitError Exiterror型の生成
func NewExitError(exitCode int, messageID string, err error) *ExitError {
	return &ExitError{
		exitCode:  exitCode,
		messageID: messageID,
		err:       err,
	}
}

// HandleExit errorハンドラー
func HandleExit(err error) int {
	if err == nil {
		return ExitCodeOK
	}

	if e, ok := err.(*ExitError); ok {
		// logging処理をして、終了コードを返却
		fmt.Printf("messageID:%v\n", e.messageID)
		fmt.Printf("err:%v\n", e.err)
		fmt.Printf("exitCode:%v\n", e.exitCode)
		return e.exitCode
	}

	return ExitCodeOK
}
