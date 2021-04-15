// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v2beta1

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-gcp-native/sdk/go/gcp"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "gcp-native:dialogflow/v2beta1:AgentEntityType":
		r = &AgentEntityType{}
	case "gcp-native:dialogflow/v2beta1:AgentEnvironmentUserSessionContext":
		r = &AgentEnvironmentUserSessionContext{}
	case "gcp-native:dialogflow/v2beta1:AgentEnvironmentUserSessionEntityType":
		r = &AgentEnvironmentUserSessionEntityType{}
	case "gcp-native:dialogflow/v2beta1:AgentIntent":
		r = &AgentIntent{}
	case "gcp-native:dialogflow/v2beta1:AgentKnowledgeBase":
		r = &AgentKnowledgeBase{}
	case "gcp-native:dialogflow/v2beta1:AgentKnowledgeBaseDocument":
		r = &AgentKnowledgeBaseDocument{}
	case "gcp-native:dialogflow/v2beta1:AgentSessionContext":
		r = &AgentSessionContext{}
	case "gcp-native:dialogflow/v2beta1:AgentSessionEntityType":
		r = &AgentSessionEntityType{}
	case "gcp-native:dialogflow/v2beta1:Conversation":
		r = &Conversation{}
	case "gcp-native:dialogflow/v2beta1:ConversationParticipant":
		r = &ConversationParticipant{}
	case "gcp-native:dialogflow/v2beta1:ConversationProfile":
		r = &ConversationProfile{}
	case "gcp-native:dialogflow/v2beta1:KnowledgeBase":
		r = &KnowledgeBase{}
	case "gcp-native:dialogflow/v2beta1:KnowledgeBaseDocument":
		r = &KnowledgeBaseDocument{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := gcp.PkgVersion()
	if err != nil {
		fmt.Println("failed to determine package version. defaulting to v1: %v", err)
	}
	pulumi.RegisterResourceModule(
		"gcp-native",
		"dialogflow/v2beta1",
		&module{version},
	)
}
