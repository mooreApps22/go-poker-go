for file in internal/poker/*.go; do
    echo
    echo "==================== $file ===================="
    cat "$file"
done
