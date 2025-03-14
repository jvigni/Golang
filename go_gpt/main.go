package main

import (
	"context"
	"fmt"
	"os"

	htgotts "github.com/hegedustibor/htgo-tts"
	"github.com/hegedustibor/htgo-tts/voices"
	openai "github.com/sashabaranov/go-openai"
)

func main() {
	MicToText()
}

func MicToText() {

}

func CallTextToSpeach() {
	speech := htgotts.Speech{
		Folder:   "audio",
		Language: voices.Spanish,
	}

	speech.Speak("Aguante boquita campeon papa no me importa nada")
	fmt.Println("Done.")
}

// speach to text
func CallWhisper() {
	client := openai.NewClient("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")
	resp, err := client.CreateTranscription(
		context.Background(),
		openai.AudioRequest{
			Model:    openai.Whisper1,
			FilePath: os.Args[1],
		},
	)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.Text)
}

// si detecta que no puede responderte, llama a un caller de verdad
func CallGPT() {
	defer fmt.Scanln()

	client := openai.NewClient("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")
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "#Reglas del negocio/mensaje de usuarios'", //1r mensaje con todo el contenido de la empresa para "cargar" al bot. 2d en adelante los mensajes de los users [speech to text]
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
	}

	fmt.Println(resp.Choices[0].Message.Content)
}
