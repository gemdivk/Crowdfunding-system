package social

import (
	"fmt"
	"net/url"
)

func GetFacebookShareLink(targetURL string) string {
	baseURL := "https://www.facebook.com/sharer/sharer.php?u="
	return fmt.Sprintf("%s%s", baseURL, url.QueryEscape(targetURL))
}
func GetTwitterShareLink(targetURL string, text string) string {
	baseURL := "https://twitter.com/intent/tweet?url="
	return fmt.Sprintf("%s%s&text=%s", baseURL, url.QueryEscape(targetURL), url.QueryEscape(text))
}
func GetLinkedInShareLink(targetURL string, title string, summary string) string {
	baseURL := "https://www.linkedin.com/shareArticle?mini=true&url="
	return fmt.Sprintf("%s%s&title=%s&summary=%s", baseURL, url.QueryEscape(targetURL), url.QueryEscape(title), url.QueryEscape(summary))
}
