package cmd

import (
	"github.com/spf13/cobra"

	log "github.com/sirupsen/logrus"

	"fmt"
	"net/url"

	"github.com/fabric8-services/fabric8-jenkins-proxy/internal/util"
)

var (
	cmdToken = &cobra.Command{
		Use:   "token",
		Short: "Prints the OSIO token for a specified user UUID.",
		Long:  `Prints the OSIO token for a specified user UUID.`,
		Run:   runCreateToken,
	}
	privateKey   string
	privateKeyID string
	uuid         string
	session      string
	validFor     int64
)

func init() {
	cmdToken.Flags().StringVarP(&privateKey, "key", "k", "", "Private key.")
	cmdToken.Flags().StringVarP(&privateKeyID, "id", "i", "", "Private key id (optional).")
	cmdToken.Flags().StringVarP(&uuid, "uuid", "u", "", "The users uuid.")
	cmdToken.Flags().StringVarP(&session, "session", "s", "", "A session state string (optional).")
	cmdToken.Flags().Int64VarP(&validFor, "valid", "v", 60, "Time token is valid in minutes (default 60).")
}

func runCreateToken(cmd *cobra.Command, args []string) {
	token, err := util.CreateOSIOToken(targetEnv, uuid, privateKey, privateKeyID, validFor, session)
	if err != nil {
		log.Fatal(err)
	}

	if encode {
		token = url.QueryEscape(token)
	}
	fmt.Println(token)
}
