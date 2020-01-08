package operations

import (
	"fmt"
	"github.com/346285234/bbs-server/data"
	"github.com/346285234/bbs-server/data/models"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

func setup() {
	url := "root:346285234@/bbs?parseTime=true"
	data.OpenDB(url)
	//clear()
}

func TestList(t *testing.T) {
	categories, err := Co.List()
	if err != nil {
		t.Error(err)
	}

	fmt.Println(categories)
}

func TestAdd(t *testing.T) {
	values := []string{"Music", "Comedy", "Film", "Entertainment", "Gaming", "Sports", "Tech",
		"Beauty & Fashion", "News", "Health"}
	var categories = make([]models.Category, len(values))
	for i, v := range values {
		categories[i] = models.Category{Value: v}
	}

	for _, v := range categories {
		err := Co.add(v)
		if err != nil {
			t.Error(err)
		}
	}
}

func clear() bool {
	categories, err := Co.List()
	if err != nil {
		return false
	}
	for _, v := range categories {
		if err := Co.remove(v.ID); err != nil {
			return false
		}
	}

	return true
}