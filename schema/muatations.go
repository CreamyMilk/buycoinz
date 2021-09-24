package schema

import (
	"log"

	"github.com/CreamyMilk/buycoinz/services"
	"github.com/graphql-go/graphql"
)

var (
	bankVerificationReponseType *graphql.Object
	mutationsTypes              *graphql.Object
)

func setupMuations() {
	bankVerificationReponseType = graphql.NewObject(graphql.ObjectConfig{
		Name: "BankVerificationReponse",
		Fields: graphql.Fields{
			"status": &graphql.Field{
				Type: graphql.Boolean,
			},
			"message": &graphql.Field{
				Type: graphql.String,
			},
		},
	})

	//user_account_number:String!, user_bank_code:String!, user_account_name:String!
	verifyAccountMutation := &graphql.Field{
		Type:        bankVerificationReponseType,
		Description: "Validate's users account number",
		Args: graphql.FieldConfigArgument{
			"user_account_number": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"user_bank_code": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},

			"user_account_name": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			payload := services.BankVerificationRequest{
				UserAccountNumber: params.Args["user_account_number"].(string),
				UserAccountName:   params.Args["user_account_name"].(string),
				UserBankCode:      params.Args["user_bank_code"].(string),
			}
			resp, err := payload.Validate()

			return resp, err
		},
	}

	mutationsTypes = graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"verifyAccount": verifyAccountMutation,
		},
	})
	log.Println("Initalized All Mutations")
}
