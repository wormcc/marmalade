package cmd

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/urfave/cli"
)

var GenSecret = cli.Command{
	Name:   "gen_secret",
	Usage:  "Gen secret for app",
	Action: generateRandomString,
}

func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}
	return b, nil
}

func generateRandomString(_ *cli.Context) {
	b, err := generateRandomBytes(32)
	if err != nil {
		fmt.Errorf("err: %v", err)
	}
	fmt.Printf("Your sercret: %s\n", base64.URLEncoding.EncodeToString(b))
}
