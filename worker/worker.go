package worker

import (
	"github.com/solthoth/like-my-wife/configuration"
	jww "github.com/spf13/jwalterweatherman"
)

func Work(config *configuration.Configurations)  {
	wife := config.SocialMediaIdentifier.Instagram
	jww.INFO.Printf("Hello %s\n", wife)
}