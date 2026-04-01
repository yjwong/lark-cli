package browser

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/yjwong/lark-cli/internal/config"
)

var (
	targetIDRe   = regexp.MustCompile(`target:\s*([A-F0-9]+)`)
	listLineRe   = regexp.MustCompile(`^([A-F0-9]{8})\s+\[.*?\]\s+.*?\s+(https?://\S+)$`)
)

func defaultCDPScriptPath() string {
	return filepath.Clean(filepath.Join(config.GetRootDir(), "..", "skills", "chrome-cdp", "scripts", "cdp.mjs"))
}

func scriptPath() string {
	if path := strings.TrimSpace(os.Getenv("LARK_CDP_SCRIPT")); path != "" {
		return path
	}
	return defaultCDPScriptPath()
}

func port() string {
	if value := strings.TrimSpace(os.Getenv("LARK_CDP_PORT")); value != "" {
		return value
	}
	return "9223"
}

func runCDP(args ...string) (string, error) {
	cmdArgs := append([]string{scriptPath(), "--port", port()}, args...)
	cmd := exec.Command("node", cmdArgs...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		msg := strings.TrimSpace(stderr.String())
		if msg == "" {
			msg = strings.TrimSpace(stdout.String())
		}
		if msg == "" {
			msg = err.Error()
		}
		return "", fmt.Errorf("%s", msg)
	}
	return strings.TrimSpace(stdout.String()), nil
}

func findTargetByURL(urlFragment string) (string, error) {
	output, err := runCDP("list")
	if err != nil {
		return "", err
	}

	for _, line := range strings.Split(output, "\n") {
		line = strings.TrimSpace(line)
		matches := listLineRe.FindStringSubmatch(line)
		if len(matches) != 3 {
			continue
		}
		if strings.Contains(matches[2], urlFragment) {
			return matches[1], nil
		}
	}
	return "", nil
}

func openTarget(url string) (string, error) {
	output, err := runCDP("open", url)
	if err != nil {
		return "", err
	}
	matches := targetIDRe.FindStringSubmatch(output)
	if len(matches) != 2 {
		return "", fmt.Errorf("could not parse target id from cdp open output: %s", output)
	}
	targetID := strings.ToUpper(matches[1])
	if len(targetID) > 8 {
		targetID = targetID[:8]
	}
	return targetID, nil
}

func eval(targetID, js string) error {
	_, err := runCDP("eval", targetID, js)
	return err
}

func WrapSheetRange(spreadsheetToken, rangeSpec string) error {
	urlFragment := "/sheets/" + spreadsheetToken
	targetID, err := findTargetByURL(urlFragment)
	if err != nil {
		return err
	}
	if targetID == "" {
		targetID, err = openTarget("https://glints.sg.larksuite.com/sheets/" + spreadsheetToken)
		if err != nil {
			return err
		}
		time.Sleep(2 * time.Second)
	}

	js := fmt.Sprintf(`(()=>{
const sh = window.spread && window.spread.sheets && window.spread.sheets[window.spread._activeSheetIndex || 0];
if (!sh) throw new Error("sheet runtime not ready");
const r = sh.rangeStringToRange(%q);
sh.setSelection(r, { row: r.startRow, col: r.startCol });
const icon = document.querySelector("svg[data-icon='WrapOutlined']");
if (!icon) throw new Error("WrapOutlined toolbar icon not found");
const button = icon.closest(".toolbar-icon-container, .button-role, .response-area");
if (!button) throw new Error("wrap toolbar button not found");
button.click();
return { ok: true };
})()`, rangeSpec)

	return eval(targetID, js)
}
