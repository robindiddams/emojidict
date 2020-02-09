package emojidict

// Definition of an emoji
type Definition struct {
	CodePoints []rune
}

// func test() {
// 	fmt.Println("getting emojis")
// 	client := scraper.Client{}
// 	if err := client.Connect(); err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	defer client.Close()
// 	versions, err := client.GetEmojiVersions()
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	latest := versions[len(versions)-1]
// 	fmt.Println("using version", latest)
// 	emojis, err := client.GetEmojis(latest)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	buf, _ := json.Marshal(emojis)
// 	ioutil.WriteFile("export/latest.json", buf, 0777)
// 	fmt.Printf("wrote %d emojis\n", len(emojis))
// }
