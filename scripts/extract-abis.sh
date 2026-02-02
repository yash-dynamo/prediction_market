#!/bin/bash
# Extract ABI files from Foundry JSON output for use with abigen

set -e

OUT_DIR="out"
CONTRACTS=("MockUSDC" "OutcomeToken" "PredictionMarket")

for contract in "${CONTRACTS[@]}"; do
    json_file="${OUT_DIR}/${contract}.sol/${contract}.json"
    abi_file="${OUT_DIR}/${contract}.sol/${contract}.abi"
    
    if [ -f "$json_file" ]; then
        if command -v jq &> /dev/null; then
            jq -r '.abi' "$json_file" > "$abi_file"
            echo "✓ Extracted ${contract}.abi"
        else
            echo "Error: jq is required but not installed. Install with: sudo apt-get install jq"
            exit 1
        fi
    else
        echo "Warning: ${json_file} not found. Run 'forge build' first."
    fi
done

echo "Done! ABI files are ready in ${OUT_DIR}/*.sol/*.abi"
