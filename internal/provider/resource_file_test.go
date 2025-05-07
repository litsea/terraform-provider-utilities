// Copyright (c) Litsea
// SPDX-License-Identifier: MIT

package provider

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDownloadFile_GET(t *testing.T) {
	// Setup mock server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "GET", r.Method)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("hello world"))
	}))
	defer ts.Close()

	// Create temp file path
	filePath := "./test_output.txt"
	defer os.Remove(filePath)

	err := downloadFile("GET", ts.URL, filePath, map[string]string{})
	require.NoError(t, err)

	content, err := os.ReadFile(filePath)
	require.NoError(t, err)
	require.Equal(t, "hello world", string(content))
}

func TestDownloadFile_POST_WithHeaders(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "POST", r.Method)
		require.Equal(t, "Bearer xyz", r.Header.Get("Authorization"))
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("posted"))
	}))
	defer ts.Close()

	filePath := "./test_post_output.txt"
	defer os.Remove(filePath)

	err := downloadFile("POST", ts.URL, filePath, map[string]string{
		"Authorization": "Bearer xyz",
	})
	require.NoError(t, err)

	content, err := os.ReadFile(filePath)
	require.NoError(t, err)
	require.Equal(t, "posted", string(content))
}

func TestDownloadFile_Failure(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	}))
	defer ts.Close()

	filePath := "./fail.txt"
	defer os.Remove(filePath)

	err := downloadFile("GET", ts.URL, "fail.txt", nil)
	require.Error(t, err)
	require.Contains(t, err.Error(), "failed to download file")
}
