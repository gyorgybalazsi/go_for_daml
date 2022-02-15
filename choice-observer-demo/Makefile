
LEDGER_ID=dczj5vgug00zejgt
DAML_LEDGER_PARTY=ledger-party-2af66bb1-9066-470e-8877-5d24a05ce852
TOKEN=eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImtpZCI6ImRhYmwtM2MzZDM2ZDUtYjQyNy00OTNiLTk5YTctZTU3ZmI4ZmZmZDQ4In0.eyJpc3MiOiJodWIuZGFtbC5jb20vbG9naW4iLCJzdWIiOiJhdXRoMHw2MTkyOWIzMjRlMmJmNTAwNmZlNTJmNjAiLCJleHAiOjE2NDQ0Mjk1NTYsImp0aSI6ImYzZGE3ZWM5LTIzMWUtNDk4ZC1hMTYzLTJkZTZjM2MwODQ2MCIsImh0dHBzOi8vZGFtbC5jb20vbGVkZ2VyLWFwaSI6eyJhY3RBcyI6WyJsZWRnZXItcGFydHktMmFmNjZiYjEtOTA2Ni00NzBlLTg4NzctNWQyNGEwNWNlODUyIl0sImFwcGxpY2F0aW9uSWQiOiJkYW1saHViIiwibGVkZ2VySWQiOiJkY3pqNXZndWcwMHplamd0IiwicmVhZEFzIjpbImxlZGdlci1wYXJ0eS0yYWY2NmJiMS05MDY2LTQ3MGUtODg3Ny01ZDI0YTA1Y2U4NTIiXX0sImxlZGdlcklkIjoiZGN6ajV2Z3VnMDB6ZWpndCIsIm93bmVyIjoidXNlci1ncmFudC01YWViYTRkYS0xYTg5LTQxNjgtOGY1Yy04OTAyZWI0OTMzNWMiLCJwYXJ0eSI6ImxlZGdlci1wYXJ0eS0yYWY2NmJiMS05MDY2LTQ3MGUtODg3Ny01ZDI0YTA1Y2U4NTIiLCJwYXJ0eU5hbWUiOiJBbGljZSJ9.YtFjcVZiA_6ddqsrAyi4ptoEz_skBLOHI5nLyrVWL7wonl3rpn0iVN4sl1n1SRj3N-AkZUs8u3jIJrKpIiQnfnZyg82p3xkypgdxfqtPesZlplbD_pIrId9eJFtUVkrvHSELaqghQJ1aqPLMdH5YTNC7S4UDOmiiJDP72q8Orxo

SANDBOX_LEDGER_ID=c9c1e1eb-0586-49eb-9d66-9ce0f32d7609
SANDBOX_PARTY=Alice

run_on_sandbox :
	grpcurl \
		-d "{\"ledgerId\":\"${SANDBOX_LEDGER_ID}\", \"filter\":{\"filtersByParty\":[{\"key\": \"${SANDBOX_PARTY}\"}]}, \"begin\": {\"boundary\": \"LEDGER_BEGIN\"}}" \
		-plaintext \
		"localhost:6865" \
		com.daml.ledger.api.v1.TransactionService/GetTransactionTrees > transactiontrees.txt

# boundary = LEDGER_BEGIN - full tx history
# boundary = LEDGER_END - txs after starting the client
run_listener :
	grpcurl \
		-keepalive-time 60 \
		-H "Authorization: Bearer ${TOKEN}" \
		-d "{\"ledgerId\":\"${LEDGER_ID}\", \"filter\":{\"filtersByParty\":[{\"key\": \"${DAML_LEDGER_PARTY}\"}]}, \"begin\": {\"boundary\": \"LEDGER_END\"}}" \
		"${LEDGER_ID}.daml.app:443" \
		com.daml.ledger.api.v1.TransactionService/GetTransactionTrees > transactiontrees.txt


OFFSET="000000000000000d"
run_listener_from_offset :
	grpcurl \
		-keepalive-time 60 \
		-H "Authorization: Bearer ${TOKEN}" \
		-d "{\"ledgerId\":\"${LEDGER_ID}\", \"filter\":{\"filtersByParty\":[{\"key\": \"${DAML_LEDGER_PARTY}\"}]}, \"begin\": {\"absolute\": \"${OFFSET}\"}}" \
		"${LEDGER_ID}.daml.app:443" \
		com.daml.ledger.api.v1.TransactionService/GetTransactionTrees > transactiontrees.txt