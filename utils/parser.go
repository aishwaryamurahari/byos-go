package utils

import (
	"bufio"
	"strconv"
	"strings"
)

// HTTPRequest represents a simplified parsed request
type HTTPRequest struct {
	Method  string
	Path    string
	Version string
	Headers map[string]string
	Body    string
}

// ParseRequest parses raw HTTP request data into an HTTPRequest struct
func ParseRequest(data string) (*HTTPRequest, error) {
	reader := bufio.NewReader(strings.NewReader(data))

	// Read request line: e.g., "GET /index.html HTTP/1.1"
	line, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}
	line = strings.TrimSpace(line)
	parts := strings.Split(line, " ")
	if len(parts) != 3 {
		return nil, err
	}

	req := &HTTPRequest{
		Method:  parts[0],
		Path:    parts[1],
		Version: parts[2],
		Headers: make(map[string]string),
	}

	// Read headers
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		line = strings.TrimSpace(line)
		if line == "" {
			// Empty line indicates end of headers
			break
		}

		headerParts := strings.SplitN(line, ":", 2)
		if len(headerParts) == 2 {
			key := strings.TrimSpace(headerParts[0])
			value := strings.TrimSpace(headerParts[1])
			req.Headers[key] = value
		}
	}

	// Read body based on Content-Length
	if lengthStr, ok := req.Headers["Content-Length"]; ok {
		length, err := strconv.Atoi(lengthStr)
		if err == nil && length > 0 {
			bodyBuf := make([]byte, length)
			reader.Read(bodyBuf)
			req.Body = string(bodyBuf)
		}
	}

	// Optionally read body
	body, _ := reader.ReadString(0) // reads till EOF
	req.Body = body

	return req, nil
}
