package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
	"time"

	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"

	"github.com/google/uuid"
	pb "github.com/gyorgybalazsi/go_daml/com/daml/ledger/api/v1"
	grpc "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

//HELPER FUNCTIONS

//Unit to be used as nil value in variants
var unit = &pb.Value{Sum: &pb.Value_Unit{Unit: &emptypb.Empty{}}}

//Empty record to be used as nil value on commands without input
var emptyRecord = &pb.Value{Sum: &pb.Value_Record{Record: &pb.Record{}}}

//Single field record
func singleFieldRecord(label string, value *pb.Value) *pb.Record {
	return &pb.Record{
		Fields: []*pb.RecordField{{Label: label, Value: value}},
	}
}

//Functions to build value object from record, int, text, party, variant

func singleFieldRecordValue(label string, value *pb.Value) *pb.Value {
	return &pb.Value{Sum: &pb.Value_Record{Record: singleFieldRecord(label, value)}}
}

func recordValue(recordMap map[string]*pb.Value) *pb.Value {
	fields := []*pb.RecordField{}
	for k, v := range recordMap {
		fields = append(fields, &pb.RecordField{
			Label: k,
			Value: v,
		})
	}
	record := &pb.Record{
		Fields: fields,
	}
	return &pb.Value{Sum: &pb.Value_Record{Record: record}}

}

func createDouble(left *pb.Value, right *pb.Value) *pb.Value {
	m := make(map[string]*pb.Value)
	m["_1"] = left
	m["_2"] = right
	return recordValue(m)
}

//func printRecordValue

func intValue(value int64) *pb.Value {
	return &pb.Value{Sum: &pb.Value_Int64{Int64: value}}
}

func printIntValue(value *pb.Value) string {
	return strconv.FormatInt(value.GetInt64(), 10)
}

func textValue(value string) *pb.Value {
	return &pb.Value{Sum: &pb.Value_Text{Text: value}}
}

func printTextValue(value *pb.Value) string {
	return value.GetText()
}

func variantValue(value string) *pb.Value {
	return &pb.Value{
		Sum: &pb.Value_Variant{
			Variant: &pb.Variant{
				Constructor: value,
				Value:       unit,
			},
		},
	}
}

func printVariantValue(value *pb.Value) string {
	return value.GetVariant().GetConstructor()
}

func partyValue(value string) *pb.Value {
	return &pb.Value{Sum: &pb.Value_Party{Party: value}}
}

func printPartyValue(value *pb.Value) string {
	return value.GetParty()
}

func optionalSomeTextValue(value string) *pb.Value {
	return &pb.Value{
		Sum: &pb.Value_Optional{
			Optional: &pb.Optional{
				Value: textValue(value),
			},
		},
	}
}

//Input builders, to be used as ChoiceArgument value input

//Integer input builder
func integerInput(label string, value int64) *pb.Value {
	return singleFieldRecordValue(label, intValue(value))
}

//Text input builder
func textInput(label string, value string) *pb.Value {
	return singleFieldRecordValue(label, textValue(value))
}

//Variant input builder
func variantInput(label string, value string) *pb.Value {
	return singleFieldRecordValue(label, variantValue(value))
}

func partyInput(label string, value string) *pb.Value {
	return singleFieldRecordValue(label, partyValue(value))
}

func optinalSomeTextInput(label string, value string) *pb.Value {
	return singleFieldRecordValue(label, optionalSomeTextValue(value))
}

//Function gets ledger id
func getLedgerID(ctx context.Context, cc *grpc.ClientConn) string {
	ledgerIDClient := pb.NewLedgerIdentityServiceClient(cc)

	ledgerID, err := ledgerIDClient.GetLedgerIdentity(ctx, &pb.GetLedgerIdentityRequest{})

	if err != nil {
		log.Fatalf("Could not get ledger ID: %v", err)
	}
	return ledgerID.LedgerId
}

//Function gets active package ids
func getPackageIDs(ctx context.Context, cc *grpc.ClientConn, ledgerID string) []string {
	packageIDClient := pb.NewPackageServiceClient(cc)
	packageIDs, err := packageIDClient.ListPackages(ctx, &pb.ListPackagesRequest{LedgerId: ledgerID})
	if err != nil {
		log.Fatalf("Could not get package IDs: %v", err)
	}
	return packageIDs.PackageIds
}

//Function prints IDs of active contracts available to a party
func printActiveContractsOfParty(ctx context.Context, cc *grpc.ClientConn, ledgerID string, party string) {

	filterMap := map[string]*pb.Filters{party: &pb.Filters{}}
	activeContractsClient := pb.NewActiveContractsServiceClient(cc)

	log.Printf("Active contracts available to %s", party)
	log.Println()

	stream, err := activeContractsClient.GetActiveContracts(
		ctx,
		&pb.GetActiveContractsRequest{
			LedgerId: ledgerID,
			Filter:   &pb.TransactionFilter{FiltersByParty: filterMap},
			Verbose:  true,
		},
	)
	if err != nil {
		log.Fatalf("Could not get active contracts: %v", err)
	}

	for {
		contractResponse, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error: %v, %v", activeContractsClient, err)
		}
		for _, contract := range contractResponse.ActiveContracts {
			log.Println(contract.GetContractId())
		}
	}
}

func submitAndGetResponse(
	ctx context.Context,
	commandClient pb.CommandServiceClient,
	ledgerID string,
	workflowID string,
	packageID string,
	contractID string,
	moduleName string,
	templateName string,
	choice string,
	input *pb.Value, //integerInput | textInput | variantInput | partyInput
	submittingParty string,
	applicationID string,
	commandID string,
	timeHopInSeconds int64) (*pb.SubmitAndWaitForTransactionResponse, error) {
	exerciseCommand := pb.ExerciseCommand{
		TemplateId: &pb.Identifier{
			PackageId:  packageID,
			ModuleName: moduleName,
			EntityName: templateName,
		},
		ContractId:     contractID,
		Choice:         choice,
		ChoiceArgument: input,
	}

	command := &pb.Command{
		Command: &pb.Command_Exercise{Exercise: &exerciseCommand},
	}

	commands := &pb.Commands{
		LedgerId:      ledgerID,
		WorkflowId:    workflowID,
		ApplicationId: applicationID,
		CommandId:     commandID,
		Party:         submittingParty,
		//LedgerEffectiveTime: &timestamppb.Timestamp{Seconds: (timeHopInSeconds + 1)},
		//MaximumRecordTime:   &timestamppb.Timestamp{Seconds: (timeHopInSeconds + 10)},
		Commands: []*pb.Command{command},
	}

	request := &pb.SubmitAndWaitRequest{
		Commands: commands,
	}

	response, err := commandClient.SubmitAndWaitForTransaction(ctx, request)

	if err != nil {
		log.Fatal("Source: submitAndGetResponse - ", err)
	}

	return response, err
}

func submitAndGetContractID(
	ctx context.Context,
	commandClient pb.CommandServiceClient,
	ledgerID string,
	workflowID string,
	packageID string,
	contractID string,
	moduleName string,
	templateName string,
	choice string,
	input *pb.Value, //integerInput | variantInput | partyInput
	submittingParty string,
	applicationID string,
	commandID string,
	timeHopInSeconds int64) (string, error) {
	exerciseCommand := pb.ExerciseCommand{
		TemplateId: &pb.Identifier{
			PackageId:  packageID,
			ModuleName: moduleName,
			EntityName: templateName,
		},
		ContractId:     contractID,
		Choice:         choice,
		ChoiceArgument: input,
	}

	command := &pb.Command{
		Command: &pb.Command_Exercise{Exercise: &exerciseCommand},
	}

	commands := &pb.Commands{
		LedgerId:      ledgerID,
		WorkflowId:    workflowID,
		ApplicationId: applicationID,
		CommandId:     commandID,
		Party:         submittingParty,
		//LedgerEffectiveTime: &timestamppb.Timestamp{Seconds: (timeHopInSeconds + 1)},
		//MaximumRecordTime:   &timestamppb.Timestamp{Seconds: (timeHopInSeconds + 10)},
		Commands: []*pb.Command{command},
	}

	request := &pb.SubmitAndWaitRequest{
		Commands: commands,
	}

	response, err := commandClient.SubmitAndWaitForTransaction(ctx, request)

	if err != nil {
		log.Fatal("Error: %v", err)
	}

	return getContractID(response.GetTransaction()), err
}

func submitExerciseAndGetContractId2(
	ctx context.Context,
	commandClient pb.CommandServiceClient,
	ledgerID string,
	workflowID string,
	packageID string,
	moduleName string,
	templateName string,
	contractId string,
	choice string,
	input *pb.Value, //integerInput | variantInput | partyInput
	submittingParty string,
	applicationID string,
	commandID string) (string, error) {
	exerciseCommand := pb.ExerciseCommand{
		TemplateId: &pb.Identifier{
			PackageId:  packageID,
			ModuleName: moduleName,
			EntityName: templateName,
		},
		ContractId:     contractId,
		Choice:         choice,
		ChoiceArgument: input,
	}
	command := &pb.Command{
		Command: &pb.Command_Exercise{Exercise: &exerciseCommand},
	}

	commands := &pb.Commands{
		LedgerId:      ledgerID,
		WorkflowId:    workflowID,
		ApplicationId: applicationID,
		CommandId:     commandID,
		Party:         submittingParty,
		//LedgerEffectiveTime: &timestamppb.Timestamp{Seconds: (timeHopInSeconds + 1)},
		//MaximumRecordTime:   &timestamppb.Timestamp{Seconds: (timeHopInSeconds + 10)},
		Commands: []*pb.Command{command},
	}

	request := &pb.SubmitAndWaitRequest{
		Commands: commands,
	}

	response, err := commandClient.SubmitAndWaitForTransaction(ctx, request)

	if err != nil {
		log.Fatal("Error: %v", err)
	}

	return getContractID(response.GetTransaction()), err
}

func printTransaction(tx *pb.Transaction) {

	var created *pb.CreatedEvent
	var archived *pb.ArchivedEvent

	for _, e := range tx.GetEvents() {
		if e.GetCreated() != nil {
			created = e.GetCreated()
		} else if e.GetArchived() != nil {
			archived = e.GetArchived()
		}
	}

	contractID := created.GetContractId()

	templateID := created.GetTemplateId()

	witnessParties := created.GetWitnessParties()

	res := []string{
		"\n\n" + contractID + " CONTRACT CREATED (" + templateID.GetEntityName() + ")",
		"\n**Contract details**\n",
		"Effective at: " + tx.GetEffectiveAt().String(),
		"Command ID: " + tx.GetCommandId(),
		"Workflow ID: " + tx.GetWorkflowId(),
		"Witness parties: " + strings.Join(witnessParties, ","),
		"Contract ID: " + contractID,
		"Template ID: " + templateID.GetEntityName() + "@" + templateID.GetPackageId(),
		"Offset: " + tx.GetOffset(),
	}

	if archived != nil {
		res = append(res, "\n**Archived contract**\n")
		res = append(res, "Contract ID: "+archived.GetContractId())
	}

	res = append(res, "\n**Creation arguments**\n")

	fields := created.GetCreateArguments().GetFields()

	for _, v := range fields {
		switch fmt.Sprintf("%T", v.GetValue().GetSum()) {
		case "*com_digitalasset_ledger_api_v1.Value_Text":
			res = append(res, v.GetLabel()+": "+printTextValue(v.GetValue()))
		case "*com_digitalasset_ledger_api_v1.Value_Party":
			res = append(res, v.GetLabel()+": "+printPartyValue(v.GetValue()))
		case "*com_digitalasset_ledger_api_v1.Value_Variant":
			res = append(res, v.GetLabel()+": "+printVariantValue(v.GetValue()))
		//Only implemented for list of parties
		case "*com_digitalasset_ledger_api_v1.Value_List":
			if fmt.Sprintf("%T", v.GetValue().GetList().GetElements()[0].GetSum()) != "*com_digitalasset_ledger_api_v1.Value_Party" {
				res = append(res, "unknown")
			}
			listElements := []string{}
			for _, element := range v.GetValue().GetList().GetElements() {
				listElements = append(listElements, printPartyValue(element))
			}
			res = append(res, v.GetLabel()+": "+strings.Join(listElements, ","))
		default:
			res = append(res, "unknown")
		}

	}

	joined := strings.Join(res, "\n")

	log.Println(joined)
}

func getContractID(tx *pb.Transaction) string {
	var event *pb.Event

	for _, e := range tx.GetEvents() {
		if e.GetCreated() != nil {
			event = e
		}
	}

	return event.GetCreated().GetContractId()
}

func printTransactionTrees(ctx context.Context, cc *grpc.ClientConn, ledgerID string, party string) {
	transactionServiceClient := pb.NewTransactionServiceClient(cc)
	filterMap := map[string]*pb.Filters{party: &pb.Filters{}}
	// TODO : make Begin a flag
	in := &pb.GetTransactionsRequest{
		LedgerId: ledgerID,
		Begin:    &pb.LedgerOffset{Value: &pb.LedgerOffset_Boundary{Boundary: pb.LedgerOffset_LEDGER_END}},
		Filter:   &pb.TransactionFilter{FiltersByParty: filterMap},
	}
	stream, err := transactionServiceClient.GetTransactionTrees(ctx, in)
	if err != nil {
		log.Fatalf("Could not get transaction trees stream: %v", err)
	}
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.GetTransactionTrees(_) = _, %v", transactionServiceClient, err)
		}
		log.Println(msg.String())
		if len(msg.Transactions) == 0 {
			log.Println("Transaction list is empty")
			break
		}
		eventsById := msg.Transactions[0].GetEventsById()
		for _, event := range eventsById {
			if event.GetCreated() != nil {
				log.Printf("Found a created event")
			}
			if event.GetExercised() != nil {
				log.Printf("Found an exercise event")
				if event.GetExercised().Choice == "CallOut" {
					log.Println("Found an exercise of `Callout`")
					log.Println(event.GetExercised().ChoiceArgument.GetRecord().Fields[0].GetValue().GetText())
				}
			}
		}
	}
}

// THE FUNCTION USED IN THE CHOICE OBSERVER DEMO

func getChoiceExerciseMessageAndAnswer(
	ctx context.Context,
	transactionServiceClient pb.TransactionServiceClient,
	commandServiceClient pb.CommandServiceClient,
	ledgerID string,
	contractId string) {

	filterMap := map[string]*pb.Filters{*party: {}}
	// TODO : make Begin a flag
	getTransactionsRequest := &pb.GetTransactionsRequest{
		LedgerId: ledgerID,
		Begin:    &pb.LedgerOffset{Value: &pb.LedgerOffset_Boundary{Boundary: pb.LedgerOffset_LEDGER_END}},
		Filter:   &pb.TransactionFilter{FiltersByParty: filterMap},
	}
	stream, err := transactionServiceClient.GetTransactionTrees(ctx, getTransactionsRequest)
	if err != nil {
		log.Fatalf("Could not get transaction trees stream: %v", err)
	}
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.GetTransactionTrees(_) = _, %v", transactionServiceClient, err)
		}
		log.Println(msg.String())
		if len(msg.Transactions) == 0 {
			log.Println("Transaction list is empty")
			break
		}
		eventsById := msg.Transactions[0].GetEventsById()
		for _, event := range eventsById {
			if event.GetCreated() != nil {
				log.Printf("Found a created event")
			}
			if event.GetExercised() != nil {
				log.Printf("Found an exercise event")
				if event.GetExercised().Choice == *call_out_choice_name {
					log.Printf("Found an exercise of %v", *call_out_choice_name)
					message := event.GetExercised().ChoiceArgument.GetRecord().Fields[0].GetValue().GetText()
					log.Printf("The incoming message is: %v\n", message)
					var command_id = uuid.New().String()
					cid, err := submitExerciseAndGetContractId2(
						ctx,                    // ctx context.Context,
						commandServiceClient,   //commandClient pb.CommandServiceClient,
						ledgerID,               //ledgerID string,
						*workflow_id,           //workflowID string,
						*package_id,            //packageID string,
						*module_name,           //moduleName string,
						*template_name,         //templateName string,
						contractId,             // contractId string,
						*call_back_choice_name, //choice string,
						optinalSomeTextInput(*call_back_input_field, *call_back_input_text), //input *pb.Value, //integerInput | variantInput | partyInput
						*party,          //submittingParty string,
						*application_id, //applicationID string,
						command_id)      //commandID string)
					if err != nil {
						log.Fatalf("Error %v", err)
					}
					if err == nil {
						log.Printf("Answer submitted, cid: %v\n", cid)
					}

				}
			}
		}
	}
}

var (
	addr                  = flag.String("addr", "localhost:6865", "Ledger address")
	party                 = flag.String("party", "Alice", "Ledger party on behalf of the app runs")
	access_token          = flag.String("access_token", "", "Access token of the party on behalf of the app runs")
	ca_file               = flag.String("ca_file", "", "The file containing the CA root cert file")
	tls                   = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	application_id        = flag.String("application_id", "damlhub", "Application id, default `damlhub")
	workflow_id           = flag.String("workflow_id", "choice_observer_demo_workflow", "Workflow id")
	package_id            = flag.String("package_id", "", "Package id of the Daml package")
	module_name           = flag.String("module_name", "", "Module name in which you want to exercise a choice")
	template_name         = flag.String("template_name", "", "Template name on which you want to exercise a choice")
	call_out_choice_name  = flag.String("call_out_choice_name", "", "Choice name used to call out")
	call_back_choice_name = flag.String("call_back_choice_name", "", "Choice name used to call back")
	call_back_input_field = flag.String("call_back_input_field", "message", "Input field name, provided that the input consist of one text field")
	call_back_input_text  = flag.String("call_back_input_text", "Hello World", "Input text, provided that the input consist of one text field")
)

type AuthToken struct {
	Token string
}

func (c AuthToken) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"Authorization": "Bearer " + c.Token,
	}, nil
}

func (c AuthToken) RequireTransportSecurity() bool {
	return false
}

func main() {
	flag.Parse()

	// SET UP CONNECTION

	var opts []grpc.DialOption
	if *tls {
		creds, err := credentials.NewClientTLSFromFile(*ca_file, "")
		if err != nil {
			log.Fatalf("Failed to create TLS credentials %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))

		auth_token := AuthToken{Token: *access_token}

		opts = append(opts, grpc.WithPerRPCCredentials(auth_token))
	} else {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}
	var kp = keepalive.ClientParameters{
		Time:                30 * time.Second,
		Timeout:             20 * time.Second,
		PermitWithoutStream: true,
	}
	opts = append(opts, grpc.WithKeepaliveParams(kp))
	conn, err := grpc.Dial(*addr, opts...)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// CREATE CONTEXT

	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	log.Println("Connected")

	//commandClient := pb.NewCommandServiceClient(cc)

	ledgerID := getLedgerID(ctx, conn)
	log.Println("Ledger Id = ", ledgerID)

	//getChoiceExerciseAndAnswer(ctx, conn, ledgerID)

	// CREATE SERCVICE CLIENTS

	activeContractsServiceClient := pb.NewActiveContractsServiceClient(conn)
	transactionServiceClient := pb.NewTransactionServiceClient(conn)
	commandServiceClient := pb.NewCommandServiceClient(conn)

	filterForAsset := map[string]*pb.Filters{
		*party: {
			Inclusive: &pb.InclusiveFilters{
				TemplateIds: []*pb.Identifier{{
					PackageId:  *package_id,
					ModuleName: *module_name,
					EntityName: *template_name,
				},
				},
			},
		},
	}

	getActiveContractSetRequest := &pb.GetActiveContractsRequest{
		LedgerId: ledgerID,
		Filter:   &pb.TransactionFilter{FiltersByParty: filterForAsset},
	}

	acl, err := activeContractsServiceClient.GetActiveContracts(ctx, getActiveContractSetRequest)

	if err != nil {
		log.Printf("Error creating get active contracts clent: %v\n", err)
	}

	acs, err := acl.Recv()

	if err != nil {
		log.Printf("Error receiving active contract set: %v\n", err)
	}

	if len(acs.ActiveContracts) == 0 {
		log.Println("There are no active Policy contracts on the ledger")
		return
	}

	contractId := acs.ActiveContracts[0].ContractId

	log.Println(contractId)

	getChoiceExerciseMessageAndAnswer(
		ctx,
		transactionServiceClient,
		commandServiceClient,
		ledgerID,
		contractId,
	)

}
