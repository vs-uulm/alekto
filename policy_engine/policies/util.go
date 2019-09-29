package policies

import (
	"flag"
	"fmt"
	"github.com/ma-zero-trust-prototype/shared_lib/structs/policy"
	"github.com/ma-zero-trust-prototype/shared_lib/yaml"
	"io/ioutil"
)

func initPolicies() {

	var policyFile []byte

	if flag.Lookup("test.v") == nil {
		policyFile, _ = ioutil.ReadFile("policies/policies.yaml")

	} else {
		policyFile, _ = ioutil.ReadFile("policies.yaml")
	}

	policies = parseYamlIntoPolicy(policyFile)
}

func parseYamlIntoPolicy(policyFile []byte) (policies []policy.ClientPolicy) {

	yaml.ParseToInterface(policyFile, &policies)

	return
}

func getNewEvaluation(clientPolicy policy.ClientPolicy) policy.PolicyEvaluation {

	return policy.PolicyEvaluation{
		Name: clientPolicy.Metadata.Name,
		Message: fmt.Sprintf("\nPolicy: %v \nDescription: %v\nError Message: ",
			clientPolicy.Metadata.Name, clientPolicy.Metadata.Description),
		Exchangeable: clientPolicy.Metadata.Exchangeable,
		Success:      true}
}

func updateEvaluation(success bool, message string, evaluation *policy.PolicyEvaluation) {

	evaluation.Success = success && evaluation.Success

	if !success {
		evaluation.Message += message + "\n"
	}
}
