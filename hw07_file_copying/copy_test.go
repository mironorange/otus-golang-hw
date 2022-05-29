package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	inputFilename := "testdata/input.txt"
	testCases := []struct {
		offset int64
		limit  int64
	}{
		{
			offset: 0,
			limit:  0,
		},
		{
			offset: 0,
			limit:  10,
		},
		{
			offset: 0,
			limit:  1000,
		},
		{
			offset: 0,
			limit:  10000,
		},
		{
			offset: 100,
			limit:  1000,
		},
		{
			offset: 6000,
			limit:  1000,
		},
	}

	for _, tc := range testCases {
		dstFilename := fmt.Sprintf("testdata/out_offset%d_limit%d.txt", tc.offset, tc.limit)
		srcFilename := fmt.Sprintf("testdata/res_offset%d_limit%d.txt", tc.offset, tc.limit)

		err := Copy(inputFilename, srcFilename, tc.offset, tc.limit)
		require.NoError(t, err)

		expected, _ := iOpenAndReadAllFile(dstFilename)
		result, _ := iOpenAndReadAllFile(srcFilename)

		require.Equal(t, expected, result)
		require.Less(t, 0, len(result))
		require.Len(t, expected, len(result))
	}

	defer func() {
		for _, tc := range testCases {
			filename := fmt.Sprintf("testdata/res_offset%d_limit%d.txt", tc.offset, tc.limit)
			os.Remove(filename)
		}
	}()
}

func TestNonExistingFile(t *testing.T) {
	dstFilename := "testdata/non_existing_file.txt"
	srcFilename := "testdata/res_non-existing_file.txt"

	err := Copy(dstFilename, srcFilename, 0, 0)
	require.Error(t, err, ErrUnsupportedFile)
}

func TestExceedingOffset(t *testing.T) {
	inputFilename := "testdata/input.txt"
	outFilename := "testdata/res.txt"

	err := Copy(inputFilename, outFilename, 1000000, 0)
	require.Error(t, err, ErrOffsetExceedsFileSize)
}
