package gorm

import (
	"fmt"
	"github.com/346285234/bbs-server/pkg/bbs"
	"os"
	"testing"
)

var co categoryOperation

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

func setup() {
	name := "mysql"
	url := "root:346285234@/bbs?parseTime=true"
	db := Open(name, url)
	//clear()

	co = newCategoryOperation(db)
}

func TestList(t *testing.T) {
	categories, err := co.list()
	if err != nil {
		t.Error(err)
	}

	fmt.Println(categories)
}

func TestAdd(t *testing.T) {
	values := []string{"Music", "Comedy", "Film", "Entertainment", "Gaming", "Sports", "Tech",
		"Beauty & Fashion", "News", "Health"}
	var categories = make([]bbs.Category, len(values))
	for i, v := range values {
		categories[i] = bbs.Category{Value: v}
	}

	for _, v := range categories {
		err := co.add(&v)
		if err != nil {
			t.Error(err)
		}
	}
}

func clear() bool {
	categories, err := co.list()
	if err != nil {
		return false
	}
	for _, v := range categories {
		if err := co.remove(v.ID); err != nil {
			return false
		}
	}

	return true
}
