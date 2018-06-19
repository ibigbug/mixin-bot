package storage

import (
	"log"
	"os"
)

var (
	accountName, accountKey string
)

var credential

func init() {
	accountName, accountKey = os.Getenv("azure_storage_account_name"), os.Getenv("azure_storage_account_key")
	if len(accountName) == 0 || len(accountKey) == 0 {
		log.Fatal("azure storage credential missing")
	}
}

func GetOneImage() ([]byte, error) {

}
