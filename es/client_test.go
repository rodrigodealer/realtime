package es

import (
	"context"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	conn := &EsClient{}
	conn.Connect()

	var ping = conn.Ping()
	assert.Equal(t, ping, 200)
}

func TestFailedConnection(t *testing.T) {
	conn := &EsClient{}
	os.Setenv("ELASTICSEARCH_URL", "localhost:9400")
	assert.Panics(t, func() { conn.Connect() })
	os.Unsetenv("ELASTICSEARCH_URL")
}

func TestIndexCreation(t *testing.T) {
	ctx := context.Background()
	conn := &EsClient{}
	conn.Connect()
	var now = strconv.Itoa(time.Now().Nanosecond())
	var indexName = "testindex_" + now
	existed, _ := conn.Client.IndexExists(indexName).Do(ctx)
	conn.IndexSetup("testindex_" + now)
	existes, _ := conn.Client.IndexExists(indexName).Do(ctx)

	assert.False(t, existed, "Index shoult not exists")
	assert.True(t, existes, "Index shoult exists")

	conn.Client.DeleteIndex(indexName)
}

func TestIndexRecreation(t *testing.T) {
	ctx := context.Background()
	conn := &EsClient{}
	conn.Connect()
	var now = strconv.Itoa(time.Now().Nanosecond())
	var indexName = "testindex_" + now
	conn.IndexSetup("testindex_" + now)
	existed, _ := conn.Client.IndexExists(indexName).Do(ctx)
	conn.IndexSetup("testindex_" + now)

	assert.True(t, existed, "Index shoult exists already")
	conn.Client.DeleteIndex(indexName)
}
