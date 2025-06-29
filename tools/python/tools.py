def read_prompt_from_file(filepath: str) -> str:
    """
    Reads a prompt from a file, stripping a potential markdown header and surrounding whitespace.
    """
    with open(filepath, "r", encoding="utf-8") as f:
        lines = f.readlines()

    # Remove leading empty lines
    while lines and not lines[0].strip():
        lines.pop(0)

    # If the first non-empty line is a header, remove it
    if lines and lines[0].strip().startswith("#"):
        lines.pop(0)

    return "".join(lines).strip()
