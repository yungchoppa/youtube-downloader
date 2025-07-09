package ytdlp

import (
	"encoding/json"
	"io"
	"os/exec"
)

type Format struct {
	FormatID   string   `json:"format_id"`
	Ext        string   `json:"ext"`
	FormatNote string   `json:"format_note"`
	Acodec     string   `json:"acodec"`
	Vcodec     string   `json:"vcodec"`
	Filesize   *int64   `json:"filesize"`
	Tbr        *float64 `json:"tbr"`
	Height     *int     `json:"height"`
}

type VideoInfo struct {
	Formats []Format `json:"formats"`
}

func GetFormats(url string) ([]Format, error) {
	cmd := exec.Command("./yt-dlp.exe", "--dump-json", url)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	var info VideoInfo
	if err := json.Unmarshal(output, &info); err != nil {
		return nil, err
	}
	var filtered []Format
	for _, f := range info.Formats {
		if (f.Ext == "mp3" || f.Ext == "mp4") && (f.Acodec != "none" || f.Vcodec != "none") {
			filtered = append(filtered, f)
		}
	}
	return filtered, nil
}

// StreamDownload запускает yt-dlp и стримит файл напрямую в http.ResponseWriter
func StreamDownload(w io.Writer, url, formatID string) error {
	cmd := exec.Command("./yt-dlp.exe", "-f", formatID, "-o", "-", url)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	if err := cmd.Start(); err != nil {
		return err
	}
	_, err = io.Copy(w, stdout)
	cmd.Wait()
	return err
}
