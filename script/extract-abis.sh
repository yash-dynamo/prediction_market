#!/bin/bash
# Extract ABI files from Foundry JSON output and generate Go bindings

set -e

OUT_DIR="out"
GO_BINDINGS_DIR="go_bindings"
CONTRACTS=("MockUSDC" "OutcomeToken" "PredictionMarket")

# Find abigen (check common locations)
ABIGEN=""
if command -v abigen &> /dev/null; then
    ABIGEN="abigen"
elif [ -f "$HOME/go/bin/abigen" ]; then
    ABIGEN="$HOME/go/bin/abigen"
elif [ -f "$GOPATH/bin/abigen" ]; then
    ABIGEN="$GOPATH/bin/abigen"
fi

# Check dependencies
if ! command -v jq &> /dev/null; then
    echo "Error: jq is required but not installed. Install with: sudo apt-get install jq"
    exit 1
fi

if [ -z "$ABIGEN" ]; then
    echo "Warning: abigen not found. Install with: go install github.com/ethereum/go-ethereum/cmd/abigen@latest"
    echo "Skipping Go bindings generation..."
    GENERATE_GO=false
else
    GENERATE_GO=true
    mkdir -p "$GO_BINDINGS_DIR"
    echo "Using abigen: $ABIGEN"
fi

# Extract ABIs and generate Go bindings
for contract in "${CONTRACTS[@]}"; do
    json_file="${OUT_DIR}/${contract}.sol/${contract}.json"
    abi_file="${OUT_DIR}/${contract}.sol/${contract}.abi"
    bin_file="${OUT_DIR}/${contract}.sol/${contract}.bin"
    go_file="${GO_BINDINGS_DIR}/${contract,,}.go"  # lowercase contract name
    
    if [ -f "$json_file" ]; then
        # Extract ABI
        jq -r '.abi' "$json_file" > "$abi_file"
        echo "✓ Extracted ${contract}.abi"
        
        # Extract bytecode
        jq -r '.bytecode.object' "$json_file" > "$bin_file"
        echo "✓ Extracted ${contract}.bin"
        
        # Generate Go bindings
        if [ "$GENERATE_GO" = true ]; then
            "$ABIGEN" \
                --abi "$abi_file" \
                --bin "$bin_file" \
                --pkg bindings \
                --type "$contract" \
                --out "$go_file" 2>&1 | grep -v "WARN" || true
            
            if [ -f "$go_file" ] && [ -s "$go_file" ]; then
                echo "✓ Generated ${go_file}"
            else
                echo "⚠ Failed to generate ${go_file}"
            fi
        fi
    else
        echo "Warning: ${json_file} not found. Run 'forge build' first."
    fi
done

echo ""
echo "Done! ABI files are ready in ${OUT_DIR}/*.sol/*.abi"
if [ "$GENERATE_GO" = true ]; then
    echo "Go bindings are ready in ${GO_BINDINGS_DIR}/"
fi