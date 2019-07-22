package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
	"time"

	pb "go_daml/com/digitalasset/ledger/api/v1"

	grpc "google.golang.org/grpc"

	protobuf "google/protobuf"
)

const (
	address = "localhost:7600"
	//Package ID needs to be updated when DAML template set changes
)

//HELPER FUNCTIONS

//Unit to be used as nil value in variants
var unit = &pb.Value{Sum: &pb.Value_Unit{Unit: &protobuf.Empty{}}}

//Empty record to be used as nil value on commands without input
var emptyRecord = &pb.Value{Sum: &pb.Value_Record{Record: &pb.Record{}}}

//Single field record
func singleFieldRecord(label string, value *pb.Value) *pb.Record {
	return &pb.Record{
		Fields: []*pb.RecordField{{Label: label, Value: value}},
	}
}

//Functions to build value object from record, int, text, party, variant

func recordValue(label string, value *pb.Value) *pb.Value {
	return &pb.Value{Sum: &pb.Value_Record{Record: singleFieldRecord(label, value)}}
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

//Input builders, to be used as ChoiceArgument value input

//Integer input builder
func integerInput(label string, value int64) *pb.Value {
	return recordValue(label, intValue(value))
}

//Text input builder
func textInput(label string, value string) *pb.Value {
	return recordValue(label, textValue(value))
}

//Variant input builder
func variantInput(label string, value string) *pb.Value {
	return recordValue(label, variantValue(value))
}

func partyInput(label string, value string) *pb.Value {
	return recordValue(label, partyValue(value))
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
	input *pb.Value, //integerInput | variantInput | partyInput
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
		LedgerId:            ledgerID,
		WorkflowId:          workflowID,
		ApplicationId:       applicationID,
		CommandId:           commandID,
		Party:               submittingParty,
		LedgerEffectiveTime: &protobuf.Timestamp{Seconds: (timeHopInSeconds + 1)},
		MaximumRecordTime:   &protobuf.Timestamp{Seconds: (timeHopInSeconds + 10)},
		Commands:            []*pb.Command{command},
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
		LedgerId:            ledgerID,
		WorkflowId:          workflowID,
		ApplicationId:       applicationID,
		CommandId:           commandID,
		Party:               submittingParty,
		LedgerEffectiveTime: &protobuf.Timestamp{Seconds: (timeHopInSeconds + 1)},
		MaximumRecordTime:   &protobuf.Timestamp{Seconds: (timeHopInSeconds + 10)},
		Commands:            []*pb.Command{command},
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
		"Template ID: " + templateID.GetName() + "@" + templateID.GetPackageId(),
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

func main() {
	// Set up a connection to the server.
	cc, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer cc.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	commandClient := pb.NewCommandServiceClient(cc)

	ledgerID := getLedgerID(ctx, cc)
	//log.Println("Ledger Id = ", ledgerId)

	pkgIDs := getPackageIDs(ctx, cc, ledgerID)
	packageID := pkgIDs[0]
	//log.Printf("Active package IDs: %v", pkgIDs)

	// HERE COME YOUR FUNCTIONS
	// ...
}
