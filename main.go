package main

import (
	"encoding/hex"
	"os/exec"
	"os"
	"io/ioutil"
	"regexp"
	"strings"
	"strconv"
)

var noteFile = "note_from_kim_bong_kun.txt"
var noteShiftRegex, _ = regexp.Compile("so here it is: ([0-9]{2})\\.")
var noteFileListRegex, _ = regexp.Compile("We encrypted these:\\s((.|\\s)+)\\sAnd we failed to encrypt these \\(if any\\):")

var keyBuffer = "kOoqCEOa4WRJxjxpVVDHfxnyyANQzjXL"
var iv = hex.EncodeToString([]byte("encryptionisfun!"))

func decodeKey(shift int) (key []byte) {
	key = make([]byte, 0x20)
	for i := range key {
		key[i] = keyBuffer[(shift+i)%32]
	}
	return key
}

func banner() {
	println()
	println("   Team")
	println(`
 ____   __ __  ______  ______    ___  ____  
|    \ |  |  ||      ||      |  /  _]|    \ 
|  o  )|  |  ||      ||      | /  [_ |  D  )
|     ||  |  ||_|  |_||_|  |_||    _]|    / 
|  O  ||  :  |  |  |    |  |  |   [_ |    \ 
|     ||     |  |  |    |  |  |     ||  .  \
|_____| \__,_|  |__|    |__|  |_____||__|\_|
`[1:])
	println("                      proudly presents...")
	println()
	println(" The medicine you have all been waiting for ;)")
	println()
	println()
}

func main() {
	banner()
	if !FileExists(noteFile) {
		println("Put me in the main directory!")
		return
	}

	bytes, _ := ioutil.ReadFile(noteFile)
	var note = string(bytes)
	var offset, _ = strconv.ParseInt(noteShiftRegex.FindStringSubmatch(note)[1], 10, 32)
	var files = SplitNewLine(noteFileListRegex.FindStringSubmatch(note)[1])

	for _, file := range files {
		if strings.Replace(file, " ", "", -1) == "" {
			continue
		}
		if !FileExists(file) || FileExists(file+"_encrypted") {
			continue
		}
		println("Decrypting: " + file)
		decryptFile(file, file+"_decrypted", int(offset))
		os.Rename(file, file+"_encrypted")
		os.Rename(file+"_decrypted", file)
	}
	println("Done!")
}

func decryptFile(inFile string, outFile string, offset int) {
	var key = hex.EncodeToString(decodeKey(offset))
	exec.Command("openssl", "enc", "-d", "-aes-256-cbc", "-in", inFile, "-out", outFile, "-K", key, "-iv", iv).Run()
}

func FileExists(file string) bool {
	_, err := os.Stat(file)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func SplitNewLine(input string) []string {
	return strings.Split(strings.Replace(input, "\r\n", "\n", -1), "\n")
}
