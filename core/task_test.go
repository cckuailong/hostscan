package core

import "testing"

func TestGetTitle(t *testing.T) {
	t1 := "<html><title>hello</title>\n<body>123</body></html>"
	tr1 := getTitle(t1)
	t.Log(tr1)

	t2 := `<html>
	<title>
		hello
	</title>
	<body>
		123
	</body>
</html>`
	tr2 := getTitle(t2)
	t.Log(tr2)
}
