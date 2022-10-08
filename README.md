# tukxdw_lambda
Example AWS Lambda imp  IHE XDW Content Creator, IHE XDW Content Consumer and IHE XDW Content Updator Actors

Required external Web Services :-
  IHE DSUB Broker
  IHE PIXm
  MySql DB - connection made using DSN or url.


The IHE XDW profile utilises the WS OASIS Human Task standard for the definition of Cross Document Workflow owners, tasks, inputs, outputs, completion behaviours, etc. DSUB Broker Subscriptions are created for each XDW task input and output that has a type of '$XDSDocumentEntryTypeCode' in the XDW definition. The resulting broker reference, NHS ID, XDW pathway, topic and expression for each subscription is persisted in the tuk event 'subscriptions' AWS Aurora DB table. This enables received notifications to be matched to a specific pathway and a specific task in that pathway. For an example of a client that reads a config folder containing XDW definition .json files, creates the required DSUB broker subscriptions and persists the XDW definitions and subscriptions with the AWS Aurora DB, refer to github.com/ipthomas/tukxdw_client

