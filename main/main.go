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
	trans := tukxdw.XDWTransaction{DSUB_BrokerURL: os.Getenv(tukcnst.AWS_ENV_DSUB_BROKER_URL), DSUB_ConsumerURL: os.Getenv(tukcnst.AWS_ENV_DSUB_CONSUMER_URL)}
	switch req.QueryStringParameters[tukcnst.QUERY_PARAM_ACTION] {
	case tukcnst.REGISTER:
		if wfdef, err := tukxdw.NewWorkflowDefinition(req.Body, nil); err == nil {
			trans.Action = tukcnst.REGISTER
			trans.WorkflowDefinition = wfdef
			trans.Pathway = wfdef.Ref
		}
	case tukcnst.XDW_CONTENT_CONSUMER:
		trans.Action = tukcnst.XDW_CONTENT_CONSUMER
		trans.Pathway = req.QueryStringParameters[tukcnst.QUERY_PARAM_PATHWAY]
	}
	return handle_Response(tukxdw.New_Transaction(&trans))
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
