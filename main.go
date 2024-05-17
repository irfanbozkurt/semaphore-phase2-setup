package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:      "setup",
		Usage:     "Use this tool to generate parameters of Groth16 via MPC",
		UsageText: "setup command [arguments...]",
		Commands: []*cli.Command{
			/* --------------------------- Phase 2 Initialize --------------------------- */
			{
				Name:        "p2n",
				Usage:       "p2n <ptauPath> <r1csPath> <phase2Path>",
				Description: "initialize phase 2 for the given circuit",
				Action:      p2n,
			},
			/* --------------------------- Phase 2 Contribute --------------------------- */
			{
				Name:        "p2c",
				Usage:       "p2c <inputPath> <outputPath>",
				Description: "contribute phase 2 randomness for Groth16",
				Action:      p2c,
			},
			/* ----------------------------- Phase 2 Verify ----------------------------- */
			{
				Name:        "p2v",
				Usage:       "p2v <inputPath> <originPath>",
				Description: "verify phase 2 contributions for Groth16",
				Action:      p2v,
			},
			/* ----------------------------- Keys Extraction ---------------------------- */
			{
				Name:        "extract-keys",
				Usage:       "extract-keys <phase1Path> <phase2Path> <phase2EvalsPath> <r1csPath>",
				Description: "extract proving and verifying keys",
				Action:      extractKeys,
			},
			{
				Name:        "sol",
				Usage:       "sol <verifyingKey>",
				Description: "export verifier smart contract from verifying key",
				Action:      exportSol,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
