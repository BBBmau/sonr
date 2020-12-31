package ui

import (
	"bytes"
	"image/png"
	"log"
	"os"

	md "github.com/sonr-io/core/internal/models"
)

// ** Icon Type for Image
type Icon int

const (
	SystemTray Icon = iota
	Close
	User
	Peer
	Invite
	iPhone
	Android
	Mac
	Windows
)

// ** Const UI Resource Path ** //
const RES_PATH = "/Users/prad/Sonr/core/pkg/res/"
const ICON_PATH = "/Users/prad/Sonr/core/pkg/res/systray.png"

func (d Icon) File() string {
	return [...]string{"systray.png", "close.png", "user.png", "peer.png", "invite.png", "iphone.png", "android.png", "mac.png", "windows.png"}[d]
}

// ^ Returns Buffer of Image by Icon Type
func GetIcon(i Icon) []byte {
	// Get Path
	path := RES_PATH + i.File()

	// Initialize
	imgBuffer := new(bytes.Buffer)
	existingImageFile, err := os.Open(path)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer existingImageFile.Close()

	// Decode PNG
	loadedImage, err := png.Decode(existingImageFile)
	if err != nil {
		log.Println(err)
		return nil
	}
	// Encode PNG into Memory
	err = png.Encode(imgBuffer, loadedImage)
	if err != nil {
		log.Println(err)
		return nil
	}
	return imgBuffer.Bytes()
}

// ^ Returns Buffer of Image from Device Type
func GetDeviceIcon(d *md.Device) []byte {
	if d.Platform == "Android" {
		return GetIcon(Android)
	} else if d.Platform == "iOS" {
		return GetIcon(iPhone)
	} else if d.Platform == "Mac" {
		return GetIcon(Mac)
	} else if d.Platform == "Windows" {
		return GetIcon(Windows)
	}
	return nil
}