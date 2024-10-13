package Basic

import (
	"fmt"
	"net/url"
	"testing"
)

func TestUrl(t *testing.T) {
	u, _ := url.JoinPath("/v1", "user", "list")
	fmt.Printf("%s\n", u)

	uEscaped := url.PathEscape("/a/b/c d")
	fmt.Println(uEscaped)

	uUnescaped, _ := url.PathUnescape(uEscaped)
	fmt.Println(uUnescaped)

	uQueryEscape := url.QueryEscape("/a/b/c d")
	fmt.Println(uQueryEscape)
	uQueryUnescape, _ := url.QueryUnescape(uQueryEscape)
	fmt.Println(uQueryUnescape)

	uQueryEscape2 := url.QueryEscape("/v1?a=1&b=2&c=3&d=4 5")
	fmt.Println(uQueryEscape2)
	uQueryUnescape2, _ := url.QueryUnescape(uQueryEscape2)
	fmt.Println(uQueryUnescape2)
}
