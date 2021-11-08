package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToCamel(t *testing.T) {
	assert.Equal(t, "UserPlan", ToCamel("user_plan"))
	assert.Equal(t, "UserPlan", ToCamel("UserPlan"))
	assert.Equal(t, "UserPlan", ToCamel("userPlan"))
	assert.Equal(t, "UserPlanV0", ToCamel("user_plan_v0"))
}
