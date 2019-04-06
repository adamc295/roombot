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
	
	b_session *discordgo.Session
	b_user *discordgo.User
	
	Conf *CoreConfig
	
)

func Init() error {
	
	var err error
	config, err := LoadConfig()
	if err != nil {
		return err
	}
	
	Conf = config
	
	fmt.Println(config.b_ID)
	fmt.Println(Conf.b_ID)
	
	fmt.Println(err)
	b_user, err = b_session.UserMe()
	if err != nil {
		return err
	}	
	
	// Apparently, not using a varible is an error in go's eyes.
	b_session.State.User = &discordgo.SelfUser {
		User: b_user,
	}
	
	return err
}