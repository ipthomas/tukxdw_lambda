package main

import (
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/ipthomas/tukcnst"
	"github.com/ipthomas/tukxdw"
)

func main() {
	lambda.Start(Handle_Request)
}
func Handle_Request(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var err error
	if wfdef, err := tukxdw.NewWorkflowDefinition(req.Body, nil); err == nil {
		trans := tukxdw.XDWTransaction{
			Action:             tukcnst.REGISTER,
			WorkflowDefinition: wfdef,
			Pathway:            wfdef.Ref,
			DSUB_BrokerURL:     os.Getenv(tukcnst.AWS_ENV_DSUB_BROKER_URL),
			DSUB_ConsumerURL:   os.Getenv(tukcnst.AWS_ENV_DSUB_CONSUMER_URL),
		}
		tukxdw.New_Transaction(&trans)
	}
	return handle_Response(err)
}
func handle_Response(err error) (*events.APIGatewayProxyResponse, error) {
	resp := events.APIGatewayProxyResponse{}
	if err == nil {
		resp.StatusCode = http.StatusOK
		resp.Body = ""
	} else {
		resp.StatusCode = http.StatusInternalServerError
		resp.Body = err.Error()
		log.Println(err.Error())
	}
	return &resp, err
}
