require "cohere"
require "dotenv/load"

co = Cohere::Client.new(api_key: ENV["COHERE_API_KEY"])

documents = [
  {
    data: {
      title: "会社概要",
      text: "私たちはAIを活用したソフトウェア開発を行う企業です。2020年に設立されました。"
    }
  },
  {
    data: {
      title: "サービス",
      text: "主なサービスはチャットボット開発、データ分析、AI導入支援です。"
    }
  }
]

body = co.chat(
  model: "command-a-03-2025",
  messages: [
    {
      role: "system",
      content: "与えられた文書に基づいて回答してください。不明な点は推測せず不明と答えてください。"
    },
    {
      role: "user",
      content: "この会社はどんなサービスを提供していますか？"
    }
  ],
  documents: documents
)

# 本文表示
content = body.dig("message", "content")
if content
  text_block = content.find { |c| c["type"] == "text" }
  puts text_block["text"] if text_block
end

# citation 表示
citations = body.dig("message", "citations")
if citations
  citations.each { |citation| p citation }
end
