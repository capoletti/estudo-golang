package arquitetura

import (
	"runtime"
	"testing"
)

func TestArquitetura(t *testing.T) {

	t.Parallel() //pode executar em paralelo
	if runtime.GOARCH == "amd64" {
		t.Skip("Não funciona na arquitetura amd64")
	}
	t.Fail()
}
