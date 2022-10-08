# tukxdw
Example AWS Lambda imp of an IHE DSUB Consumer Actor grouped with, IHE XDW Content Creator, IHE XDW Content Consumer and IHE XDW Content Updator Actors

Required import github.com/ipthomas/tukint

Required external Web Services :-
  IHE DSUB Broker
  IHE PIXm
  AWS Aurora DB

Processes IHE DSUB Notify messages. When a DSUB Notification is received, the regional patient ID contained in the Notify message is used to obtain the patients NHS ID using the configured IHE PIXm Service. The DSUB notification is then associated with a clinical pathway and an event is then persisted in the 'events' AWS Aurora DB table, containing the XDS meta values provided in the Notify message along with the relevent patients' NHS ID and clinical pathway. Using the associated pathway, a check is made for a current XDW document for the pathway and patient. If no XDW document is present, an new XDW document is created for the pathway and patient and persisted in the 'workflows' AWS Aurora DB table. If an XDW document exists, it is updated with the event details according to the XDW definition for the workflow and a new version is persisted.

XDW Definitions must be created and persisted in the 'xdws' AWS Aurora DB, for each clinical Pathway to be supported. The IHE XDW profile utilises the WS OASIS Human Task standard for the definition of Cross Document Workflow owners, tasks, inputs, outputs, completion behaviours, etc. DSUB Broker Subscriptions must be created for each XDW task output that has a type of '$XDSDocumentEntryTypeCode' in the XDW definition. The resulting broker reference, NHS ID, XDW pathway, topic and expression for each subscription must be persisted in 'subscriptions' AWS Aurora DB table. This enables received notifications to be matched to a specific pathway and a specific task in that pathway. For an example implimentation of a client that reads a config folder containing XDW definition .json files, creates the required DSUB broker subscriptions and persists the XDW definitions and subscriptions with the AWS Aurora DB, refer to github.com/ipthomas/xdwclient# tukxdw_lambda
# tukxdw_lambda
