package schema

import (
	"log"

	"github.com/CreamyMilk/buycoinz/paystack"
	"github.com/CreamyMilk/buycoinz/services"
	"github.com/CreamyMilk/buycoinz/utils"
	"github.com/graphql-go/graphql"
)

var (
	getAccountNameType *graphql.Object
	bankType           *graphql.Object
	queryFields        graphql.Fields
)

func setupQueries() {
	getAccountNameType = graphql.NewObject(graphql.ObjectConfig{
		Name: "AccountNameResponse",
		Fields: graphql.Fields{
			"status": &graphql.Field{
				Type:        graphql.Boolean,
				Description: "States whether the supplied Account Details are valid",
			},
			"message": &graphql.Field{
				Type:        graphql.String,
				Description: "Gives a human readable response to the user",
			},
			"accountName": &graphql.Field{
				Type:        graphql.String,
				Description: "Stored Account Name formated in Sentence Case",
			},
		},
	})

	bankType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "Bank",
		Description: "A represention of the bank object",
		Fields: graphql.Fields{
			"name": &graphql.Field{
				Type:        graphql.String,
				Description: "Gives the full name of the bank",
			},
			"code": &graphql.Field{
				Type:        graphql.String,
				Description: "Bank Code is a number used to identify a bank",
			},
			"active": &graphql.Field{
				Type:        graphql.Boolean,
				Description: "States whether the bank is active",
			},
		},
	})

	getPaystackBankList := &graphql.Field{
		Type:        graphql.NewList(bankType),
		Description: "Returns a list of banks with their respective metadata",
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			banks, err := paystack.GetBanksList()
			return banks, err
		},
	}

	getPreferedAccountNameQuery := &graphql.Field{
		Type:        getAccountNameType,
		Description: "Returns the preferred user Account Name",
		Args: graphql.FieldConfigArgument{
			"account_number": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"bank_code": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			payload := services.PreferedAccountNameRequest{
				AccountNumber: params.Args["account_number"].(string),
				BankCode:      params.Args["bank_code"].(string),
			}
			resp, err := payload.GetPreferedAccountName()

			//Formats the Name back to Sentense case
			resp.AccountName = utils.ToSentenceCase(resp.AccountName)
			return resp, err
		},
	}

	queryFields = graphql.Fields{
		"getBanks":               getPaystackBankList,
		"getPreferedAccountName": getPreferedAccountNameQuery,
	}
	log.Println("Initalized All Queries")
}
