package worker

import (
	"log"

	"github.com/solthoth/like-my-wife/configuration"
)

func Work(config *configuration.Configurations)  {
	var wife = config.SocialMediaIdentifier.Instagram
	log.Printf("Hello %s\n", wife)
}