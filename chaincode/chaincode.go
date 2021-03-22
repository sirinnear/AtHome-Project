package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type Certification struct {
	// docType is used to distinguish the various types of objects in state database
	ObjectType string `json:"docType"`
	// the fieldtags are needed to keep case from bouncing around
	Name   string `json:"name"`
	Issuer string `json:"issuer"`
	Date   string `json:"date"`
}

type CertificationPrivateDetails struct {
	// docType is used to distinguish the various types of objects in state database
	ObjectType string `json:"docType"`
	// the fieldtags are needed to keep case from bouncing around
	Name  string `json:"name"`
	Owner string `json:"owner"`
}

type SmartContract struct {
	contractapi.Contract
}

func (s *SmartContract) InitCertification(ctx contractapi.TransactionContextInterface) error {
	transMap, err := ctx.GetStub().GetTransient()
	if err != nil {
		return fmt.Errorf("Error getting transient: " + err.Error())
	}

	transientCertificationJSON, ok := transMap["certification"]
	if !ok {
		return fmt.Errorf("certification not found in the transient map")
	}

	type certificationTransientInput struct {
		Name   string `json:"name"`
		Issuer string `json:"issuer"`
		Date   string `json:"Date"`
		Owner  string `json:"owner"`
	}

	var certificationInput certificationTransientInput
	err = json.Unmarshal(transientCertificationJSON, &certificationInput)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %s", err.Error())
	}
	if len(certificationInput.Name) == 0 {
		return fmt.Errorf("name field must be a non-empty string")
	}
	if len(certificationInput.Issuer) == 0 {
		return fmt.Errorf("issuer field must be a non-empty string")
	}
	if len(certificationInput.Date) == 0 {
		return fmt.Errorf("date field must be a non-empty string")
	}
	if len(certificationInput.Owner) == 0 {
		return fmt.Errorf("owner field must be a non-empty string")
	}

	certificationAsBytes, err := ctx.GetStub().GetPrivateData("collectionCertifications", certificationInput.Name)
	if err != nil {
		return fmt.Errorf("Failed to get certification: " + err.Error())
	} else if certificationAsBytes != nil {
		fmt.Println("This certification already exists: " + certificationInput.Name)
		return fmt.Errorf("This certification already exists: " + certificationInput.Name)
	}

	certification := &Certification{
		ObjectType: "Certification",
		Name:       certificationInput.Name,
		Issuer:     certificationInput.Issuer,
		Date:       certificationInput.Date,
	}

	certificationJSONasBytes, err := json.Marshal(certification)
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	err = ctx.GetStub().PutPrivateData("collectionCertifications", certificationInput.Name, certificationJSONasBytes)
	if err != nil {
		return fmt.Errorf("failed to put Certification: %s", err.Error())
	}

	certificationPrivateDetails := &CertificationPrivateDetails{
		ObjectType: "CertificationPrivateDetails",
		Name:       certificationInput.Name,
		Owner:      certificationInput.Owner,
	}

	certificationPrivateDetailsAsBytes, err := json.Marshal(certificationPrivateDetails)
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	err = ctx.GetStub().PutPrivateData("collectionCertificationPrivateDetails", certificationInput.Name, certificationPrivateDetailsAsBytes)
	if err != nil {
		return fmt.Errorf("failed to put Certification private details: %s", err.Error())
	}

	indexName := "issuer~name"
	issuerNameIndexKey, err := ctx.GetStub().CreateCompositeKey(indexName, []string{certification.Issuer, certification.Name})
	if err != nil {
		return err
	}

	value := []byte{0x00}
	_ = ctx.GetStub().PutPrivateData("collectionCertifications", issuerNameIndexKey, value)

	return nil
}

func (s *SmartContract) ReadCertification(ctx contractapi.TransactionContextInterface, certificationID string) (*Certification, error) {

	certificationJSON, err := ctx.GetStub().GetPrivateData("collectionCertifications", certificationID)
	if err != nil {
		return nil, fmt.Errorf("failed to read from certification %s", err.Error())
	}
	if certificationJSON == nil {
		return nil, fmt.Errorf("%s does not exist", certificationID)
	}

	certification := new(Certification)
	_ = json.Unmarshal(certificationJSON, certification)

	return certification, nil
}

func (s *SmartContract) ReadCertificationPrivateDetails(ctx contractapi.TransactionContextInterface, certificationID string) (*CertificationPrivateDetails, error) {

	certificationDetailsJSON, err := ctx.GetStub().GetPrivateData("collectionCertificationPrivateDetails", certificationID)
	if err != nil {
		return nil, fmt.Errorf("failed to read from certification details %s", err.Error())
	}
	if certificationDetailsJSON == nil {
		return nil, fmt.Errorf("%s does not exist", certificationID)
	}

	certificationDetails := new(CertificationPrivateDetails)
	_ = json.Unmarshal(certificationDetailsJSON, certificationDetails)

	return certificationDetails, nil
}

func (s *SmartContract) Delete(ctx contractapi.TransactionContextInterface) error {

	transMap, err := ctx.GetStub().GetTransient()
	if err != nil {
		return fmt.Errorf("Error getting transient: " + err.Error())
	}

	transientDeleteCertificationJSON, ok := transMap["certification_delete"]
	if !ok {
		return fmt.Errorf("certification to delete not found in the transient map")
	}

	type certificationDelete struct {
		Name string `json:"name"`
	}

	var certificationDeleteInput certificationDelete
	err = json.Unmarshal(transientDeleteCertificationJSON, &certificationDeleteInput)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %s", err.Error())
	}

	if len(certificationDeleteInput.Name) == 0 {
		return fmt.Errorf("name field must be a non-empty string")
	}

	valAsBytes, err := ctx.GetStub().GetPrivateData("collectionCertifications", certificationDeleteInput.Name)
	if err != nil {
		return fmt.Errorf("failed to read certification: %s", err.Error())
	}
	if valAsBytes == nil {
		return fmt.Errorf("certification private details does not exist: %s", certificationDeleteInput.Name)
	}

	var certificationToDelete Certification
	err = json.Unmarshal([]byte(valAsBytes), &certificationToDelete)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %s", err.Error())
	}

	err = ctx.GetStub().DelPrivateData("collectionCertifications", certificationDeleteInput.Name)
	if err != nil {
		return fmt.Errorf("Failed to delete state:" + err.Error())
	}

	indexName := "issuer~name"
	colorNameIndexKey, err := ctx.GetStub().CreateCompositeKey(indexName, []string{certificationToDelete.Issuer, certificationToDelete.Name})
	if err != nil {
		return err
	}

	err = ctx.GetStub().DelPrivateData("collectionCertifications", colorNameIndexKey)
	if err != nil {
		return fmt.Errorf("Failed to delete certification:" + err.Error())
	}

	err = ctx.GetStub().DelPrivateData("collectionCertificationPrivateDetails", certificationDeleteInput.Name)
	if err != nil {
		return err
	}

	return nil
}

func (s *SmartContract) GetCertificationsByRange(ctx contractapi.TransactionContextInterface, startKey string, endKey string) ([]Certification, error) {

	resultsIterator, err := ctx.GetStub().GetPrivateDataByRange("collectionCertifications", startKey, endKey)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	results := []Certification{}

	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		newCertification := new(Certification)

		err = json.Unmarshal(response.Value, newCertification)
		if err != nil {
			return nil, err
		}

		results = append(results, *newCertification)
	}

	return results, nil
}

func (s *SmartContract) getQueryResultForQueryString(ctx contractapi.TransactionContextInterface, queryString string) ([]Certification, error) {

	resultsIterator, err := ctx.GetStub().GetPrivateDataQueryResult("collectionCertifications", queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	results := []Certification{}

	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		newCertification := new(Certification)

		err = json.Unmarshal(response.Value, newCertification)
		if err != nil {
			return nil, err
		}

		results = append(results, *newCertification)
	}
	return results, nil
}

func (s *SmartContract) QueryCertificationsByOwner(ctx contractapi.TransactionContextInterface, owner string) ([]Certification, error) {

	ownerString := strings.ToLower(owner)

	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"certification\",\"owner\":\"%s\"}}", ownerString)

	queryResults, err := s.getQueryResultForQueryString(ctx, queryString)
	if err != nil {
		return nil, err
	}
	return queryResults, nil
}

func (s *SmartContract) QueryCertifications(ctx contractapi.TransactionContextInterface, queryString string) ([]Certification, error) {

	queryResults, err := s.getQueryResultForQueryString(ctx, queryString)
	if err != nil {
		return nil, err
	}
	return queryResults, nil
}

func (s *SmartContract) GetCertificationHash(ctx contractapi.TransactionContextInterface, collection string, certificationID string) (string, error) {

	hashAsBytes, err := ctx.GetStub().GetPrivateDataHash(collection, certificationID)
	if err != nil {
		return "", fmt.Errorf("Failed to get public data hash for certification:" + err.Error())
	} else if hashAsBytes == nil {
		return "", fmt.Errorf("Certification does not exist: " + certificationID)
	}

	return string(hashAsBytes), nil
}

func main() {

	chaincode, err := contractapi.NewChaincode(new(SmartContract))

	if err != nil {
		fmt.Printf("Error creating private certifications chaincode: %s", err.Error())
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting private certifications chaincode: %s", err.Error())
	}
}
