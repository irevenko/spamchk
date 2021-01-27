package spamchk

import (
	"testing"
)

func TestIsStringSpam_WithSpam(t *testing.T) {
	expectation := true
	actual := IsStringSpam("as one experienced e - supplier , our cyber corporation brings a lawful access to medis . we supply medicals on severe ache , swelling , muscular relaxation , tensions , over wt . and insomnia .")
	if actual != expectation {
		t.Errorf("Expected %v but got %v", expectation, actual)
	}
}

func TestIsStringSpam_WithHam(t *testing.T) {
	expectation := false
	actual := IsStringSpam("dear vince , i enjoyed very much meeting you at the eprm conference . as agreed , i am attaching both my resume and the paper on my dissertation work , which i will be presenting at the real options conference in ucla next month .")
	if actual != expectation {
		t.Errorf("Expected %v but got %v", expectation, actual)
	}
}
