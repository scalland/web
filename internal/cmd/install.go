package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"text/template"

	"github.com/spf13/cobra"
)

const systemdTemplate = `[Unit]
Description=
After=network.target

[Service]
Type=simple
User={{ .User }}
WorkingDirectory={{ .WorkDir }}
ExecStart={{ .BinPath }} serve --config={{ .ConfigPath }}
Restart=on-failure
RestartSec=5

[Install]
WantedBy=multi-user.target
`

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install as a systemd service (Linux only)",
	RunE: func(cmd *cobra.Command, args []string) error {
		if runtime.GOOS != "linux" {
			return fmt.Errorf("install command is only available on Linux")
		}

		binPath, _ := os.Executable()
		cwd, _ := os.Getwd()

		data := map[string]string{
			"User":       os.Getenv("USER"),
			"WorkDir":    cwd,
			"BinPath":    binPath,
			"ConfigPath": cwd + "/configs/app.yml",
		}

		servicePath := fmt.Sprintf("/etc/systemd/system/%s.service", "")
		f, err := os.Create(servicePath)
		if err != nil {
			return fmt.Errorf("create service file: %w (try running with sudo)", err)
		}
		defer f.Close()

		tmpl, _ := template.New("service").Parse(systemdTemplate)
		if err := tmpl.Execute(f, data); err != nil {
			return fmt.Errorf("write service file: %w", err)
		}

		// Reload systemd
		if err := exec.Command("systemctl", "daemon-reload").Run(); err != nil {
			return fmt.Errorf("systemctl daemon-reload: %w", err)
		}
		if err := exec.Command("systemctl", "enable", "").Run(); err != nil {
			return fmt.Errorf("systemctl enable: %w", err)
		}

		fmt.Printf("Service installed: %s\n", servicePath)
		fmt.Printf("Run: sudo systemctl start %s\n", "")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}
