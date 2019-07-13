package internal_test

import (
	"bytes"
	"io"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/slonegd-otus-go/10_envdir/internal"
)

func TestExecute(t *testing.T) {
	if err := os.Chmod("./test/D", 0000); err != nil {
		log.Fatal(err)
	}

	tests := []struct {
		name     string
		in       io.Reader
		args     []string
		wantOut  string
		wantErrw string
		wantErr  string
	}{
		{
			name:    "env with ./test",
			in:      strings.NewReader(""),
			args:    []string{"./test", "env"},
			wantOut: "A=12\nB=34\n",
			// C=56 находиться во внтуреннем каталоге, не проходим рекурсивно
			// D=78 закрыт для чтения, не паникуем
		},
		{
			name:    "env with wrong path",
			in:      strings.NewReader(""),
			args:    []string{"./test2", "env"},
			wantErr: "open ./test2: no such file or directory",
		},
		{
			name:    "dont exist prog",
			in:      strings.NewReader(""),
			args:    []string{"./test", "envenv"},
			wantErr: "Запуск программы завершился с ошибкой: exec: \"envenv\": executable file not found in $PATH",
		},
		{
			name:    "wrong args",
			in:      strings.NewReader(""),
			args:    []string{"envenv"},
			wantErr: "должно быть 2 аргумента: путь до каталога и имя программы",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			errw := &bytes.Buffer{}
			err := internal.Execute(tt.in, out, errw, tt.args)
			if err != nil {
				assert.Equal(t, tt.wantErr, err.Error())
			}
			assert.Equal(t, tt.wantOut, out.String())
			assert.Equal(t, tt.wantErrw, errw.String())
		})
	}

	if err := os.Chmod("./test/D", 0666); err != nil {
		log.Fatal(err)
	}
}
