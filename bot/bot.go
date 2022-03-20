package bot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Dipankar-Medhi/goDiscordBot/config"
	"github.com/bwmarrin/discordgo"
)

var BotId string
var goBot *discordgo.Session

type apiKey struct {
	ApiKey string `json: "apiKey"`
}

type addressDetails struct {
	Data struct {
		Item struct {
			ConfirmedBalance struct {
				Amount string `json: "amount"`
				Unit   string `json: "uinit"`
			} `json: "confirmedBalance"`
		} `json "item"`
	} `json: "data"`
}

func getApiKey() string {
	data, err := ioutil.ReadFile("./apiConfig.json")
	if err != nil {
		fmt.Println(err.Error())
	}
	var d apiKey
	err = json.Unmarshal(data, &d)
	if err != nil {
		fmt.Println(err.Error())
	}
	return d.ApiKey
}

func getAddressDetails(address string, apiKey string) (string, string) {

	url := "https://rest.cryptoapis.io/v2/blockchain-data/bitcoin/testnet/addresses/" + address + "?context=yourExampleString"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-API-Key", "ebc8a1a3d4a7b72814949e366b5e3fa670c2fc28")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	d := &addressDetails{}

	err := json.NewDecoder(bytes.NewReader(body)).Decode(d)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Here is the data")
	fmt.Println("Balance: " + d.Data.Item.ConfirmedBalance.Amount + ", Unit: " + d.Data.Item.ConfirmedBalance.Unit)

	balance := d.Data.Item.ConfirmedBalance.Amount
	unit := d.Data.Item.ConfirmedBalance.Unit

	return balance, unit
}

func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	user, err := goBot.User("@me")
	if err != nil {
		fmt.Println(err.Error())
	}

	BotId = user.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot is running...")

	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	key := getApiKey()
	if m.Author.ID == BotId {
		return
	}
	if m.Content == "mybalance" {

		balance, unit := getAddressDetails("mzYijhgmzZrmuB7wBDazRKirnChKyow4M3", key)

		_, _ = s.ChannelMessageSend(m.ChannelID, "Balance: "+balance+", Uint: "+unit)
		fmt.Println()

	}
}
