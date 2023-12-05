package helper

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
)

func GetBasePrompt(prompt string) string {
	return "Membuat panduan kursus dari `" + prompt + "`terdiri dari judul utama, sub-judul, ringkasan singkat dari kursus tersebut yang terdiri dari 10 - 25 kata, lama durasi dari kursus tersebut, jumlah sub-judul yang ada, tema kegiatan dari kursus tersebut seperti pendidikan, olahraga, teknologi dan sejenisya (lower-case), tipe kegiatan dari kursus tersebut dengan pilihan terbatas pada indoor, outdoor, hybrid (lower-case). Setiap sub-judul memiliki sebuah materi yang berisi kalimat pembuka yang menjelaskan tentang sub-judul tersebut, panduan langkah demi langkah, kalimat penutup yang dipisahkan oleh baris. Setiap sub-judul memiliki deksripsi yang menjelaskan tentang sub-judul tersebut.Setiap langkah pada panduan dijelaskan secara rinci dan jelas dengan minimal tiap langkah memiliki satu paragraf. Setiap sub-judul memiliki panduan dipisahkan menggunakan baris. Pastikan Setiap kursus memiliki minimal 3 sub-judul. Setiap panduan didalam sub-judul hanya memilki satu panduan didalamnya dan memiliki minimal 3 langkah didalamnya dengan setiap langkah hanya berada pada satu panduan. Setiap langkah pada panduan memiliki nomor serta dipisahkan oleh baris. Buatkan dalam format json dengan title, ringkasan singkat dan sub-judul terpisah. Serta Pastikan nama key durasi sama dengan `duration`, jumlah sub-judul sama dengan `lessons`, jenis kegiatan sama dengan `type_activity` , tipe kegiatan sama dengan `theme_activity` ,judul sama dengan title, ringkasan singkat sama dengan `desc`, sub-judul sama dengan `subtitles` berbentuk list yang didalamnya ada `topic` `content`, judul didalam `subtitles` sama dengan `topic`, deskripsi didalam sub-judul bernama `shortdesc` ,panduan sama dengan `content` benbentuk object (dictionary) berada didalam subtitles, kalimat pembuka pada panduan dengan nama `opening`, langkah pada panduan dengan nama `step` berbentuk list [str, str, ...], kalimat penutup pada panduan dengan nama `closing`. Pastikan output selalu konsisten dan memiliki format json yang selalu sama, nama key json sama dengan nama yang sudah diberikan. output hanya berupa json saja."
}

func GetBasePrompt2(prompt string) string {
	return "menbuat panduan kursus dari `" + prompt + "`. untuk type_activity hanya `indoor` dan `outdoor`. untuk theme_activity menyesuaikan. terdiri dari title, desc, duration, subtitles[], didalam subtitles ada topic, shortdesc, content, didalam content ada opening, step[], closing"
}
func GenerateCourse(promptRequest string) (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading env file")
	}

	apiKey := os.Getenv("OPENAI_API_KEY")

	client := openai.NewClient(apiKey)

	req := openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    "user",
				Content: GetBasePrompt(promptRequest),
			},
		},
		MaxTokens:   2000,
		Temperature: 0.1,
		N:           5,
	}
	res, err := client.CreateChatCompletion(context.Background(), req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	strjson := res.Choices[0].Message.Content

	fmt.Println(strjson)

	return strjson, nil
}
