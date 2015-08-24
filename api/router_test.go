
package api

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
)

func TestGetJobs( t *testing.T ){
	url := "localhost:2345/api/v1/jobs"
	r, e := makeRequest("GET", url, nil );

	assert := assert.New(t)
	assert.Equal( r , "1" )

	if e != nil {
		fmt.Println( r )
	} else {
		fmt.Println(e)
	}
}


