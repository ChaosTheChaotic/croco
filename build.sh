#!/usr/bin/env bash

set -e

# Base directories
ANDROID_NDK_HOME="/opt/android-sdk/ndk"
NDK_TOOLCHAIN="$ANDROID_NDK_HOME/27.1.12297006/toolchains/llvm/prebuilt/linux-x86_64/bin"
OUTPUT_BASE="build"

# Create output base directory
mkdir -p "$OUTPUT_BASE"

# Array of target architectures with their corresponding CC and GOARCH values
architectures=(
    "arm64:aarch64-linux-android33-clang:arm64"
    "arm:armv7a-linux-androideabi33-clang:arm"
    "x86_64:x86_64-linux-android33-clang:amd64"
    "x86:i686-linux-android33-clang:386"
)

echo "Building for all Android architectures..."

for arch_info in "${architectures[@]}"; do
    # Split the architecture info into separate variables
    IFS=':' read -r arch_name cc_binary goarch <<< "$arch_info"
    
    output_dir="$OUTPUT_BASE/$arch_name"
    mkdir -p "$output_dir"
    
    echo "Building for $arch_name..."
    
    # Set environment variables and build
    CC="$NDK_TOOLCHAIN/$cc_binary" \
    GOOS="android" \
    GOARCH="$goarch" \
    CGO_ENABLED=1 \
    go build -buildmode=c-shared -o "$output_dir/libcroco.so" ./croco_clib.go
    
    echo "✓ Built $output_dir/libcroco.so"
done

echo "All builds completed successfully!"
echo "Output files are organized in: $OUTPUT_BASE/"
