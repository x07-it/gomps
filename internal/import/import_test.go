package import_test

import (
	"testing"

	. "github.com/x07-it/gomps"
	. "github.com/x07-it/gomps/components"
	. "github.com/x07-it/gomps/html"
	. "github.com/x07-it/gomps/http"
)

func TestImports(t *testing.T) {
	t.Run("this is just a test that does nothing, but I need the dot imports above", func(t *testing.T) {
		_ = El("div")
		_ = A()
		_ = HTML5(HTML5Props{})
		_ = Adapt(nil)
	})
}
