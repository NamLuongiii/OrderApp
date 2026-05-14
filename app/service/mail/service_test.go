package mail

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadTemplates(t *testing.T) {
	templates, e := loadTemplates()
	assert.NoError(t, e)
	assert.NotEmpty(t, templates)
}
