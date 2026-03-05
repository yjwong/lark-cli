package api

import (
	"fmt"
	"io"
	"net/url"
)

// TranscriptOptions contains options for exporting a transcript
type TranscriptOptions struct {
	NeedSpeaker   bool   // Include speaker names
	NeedTimestamp bool   // Include timestamps
	FileFormat    string // "txt" or "srt"
}

// GetMinute retrieves metadata for a minutes recording
// minuteToken: the minute token (24 characters, from the minutes URL)
func (c *Client) GetMinute(minuteToken string) (*Minute, error) {
	path := fmt.Sprintf("/minutes/v1/minutes/%s", url.PathEscape(minuteToken))

	var resp MinuteResponse
	if err := c.Get(path, &resp); err != nil {
		return nil, err
	}

	if err := resp.Err(); err != nil {
		return nil, err
	}

	return resp.Data.Minute, nil
}

// GetMinuteTranscript exports the transcript of a minutes recording
// Returns the raw transcript content as bytes
func (c *Client) GetMinuteTranscript(minuteToken string, opts TranscriptOptions) ([]byte, error) {
	params := url.Values{}
	if opts.NeedSpeaker {
		params.Set("need_speaker", "true")
	}
	if opts.NeedTimestamp {
		params.Set("need_timestamp", "true")
	}
	if opts.FileFormat != "" {
		params.Set("file_format", opts.FileFormat)
	}

	path := fmt.Sprintf("/minutes/v1/minutes/%s/transcript", url.PathEscape(minuteToken))
	if len(params) > 0 {
		path += "?" + params.Encode()
	}

	// Use Download since this returns binary data
	body, _, err := c.Download(path)
	if err != nil {
		return nil, err
	}
	defer body.Close()

	content, err := io.ReadAll(body)
	if err != nil {
		return nil, fmt.Errorf("failed to read transcript: %w", err)
	}

	return content, nil
}

// GetMinuteMediaURL retrieves the download URL for the audio/video file
// The returned URL is valid for 24 hours
func (c *Client) GetMinuteMediaURL(minuteToken string) (string, error) {
	path := fmt.Sprintf("/minutes/v1/minutes/%s/media", url.PathEscape(minuteToken))

	var resp MinuteMediaResponse
	if err := c.Get(path, &resp); err != nil {
		return "", err
	}

	if err := resp.Err(); err != nil {
		return "", err
	}

	return resp.Data.DownloadURL, nil
}
