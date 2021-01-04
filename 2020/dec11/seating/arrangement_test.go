package seating

import (
	"reflect"
	"testing"

	"github.com/ablqk/adventofcode/libs/fileread"
	"github.com/stretchr/testify/assert"
)

func TestArrangement_ParseLine(t *testing.T) {
	a := Arrangement{}
	fileread.ReadAndApply("../testdata/in.txt", a.ParseLine)
	assert.True(t, a.room[5][5].hasSeat)
}

func TestArrangement_ForkRoom(t *testing.T) {
	a := Arrangement{
		width:  3,
		height: 3,
		room: [][]square{
			{
				{hasPerson: true, hasSeat: true},
				{hasPerson: true},
				{},
			},
			{
				{},
				{hasSeat: true},
				{},
			},
			{
				{hasSeat: true},
				{},
				{hasPerson: true, hasSeat: true},
			},
		},
	}
	b := a.ForkRoom()
	assert.True(t, reflect.DeepEqual(a.room, b))
}

func TestArrangement_CountPeople(t *testing.T) {
	a := Arrangement{
		width:  3,
		height: 3,
		room: [][]square{
			{{hasPerson: true}, {}, {}},
			{{}, {hasPerson: true}, {}},
			{{hasPerson: true}, {}, {hasPerson: true}},
		},
	}
	assert.Equal(t, 4, a.CountPeople())
}
