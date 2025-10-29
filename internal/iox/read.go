package iox

import(
	"io"
	"os"
)
 
func ReadInput(arg string, inPath string, stdin io.Reader) ([]byte, error){
	//引数があれば最優先で返す
	if arg != ""{
		return []byte(arg), nil
	}

	//入力ファイルパスがあれば優先して返す
	if inPath != ""{
		return os.ReadFile(inPath)
	}

	//上記すべてが無ければ、io.Readerから読み込む
	return io.ReadAll(stdin)
}