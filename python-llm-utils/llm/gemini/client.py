import os

from dotenv import find_dotenv, load_dotenv
from google import genai


def create_genai_client() -> genai.Client:
    """
    Initializes and returns a Google GenAI client using the GEMINI_API_KEY from .env.

    Raises:
        ValueError: If the GEMINI_API_KEY is not set in environment variables.
    """
    load_dotenv(find_dotenv())

    api_key = os.getenv("GEMINI_API_KEY")
    if not api_key:
        raise ValueError("Missing GEMINI_API_KEY in environment variables.")

    return genai.Client(api_key=api_key)
