package common

import (
	"fmt"
	"github.com/jonas747/discordgo"
)

const (
	v_major = 0
	v_minor = 0
	v_patch = 1
)

var (
	versionnum = fmt.Sprintf("%d.%d.%d", v_major, v_minor, v_patch)
	version = versionnum
	
	b_session = *discordgo.Session
	b_user = *discordgo.User
)

func Init() error {
	
	return
}