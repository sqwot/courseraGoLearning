package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"

	tgbotapi "gopkg.in/telegram-bot-api.v4"
)

func getToken() string {
	data, err := ioutil.ReadFile("tgtoken.secret")
	if err != nil {
		panic(err)
	}
	return string(data)

}

var (
	BotToken  string = getToken()
	WebhokURL        = "https://9530d9357941.ngrok.io"
	//.\ngrok.exe http 8080
)

var rss = map[string]string{
	"Habr": "https://habr.ru/rss/best",
}

type RSS struct {
	Items []Item `xml:"channel>item"`
}
type Item struct {
	URL   string `xml:"guid"`
	Title string `xml:"title"`
}

func getNews(url string) (*RSS, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	rss := new(RSS)
	err = xml.Unmarshal(body, rss)
	if err != nil {
		return nil, err
	}

	return rss, nil
}

func main() {
	colorReset := "\033[0m"
	colorGreen := "\033[32m"
	fmt.Println(string(colorGreen), "***************TG Bot***************")
	fmt.Println(string(colorReset), "")

	bot, err := tgbotapi.NewBotAPI(BotToken)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Autorized on account %s\n", bot.Self.UserName)

	_, err = bot.SetWebhook(tgbotapi.NewWebhook(WebhokURL))
	if err != nil {
		panic(err)
	}

	updates := bot.ListenForWebhook("/")
	go http.ListenAndServe(":8080", nil)
	fmt.Println("start listen :8080")

	for update := range updates {
		if url, ok := rss[update.Message.Text]; ok {
			rss, err := getNews(url)
			if err != nil {
				bot.Send(tgbotapi.NewMessage(
					update.Message.Chat.ID,
					"sorry, error happend",
				))
			}
			for _, item := range rss.Items {
				bot.Send(tgbotapi.NewMessage(
					update.Message.Chat.ID,
					item.URL+"\n"+item.Title,
				))
			}

		} else {
			bot.Send(tgbotapi.NewMessage(
				update.Message.Chat.ID,
				"there is only Habr feed availible",
			))
		}
	}
}
