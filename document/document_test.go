package document

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// begin
	code := m.Run()
	// end
	os.Exit(code)

}

func TestIndex(t *testing.T) {
	Index()
}

func TestGet(t *testing.T) {
	t.Run("found", func(t *testing.T) {
		user, err := Get("1")
		if err != nil {
			t.Log(err)
		}
		t.Log(user)
	})
	t.Run("not found", func(t *testing.T) {
		user, err := Get("na")
		if err != nil {
			t.Log(err)
		}
		t.Log(user)
	})
}

func TestUpdate(t *testing.T) {
	err := Update("1")
	if err != nil {
		t.Log(err)
	}
}
func TestDelete(t *testing.T) {
	Delete("1")
}
