package internal

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func md5Encode(data string) string {
	cpt := md5.New()
	cpt.Write([]byte(data))
	result := hex.EncodeToString(cpt.Sum(nil))
	fmt.Println(result)
	return result
}
