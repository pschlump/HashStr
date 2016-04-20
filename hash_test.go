package HashStr

//
// Hash string into number or into short name.
//
// (C) Philip Schlump, 2013-2015.
// MIT Licensed
//

import (
	"encoding/hex"
	"testing"
)

//func HashStr(s []byte) (n int) {
//func HashStrToName(s string) (s string) {
func Test_HashFunctions(t *testing.T) {
	n := HashStr([]byte("abc"))
	// fmt.Printf("n=%d\n", n)
	if n != 2565023054915971881 {
		t.Errorf("Unexpecd value for hash")
	}
	s := HashStrToName("select 12 from dual")
	// fmt.Printf("s=>%s< %x\n", s, s)
	exp, _ := hex.DecodeString("0a54c2bdc38d22")
	if s != string(exp) {
		t.Errorf("Unexpecd value for hash")
	}
}

/*

// -----------------------------------------------------------------------------------------------------------------------------------------------
func Test_JsonPPathServer(t *testing.T) {
	tests := []struct {
		url          string
		expectedBody string
	}{
		{
			"http://example.com/img/foo.jpg",
			`{"abc":"def"}`,
		},
		{
			"http://example.com/api/status",
			`{"abc":"def"}`,
		},
		{
			"http://example.com/api/status?callback=j1232131231",
			`j1232131231({"abc":"def"});`,
		},
	}
	// ct := h.Get("Content-Type")
	// if rw.StatusCode == http.StatusOK && strings.HasPrefix(ct, "application/json") {
	bot := mid.NewConstHandler(`{"abc":"def"}`, "Content-Type", "application/json")
	ms := NewJSONPServer(bot, []string{"/api/status"}, `^[a-zA-Z\$_][a-zA-Z0-9\$_]*$`)
	var err error
	lib.SetupTestCreateDirs()

	for ii, test := range tests {

		rec := httptest.NewRecorder()

		wr := goftlmux.NewMidBuffer(rec)

		var req *http.Request

		req, err = http.NewRequest("GET", test.url, nil)
		if err != nil {
			t.Fatalf("Test %d: Could not create HTTP request: %v", ii, err)
		}
		lib.SetupTestMimicReq(req, "example.com")

		ms.ServeHTTP(wr, req)
		wr.FinalFlush()

		b := string(rec.Body.Bytes())
		if b != test.expectedBody {
			t.Errorf("Error %2d, reject error got: %s, expected %s\n", ii, b, test.expectedBody)
		}

	}

}

*/

/* vim: set noai ts=4 sw=4: */
