package v1

import (
	"testing"

	"gotest.tools/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/validation/field"
)

func Test_Policy_Name(t *testing.T) {
	subject := Policy{
		ObjectMeta: metav1.ObjectMeta{
			Name: "this-is-a-way-too-long-policy-name-that-should-trigger-an-error-when-calling-the-policy-validation-method",
		},
	}
	errs := subject.Validate(true, nil)
	assert.Assert(t, len(errs) == 1)
	assert.Equal(t, errs[0].Field, "name")
	assert.Equal(t, errs[0].Type, field.ErrorTypeTooLong)
	assert.Equal(t, errs[0].Detail, "must have at most 63 bytes")
	assert.Equal(t, errs[0].Error(), "name: Too long: must have at most 63 bytes")
}

func Test_Policy_IsNamespaced(t *testing.T) {
	namespaced := Policy{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "this-is-a-way-too-long-policy-name-that-should-trigger-an-error-when-calling-the-policy-validation-method",
			Namespace: "abcd",
		},
	}
	notNamespaced := Policy{
		ObjectMeta: metav1.ObjectMeta{
			Name: "this-is-a-way-too-long-policy-name-that-should-trigger-an-error-when-calling-the-policy-validation-method",
		},
	}
	assert.Equal(t, namespaced.IsNamespaced(), false)
	assert.Equal(t, notNamespaced.IsNamespaced(), false)
}
