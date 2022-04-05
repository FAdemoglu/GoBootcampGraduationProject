package pagination

import "testing"

func TestNew(t *testing.T) {
	tests := []struct {
		tag                                                                     string
		page, pageSize, total                                                   int
		expectedPage, expectedPageSize, expectedTotal, pageCount, offset, limit int
	}{
		// varying page
		{"t1", 1, 20, 50, 1, 20, 50, 3, 0, 20},
		{"t2", 2, 20, 50, 2, 20, 50, 3, 20, 20},
		{"t3", 3, 20, 50, 3, 20, 50, 3, 40, 20},
		{"t4", 3, 20, 50, 3, 20, 50, 3, 40, 20},
		{"t5", 0, 20, 50, 1, 20, 50, 3, 0, 20},

		// varying pageSize
		{"t6", 1, 0, 50, 1, 100, 50, 1, 0, 100},
		{"t7", 1, -1, 50, 1, 100, 50, 1, 0, 100},
		{"t8", 1, 100, 50, 1, 100, 50, 1, 0, 100},
		{"t9", 1, 1001, 50, 1, 1000, 50, 1, 0, 1000},

		// varying total
		{"t10", 1, 20, 0, 1, 20, 0, 0, 0, 20},
		{"t11", 1, 20, -1, 1, 20, -1, -1, 0, 20},
	}

	for _, test := range tests {
		p := New(test.page, test.pageSize, test.total)
		if test.expectedPage != p.Page {
			t.Errorf("Test failed [%s] : %v was given, %v want, %v got", test.tag, test.page, test.expectedPage, p.Page)
		}
		if test.expectedPageSize != p.PageSize {
			t.Errorf("Test failed [%s] : %v was given, %v want, %v got", test.tag, test.pageSize, test.expectedPageSize, p.PageSize)
		}

		if test.expectedPageSize != p.PageSize {
			t.Errorf("Test failed [%s] : %v was given, %v want, %v got", test.tag, test.pageSize, test.expectedPageSize, p.PageSize)
		}
		if test.expectedTotal != p.TotalCount {
			t.Errorf("Test failed [%s] : %v was given, %v want, %v got", test.tag, test.total, test.expectedTotal, p.TotalCount)
		}
		if test.pageCount != p.PageCount {
			t.Errorf("Test failed [%s] : %v want, %v got", test.tag, test.pageCount, p.PageCount)
		}
		if test.offset != p.Offset() {
			t.Errorf("Test failed [%s] : %v want, %v got", test.tag, test.offset, p.Offset())
		}
		if test.limit != p.Limit() {
			t.Errorf("Test failed [%s] : %v want, %v got", test.tag, test.limit, p.Limit())
		}
	}
}
