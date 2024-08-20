import os


def count_lines_of_code(directory, extensions):
    total_lines = 0
    for root, _, files in os.walk(directory):
        if "node_modules" in root:
            continue  # Skip the node_modules directory
        for file in files:
            if any(file.endswith(ext) for ext in extensions):
                with open(os.path.join(root, file), "r") as f:
                    total_lines += sum(1 for line in f)
    return total_lines


# Usage
extensions = [
    ".py",
    ".js",
    ".html",
    ".ts",
    ".tsx",
    ".go",
]  # Add your file extensions here
directory = "."  # Start from the current directory
print(f"Total lines of code: {count_lines_of_code(directory, extensions)}")
