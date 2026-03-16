import os
import cohere
from dotenv import load_dotenv

load_dotenv()

co = cohere.ClientV2(os.environ["COHERE_API_KEY"])

documents = [
    {
        "data": {
            "title": "会社概要",
            "text": "私たちはAIを活用したソフトウェア開発を行う企業です。2020年に設立されました。"
        }
    },
    {
        "data": {
            "title": "サービス",
            "text": "主なサービスはチャットボット開発、データ分析、AI導入支援です。"
        }
    },
]

response = co.chat(
    model="command-a-03-2025",
    messages=[
        {
            "role": "system",
            "content": "与えられた文書に基づいて回答してください。不明な点は推測せず不明と答えてください。"
        },
        {
            "role": "user",
            "content": "この会社はどんなサービスを提供していますか？"
        }
    ],
    documents=documents,
)

# 本文表示
if response.message.content:
    print(response.message.content[0].text)

# citation 表示
if getattr(response.message, "citations", None):
    for citation in response.message.citations:
        print(citation)