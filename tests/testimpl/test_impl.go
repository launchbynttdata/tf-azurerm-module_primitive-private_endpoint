package testimpl

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/launchbynttdata/lcaf-component-terratest/types"
	"github.com/stretchr/testify/assert"
)

func TestComposablePrivateEndpointComplete(t *testing.T, ctx types.TestContext) {

	t.Run("TestPrivateEndpoint", func(t *testing.T) {
		privateEndpointId := terraform.OutputContext(t, t.Context(), ctx.TerratestTerraformOptions(), "private_endpoint_id")
		assert.NotEmpty(t, privateEndpointId, "Private endpoint ID must not be empty")
	})
}
