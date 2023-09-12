package linkparser

import "testing"

func TestLinkParser(t *testing.T) {
	htmlString := `
	<div><a href="https://www.google.com/">Google</a></div>
	<div><a href="https://www.youtube.com/">Youtube</a></div>
	`
	parsed, err := Parse(htmlString)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	linkOne := Link{
		Href: "https://www.google.com/",
		Text: "Google",
	}
	linkTwo := Link{
		Href: "https://www.youtube.com/",
		Text: "Youtube",
	}
	testLinkArray := [2]Link{linkOne, linkTwo}

	for i := 0; i < len(testLinkArray); i++ {
		if testLinkArray[i].Href != parsed[i].Href  {
			t.Errorf("Want %v, got %v", testLinkArray[i].Href, parsed[i].Href)
		}
		if testLinkArray[i].Text != parsed[i].Text {
			t.Errorf("Want %v, got %v", testLinkArray[i].Text, parsed[i].Text)
		}
	}
}