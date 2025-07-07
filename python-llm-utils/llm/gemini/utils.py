from pathlib import Path
from typing import Union

from google.genai.types import GenerateContentResponse


def read_prompt_from_file(filepath: str) -> str:
    """Reads a prompt from a file, stripping a potential markdown header.

    This function reads a file, removes any leading blank lines, and if the
    first line of content is a markdown header (starts with '#'), it removes
    that line as well. The remaining content is returned as a single string
    with leading/trailing whitespace stripped.

    Args:
        filepath: The path to the prompt file.

    Returns:
        The processed content of the prompt file.

    Raises:
        FileNotFoundError: If the file is not found.
        IOError: If any other I/O error occurs.
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


def read_text_from_file(filepath: Union[str, Path]) -> str:
    """Reads the entire content of a text file and returns it as a string.

    Args:
        filepath: The path to the file to be read.

    Returns:
        The content of the file as a string.

    Raises:
        FileNotFoundError: If the file is not found.
        IOError: If any other I/O error occurs.
    """
    return Path(filepath).read_text(encoding="utf-8")


def write_gemini_text_to_markdown(
    response: GenerateContentResponse, output_path: str
) -> None:
    """Extracts text from a Gemini response and writes it to a markdown file.

    Args:
        response: The GenerateContentResponse object from the Gemini API.
        output_path: The path to the output markdown file.

    Raises:
        ValueError: If the response from the model is invalid (e.g., blocked),
            empty, or contains no text. This is typically raised by the
            `response.text` property.
        IOError: If there is an error writing to the file.
    """
    try:
        raw_text = response.text
        with open(output_path, "w", encoding="utf-8") as f:
            f.write(raw_text)
    except ValueError as e:
        raise ValueError(
            "Invalid or empty response from model, cannot extract text."
        ) from e
    except IOError as e:
        raise IOError(f"Failed to write markdown file to '{output_path}'") from e
