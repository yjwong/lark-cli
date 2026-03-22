package main

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestSlugify(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"Create a user", "create-a-user"},
		{"API Call Guide", "api-call-guide"},
		{"Process overview", "process-overview"},
		{"Hello / World", "hello-world"},
		{"test__double", "test-double"},
		{"", "index"},
		{"MiXeD CaSe 123", "mixed-case-123"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := slugify(tt.input)
			if got != tt.want {
				t.Errorf("slugify(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestDeriveOutputPath(t *testing.T) {
	tests := []struct {
		name       string
		page       ScrapedPage
		want       string
	}{
		{
			name: "new-style path",
			page: ScrapedPage{
				TargetPath: "/server-docs/contact-v3/user/create",
			},
			want: "server-docs/contact-v3/user/create.md",
		},
		{
			name: "old-style with service version",
			page: ScrapedPage{
				Path: "/uAjLw4CM/ukTMukTMukTM/minutes-v1/minute/get",
			},
			want: "server-docs/minutes-v1/minute/get.md",
		},
		{
			name: "fallback with breadcrumb",
			page: ScrapedPage{
				Path:       "/uAjLw4CM/unknown",
				ParentPath: []string{"API Call Guide", "Calling Process"},
				Name:       "Process overview",
			},
			want: "server-docs/api-call-guide/calling-process/process-overview.md",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := deriveOutputPath(&tt.page)
			if got != tt.want {
				t.Errorf("deriveOutputPath() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestParseServiceResource(t *testing.T) {
	tests := []struct {
		name         string
		docPath      string
		wantService  string
		wantResource string
	}{
		{
			name:         "new-style contact path",
			docPath:      "/server-docs/contact-v3/user/create",
			wantService:  "contact-v3",
			wantResource: "user",
		},
		{
			name:         "old-style minutes path",
			docPath:      "/uAjLw4CM/ukTMukTMukTM/minutes-v1/minute/get",
			wantService:  "minutes-v1",
			wantResource: "minute",
		},
		{
			name:         "junk path returns empty",
			docPath:      "/uAjLw4CM/ukTMukTMukTM/some/path",
			wantService:  "",
			wantResource: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotService, gotResource := parseServiceResource(tt.docPath)
			if gotService != tt.wantService {
				t.Errorf("parseServiceResource(%q) service = %q, want %q", tt.docPath, gotService, tt.wantService)
			}
			if gotResource != tt.wantResource {
				t.Errorf("parseServiceResource(%q) resource = %q, want %q", tt.docPath, gotResource, tt.wantResource)
			}
		})
	}
}

func TestLooksLikeJunkPath(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"uAjLw4CM", true},
		{"ukTMukTMukTM", true},
		{"contact-v3", false},
		{"user", false},
		{"user_task", false},
		{"im-v1", false},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := looksLikeJunkPath(tt.input)
			if got != tt.want {
				t.Errorf("looksLikeJunkPath(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestDeriveOutputPathNoTraversal(t *testing.T) {
	// deriveOutputPath does not validate traversal itself (that's done in scrapePage),
	// but we can verify that the output path it returns never starts with ".."
	// even for unusual inputs, and that filepath.Join + filepath.Rel would catch traversal.
	outputDir := t.TempDir()

	cases := []struct {
		name string
		page ScrapedPage
	}{
		{
			name: "normal new-style path",
			page: ScrapedPage{TargetPath: "/server-docs/contact-v3/user/create"},
		},
		{
			name: "normal old-style path",
			page: ScrapedPage{Path: "/uAjLw4CM/ukTMukTMukTM/minutes-v1/minute/get"},
		},
		{
			name: "breadcrumb fallback",
			page: ScrapedPage{
				Path:       "/uAjLw4CM/unknown",
				ParentPath: []string{"API Call Guide", "Calling Process"},
				Name:       "Process overview",
			},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			relPath := deriveOutputPath(&tt.page)

			// The relative path should never start with ".."
			if strings.HasPrefix(relPath, "..") {
				t.Errorf("deriveOutputPath() returned a path starting with '..': %q", relPath)
			}

			// Simulate the validation logic from scrapePage
			outFile := filepath.Join(outputDir, relPath)
			absOutFile, err := filepath.Abs(outFile)
			if err != nil {
				t.Fatalf("filepath.Abs failed: %v", err)
			}
			absOutput, err := filepath.Abs(outputDir)
			if err != nil {
				t.Fatalf("filepath.Abs failed: %v", err)
			}
			rel, err := filepath.Rel(absOutput, absOutFile)
			if err != nil || strings.HasPrefix(rel, "..") {
				t.Errorf("path traversal would be detected for relPath %q: rel=%q err=%v", relPath, rel, err)
			}
		})
	}
}
