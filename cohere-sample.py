
import os
import cohere
from dotenv import load_dotenv

load_dotenv()
co = cohere.Client(os.environ['COHERE_API_KEY'])
response = co.chat(
    model='command-a-03-2025',
    message='AIでできることを教えて',
    max_tokens=100
)
print(response.text)