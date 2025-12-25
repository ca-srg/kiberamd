package main

import (
	"fmt"
	"os"

	"github.com/ca-srg/kiberamd/internal/config"
	"github.com/ca-srg/kiberamd/internal/export"
	"github.com/ca-srg/kiberamd/internal/kibela"
	"github.com/spf13/cobra"
)

var outputDir string

var rootCmd = &cobra.Command{
	Use:   "kiberag-export",
	Short: "Export all notes from Kibela to markdown files",
	Long:  `Export all notes from Kibela using GraphQL API and save them as markdown files.`,
	RunE:  runExport,
}

func init() {
	rootCmd.Flags().StringVarP(&outputDir, "output", "o", "markdown", "output directory for markdown files")
}

func runExport(cmd *cobra.Command, args []string) error {
	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	client := kibela.NewClient(cfg.KibelaTeam, cfg.KibelaToken)

	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	exporter := export.New(client)

	fmt.Printf("Starting export of all Kibela notes to '%s'...\n", outputDir)

	err = exporter.ExportAllNotes(outputDir)
	if err != nil {
		return fmt.Errorf("failed to export notes: %w", err)
	}

	fmt.Println("Export completed successfully!")
	return nil
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
