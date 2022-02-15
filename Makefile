# Ledger address
ADDRESS=h8pv3bzsgre4c8gi.daml.app:443
# Party on behalf of which the app runs, can be copied from Daml Hub.
# Admin in the POC choice observer demo.
PARTY=ledger-party-d784fcf1-eaf4-4ad3-8482-47f3905d0749
# Access token for the party on behalf of which the app runs.
# Can be copied from Daml Hub, expires every 24 hours.
JWT=eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImtpZCI6ImRhYmwtYzFlNzUwMmEtODU1MC00MTRjLTg4YmQtMWVmNDIxODBiOTdlIn0.eyJpc3MiOiJodWIuZGFtbC5jb20vbG9naW4iLCJzdWIiOiJhdXRoMHw2MTkyOWIzMjRlMmJmNTAwNmZlNTJmNjAiLCJleHAiOjE2NDQ3NTA3OTcsImp0aSI6IjNiNDUyYzhhLTE1MDgtNDcwZS04ZDgwLTdmZTkzMmU4MTkyOSIsImh0dHBzOi8vZGFtbC5jb20vbGVkZ2VyLWFwaSI6eyJhY3RBcyI6WyJsZWRnZXItcGFydHktZDc4NGZjZjEtZWFmNC00YWQzLTg0ODItNDdmMzkwNWQwNzQ5Il0sImFwcGxpY2F0aW9uSWQiOiJkYW1saHViIiwibGVkZ2VySWQiOiJoOHB2M2J6c2dyZTRjOGdpIiwicmVhZEFzIjpbImxlZGdlci1wYXJ0eS1kNzg0ZmNmMS1lYWY0LTRhZDMtODQ4Mi00N2YzOTA1ZDA3NDkiXX0sImxlZGdlcklkIjoiaDhwdjNienNncmU0YzhnaSIsIm93bmVyIjoidXNlci1ncmFudC05N2EyMjAyNi1jYzZhLTQ1YWQtOTM1Ny0xNjA2MDcyYTUxMzQiLCJwYXJ0eSI6ImxlZGdlci1wYXJ0eS1kNzg0ZmNmMS1lYWY0LTRhZDMtODQ4Mi00N2YzOTA1ZDA3NDkiLCJwYXJ0eU5hbWUiOiJBZG1pbiJ9.d77Bc6CmS7d3m3NV3soi4NKXQYmPSDDRSMwgObKu5zHv5Sv54NPlvhylT4WYe3_hbIKgd1mQO_xqeGjSyJ1q759LF29VuqDJXZ0M3_wzxPnJHzkdHmcqZfCW6U118AS-D6DIFLR_Dyi-gWy_yhvelz_L75mq8sIxV8zBv1l5eL0
# TLS cert for Daml Hub, can be downloaded from here: https://letsencrypt.org/certs/isrgrootx1.pem
CA_FILE=./isrgrootx1.pem
# Package id of the DALF package uploaded to Daml Hub.
PACKAGE_ID=cceb368817c7ceefca94498392cb2e06bd194ad3280aa1d45d0f96e35bfdc55e
# The module name for the choice
MODULE_NAME=Workflow
# The template name for the choice
TEMPLATE_NAME=Policy
# Choice name used to call out
CALL_OUT_CHOICE_NAME=Owner_GetPhoneFromExternal
CALL_BACK_CHOICE_NAME=Admin_UpdatePhoneFromExternal
CALL_BACK_INPUT_FIELD=phone
# Will be submitted to the ledger as `Some newPhone` of type `Optional Text`
CALL_BACK_INPUT_TEXT=newPhone

start :
	go run go_for_daml.go \
		-access_token=${JWT} \
		-addr=${ADDRESS} \
		-application_id=damlhub \
		-ca_file=${CA_FILE} \
		-module_name=${MODULE_NAME} \
		-package_id=${PACKAGE_ID} \
		-party=${PARTY} \
		-template_name=${TEMPLATE_NAME} \
		-tls=true \
		-workflow_id=choice_observer_demo_workflow \
		-call_out_choice_name=${CALL_OUT_CHOICE_NAME} \
		-call_back_choice_name=${CALL_BACK_CHOICE_NAME} \
		-call_back_input_field=${CALL_BACK_INPUT_FIELD} \
		-call_back_input_text=${CALL_BACK_INPUT_TEXT}

# Will print out the application flags
help : 
	go run go_for_daml.go --help
