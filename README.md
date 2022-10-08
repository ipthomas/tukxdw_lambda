# tukxdw_lambda
Example AWS Lambda imp  IHE XDW Content Creator, IHE XDW Content Consumer and IHE XDW Content Updator Actors

Required external Web Services :-
  IHE DSUB Broker
  IHE PIXm
  MySql DB - connection made using DSN or url.

The IHE XDW profile utilises the WS OASIS Human Task standard for the definition of Cross Document Workflow owners, tasks, inputs, outputs, completion behaviours, etc.

The current implementation supports the registering of a XDW definition with the TUK Event Service. The registering process creates DSUB Broker Subscriptions for each XDW task input and output that has a type of '$XDSDocumentEntryTypeCode' in the XDW definition. The resulting broker reference, NHS ID, XDW pathway, topic and expression for each subscription is persisted in the tuk event 'subscriptions' TUK Event service subscriptions DB table. This enables received notifications from a DSUB Broker to be matched to a specific pathway and a specific task in that pathway.

For an example implementation of a DSUB Broker Event Consumer that persists IHE DSUB Notify messages to the TUK Event Service database refer to github.com/ipthomas/tukdsub and github/ipthomas/tukdsub_lambda.
