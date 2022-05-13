package utils

import (
	"github.com/briandowns/spinner"
	"github.com/joho/godotenv"

	"os"
	"path"
	"time"
)

func GetTokenFromLocalEnv() (string, error) {
	p := path.Join(path.Join(os.Getenv("GOPATH"), "src/git.garena.com/russell.chanxl/dinner/.env"))
	err := godotenv.Load(p)
	if err != nil {
		return "", err
	}
	return os.Getenv("TOKEN"), nil
}

func NewSpinner() *spinner.Spinner {
	s := spinner.New(spinner.CharSets[54], 100*time.Millisecond)
	s.FinalMSG = "✔\n"
	return s
}
