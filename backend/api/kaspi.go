package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"net/http/cookiejar"
	"strconv"
	"strings"
	"unicode"
)

func (app *application) kaspiParser() {
	url := "https://promotion.kaspi.kz/"

	client := &http.Client{}

	cookieJar, _ := cookiejar.New(nil)
	client.Jar = cookieJar

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	cookieStrings := []string{
		"installId=1FFC815F-CE3C-4EAB-BDEF-72818FA5D150",
		"is_mobile_app=true",
		"kaspi-payment-region=19",
		"locale=ru-RU",
		"ma_platform_type=IOS",
		"ma_platform_ver=16.6.1",
		"ma_ver=5.39.2",
		"mobapp_version=37",
		"pd=AkgeqP3ohroi9PjjhQi1CE",
		"_ga_H6ETDBR5S9=GS1.1.1712991367.3.0.1712991371.56.0.371732822",
		"_gcl_au=1.1.1926414148.1712836524",
		"_fbp=fb.1.1712836526236.1116366220",
		"amp_3c7b23=E476DA64-382C-4A69-AA0C-34E43F104CEE...1hrb3rss4.1hrb3rss4.0.0.0",
		"ssaid=61bf0f60-f7fa-11ee-a4b2-01f77a1e9eb2",
		"_ga=GA1.1.1194434315.1712836525",
		"_hjSessionUser_283363=eyJpZCI6IjY4YTJjY2QyLTkwZGUtNWQyMi05MDdkLTFhN2Y2ZGI2YTRkZCIsImNyZWF0ZWQiOjE3MTI4MzY1MjUxOTQsImV4aXN0aW5nIjp0cnVlfQ==",
		"_ym_isad=2",
		"_tt_enable_cookie=1",
		"_ttp=CGEqqIbQWJ0LykbJnXbAIOcNYQk",
		"_ym_d=1712836526",
		"_ym_uid=1712836526505103663",
		"test.user.group=41",
	}

	for _, cookieStr := range cookieStrings {
		parts := strings.SplitN(cookieStr, "=", 2)
		cookie := &http.Cookie{Name: parts[0], Value: parts[1]}
		req.AddCookie(cookie)
	}

	req.Header.Set("X-Mobile-App", "true")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var parse func(*html.Node)
	parse = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "data-id" && checkBonus(attr.Val) {
					endIndex := strings.Index(attr.Val, "B_")
					temp := attr.Val[:endIndex]
					startIndex := strings.LastIndex(temp, "_") + 1
					bonus := temp[startIndex:]
					dataId := attr.Val
					f, err := strconv.ParseFloat(bonus, 64)
					if err != nil {
						fmt.Println("Ошибка при преобразовании строки в число:", err)
						return
					}
					var title string
					for _, attr := range n.Attr {
						if attr.Key == "class" && strings.Contains(attr.Val, "promotion-item") {
							for c := n.FirstChild; c != nil; c = c.NextSibling {
								if c.Type == html.ElementNode && c.Data == "div" {
									for _, attr := range c.Attr {
										if attr.Key == "class" {
											if attr.Val == "promotion-item__title" {
												title = getTextContent(c)
											}
										}
									}
								}
							}
						}
					}

					for _, attr := range n.Attr {
						if attr.Key == "href" {
							if strings.Contains(attr.Val, "https://kaspi.kz/shop/c/") {
								startIndex := strings.Index(attr.Val, "/c/") + len("/c/")
								temp := attr.Val[startIndex:]
								endIndex := strings.Index(temp, "/")
								category := temp[:endIndex]
								app.promos.AddKaspi(title, attr.Val, "Kaspi Bank", "Promo", spaceReadable(category), f)
							} else {
								startIndex := strings.Index(dataId, "_") + 1
								temp := dataId[startIndex:]
								endIndex := strings.Index(temp, "_")
								app.promos.AddKaspi(title, attr.Val, "Kaspi Bank", "Promo", spaceReadable(strings.ToLower(temp[:endIndex])), f)
							}
						}
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			parse(c)
		}
	}

	parse(doc)
}

func isDigit(str string) bool {
	for _, char := range str {
		if !unicode.IsDigit(char) {
			return false
		}
	}
	return true
}

func checkBonus(str string) bool {
	if strings.Contains(str, "B_") {
		Index := strings.Index(str, "B_")
		temp := str[Index-1 : Index]
		if isDigit(temp) {
			return true
		}
	}

	return false
}

func spaceReadable(str string) string {
	new := strings.Map(func(r rune) rune {
		if unicode.IsDigit(r) {
			return -1
		}
		return r
	}, str)
	new = strings.Replace(new, "%", " ", -1)

	return new
}

func getTextContent(n *html.Node) string {
	var textContent string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.TextNode {
			textContent += c.Data
		}
	}
	return textContent
}
