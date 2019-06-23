package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

var logger = shim.NewLogger("test-main")

// Test struct
type Test struct{}

// Invoke - Called on chaincode invoke. Takes a function name passed and calls that function. Passes the initial arguments passed are passed on to the called function.
func (test *Test) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	transactionName, _ := stub.GetFunctionAndParameters()

	logger.Infof("Executing transaction '%s'", transactionName)

	return shim.Success(nil)
}

// Main - main - starts up the chaincode
func main() {
	// LogDebug, LogInfo, LogNotice, LogWarning, LogError, LogCritical (Default: LogDebug)
	logger.SetLevel(shim.LogDebug)

	err := shim.Start(new(Test))
	if err != nil {
		logger.Errorf("Error starting Test: %s", err)
	}
}

// Init Function - Called when the user deploys the chaincode
func (test *Test) Init(stub shim.ChaincodeStubInterface) pb.Response {
	logger.Infof("Test chaincode initialized!")

	return shim.Success(nil)
}
