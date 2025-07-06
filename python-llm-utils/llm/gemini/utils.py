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


def read_text_from_file(filepath: str) -> str:
    """
    Reads the entire content of a text file and returns it as a string.

    Args:
        filepath: The path to the file to be read.

    Returns:
        The content of the file as a string.

    Raises:
        FileNotFoundError: If the file is not found.
        IOError: If any other I/O error occurs.
    """
    with open(filepath, "r", encoding="utf-8") as f:
        return f.read()
