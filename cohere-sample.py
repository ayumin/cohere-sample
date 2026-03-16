import os
import cohere
from dotenv import load_dotenv

load_dotenv()
co = cohere.Client(os.environ['COHERE_API_KEY'])

# 参照させたいドキュメントを用意
documents = [
    {"title": "会社概要", "snippet": "私たちはAIを活用したソフトウェア開発を行う企業です。2020年に設立されました。"},
    {"title": "サービス", "snippet": "主なサービスはチャットボット開発、データ分析、AI導入支援です。"},
]

response = co.chat(
    model='command-a-03-2025',
    message='この会社はどんなサービスを提供していますか？',
    documents=documents,
)

print(response.text)

# 引用元も確認できる
for citation in response.citations:
    print(f"引用: {citation}")