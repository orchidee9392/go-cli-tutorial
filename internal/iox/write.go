package iox

import (
	"io"
	"os"
)

func WriteOutput(outPath string, stdout io.Writer, b []byte) error {
	//出力ファイルパスがあれば優先してそこに出力する
	if outPath != "" {
		return os.WriteFile(outPath, b, 0o644)
	}

	//上記のすべてが無ければ、io.Writerに出力する
	_, err := stdout.Write(b)
	return err
}
